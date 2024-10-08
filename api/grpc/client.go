package grpc

import (
	"context"
	"github.com/krls256/card-validator/card"
	"github.com/krls256/card-validator/errors"
	transportGRPC "github.com/krls256/card-validator/pkg/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ErrInternalError = errors.NewErrorWithCode("internal error", 10)

func NewClient(cfg transportGRPC.Config) (*Client, error) {
	conn, err := grpc.NewClient(cfg.DNS(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcClient: NewCardValidatorServiceClient(conn),
	}, nil
}

type Client struct {
	grpcClient CardValidatorServiceClient
}

func (cl *Client) Validate(ctx context.Context, c card.Card) error {
	r, err := cl.grpcClient.Validate(ctx, &Card{
		Number: c.Number,
		Month:  c.Month,
		Year:   c.Year,
	})

	if err != nil {
		return err
	}

	if r.Valid {
		return nil
	}

	if r.Error != nil {
		return errors.NewErrorWithCode(r.Error.Message, int(r.Error.Code))
	}

	return ErrInternalError
}
