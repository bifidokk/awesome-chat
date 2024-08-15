package app

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/auth/internal/api/user"
	userApi "github.com/bifidokk/awesome-chat/auth/internal/api/user"
	"github.com/bifidokk/awesome-chat/auth/internal/client/db"
	"github.com/bifidokk/awesome-chat/auth/internal/client/db/pg"
	"github.com/bifidokk/awesome-chat/auth/internal/client/db/transaction"
	"github.com/bifidokk/awesome-chat/auth/internal/closer"
	"github.com/bifidokk/awesome-chat/auth/internal/config"
	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	userRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user"
	"github.com/bifidokk/awesome-chat/auth/internal/service"
	userService "github.com/bifidokk/awesome-chat/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository

	userService service.UserService

	userAPI *user.API
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) PgConfig() config.PGConfig {
	if sp.pgConfig == nil {
		pgConfig, err := config.NewPGConfig()

		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		sp.pgConfig = pgConfig
	}

	return sp.pgConfig
}

func (sp *serviceProvider) GrpcConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		grpcConfig, err := config.NewGRPCConfig()

		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		sp.grpcConfig = grpcConfig
	}

	return sp.grpcConfig
}

func (sp *serviceProvider) HTTPConfig() config.HTTPConfig {
	if sp.httpConfig == nil {
		httpConfig, err := config.NewHTTPConfig()

		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		sp.httpConfig = httpConfig
	}

	return sp.httpConfig
}

func (sp *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if sp.swaggerConfig == nil {
		swaggerConfig, err := config.NewSwaggerConfig()

		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		sp.swaggerConfig = swaggerConfig
	}

	return sp.swaggerConfig
}

func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		client, err := pg.New(ctx, sp.PgConfig().DSN())

		if err != nil {
			log.Fatalf("failed to create DB client: %v", err)
		}

		err = client.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		closer.Add(client.Close)

		sp.dbClient = client
	}

	return sp.dbClient
}

func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		txManager := transaction.NewTransactionManager(sp.DBClient(ctx).DB())

		sp.txManager = txManager
	}

	return sp.txManager
}

func (sp *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.userRepository
}

func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewUserService(sp.UserRepository(ctx))
	}

	return sp.userService
}

func (sp *serviceProvider) UserAPI(ctx context.Context) *user.API {
	if sp.userAPI == nil {
		sp.userAPI = userApi.NewUserAPI(sp.UserService(ctx))
	}

	return sp.userAPI
}
