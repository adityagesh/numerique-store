package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	gracefulShutDown()
	run()
}

func gracefulShutDown() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdown
		log.Printf("Gracefully shutting down server...")
		// Handle graceful shutdown
		os.Exit(0)
	}()
}

func run() {
	lis, err := net.Listen("tcp", ":9000")
	defer lis.Close()

	if err != nil {
		log.Fatalf("Error starting server %v", err)
	}
	log.Printf("Server started in port %v", lis.Addr().String())

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error binding GRPC server %v", err)
	}
}
