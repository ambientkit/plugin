// Package scssession is an Ambient plugin that provides session management using SCS.
package scssession

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/pkg/aesdata"
	"github.com/ambientkit/plugin/pkg/jsonstore"
	"github.com/ambientkit/plugin/sessionmanager/scssession/websession"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase

	sessionManager *scs.SessionManager
	sess           *websession.Session

	sessionKey string
}

// New returns an Ambient plugin that provides session management using SCS.
func New(sessionKey string) *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
		sessionKey: sessionKey,
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName() string {
	return "scssession"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{
		{Grant: ambient.GrantRouterMiddlewareWrite, Description: "Access to read and write session data for the user."},
	}
}

const (
	// SessionKey allows user to set the session key.
	SessionKey = "Session Key"
)

// Settings returns a list of user settable fields.
func (p *Plugin) Settings() []ambient.Setting {
	return []ambient.Setting{
		{
			Name:    SessionKey,
			Type:    ambient.InputPassword,
			Default: p.sessionKey,
			Hide:    true,
		},
	}
}

// Middleware returns router middleware.
func (p *Plugin) Middleware() []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		p.sessionManager.LoadAndSave,
	}
}

// SessionManager returns the session manager.
func (p *Plugin) SessionManager(logger ambient.Logger, ss ambient.SessionStorer) (ambient.AppSession, error) {
	// Set up the session storage provider.
	en := aesdata.NewEncryptedStorage(p.sessionKey)
	store, err := jsonstore.NewJSONSession(ss, en)
	if err != nil {
		return nil, err
	}

	sessionName := "session"

	p.sessionManager = scs.New()
	p.sessionManager.Lifetime = 24 * time.Hour
	// This must be false for RememberMe() to work.
	p.sessionManager.Cookie.Persist = false
	p.sessionManager.Store = store
	p.sess = websession.New(sessionName, p.sessionManager)

	return p.sess, nil
}
