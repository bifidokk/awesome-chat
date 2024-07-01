package main

import (
	"context"
	"flag"
	"github.com/bifidokk/awesome-chat/chat-server/internal/config"
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository"
	"github.com/bifidokk/awesome-chat/chat-server/internal/repository/chat"
	"log"
	"net"

	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

var configPath string

type server struct {
	desc.UnimplementedChatV1Server
	chatRepository repository.ChatRepository
}

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	listener, err := net.Listen("tcp", grpcConfig.Address())

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	chatRepository := chat.NewRepository(pool)
	desc.RegisterChatV1Server(s, &server{chatRepository: chatRepository})

	log.Printf("Server listening at %v", listener.Addr())

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create a new chat: %v", req)

	chatId, err := s.chatRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send a message: %v", req)

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat: %v", req)

	err := s.chatRepository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
