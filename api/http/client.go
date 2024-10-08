package http

import (
	"context"
	"github.com/krls256/card-validator/card"
	"github.com/krls256/card-validator/errors"
	transportHTTP "github.com/krls256/card-validator/pkg/transport/http"
)

var ErrInternalError = errors.NewErrorWithCode("internal error", 20)

const ValidatePath = "/validate"

func NewClient(cfg transportHTTP.Config) *Client {
	return &Client{cfg: cfg}
}

type Client struct {
	cfg transportHTTP.Config
}

func (cl *Client) Validate(ctx context.Context, c card.Card) error {
	httpReq, err := cl.NewPostRequest(ctx, ValidatePath, c)
	if err != nil {
		return err
	}

	resp, err := HandleResponse[ValidationResponse](httpReq)
	if err != nil {
		return err
	}

	if resp.Valid {
		return nil
	}

	if resp.Error != nil {
		return errors.NewErrorWithCode(resp.Error.Message, resp.Error.Code)
	}

	return ErrInternalError
}
