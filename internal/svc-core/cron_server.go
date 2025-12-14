package svcCore

import (
	"encoding/base64"
	"strings"

	"github.com/valyala/fasthttp"
)

func (s *Server) validateCronAuth(ctx *fasthttp.RequestCtx) bool {
	auth := string(ctx.Request.Header.Peek("Authorization"))
	if auth == "" {
		return false
	}

	const prefix = "Basic "
	if len(auth) < len(prefix) || auth[:len(prefix)] != prefix {
		return false
	}

	payload, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == s.config.CronUsername && pair[1] == s.config.CronPassword
}

func (s *Server) HandleCronTrigger(ctx *fasthttp.RequestCtx) {
	if !s.validateCronAuth(ctx) {
		ctx.Error("Unauthorized", fasthttp.StatusUnauthorized)
		return
	}

	cronType := ctx.UserValue("cronType").(string)
	s.log.Infof("Received cron trigger: %s", cronType)

	switch cronType {
	case "example":
		s.log.Info("Running example cron job...")
	default:
		s.log.Warnf("Unknown cron type: %s", cronType)
		ctx.Error("Unknown cron type", fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("OK")
}
