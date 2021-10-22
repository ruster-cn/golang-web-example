package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/op-server/pkg/server/config"

	"github.com/op-server/pkg/server/globalrouter"

	dao2 "github.com/op-server/pkg/server/dao"

	"github.com/op-server/pkg/server/orm"

	"github.com/op-server/pkg/logger"
)

type PaasHTTPServer struct {
	httpServerConfiguration *config.HTTPServerConfiguration
	ormConfiguration        *orm.MysqlConnectConfiguration
}

func NewPaasHTTPSever(httpConfig *config.HTTPServerConfiguration, ormConfig *orm.MysqlConnectConfiguration) *PaasHTTPServer {
	return &PaasHTTPServer{httpServerConfiguration: httpConfig, ormConfiguration: ormConfig}
}

func (server *PaasHTTPServer) Start(stopCh <-chan struct{}) error {
	ormEngine, err := orm.NewMysqlEngine(server.ormConfiguration)
	if err != nil {
		return fmt.Errorf("new mysql engine fail,%v", err)
	}
	dao := dao2.NewBaseDao(ormEngine)
	routers := globalrouter.NewGinEngine(dao, server.httpServerConfiguration)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", server.httpServerConfiguration.Addr, server.httpServerConfiguration.Port),
		Handler: routers,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-stopCh

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), server.httpServerConfiguration.GracefulTimeOut)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: ", err)
	}

	logger.Info("Server exiting")
	return nil
}
