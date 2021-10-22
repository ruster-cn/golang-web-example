package main

import (
	"math/rand"
	_ "net/http/pprof"
	"time"

	"github.com/op-server/pkg/logger"

	"github.com/op-server/cmd/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewPaasServerCommand()
	if err := command.Execute(); err != nil {
		logger.Fatalf("paas server start fail,%v \n", err)
	}
}
