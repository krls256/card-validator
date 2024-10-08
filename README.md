# Card Validator

ISO/IEC 7812 card validation algorithm. Supported to use as library but HTTP and gRPC services and clients also included.

```shell
    go get github.com/krls256/card-validator
```

## As Library

```go
package main

import (
	"fmt"
	"github.com/krls256/card-validator/card"
)

func main() {
	c := card.Card{Number: "4111111111111111", Month: "01", Year: "2028"}

	fmt.Println(c.IsValid())
} 
```

## HTTP client

```go
package main

import (
	"context"
	"fmt"
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

	httpClient := apiHTTP.NewClient(cfg.HTTPConfig)

	c := card.NewCard("4111111111111111", "01", "2028")
	
	fmt.Printf("http call: err=%v\n", httpClient.Validate(context.Background(), c))
}

```

## gRPC client

```go
package main

import (
	"context"
	"fmt"
	apiGRPC "github.com/krls256/card-validator/api/grpc"
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

	c := card.NewCard("4111111111111111", "01", "2028")

	fmt.Printf("grpc call: err=%v\n", grpcClient.Validate(context.Background(), c))
}

```