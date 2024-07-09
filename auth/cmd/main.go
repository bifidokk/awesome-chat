package main

import (
	"context"
	"flag"
	"github.com/bifidokk/awesome-chat/auth/internal/converter"
	userRepository "github.com/bifidokk/awesome-chat/auth/internal/repository/user"
	"github.com/bifidokk/awesome-chat/auth/internal/service"
	userService "github.com/bifidokk/awesome-chat/auth/internal/service/user"
	"log"
	"net"

	"github.com/bifidokk/awesome-chat/auth/internal/config"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

var configPath string

type server struct {
	desc.UnimplementedAuthV1Server
	userService service.UserService
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

	userRepo := userRepository.NewRepository(pool)
	userServ := userService.NewUserService(userRepo)
	desc.RegisterAuthV1Server(s, &server{userService: userServ})

	log.Printf("Server listening at %v", listener.Addr())

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create a new user: %v", req)

	userId, err := s.userService.Create(ctx, converter.ToCreateUserFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: userId,
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get user: %v", req)

	user, err := s.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return converter.ToGetUserResponseFromUser(user), nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user: %v", req)

	err := s.userService.Update(ctx, converter.ToUpdateUserFromUpdateRequest(req))

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user: %v", req)

	err := s.userService.Delete(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
