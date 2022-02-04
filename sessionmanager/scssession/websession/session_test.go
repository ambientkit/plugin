package websession_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/aesdata"
	"github.com/ambientkit/plugin/sessionmanager/scssession/websession"
	"github.com/ambientkit/plugin/storage/localstorage/store"
	"github.com/stretchr/testify/assert"
)

var storageFile = "data.bin"

func setup(t *testing.T) (ambient.AppSession, func(next http.Handler) http.Handler) {
	// Set up the session storage provider.
	err := ioutil.WriteFile(storageFile, []byte(""), 0644)
	assert.NoError(t, err)
	ss := store.NewLocalStorage(storageFile)
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	store, err := websession.NewJSONSession(ss, en)
	assert.NoError(t, err)

	// Initialize a new session manager and configure the session lifetime.
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = false
	sessionManager.Store = store
	sess := websession.New("session", sessionManager)

	return sess, sessionManager.LoadAndSave
}

func teardown() {
	// Clean up.
	os.Remove(storageFile)
}

func TestNewSession(t *testing.T) {
	sess, sessHandler := setup(t)
	defer teardown()

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

	// Test session persistence.
	r = httptest.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Test user
		u := "foo"
		sess.Login(r, u)

		// Ensure the expiration is set properly.
		sess.Persist(r, true)
	})

	mw = sessHandler(mux)
	mw.ServeHTTP(w, r)
	assert.False(t, w.Result().Cookies()[0].Expires.IsZero())
}
