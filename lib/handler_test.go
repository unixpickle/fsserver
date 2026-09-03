package fsserver

import (
	"net/http/httptest"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	h := &Handler{BasicAuth: "user:password:with:colons"}

	for _, tc := range []struct {
		name     string
		username string
		password string
		want     bool
	}{
		{"correct", "user", "password:with:colons", true},
		{"wrong username", "other", "password:with:colons", false},
		{"wrong password", "user", "other", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", nil)
			r.SetBasicAuth(tc.username, tc.password)
			if got := h.authenticate(r); got != tc.want {
				t.Fatalf("authenticate() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestAuthenticateRejectsMalformedConfiguration(t *testing.T) {
	h := &Handler{BasicAuth: "missing-separator"}
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("missing-separator", "")
	if h.authenticate(r) {
		t.Fatal("authenticate() accepted malformed configuration")
	}
}
