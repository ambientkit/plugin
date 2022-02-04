package routertestsuite

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/stretchr/testify/assert"
)

// TestSuite performs standard tests.
type TestSuite struct{}

// New returns a router test suite.
func New() *TestSuite {
	return new(TestSuite)
}

// Run all the tests.
func (ts *TestSuite) Run(t *testing.T, mux func() ambient.AppRouter) {
	ts.TestParams(t, mux())
	ts.TestInstance(t, mux())
	ts.TestPostForm(t, mux())
	ts.TestPostJSON(t, mux())
	ts.TestGet(t, mux())
	ts.TestDelete(t, mux())
	ts.TestHead(t, mux())
	ts.TestOptions(t, mux())
	ts.TestPatch(t, mux())
	ts.TestPut(t, mux())
	ts.Test404(t, mux())
	ts.Test500NoError(t, mux())
	ts.Test500WithError(t, mux())
	ts.Test400(t, mux())
	ts.TestNotFound(t, mux())
	ts.TestBadRequest(t, mux())
	ts.TestClear(t, mux())
}

// defaultServeHTTP is the default ServeHTTP function that receives the status and error from
// the function call.
var defaultServeHTTP = func(w http.ResponseWriter, r *http.Request, status int,
	err error) {
	if status >= 400 {
		if err != nil {
			http.Error(w, err.Error(), status)
		} else {
			http.Error(w, "", status)
		}
	}
}

// TestParams .
func (ts *TestSuite) TestParams(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	outParam := ""
	mux.Get("/user/{name}", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		outParam = mux.Param(r, "name")
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("GET", "/user/john", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "john", outParam)
}

// TestInstance .
func (ts *TestSuite) TestInstance(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	outParam := ""
	mux.Get("/user/{name}", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		outParam = mux.Param(r, "name")
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("GET", "/user/john", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "john", outParam)
}

// TestPostForm .
func (ts *TestSuite) TestPostForm(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	form := url.Values{}
	form.Add("username", "jsmith")

	outParam := ""
	mux.Post("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		r.ParseForm()
		outParam = r.FormValue("username")
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("POST", "/user", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "jsmith", outParam)
}

// TestPostJSON .
func (ts *TestSuite) TestPostJSON(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	j, err := json.Marshal(map[string]interface{}{
		"username": "jsmith",
	})
	assert.Nil(t, err)

	outParam := ""
	mux.Post("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		b, err := ioutil.ReadAll(r.Body)
		assert.Nil(t, err)
		r.Body.Close()
		outParam = string(b)
		assert.Equal(t, `{"username":"jsmith"}`, string(b))
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("POST", "/user", bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, `{"username":"jsmith"}`, outParam)
}

// TestGet .
func (ts *TestSuite) TestGet(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Get("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

// TestDelete .
func (ts *TestSuite) TestDelete(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Delete("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("DELETE", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

// TestHead .
func (ts *TestSuite) TestHead(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Head("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("HEAD", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

// TestOptions .
func (ts *TestSuite) TestOptions(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Options("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("OPTIONS", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

// TestPatch .
func (ts *TestSuite) TestPatch(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Patch("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("PATCH", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

// TestPut .
func (ts *TestSuite) TestPut(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Put("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("PUT", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
}

// Test404 .
func (ts *TestSuite) Test404(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Get("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("GET", "/badroute", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, false, called)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test500NoError .
func (ts *TestSuite) Test500NoError(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := true

	mux.Get("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusInternalServerError, nil
	})

	r := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Test500WithError .
func (ts *TestSuite) Test500WithError(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := true
	specificError := errors.New("specific error")

	mux.Get("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusInternalServerError, specificError
	})

	r := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, true, called)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, w.Body.String(), specificError.Error()+"\n")
}

// Test400 .
func (ts *TestSuite) Test400(t *testing.T, mux ambient.AppRouter) {
	notFound := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	},
	)

	mux.SetServeHTTP(defaultServeHTTP)
	mux.SetNotFound(notFound)

	r := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestNotFound .
func (ts *TestSuite) TestNotFound(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	r := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()
	mux.Error(http.StatusNotFound, w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestBadRequest .
func (ts *TestSuite) TestBadRequest(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	r := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()
	mux.Error(http.StatusBadRequest, w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestClear .
func (ts *TestSuite) TestClear(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	called := false

	mux.Get("/user", func(w http.ResponseWriter, r *http.Request) (status int, err error) {
		called = true
		return http.StatusOK, nil
	})

	r := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, called)

	called = false
	mux.Clear("GET", "/user")

	r = httptest.NewRequest("GET", "/user", nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp = w.Result()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.False(t, called)
}
