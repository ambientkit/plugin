package websession_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/aesdata"
	"github.com/ambientkit/plugin/pkg/filestore"
	"github.com/ambientkit/plugin/pkg/jsonstore"
	"github.com/ambientkit/plugin/pkg/sessmgrtestsuite"
	"github.com/ambientkit/plugin/sessionmanager/scssession/websession"
	"github.com/stretchr/testify/assert"
)

var storageFile = "data.bin"

// Run the standard router test suite.
func TestMain(t *testing.T) {
	ts := sessmgrtestsuite.New()

	ts.Run(t, func() (ambient.AppSession, func(next http.Handler) http.Handler) {
		return setup(t)
	})

	teardown()
}

func setup(t *testing.T) (ambient.AppSession, func(next http.Handler) http.Handler) {
	// Set up the session storage provider.
	err := ioutil.WriteFile(storageFile, []byte(""), 0644)
	assert.NoError(t, err)
	ss := filestore.New(storageFile)
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	store, err := jsonstore.NewJSONSession(ss, en)
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
