package sessmgrtestsuite

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/stretchr/testify/assert"
)

// TestSuite performs standard tests.
type TestSuite struct{}

// New returns a session manager test suite.
func New() *TestSuite {
	return new(TestSuite)
}

// Run all the tests.
func (ts *TestSuite) Run(t *testing.T, sess func() (ambient.AppSession, func(next http.Handler) http.Handler)) {
	s, sessHandler := sess()
	ts.TestSessions(t, s, sessHandler)

	s, sessHandler = sess()
	ts.TestPersist(t, s, sessHandler)
}

// TestSessions .
func (ts *TestSuite) TestSessions(t *testing.T, sess ambient.AppSession, sessHandler func(next http.Handler) http.Handler) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test user
		u := "foo"
		sess.Login(r, u)
		user, err := sess.AuthenticatedUser(r)
		assert.True(t, err == nil)
		assert.Equal(t, u, user)

		// Test Logout
		sess.Logout(r)
		_, err = sess.AuthenticatedUser(r)
		assert.False(t, err == nil)

		// Test CSRF
		assert.False(t, sess.CSRF(r, r.FormValue("token")))
		token := sess.SetCSRF(r)
		r.Form = url.Values{}
		r.Form.Set("token", token)
		assert.True(t, sess.CSRF(r, r.FormValue("token")))
	})

	mw := sessHandler(mux)
	mw.ServeHTTP(w, r)

	// Ensure the expiration is set properly.
	assert.True(t, w.Result().Cookies()[0].Expires.IsZero())
}

// TestPersist .
func (ts *TestSuite) TestPersist(t *testing.T, sess ambient.AppSession, sessHandler func(next http.Handler) http.Handler) {
	// Test session persistence.
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test user
		u := "foo"
		sess.Login(r, u)

		// Ensure the expiration is set properly.
		sess.Persist(r, true)
	})

	mw := sessHandler(mux)
	mw.ServeHTTP(w, r)
	assert.False(t, w.Result().Cookies()[0].Expires.IsZero())
}
