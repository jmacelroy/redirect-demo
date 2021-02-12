package client

import (
	"context"
	"errors"
	"net/http"
)

const divertHeaderName = "x-okteto-redirect"

type divertHeaderKey struct{}

// DivertHeaderCtxKey is the unique key value for diver header
// value injected into context.
var DivertHeaderCtxKey = divertHeaderKey{}

// PropagateFromContext will retrieve divert headers from the context provided
// and add the correct headers to the provided http request so that they can
// be propagated to the subsequent http request in a service chain.
func PropagateFromContext(ctx context.Context, r *http.Request) *http.Request {
	divertHeaderValue, _ := FromContext(r.Context())
	if divertHeaderValue != "" {
		r.Header.Set(divertHeaderName, divertHeaderValue)
	}
	return r
}

// FromContext provides the divert header values stored in context.
func FromContext(ctx context.Context) (string, error) {
	value, ok := ctx.Value(DivertHeaderCtxKey).(string)
	if !ok {
		return "", errors.New("unable to convert header value")
	}
	return value, nil
}

// FromHeaders extracts divert headers from an http request
// and provides the value. If missing then empty string
// is provided.
func FromHeaders(r *http.Request) string {
	return r.Header.Get(divertHeaderName)
}

// InjectDivertHeaderContext is an http middleware handler for
// adding Okteto divert headers into context.
func InjectDivertHeaderContext() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), DivertHeaderCtxKey, FromHeaders(r))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
