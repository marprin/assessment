package http

import (
	"github.com/marprin/assessment/fetchapp/internal/config"
	"github.com/marprin/assessment/fetchapp/internal/domain/storage/repository/currency"
	"github.com/marprin/assessment/fetchapp/internal/domain/storage/repository/gateway"
	storageSvc "github.com/marprin/assessment/fetchapp/internal/domain/storage/service"
	"github.com/marprin/assessment/fetchapp/internal/domain/user/service"
	"github.com/marprin/assessment/fetchapp/internal/handler"
	"github.com/marprin/assessment/fetchapp/internal/middleware"
	"github.com/marprin/assessment/fetchapp/internal/route"
	"github.com/marprin/assessment/fetchapp/pkg/cache"
	"github.com/marprin/assessment/fetchapp/pkg/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:   "serve",
		Short: "Fetch App HTTP Service",
		RunE:  run,
	}
)

func run(cmd *cobra.Command, args []string) error {
	logrus.Infoln("Starting initialize http service")

	cfg := &config.Config{}
	config.ReadConfig(cfg, "main")

	jwt := jwt.New(cfg.Token.Secret, cfg.Token.Issuer)
	cache := cache.New()
	middlewareRepo := middleware.New(&middleware.Options{
		JwtRepo: jwt,
	})
	gatewayRepo := gateway.New(cfg.Gateway.URL)
	currencyRepo := currency.New(cfg.Currency.URL, cfg.Currency.APIKey)
	userService := service.New(cfg, jwt)
	storageService := storageSvc.New(&storageSvc.Options{
		GatewayRepo:  gatewayRepo,
		CurrencyRepo: currencyRepo,
		Cache:        cache,
	})

	handler := handler.New(&handler.Options{
		StorageService: storageService,
		UserService:    userService,
	})

	server := route.New(&route.Options{
		Config:     cfg,
		Handler:    handler,
		Middleware: middlewareRepo,
	})
	server.Run()

	logrus.Infoln("Terminate http service")

	return nil
}

// ServeHTTPCmd is the exposed function
func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}
