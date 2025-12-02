package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Konsultin/project-goes-here/config"
	"github.com/Konsultin/project-goes-here/dto"
	"github.com/Konsultin/project-goes-here/internal/middleware"
	svcCore "github.com/Konsultin/project-goes-here/internal/svc-core"
	"github.com/Konsultin/project-goes-here/libs/errk"
	"github.com/Konsultin/project-goes-here/libs/logk"
	logkOption "github.com/Konsultin/project-goes-here/libs/logk/option"
	"github.com/Konsultin/project-goes-here/libs/routek"
	"github.com/valyala/fasthttp"
)

type responder struct {
	debug bool
}

func newResponder(debug bool) responder {
	return responder{debug: debug}
}

func (r responder) success(ctx *fasthttp.RequestCtx, status int, code dto.Code, message string, data any) {
	r.write(ctx, status, code, message, data)
}

func (r responder) error(ctx *fasthttp.RequestCtx, status int, code dto.Code, message string, err error) {
	var data any
	if r.debug && err != nil {
		data = map[string]any{"error": err.Error()}
	}
	r.write(ctx, status, code, message, data)
}

func (r responder) write(ctx *fasthttp.RequestCtx, status int, code dto.Code, message string, data any) {
	resp := dto.Response[any]{
		Message:   message,
		Code:      code,
		Data:      data,
		Timestamp: time.Now().UTC().UnixMilli(),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"message":"internal server error","code":"INTERNAL_ERROR","data":null,"timestamp":0}`)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(status)
	ctx.SetBody(body)
}

func konsultinAscii() string {
	return `
'     __  _   ___   ____   _____ __ __  _     ______  ____  ____       ___      ___ __ __ 
'    |  |/ ] /   \ |    \ / ___/|  |  || |   |      ||    ||    \     |   \    /  _]  |  |
'    |  ' / |     ||  _  (   \_ |  |  || |   |      | |  | |  _  |    |    \  /  [_|  |  |
'    |    \ |  O  ||  |  |\__  ||  |  || |___|_|  |_| |  | |  |  |    |  D  ||    _]  |  |
'    |     ||     ||  |  |/  \ ||  :  ||     | |  |   |  | |  |  | __ |     ||   [_|  :  |
'    |  .  ||     ||  |  |\    ||     ||     | |  |   |  | |  |  ||  ||     ||     |\   / 
'    |__|\_| \___/ |__|__| \___| \__,_||_____| |__|  |____||__|__||__||_____||_____| \_/  
'      
'    Boilerplate created by Kenly Krisaguino - @kenly.krisaguino on Instagram
'	 Version: 1.0.0
'                                                                                         
	`
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		logk.Get().Fatal("Failed to load config", logkOption.Error(errk.Trace(err)))
	}
	startedAt := time.Now()
	rootLog := logk.Get().NewChild(logkOption.WithNamespace("api"))
	rootLog.Infof("API starting... env=%s", cfg.Env)

	fmt.Println(konsultinAscii())

	coreServer, err := svcCore.New(cfg, startedAt)
	if err != nil {
		rootLog.Fatal("Failed to init core server", logkOption.Error(errk.Trace(err)))
	}
	defer func() {
		if err := coreServer.Close(); err != nil {
			rootLog.Error("Failed to close resources", logkOption.Error(errk.Trace(err)))
		}
	}()

	rt, err := routek.NewRouter(routek.Config{
		Handlers: map[string]any{
			"core": coreServer,
		},
	})
	if err != nil {
		rootLog.Fatal("Failed to init router", logkOption.Error(errk.Trace(err)))
	}

	responder := newResponder(cfg.Debug)
	handler, err := middleware.Init(middleware.Config{
		Handler:          rt.Handler,
		Logger:           rootLog,
		OnError:          responder.error,
		RateLimitRPS:     cfg.RateLimitRPS,
		RateLimitBurst:   cfg.RateLimitBurst,
		CORSAllowOrigins: cfg.CORSAllowOrigins,
	})
	if err != nil {
		rootLog.Fatal("Failed to init middleware", logkOption.Error(errk.Trace(err)))
	}

	server := &fasthttp.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.HTTPReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.HTTPWriteTimeoutSeconds) * time.Second,
		IdleTimeout:  time.Duration(cfg.HTTPIdleTimeoutSeconds) * time.Second,
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	errCh := make(chan error, 1)

	go func() {
		rootLog.Infof("Listening on %s", addr)
		if err := server.ListenAndServe(addr); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-stop:
		rootLog.Infof("Received signal %s, shutting down", sig)
	case err := <-errCh:
		if err != nil {
			rootLog.Fatal("Server error", logkOption.Error(errk.Trace(err)))
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.ShutdownWithContext(shutdownCtx); err != nil {
		rootLog.Error("Graceful shutdown failed", logkOption.Error(errk.Trace(err)))
	} else {
		rootLog.Info("Server stopped gracefully")
	}
}
