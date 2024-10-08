package main

import (
	"context"
	"fmt"
	apiGRPC "github.com/krls256/card-validator/api/grpc"
	apiHTTP "github.com/krls256/card-validator/api/http"
	"github.com/krls256/card-validator/card"
	"github.com/krls256/card-validator/pkg/config"
	"log"
)

func main() {
	cfg, err := config.New("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	grpcClient, err := apiGRPC.NewClient(cfg.GRPCConfig)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := apiHTTP.NewClient(cfg.HTTPConfig)

	c := card.NewCard("4111111111111111", "01", "2028")

	fmt.Printf("grpc call: err=%v\n", grpcClient.Validate(context.Background(), c))
	fmt.Printf("http call: err=%v\n", httpClient.Validate(context.Background(), c))
}
