package main

import (
	"net/http"
	"testing"

	"snippetbox.stwn.dev/internal/assert"
)

// func TestPing(t *testing.T) {
// 	rr := httptest.NewRecorder()

// 	r, err := http.NewRequest(http.MethodGet, "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ping(rr, r)

// 	rs := rr.Result()
// 	assert.Equal(t, rs.StatusCode, http.StatusOK)

// 	defer rs.Body.Close()
// 	body, err := io.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	body = bytes.TrimSpace(body)

// 	assert.Equal(t, string(body), "OK")
// }

// end to end test sample
// func TestPing(t *testing.T) {
// 	// Create a new instance of our application struct.
// 	app := &application{
// 		logger: slog.New(slog.DiscardHandler),
// 	}

// 	ts := httptest.NewTLSServer(app.routes())
// 	defer ts.Close()

// 	// The network address that the test server is listening on is contained in
// 	// the ts.URL field. We can use this along with the ts.Client().Get() method
// 	// to make a GET /ping request against the test server. This returns a
// 	// http.Response struct containing the response.
// 	rs, err := ts.Client().Get(ts.URL + "/ping")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// We can then check the value of the response status code and body using
// 	// the same pattern as before.
// 	assert.Equal(t, rs.StatusCode, http.StatusOK)

// 	defer rs.Body.Close()
// 	body, err := io.ReadAll(rs.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	body = bytes.TrimSpace(body)
// 	assert.Equal(t, string(body), "OK")
// }

// test using testutils helper
func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
