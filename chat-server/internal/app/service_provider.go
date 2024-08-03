package app

import (
	"context"
	"log"

	"github.com/bifidokk/awesome-chat/chat-server/internal/api/chat"
	"github.com/bifidokk/awesome-chat/chat-server/internal/closer"
	"github.com/bifidokk/awesome-chat/chat-server/internal/config"
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository"
	chatRepository "github.com/bifidokk/awesome-chat/chat-server/internal/repository/chat"
	"github.com/bifidokk/awesome-chat/chat-server/internal/service"
	chatService "github.com/bifidokk/awesome-chat/chat-server/internal/service/chat"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	pgPool     *pgxpool.Pool

	chatRepository repository.ChatRepository

	chatService service.ChatService

	chatAPI *chat.API
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

func (sp *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		sp.chatRepository = chatRepository.NewRepository(sp.PgPool(ctx))
	}

	return sp.chatRepository
}

func (sp *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if sp.chatService == nil {
		sp.chatService = chatService.NewChatService(sp.ChatRepository(ctx))
	}

	return sp.chatService
}

func (sp *serviceProvider) ChatAPI(ctx context.Context) *chat.API {
	if sp.chatAPI == nil {
		sp.chatAPI = chat.NewChatAPI(sp.ChatService(ctx))
	}

	return sp.chatAPI
}
