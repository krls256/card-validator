package handlers

import (
	stdErrors "errors"
	"github.com/gofiber/fiber/v3"
	apiHTTP "github.com/krls256/card-validator/api/http"
	"github.com/krls256/card-validator/card"
	"github.com/krls256/card-validator/errors"
	"github.com/krls256/card-validator/pkg/transport/http"
)

func NewCardHTTPValidatorHandler() CardHTTPValidatorHandler {
	return CardHTTPValidatorHandler{}
}

type CardHTTPValidatorHandler struct{}

func (h CardHTTPValidatorHandler) Register(router fiber.Router) {
	router.Post("validate", h.validate)
}

func (h CardHTTPValidatorHandler) validate(ctx fiber.Ctx) error {
	c := card.Card{}

	if err := ctx.Bind().JSON(&c); err != nil {
		return http.BadRequest(ctx, err)
	}

	err := c.IsValid()

	resp := apiHTTP.ValidationResponse{}

	if err != nil {
		var ewc errors.ErrorWithCode
		ok := stdErrors.As(err, &ewc)

		code := 0
		if ok {
			code = ewc.Code()
		}

		resp.Error = &apiHTTP.ValidationError{
			Code:    code,
			Message: err.Error(),
		}
	} else {
		resp.Valid = true
	}

	return http.OK(ctx, resp)
}
