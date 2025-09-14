package main

// 247

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/v5/stdlib"
	"snippetbox.stwn.dev/internal/models"
)

type application struct {
	env            string
	logger         *slog.Logger
	snippets       *models.SnippetModel
	users          *models.UserModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	// parse command parameter
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://web:snippet123web@localhost:5432/snippetbox?sslmode=disable", "PostgreSQL data source")
	flag.Parse()

	// standard go logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// connect db
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Connected to PostgreSQL Database.")
	defer db.Close()

	// init session redis for session manager
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	sessionManager := scs.New()
	sessionManager.Store = redisstore.New(redisPool)
	sessionManager.Lifetime = 12 * time.Hour

	// init template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// initialize a form decoder instance
	formDecoder := form.NewDecoder()

	// init application struct by DI
	app := &application{
		env:            os.Getenv("APP_ENV"),
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db}, // use constructor function instead
		users:          &models.UserModel{DB: db},    // use constructor function instead
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// server configuration
	srv := &http.Server{
		Addr:      *addr,
		Handler:   app.routes(),
		ErrorLog:  slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: tlsConfig,
		// idle, read and write timeout to server
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// maximum header size 0.5MB
		MaxHeaderBytes: 524288,
	}

	logger.Info("starting server", "addr", *addr)

	err = srv.ListenAndServeTLS("tls/cert.pem", "tls/key.pem")
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
