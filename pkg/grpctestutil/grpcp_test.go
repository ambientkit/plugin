package grpctestutil_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/ambient/pkg/ambientapp"
	"github.com/ambientkit/plugin/pkg/grpctestutil"
	"github.com/stretchr/testify/assert"
)

func TestGRPC(t *testing.T) {
	// Setup gRPC server.
	app := grpcSetup(t)
	// Stop plugins when done.
	defer app.StopGRPCClients()

	tests(t, app, setGrants(t, app))
}

func TestStandard(t *testing.T) {
	// Setup standard server.
	app := standardSetup(t)

	tests(t, app, setGrants(t, app))
}

func grpcSetup(t assert.TestingT) *ambientapp.App {
	// Set the test relative to the project directory since the plugin path
	// is relative to that.
	path, _ := os.Getwd()
	basePath := strings.TrimSuffix(path, "/pkg/grpctestutil")
	if err := os.Chdir(basePath); err != nil {
		assert.FailNow(t, err.Error())
	}

	// Set up the application.
	app, err := grpctestutil.GRPCSetup(false)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	return app
}

func standardSetup(t assert.TestingT) *ambientapp.App {
	// Set the test relative to the project directory since the plugin path
	// is relative to that.
	path, _ := os.Getwd()
	basePath := strings.TrimSuffix(path, "/pkg/grpctestutil")
	if err := os.Chdir(basePath); err != nil {
		assert.FailNow(t, err.Error())
	}

	// Set up the application.
	app, err := grpctestutil.StandardSetup(false)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	return app
}

func doRequest(t assert.TestingT, mux http.Handler, r *http.Request) (*http.Response, string) {
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	return resp, string(body)
}

func setGrants(t assert.TestingT, app *ambientapp.App) http.Handler {
	ps := app.PluginSystem()
	assert.NoError(t, ps.SetEnabled("hello", true))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantRouterRouteWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantUserAuthenticatedRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantUserAuthenticatedWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginNeighborGrantRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginNeighborGrantWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePluginRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePluginDelete))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePluginEnable))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePluginDisable))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePostWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePostRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSitePostDelete))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginNeighborRouteRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantUserPersistWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantAllUserAuthenticatedWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginNeighborSettingRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginSettingWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginSettingRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginNeighborSettingWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantPluginTrustedRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteTitleWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteTitleRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteSchemeWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteSchemeRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteURLWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteURLRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteUpdatedRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteContentRead))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteContentWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteFuncMapWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantSiteAssetWrite))
	assert.NoError(t, ps.SetGrant("hello", ambient.GrantRouterMiddlewareWrite))

	mux, err := app.Handler()
	if err != nil {
		t.Errorf("could not create handler: %v", err.Error())
	}

	return mux
}

