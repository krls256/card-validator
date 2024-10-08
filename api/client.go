package api

import (
	"context"
	"github.com/krls256/card-validator/card"
)

type Client interface {
	Validate(ctx context.Context, c card.Card) error
}
