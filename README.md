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

```
TBD
```

## gRPC client

```
TBD
```