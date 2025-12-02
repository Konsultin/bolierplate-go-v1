package middleware

import (
	"strings"

	"github.com/valyala/fasthttp"
)

func CORS(allowedOrigins []string) func(fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			origin := string(ctx.Request.Header.Peek("Origin"))

			if originAllowed(origin, allowedOrigins) {
				if origin == "" && len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
					ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
				} else {
					ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
					ctx.Response.Header.Set("Vary", "Origin")
				}
				ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			}

			ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization,Content-Type")

			if ctx.IsOptions() {
				ctx.SetStatusCode(fasthttp.StatusNoContent)
				return
			}

			next(ctx)
		}
	}
}

func originAllowed(origin string, allowed []string) bool {
	if origin == "" {
		return true
	}
	for _, allowedOrigin := range allowed {
		if allowedOrigin == "*" || strings.EqualFold(allowedOrigin, origin) {
			return true
		}
	}
	return false
}
