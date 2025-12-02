package middleware

import "github.com/valyala/fasthttp"

func Chain(final fasthttp.RequestHandler, middlewares ...func(fasthttp.RequestHandler) fasthttp.RequestHandler) fasthttp.RequestHandler {
	handler := final
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
