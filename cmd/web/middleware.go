package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

// const contextUserKey = contextKey("user_ip")

const contextUserKey contextKey = "user_ip"

func (app *application) ipFromContext(ctx context.Context) string {
	ip, ok := ctx.Value(contextUserKey).(string)
	if !ok {
		return ""
	}
	return ip
}

func (app *application) addIPToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ip, err := getIP(r)
		if err != nil {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}

			ctx = context.WithValue(ctx, contextUserKey, ip)
		} else {
			ctx = context.WithValue(ctx, contextUserKey, ip)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getIP(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unkown", err
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}

	// if the header X-Forwarded-For is set, then the request came through a proxy
	forward := r.Header.Get("X-Forwarded-For")
	if len(forward) > 0 {
		// this will be the first IP address in the list, which is the original client IP
		ip = forward
	}

	if len(ip) == 0 {
		// if the IP address is empty we return something
		ip = "forward"
	}

	return ip, nil
}
