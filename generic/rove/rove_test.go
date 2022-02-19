package rove_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var err error
	// docker run --name=mysql57 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:5.7
	// docker rm mysql57 -f
	// os.Setenv("DB_USERNAME", "root")
	// os.Setenv("DB_PASSWORD", "password")
	// os.Setenv("DB_HOSTNAME", "localhost")
	// os.Setenv("DB_PORT", "3306")
	// os.Setenv("DB_NAME", "main")
	// p := rove.New(nil)
	// err = p.Enable(nil)
	assert.NoError(t, err)
}
