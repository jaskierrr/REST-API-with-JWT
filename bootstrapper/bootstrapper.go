package bootstrapper

import (
	"context"
	"log"
	"log/slog"
	"main/config"
	"main/internal/controller"
	"main/internal/database"
	"main/internal/handlers"
	logger "main/logger"
	repo "main/internal/repositories"
	"main/restapi"
	"main/restapi/operations"
	"main/internal/service"

	"github.com/go-openapi/loads"
	"github.com/go-playground/validator/v10"
)

type RootBootstrapper struct {
	Infrastructure struct {
		Logger *slog.Logger
		Server *restapi.Server
		DB     database.DB
	}
	Controller controller.Controller
	Config     *config.Config
	Handlers   handlers.Handlers
	Repository repo.Repository
	Service    service.Service

	Validator *validator.Validate
}

type RootBoot interface {
	registerRepositoriesAndServices(ctx context.Context, db database.DB)
	registerAPIServer(cfg config.Config) error
	RunAPI() error
}

func New() RootBoot {
	return &RootBootstrapper{
		Config: config.NewConfig(),
	}
}

func (r *RootBootstrapper) RunAPI() error {
	ctx := context.Background()
	r.Infrastructure.Logger = logger.NewLogger()

	r.registerRepositoriesAndServices(ctx, r.Infrastructure.DB)
	err := r.registerAPIServer(*r.Config)
	if err != nil {
		log.Fatal("cant start server")
	}

	return nil
}

func (r *RootBootstrapper) registerRepositoriesAndServices(ctx context.Context, db database.DB) {
	logger := r.Infrastructure.Logger
	r.Infrastructure.DB = database.NewDB().NewConn(ctx, *r.Config, logger)
	r.Repository = repo.NewUserRepo(r.Infrastructure.DB, logger)
	r.Service = service.New(r.Repository, logger)
}

func (r *RootBootstrapper) registerAPIServer(cfg config.Config) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewCryptoAPI(swaggerSpec)

	logger := r.Infrastructure.Logger

	r.Controller = controller.New(r.Service, logger)

	r.Validator = validator.New(validator.WithRequiredStructEnabled())

	r.Handlers = handlers.New(r.Controller, r.Validator, logger)
	r.Handlers.Link(api)
	if r.Handlers == nil {
		log.Fatal("handlers initialization failed")
	}

	r.Infrastructure.Server = restapi.NewServer(api)
	r.Infrastructure.Server.Port = cfg.ServerPort
	r.Infrastructure.Server.ConfigureAPI()
	if err := r.Infrastructure.Server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return nil
}
