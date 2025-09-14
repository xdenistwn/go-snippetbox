package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `
        INSERT INTO users
        (
            name, email, hashed_password, created
        ) 
        VALUES
        (
            $1, $2, $3, NOW()
        )
    `
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var pgSQLError *pgconn.PgError

		if errors.As(err, &pgSQLError) {
			if pgSQLError.Code == "23505" && strings.Contains(pgSQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}

		return err
	}

	return nil
}

func (m *UserModel) Authenticate(name, email, password string) error {
	return nil
}

func (m *UserModel) Exists(name, email, password string) error {
	return nil
}
