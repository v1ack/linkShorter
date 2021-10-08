package main

import (
	"github.com/v1ack/linkShorter/internal/app"
	"github.com/v1ack/linkShorter/internal/providers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"os"

	pkg "github.com/v1ack/linkShorter/pkg"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()

	var provider providers.DataProvider

	//dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	dbConnectionString := "dbname=shorter password=admin user=admin sslmode=disable"

	if dbConnectionString != "" {
		provider, err = providers.CreatePgProvider(dbConnectionString)
		if err != nil {
			log.Fatalln("Failed to init db:", err)
		}
	} else {
		log.Info("No connection string provided, using inmemory")
		provider = providers.CreateInMemoryProvider()
	}

	defer provider.Close()

	service, err := app.NewService(provider)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	pkg.RegisterShorterServer(s, service)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = Run("dns:///" + addr)
	log.Fatalln(err)
}
