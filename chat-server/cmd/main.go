package main

import (
	"context"
	"log"

	application "github.com/bifidokk/awesome-chat/chat-server/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := application.NewApp(ctx)

	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = app.Run()

	if err != nil {
		log.Fatalf("failed to run: %s", err.Error())
	}
}