func tests(t *testing.T, app *ambientapp.App, mux http.Handler) {
	resp, body := doRequest(t, mux, httptest.NewRequest("GET", "/", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "hello world", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/another", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "hello world - another", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/name/foo", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "hello: foo", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/name/bar", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "hello: bar", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/nameold/foo", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "hello: foo", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/error", nil))
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	assert.Equal(t, "Forbidden\n", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/created", nil))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "created: ", string(body))

	r := httptest.NewRequest("GET", "/headers", nil)
	r.Header.Set("foo", "123")
	r.Header.Set("bar", "who")
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "headers: 2", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/form", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "\n<!DOCTYPE html>\n<html lang=\"en\">\n<head></head>\n<body>\n\t<form method=\"post\">\n\t<label for=\"fname\">First name:</label>\n\t<input type=\"text\" id=\"fname\" name=\"fname\" value=\"a\"><br><br>\n\t<label for=\"lname\">Last name:</label>\n\t<input type=\"text\" id=\"lname\" name=\"lname\" value=\"b\"><br><br>\n\t<input type=\"submit\" value=\"Submit\">\n\t</form>\n</body>\n</html>\n", string(body))

	form := url.Values{}
	form.Add("a", "foo")
	form.Add("b", "bar")
	r = httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, `body: "a=foo&b=bar"`, string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/loggedin", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "login: () (user not found)", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/login", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "login: (<nil>) (username) (<nil>)", string(body))

	// Test with authenticated cookie.
	r = httptest.NewRequest("GET", "/loggedin", nil)
	for _, v := range resp.Cookies() {
		r.AddCookie(v)
	}
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "login: (username) (<nil>)", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/errors", nil))
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	assert.Equal(t, "request does not exist for the grant\n", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginGrantList", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Grants: 18", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginGrantListBad", nil))
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "item was not found\n", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginGrants", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Grants: 18", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginGranted", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Granted: false", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginGrantedBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Granted: false", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/setNeighborPluginGrantFalse", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Granted: false", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/setNeighborPluginGrantTrue", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Granted: true", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginGranted", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Granted: true", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginRequestedGrant", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Requested: true", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/neighborPluginRequestedGrantBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Requested: false", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/plugins", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Plugins: 4", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/pluginNames", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Plugin names: 4", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("DELETE", "/deletePlugin", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Delete plugin: <nil>", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("DELETE", "/deletePluginBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Delete plugin: plugin name not found", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/enablePlugin", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Enable plugin: <nil>", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/enablePluginBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Enable plugin: item was not found", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/disablePlugin", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Disable plugin: <nil>", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/disablePluginBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Disable plugin: item was not found", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("POST", "/savePost", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Posts are the same.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/publishedPosts", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Posts are the same.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/publishedPages", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Pages are the same.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/postBySlug", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Pages are the same.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/postBySlugBad", nil))
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "item was not found\n", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/postByID", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Pages are the same.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/postByIDBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Pages are the same.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("DELETE", "/deletePostByID", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Works.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/pluginNeighborRoutesList", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Routes: 1", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/pluginNeighborRoutesListBad", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Routes: 0", string(body))

	resp, _ = doRequest(t, mux, httptest.NewRequest("GET", "/userPersist", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 86400, resp.Cookies()[0].MaxAge)

	resp, _ = doRequest(t, mux, httptest.NewRequest("GET", "/userPersistFalse", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 0, resp.Cookies()[0].MaxAge)

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/grantRequests", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Grant requests: 19", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/userLogout", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "User cleared.", string(body))
	assert.Equal(t, "", resp.Cookies()[0].Value)
	assert.Equal(t, -1, resp.Cookies()[0].MaxAge)

	// Login user.
	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/login", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "login: (<nil>) (username) (<nil>)", string(body))

	// Test with authenticated cookie.
	r = httptest.NewRequest("GET", "/loggedin", nil)
	for _, v := range resp.Cookies() {
		r.AddCookie(v)
	}
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "login: (username) (<nil>)", string(body))

	// Destroy users.
	r = httptest.NewRequest("GET", "/logoutAllUsers", nil)
	for _, v := range resp.Cookies() {
		r.AddCookie(v)
	}
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Users cleared.", string(body))

	// Test with authenticated cookie again.
	r = httptest.NewRequest("GET", "/loggedin", nil)
	for _, v := range resp.Cookies() {
		r.AddCookie(v)
	}
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "login: () (user not found)", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/csrf", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 32, len(body))

	form = url.Values{}
	form.Add("token", body)
	r = httptest.NewRequest("POST", "/csrf", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range resp.Cookies() {
		r.AddCookie(v)
	}
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Token is valid.", string(body))

	// Try same request again and should fail.
	r = httptest.NewRequest("POST", "/csrf", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range resp.Cookies() {
		r.AddCookie(v)
	}
	resp, body = doRequest(t, mux, r)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "token is not valid\n", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/sessionValue", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Session value works.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/PluginNeighborSettingsList", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Neighbor settings: 10", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/setPluginSetting", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Plugin setting works.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/setNeighborPluginSetting", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Plugin neighbor setting works.", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/pluginTrusted", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Plugin trusted: false true", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/title", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Site title: foo", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/scheme", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Site scheme: https", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/url", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Site URL: bar | Full URL: https://bar", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/updated", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Site updated: "+time.Now().UTC().Format("20060102"), string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/content", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Site content: foo bar", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/tags", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Site tags: tag1 "+time.Now().UTC().Format("20060102"), string(body))

	// resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/assets", nil))
	// assert.Equal(t, http.StatusOK, resp.StatusCode)
	// assert.Equal(t, "Site assets: []ambient.Asset{ambient.Asset{Filetype:\"generic\", Location:\"head\", Auth:\"\", Attributes:[]ambient.Attribute(nil), LayoutOnly:[]ambient.LayoutType(nil), TagName:\"title\", ClosingTag:false, External:false, Inline:true, SkipExistCheck:false, Path:\"\", Content:\"{{if .pagetitle}}{{.pagetitle}} | foo{{else}}foo{{end}}\", Replace:[]ambient.Replace(nil)}}", string(body))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/assetsHello", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, strings.Contains(body, "FuncMap: hello: Foo"))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/assetsError", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, strings.Contains(body, "this is an error"))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/pageHello", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, strings.Contains(body, "FuncMap: hello: Foo"), body)
	assert.True(t, strings.Contains(body, `<link rel="canonical" href="cool">`), body)

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/api/healthcheck", nil))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, `{"message":"ok"}`, body)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	// resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/context", nil))
	// assert.Equal(t, http.StatusOK, resp.StatusCode)
	// assert.Equal(t, "context is foundf", body)
	// assert.Equal(t, "foo", hello.Get(r))

	resp, body = doRequest(t, mux, httptest.NewRequest("GET", "/redirect", nil))
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, `<a href="/redirectTo">Found</a>.`+"\n\n", body)
}
