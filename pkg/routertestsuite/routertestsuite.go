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

// RouterTest performance standard router tests.
type RouterTest struct{}

// New returns a router test suite.
func New() *RouterTest {
	return new(RouterTest)
}

// Run all the tests.
func (rt *RouterTest) Run(t *testing.T, mux func() ambient.AppRouter) {
	rt.TestParams(t, mux())
	rt.TestInstance(t, mux())
	rt.TestPostForm(t, mux())
	rt.TestPostJSON(t, mux())
	rt.TestGet(t, mux())
	rt.TestDelete(t, mux())
	rt.TestHead(t, mux())
	rt.TestOptions(t, mux())
	rt.TestPatch(t, mux())
	rt.TestPut(t, mux())
	rt.Test404(t, mux())
	rt.Test500NoError(t, mux())
	rt.Test500WithError(t, mux())
	rt.Test400(t, mux())
	rt.TestNotFound(t, mux())
	rt.TestBadRequest(t, mux())
	rt.TestClear(t, mux())
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
func (rt *RouterTest) TestParams(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestInstance(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestPostForm(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestPostJSON(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestGet(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestDelete(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestHead(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestOptions(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestPatch(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestPut(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) Test404(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) Test500NoError(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) Test500WithError(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) Test400(t *testing.T, mux ambient.AppRouter) {
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
func (rt *RouterTest) TestNotFound(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	r := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()
	mux.Error(http.StatusNotFound, w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestBadRequest .
func (rt *RouterTest) TestBadRequest(t *testing.T, mux ambient.AppRouter) {
	mux.SetServeHTTP(defaultServeHTTP)

	r := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()
	mux.Error(http.StatusBadRequest, w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestClear .
func (rt *RouterTest) TestClear(t *testing.T, mux ambient.AppRouter) {
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
