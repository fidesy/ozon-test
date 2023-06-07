package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/fidesy/ozon-test/internal/config"
	"github.com/fidesy/ozon-test/internal/delivery/grpc"
	"github.com/fidesy/ozon-test/internal/delivery/handler"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
	"github.com/fidesy/ozon-test/internal/service"
)

func main() {
	conf, err := config.Load("./configs/config.yml")
	checkError(err)

	db := flag.String("db", "", "database to store URLs")
	flag.Parse()
	if *db != "" {
		conf.Database = *db
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	defer cancel()

	repos, err := persistence.NewRepository(ctx, conf)
	checkError(err)
	defer func() {
		err = repos.Close()
		log.Println("error when closing pool connection:", err.Error())
	}()

	service := service.NewService(conf, repos)
	handlers := handler.New(service)
	routes := handlers.InitRoutes()

	// run http server
	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), routes)
		checkError(err)
	}()

	grpcServer := grpc.NewServer(service)
	go func() {
		err = grpcServer.Start(fmt.Sprintf(":%s", conf.GrpcPort))
		checkError(err)
	}()

	<-ctx.Done()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
