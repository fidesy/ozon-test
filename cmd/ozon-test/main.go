package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/fidesy/ozon-test/internal/delivery/grpc"
	"github.com/fidesy/ozon-test/internal/delivery/handler"
	"github.com/fidesy/ozon-test/internal/infrastructure/persistence"
	"github.com/fidesy/ozon-test/internal/usecase"
	"github.com/fidesy/ozon-test/pkg/utils"
)

func main() {
	conf, err := utils.LoadConfig("./configs/config.yml")
	checkError(err)

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

	usecases := usecase.NewUsecase(conf, repos)
	handlers := handler.New(usecases)
	routes := handlers.InitRoutes()

	// run http server
	go func() {
		err = http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), routes)
		checkError(err)
	}()

	grpcServer := grpc.NewServer(usecases)
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
