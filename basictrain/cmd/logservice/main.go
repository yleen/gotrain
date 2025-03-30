package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/allen/Go/log"
	"github.com/allen/Go/registry"
	"github.com/allen/Go/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"

	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddr,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddr + "/services",
		HeartbeatURL:     serviceAddr + "/heartbeat",
	}

	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers)

	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
