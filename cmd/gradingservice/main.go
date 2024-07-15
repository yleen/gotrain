package main

import (
	"context"
	"fmt"

	stlog "log"

	"github.com/allen/Go/grades"
	"github.com/allen/Go/log"
	"github.com/allen/Go/registry"
	"github.com/allen/Go/service"
)

func main() {
	host, port := "localhost", "6000"

	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddr,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddr + "/services",
		HeartbeatURL:     serviceAddr + "/heartbeat",
	}

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandlers)

	if err != nil {
		stlog.Fatalln(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service.")
}
