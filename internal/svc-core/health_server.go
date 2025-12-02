package svcCore

import (
	"time"

	"github.com/Konsultin/project-goes-here/dto"
	f "github.com/valyala/fasthttp"
)

func (s *Server) HealthCheck(ctx *f.RequestCtx) (*dto.HealthData, error) {
	uptime := time.Since(s.startedAt).String()

	data := dto.HealthData{
		Status:       "ok",
		Uptime:       uptime,
		Started:      s.startedAt.UTC().Format(time.RFC3339),
		Env:          s.config.Env,
		Hostname:     string(ctx.Request.URI().Host()),
		Dependencies: map[string]string{},
	}

	s.log.Debugf("Ran Health Check: %+v", data)

	return &data, nil
}
