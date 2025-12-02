package svcCore

import (
	"encoding/json"

	"github.com/Konsultin/project-goes-here/libs/errk"
	logkOption "github.com/Konsultin/project-goes-here/libs/logk/option"
	f "github.com/valyala/fasthttp"
)

// s.response writes a JSON response with a consistent header and handles marshal errors gracefully.
func (s *Server) response(ctx *f.RequestCtx, statusCode int, payload any) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(statusCode)

	body, err := json.Marshal(payload)
	if err != nil {
		if s.log != nil {
			s.log.Error("failed to marshal response", logkOption.Error(errk.Trace(err)))
		}
		ctx.SetStatusCode(f.StatusInternalServerError)
		ctx.SetBodyString(`{"message":"internal server error","code":"INTERNAL_ERROR","data":null,"timestamp":0}`)
		return
	}

	ctx.SetBody(body)
}
