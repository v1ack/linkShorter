package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/v1ack/linkShorter/internal/app"
	"github.com/v1ack/linkShorter/internal/store"
	pkg "github.com/v1ack/linkShorter/pkg"
	migration "github.com/v1ack/linkShorter/sql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	config := ParseEnv()

	go ListenForClose(log)

	addr := fmt.Sprintf("0.0.0.0:%v", config.grpcPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	DeferClose(l.Close)

	s := grpc.NewServer()

	var provider store.DataStore

	if config.dbConnection != "" {
		// Connect to db
		connection, err := sql.Open("postgres", config.dbConnection)
		if err != nil {
			log.Fatalln("Failed to init db:", err)
		}

		migration.RunMigrations(connection)

		// Set provider
		provider = store.CreateDbProvider(connection)
		DeferClose(connection.Close)
	} else {
		log.Info("No connection string provided, using inmemory")
		provider = store.CreateInMemoryProvider()
	}

	service, err := app.NewService(provider)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	pkg.RegisterShorterServer(s, service)

	go func() {
		// Serve gRPC Server
		log.Info("Serving gRPC on https://", addr)
		DeferClose(func() error {
			log.Info("closing grpc")
			s.GracefulStop()
			return nil
		})
		log.Fatal(s.Serve(l))
	}()

	err = Run("dns:///"+addr, config)
	log.Fatalln(err)
}
