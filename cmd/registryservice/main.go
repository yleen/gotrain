package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/allen/Go/registry"
)

func main() {
	registry.SetupRegistryService()
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
