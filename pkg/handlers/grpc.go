package handlers

import (
	"context"
	stdErrors "errors"
	"github.com/krls256/card-validator/api/grpc"
	"github.com/krls256/card-validator/card"
	"github.com/krls256/card-validator/errors"
)

func NewGRPCCardValidatorHandler() CardGRPCValidatorHandler {
	return CardGRPCValidatorHandler{}
}

type CardGRPCValidatorHandler struct{}

func (h CardGRPCValidatorHandler) Validate(ctx context.Context, c *grpc.Card) (*grpc.ValidateResult, error) {
	err := card.NewCard(c.Number, c.Month, c.Year).IsValid()

	if err != nil {
		var ewc errors.ErrorWithCode
		ok := stdErrors.As(err, &ewc)

		code := 0
		if ok {
			code = ewc.Code()
		}

		return &grpc.ValidateResult{
			Valid: false,
			Error: &grpc.Error{
				Code:    int32(code),
				Message: err.Error(),
			},
		}, nil
	}

	return &grpc.ValidateResult{
		Valid: true,
	}, nil
}
