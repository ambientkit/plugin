package gorillastore

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ambientkit/plugin/pkg/aesdata"
	"github.com/ambientkit/plugin/pkg/filestore"
	"github.com/ambientkit/plugin/pkg/jsonstore"

	"github.com/stretchr/testify/assert"
)

// Test for GH-8 for FilesystemStore
func TestGH8FilesystemStore(t *testing.T) {
	// Use local filesytem when developing.
	f := "data.bin"
	err := ioutil.WriteFile(f, []byte(""), 0644)
	assert.NoError(t, err)
	ss := filestore.New(f)

	// Set up the session storage provider.
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	jsonstore, err := jsonstore.NewJSONSession(ss, en)
	assert.NoError(t, err)
	defer os.Remove(f)

	store := NewGenericStore(jsonstore, []byte(""))

	originalPath := "/"
	store.Options.Path = originalPath
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}

	session, err := store.New(req, "hello")
	if err != nil {
		t.Fatal("failed to create session", err)
	}

	store.Options.Path = "/foo"
	if session.Options.Path != originalPath {
		t.Fatalf("bad session path: got %q, want %q", session.Options.Path, originalPath)
	}
}

// Test for GH-2.
func TestGH2MaxLength(t *testing.T) {
	// Use local filesytem when developing.
	f := "data.bin"
	err := ioutil.WriteFile(f, []byte(""), 0644)
	assert.NoError(t, err)
	ss := filestore.New(f)

	// Set up the session storage provider.
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	jsonstore, err := jsonstore.NewJSONSession(ss, en)
	assert.NoError(t, err)
	defer os.Remove(f)

	store := NewGenericStore(jsonstore, []byte("some key"))

	//store := NewGenericStore("", []byte("some key"))
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	w := httptest.NewRecorder()

	session, err := store.New(req, "my session")
	if err != nil {
		t.Fatal("failed to create session", err)
	}

	session.Values["big"] = make([]byte, base64.StdEncoding.DecodedLen(4096*2))
	err = session.Save(req, w)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	store.MaxLength(4096 * 3) // A bit more than the value size to account for encoding overhead.
	err = session.Save(req, w)
	if err != nil {
		t.Fatal("failed to Save:", err)
	}
}

// Test delete filesystem store with max-age: -1
func TestGH8FilesystemStoreDelete(t *testing.T) {
	// Use local filesytem when developing.
	f := "data.bin"
	err := ioutil.WriteFile(f, []byte(""), 0644)
	assert.NoError(t, err)
	ss := filestore.New(f)

	// Set up the session storage provider.
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	jsonstore, err := jsonstore.NewJSONSession(ss, en)
	assert.NoError(t, err)
	defer os.Remove(f)

	store := NewGenericStore(jsonstore, []byte("some key"))

	//store := NewGenericStore("", []byte("some key"))
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	w := httptest.NewRecorder()

	session, err := store.New(req, "hello")
	if err != nil {
		t.Fatal("failed to create session", err)
	}

	err = session.Save(req, w)
	if err != nil {
		t.Fatal("failed to save session", err)
	}

	session.Options.MaxAge = -1
	err = session.Save(req, w)
	if err != nil {
		t.Fatal("failed to delete session", err)
	}
}

// Test delete filesystem store with max-age: 0
func TestGH8FilesystemStoreDelete2(t *testing.T) {
	// Use local filesytem when developing.
	f := "data.bin"
	err := ioutil.WriteFile(f, []byte(""), 0644)
	assert.NoError(t, err)
	ss := filestore.New(f)

	// Set up the session storage provider.
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	jsonstore, err := jsonstore.NewJSONSession(ss, en)
	assert.NoError(t, err)
	defer os.Remove(f)

	store := NewGenericStore(jsonstore, []byte("some key"))

	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	w := httptest.NewRecorder()

	session, err := store.New(req, "hello")
	if err != nil {
		t.Fatal("failed to create session", err)
	}

	err = session.Save(req, w)
	if err != nil {
		t.Fatal("failed to save session", err)
	}

	session.Options.MaxAge = 0
	err = session.Save(req, w)
	if err != nil {
		t.Fatal("failed to delete session", err)
	}
}
