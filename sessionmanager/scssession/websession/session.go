// Package websession provides session management.
package websession

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexedwards/scs/v2"
	"github.com/ambientkit/plugin/pkg/uuid"
)

// Session stores session level information
type Session struct {
	Name    string
	manager *scs.SessionManager
}

// New returns a new session cookie store.
func New(name string, manager *scs.SessionManager) *Session {
	return &Session{
		manager: manager,
		Name:    name,
	}
}

// Persist sets session to persist after browser is closed.
func (s *Session) Persist(r *http.Request, persist bool) {
	s.manager.RememberMe(r.Context(), persist)
}

// Logout and destroy session.
func (s *Session) Logout(r *http.Request) {
	s.manager.Destroy(r.Context())
}

// LogoutAll destroys all sessions.
func (s *Session) LogoutAll(r *http.Request) error {
	return s.manager.Iterate(r.Context(), func(ctx context.Context) error {
		userID := s.manager.GetString(ctx, "userID")

		// Only destroy authenticated sessions.
		if len(userID) > 0 {
			return s.manager.Destroy(ctx)
		}

		return nil
	})
}

// AuthenticatedUser returns the user ID if authenticated or an error.
func (s *Session) AuthenticatedUser(r *http.Request) (string, error) {
	userID := s.manager.GetString(r.Context(), "userID")

	if len(userID) == 0 {
		return "", errors.New("user not found")
	}

	return userID, nil
}

// Login user by storing user ID in request context.
func (s *Session) Login(r *http.Request, value string) {
	s.manager.Put(r.Context(), "userID", value)
}

// SessionValue returns a value stored in the user session.
func (s *Session) SessionValue(r *http.Request, name string) string {
	return s.manager.GetString(r.Context(), name)
}

// SetSessionValue sets a value in the user session or returns an error.
func (s *Session) SetSessionValue(r *http.Request, name string, value string) error {
	if strings.EqualFold("userID", name) {
		return fmt.Errorf("cannot set reserved 'userID' value using SetSessionValue")
	}

	s.manager.Put(r.Context(), name, value)

	return nil
}

// DeleteSessionValue deletes a value in the user session.
func (s *Session) DeleteSessionValue(r *http.Request, name string) {
	s.manager.Remove(r.Context(), name)
}

// SetCSRF sets a cross site request forgery token for the current request to
// allow for validation during form submission.
func (s *Session) SetCSRF(r *http.Request) string {
	token := uuid.RandomString(32)
	path := "csrf_" + r.URL.Path
	s.SetSessionValue(r, path, token)
	return token
}

// CSRF return true if the cross site request forgery token matches what was
// stored before form submission.
func (s *Session) CSRF(r *http.Request, token string) bool {
	path := "csrf_" + r.URL.Path
	v := s.SessionValue(r, path)

	if len(v) > 0 {
		s.manager.Remove(r.Context(), path)
		if v == token && len(token) > 0 {
			return true
		}
	}

	return false
}
