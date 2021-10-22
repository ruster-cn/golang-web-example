package pkg

import (
	"github.com/hashicorp/go-multierror"
	"github.com/op-server/pkg/server"
)

type PaasServerMain struct {
	paasHTTPServer *server.PaasHTTPServer
}

func NewPaasServerMain(config *Configuration) (*PaasServerMain, error) {
	paasHTTPServer := server.NewPaasHTTPSever(config.PaasServer, config.Orm)

	return &PaasServerMain{paasHTTPServer: paasHTTPServer}, nil
}

func (main *PaasServerMain) Start(stop <-chan struct{}) error {
	var errs error

	if err := main.paasHTTPServer.Start(stop); err != nil {
		errs = multierror.Append(errs, err)
	}

	return errs
}
