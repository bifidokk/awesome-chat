package app

import (
	"context"
	"log"
	"net"

	"github.com/bifidokk/awesome-chat/chat-server/internal/closer"
	"github.com/bifidokk/awesome-chat/chat-server/internal/config"
	desc "github.com/bifidokk/awesome-chat/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App is the main application struct
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates a new instance of App.
func NewApp(ctx context.Context) (*App, error) {
	application := &App{}
	err := application.initDependencies(ctx)

	if err != nil {
		return nil, err
	}

	return application, nil
}

// Run starts the gRPC server for the application.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDependencies(ctx context.Context) error {
	a.serviceProvider = &serviceProvider{}

	inits := []func(context context.Context) error{
		a.iniConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, initFunction := range inits {
		if err := initFunction(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) iniConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatAPI(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GrpcConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)

	if err != nil {
		return err
	}

	return nil
}
