package cmd

import (
	"github.com/spf13/cobra"
	"webhook/internal/clients"
	hookHandler2 "webhook/internal/hookHandler"
	"webhook/internal/mongodb"
	"webhook/internal/redis"
	"webhook/pkg/log"
)

type hookHandler struct {
	instance hookHandler2.HookRootService
	command  *cobra.Command
	env      string
}

var hookHandlerRoot *hookHandler

func init() {
	hookHandlerRoot = &hookHandler{
		command: &cobra.Command{
			Use:   "hook-handler",
			Short: "hook handler",
			Long:  `hook handler`,
			RunE:  startupHandler,
		},
		//default values
		env: "test",
	}
	hookHandlerRoot.command.Flags().StringVarP(&hookHandlerRoot.env, "env", "e", hookHandlerRoot.env, "select your env.")
	RootCmd.AddCommand(hookHandlerRoot.command) // RootCmd kullanılabilir olmalı
}

func startupHandler(cmd *cobra.Command, args []string) error {
	logger := log.SetupLogger()
	logger.Infof("hook handler is starting")

	redisClient := redis.NewRedisClient(&logger)
	hookApiClient := clients.NewHookApiClient()
	hookRepository := mongodb.NewHookRepository(&logger)
	hookService := hookHandler2.NewHookHandlerService(hookRepository, hookApiClient, redisClient, &logger)
	hookRootService := hookHandler2.NewHookRootService(&logger, hookService)
	hookHandlerRoot.instance = hookRootService
	if err := hookHandlerRoot.instance.Start(); err != nil {
		logger.Errorf("hook handler failed to start: | Error: %s", err.Error())
	}
	return nil
}
