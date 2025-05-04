package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zeshi09/zeshifyd/internal/bus"
	"github.com/zeshi09/zeshifyd/internal/model"
	"github.com/zeshi09/zeshifyd/internal/storage"
)

func main() {
	handler := &bus.NotificationHandler{
		OnNotify: func(n model.Notification) {
			if err := storage.SaveNotification(n); err != nil {
				log.Printf("failed to save notification: %v", err)
			}
			fmt.Fprintf(os.Stdout, "%s: %s\n", n.Summary, n.Body)
		},
	}

	if err := handler.Start(); err != nil {
		log.Fatalf("start daemon error: %v", err)
	}
}
