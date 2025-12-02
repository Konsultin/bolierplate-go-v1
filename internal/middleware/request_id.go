package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/valyala/fasthttp"
)

const requestIDKey = "requestID"

func RequestID() func(fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			id := incomingRequestID(ctx)
			if id == "" {
				id = newRequestID()
			}

			ctx.SetUserValue(requestIDKey, id)
			ctx.Response.Header.Set("X-Request-ID", id)

			next(ctx)
		}
	}
}

func RequestIDFromContext(ctx *fasthttp.RequestCtx) string {
	if v := ctx.UserValue(requestIDKey); v != nil {
		if id, ok := v.(string); ok {
			return id
		}
	}
	return ""
}

func incomingRequestID(ctx *fasthttp.RequestCtx) string {
	h := strings.TrimSpace(string(ctx.Request.Header.Peek("X-Request-ID")))
	if h == "" {
		return ""
	}
	return h
}

func newRequestID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	return hex.EncodeToString(b[:])
}
