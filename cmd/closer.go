package main

import (
	"google.golang.org/grpc/grpclog"
	"os"
	"os/signal"
	"syscall"
)

var arr []func() error

func DeferClose(f func() error) {
	arr = append([]func() error{f}, arr...)
}

func CloseAll(log grpclog.LoggerV2) {
	for _, fun := range arr {
		err := fun()
		if err != nil {
			log.Fatalf("close error %v", err.Error())
		}
	}
}

func ListenForClose(log grpclog.LoggerV2) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		log.Infoln("stopping")
		CloseAll(log)

		os.Exit(0)
	}()
}
