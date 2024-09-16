package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"
	hookApi2 "webhook/docs/hookApi"
	"webhook/internal/contentApi"
	"webhook/internal/hookApi"
	"webhook/internal/mongodb"
	"webhook/pkg/echoExtension"
	"webhook/pkg/log"
)

type hookApiManager struct {
	instance *echo.Echo
	command  *cobra.Command
	env      string
	port     string
}

var hookApiRoot *hookApiManager

func init() {
	hookApiRoot = &hookApiManager{
		command: &cobra.Command{
			Use:   "hook-api",
			Short: "hook api",
			Long:  `hook api`,
			RunE:  startupApi,
		},
		//default values
		env:  "test",
		port: "5030",
	}
	hookApiRoot.command.Flags().StringVarP(&hookApiRoot.env, "env", "e", hookApiRoot.env, "select your env.")
	hookApiRoot.command.Flags().StringVarP(&hookApiRoot.port, "port", "p", hookApiRoot.port, "service port")
	RootCmd.AddCommand(hookApiRoot.command)
}

// @title         Mehti Umit - WebHook Api
// @version       1.0
// @description   Mehti Umit - WebHook Api
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes       http
// @BasePath      /webhook/api
func startupApi(cmd *cobra.Command, args []string) error {
	logger := log.SetupLogger()
	logger.Infof("hook api is starting on port: %s", hookApiRoot.port)
	hookApiRoot.instance = echo.New()
	hookApiRoot.instance.Debug = false
	hookApiRoot.instance.HideBanner = false
	hookApiRoot.instance.HidePort = false
	hookApiRoot.instance.Logger = &logger
	hookApi2.SwaggerInfohookApi.Host = "localhost:" + hookApiRoot.port
	echoExtension.RegisterGlobalMiddlewares(hookApiRoot.instance, "/hook-api", make(map[string]string))
	contentRepository := mongodb.NewContentRepository(&logger)
	hookRepository := mongodb.NewHookRepository(&logger)
	hookService := hookApi.NewHookService(hookRepository, contentRepository, &logger)
	contentService := contentApi.NewContentService(contentRepository, &logger)
	contentApi.NewContentHandler(hookApiRoot.instance, contentService, &logger)
	hookApi.NewHookHandler(hookApiRoot.instance, hookService, &logger)
	hookApiRoot.instance.GET("/webhook/swagger/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("hookApi")))
	logger.Infof("swagger url: http://localhost:%s/webhook/swagger/index.html", hookApiRoot.port)
	go func() {
		if err := hookApiRoot.instance.Start(":" + hookApiRoot.port); err != nil {
			logger.Errorf("failed to start hook api: %s", err.Error())
		}
	}()
	echoExtension.Shutdown(hookApiRoot.instance, 2*time.Second)
	return nil
}
