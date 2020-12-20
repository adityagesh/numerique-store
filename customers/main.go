package main

import (
	"database/sql"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/adityagesh/numerique-store/customers/internal/handler"
	"github.com/adityagesh/numerique-store/customers/internal/rpc/register"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var db *sql.DB

func main() {
	gracefulShutDown()
	run()
}

func run() {
	lis, err := net.Listen("tcp", ":9000")
	defer lis.Close()

	db = createDatabaseConnection()
	if err != nil {
		log.Fatalf("Error starting server %v", err)
	}
	log.Printf("Server started in port %v", lis.Addr().String())

	grpcServer := grpc.NewServer()
	registerGRPCServers(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error binding GRPC server %v", err)
	}
}

func gracefulShutDown() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdown
		log.Info("Gracefully shutting down server...")
		db.Close()
		// Handle graceful shutdown
		// close http server and db connection
		os.Exit(0)
	}()
}

func registerGRPCServers(grpcServer *grpc.Server) {
	register.RegisterRegisterServiceServer(grpcServer, &handler.RegisterServer{
		DB: db,
	})

}

func createDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/customers")
	if err != nil {
		log.Fatalf("Error connect to db")
	}
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)
	return db
}
