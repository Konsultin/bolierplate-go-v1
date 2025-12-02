package svcCore

import (
	"time"

	"github.com/Konsultin/project-goes-here/dto"
	f "github.com/valyala/fasthttp"
)

func (s *Server) HealthCheck(ctx *f.RequestCtx) {
	uptime := time.Since(s.startedAt).String()

	data := dto.HealthData{
		Status:       "ok",
		Uptime:       uptime,
		Started:      s.startedAt.UTC().Format(time.RFC3339),
		Env:          s.config.Env,
		Hostname:     string(ctx.Request.URI().Host()),
		Dependencies: map[string]string{},
	}

	resp := dto.Response[dto.HealthData]{
		Message:   "liveness ok",
		Code:      dto.CodeOK,
		Data:      data,
		Timestamp: time.Now().UTC().UnixMilli(),
	}

	s.log.Debugf("Ran Health Check: %+v", resp)

	s.response(ctx, f.StatusOK, resp)
}
