package main

import (
	"fmt"

	"http3-server-poc/cmd/api/bootstrap"
	"http3-server-poc/cmd/api/config"

	"go.uber.org/zap"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	httpApi := bootstrap.Api(logger)

	logger.Info(fmt.Sprintf("Starting HTTP3 server on %v %v", httpApi.Addr, &httpApi.Port))

	err = httpApi.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
