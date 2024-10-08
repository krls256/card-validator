package main

import (
	"fmt"
	"github.com/krls256/card-validator/card"
)

func main() {
	c := card.Card{Number: "4111111111111111", Month: "01", Year: "2028"}

	fmt.Println(c.IsValid())
}
