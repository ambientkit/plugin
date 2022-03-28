package grpctestutil_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/mock"
	"github.com/go-playground/assert/v2"
)

func TestLoad(t *testing.T) {
	// Load the ambient application.
	app := standardSetup(t)
	ps := app.PluginSystem()
	mux, err := app.Handler()
	if err != nil {
		t.Fatal(err.Error())
	}
	ss := app.SecureSite()

	// Load a new plugin.
	pluginName := "test"
	p := mock.NewPlugin(pluginName, "1.0.0")
	p.MockRoutes = func(m *ambient.PluginBase) {
		m.Mux.Get("/", func(w http.ResponseWriter, r *http.Request) error {
			fmt.Fprint(w, "cool dude!")
			return nil
		})
	}
	ps.LoadPlugin(p, true, false)
	ss.LoadSinglePluginPages(pluginName)
	ps.SetEnabled(pluginName, true)
	ps.SetGrant(pluginName, ambient.GrantRouterRouteWrite)

	// Test the endpoint.
	resp, body := doRequest(t, mux, httptest.NewRequest("GET", "/", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "cool dude!", string(body))

}
