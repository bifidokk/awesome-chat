package app

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/auth/internal/api/user"
	userApi "github.com/bifidokk/awesome-chat/auth/internal/api/user"
	"github.com/bifidokk/awesome-chat/auth/internal/closer"
	"github.com/bifidokk/awesome-chat/auth/internal/config"
	"github.com/bifidokk/awesome-chat/auth/internal/repository"
	userRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user"
	"github.com/bifidokk/awesome-chat/auth/internal/service"
	userService "github.com/bifidokk/awesome-chat/auth/internal/service/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	pgPool     *pgxpool.Pool

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

func (sp *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if sp.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, sp.PgConfig().DSN())

		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		closer.Add(
			func() error {
				pool.Close()
				return nil
			})

		sp.pgPool = pool
	}

	return sp.pgPool
}

func (sp *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepository.NewRepository(sp.PgPool(ctx))
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
