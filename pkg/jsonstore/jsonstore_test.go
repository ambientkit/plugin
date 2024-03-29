package jsonstore

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/ambientkit/plugin/pkg/aesdata"
	"github.com/ambientkit/plugin/pkg/filestore"
	"github.com/stretchr/testify/assert"
)

func TestNewJSONSession(t *testing.T) {
	// Use local filesytem when developing.
	f := "data.bin"
	err := ioutil.WriteFile(f, []byte(""), 0644)
	assert.NoError(t, err)
	ss := filestore.New(f)

	// Set up the session storage provider.
	secretkey := "82a18fbbfed2694bb15d512a70c53b1a088e669966918d3d474564b2ac44349b"
	en := aesdata.NewEncryptedStorage(secretkey)
	store, err := NewJSONSession(ss, en)
	assert.NoError(t, err)

	token := "abc"
	data := "hello"

	// Set date in the future.
	err = store.Commit(token, []byte(data), time.Now().AddDate(0, 0, 1))
	assert.NoError(t, err)

	b, exists, err := store.Find(token)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, data, string(b))

	err = store.Delete(token)
	assert.NoError(t, err)

	_, exists, err = store.Find(token)
	assert.NoError(t, err)
	assert.False(t, exists)

	// Set date in the past.
	err = store.Commit(token, []byte(data), time.Now().AddDate(0, 0, -1))
	assert.NoError(t, err)
	b, exists, err = store.Find(token)
	assert.Nil(t, b)
	assert.False(t, exists)
	assert.NoError(t, err)

	os.Remove(f)
}
