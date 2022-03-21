package main

import (
	"math/rand"
	_ "net/http/pprof"
	"time"

	"github.com/op-server/cmd/app"
	log "github.com/ruster-cn/zap-log-wrapper"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewPaasServerCommand()
	if err := command.Execute(); err != nil {
		log.Fatalf("paas server start fail,%v \n", err)
	}
}
