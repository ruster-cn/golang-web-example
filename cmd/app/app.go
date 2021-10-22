package app

import (
	"fmt"

	"github.com/op-server/pkg/logger"

	"github.com/op-server/pkg"
	"k8s.io/sample-controller/pkg/signals"

	"github.com/spf13/cobra"
)

var configPath string

//NewPaasServerCommand make paas server cobra's command
func NewPaasServerCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "op-server",
		Short: "op-server is multi cloud and multi kubernetes's cluster management platform ",
		RunE: func(cmd *cobra.Command, args []string) error {
			mianServer, err := setup()
			if err != nil {
				return fmt.Errorf("setup paas server fail,%v", err)
			}
			if err := run(mianServer); err != nil {
				return fmt.Errorf("run paas server fail,%v", err)
			}
			return nil
		},
	}
	flags(command)
	return command
}

//flags parsing command line parameters
func flags(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&configPath, "conf", "c", "./config.yaml", "the path of config")
}

//setup init op-server
//TODO: init paasServerMain
func setup() (*pkg.PaasServerMain, error) {
	config, err := pkg.NewConfigurationFromFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := logger.NewLogger(config.Log); err != nil {
		return nil, fmt.Errorf("new logger fail,%v", err)
	}

	serverMain, err := pkg.NewPaasServerMain(config)
	if err != nil {
		return nil, err
	}
	return serverMain, nil

}

//run start op-server
//TODO: start paasServerMain
func run(main *pkg.PaasServerMain) error {
	stopCh := signals.SetupSignalHandler()

	if err := main.Start(stopCh); err != nil {
		return err
	}

	return nil
}
