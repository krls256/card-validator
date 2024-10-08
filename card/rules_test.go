package card

import (
	stdErrors "errors"
	"github.com/krls256/card-validator/errors"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

var (
	validNumber = "4111111111111111"
	validMonth  = "01"
	validYear   = strconv.Itoa(time.Now().Year() + 1)
)

type IsValidTestCase struct {
	Card       Card
	ExpectCode int
}

var isValidTestCases = []IsValidTestCase{
	{Card: NewCard(validNumber, validMonth, validYear), ExpectCode: 0},
	{Card: NewCard(validNumber, "041", validYear), ExpectCode: 1},
	{Card: NewCard(validNumber, "ff", validYear), ExpectCode: 1},
	{Card: NewCard(validNumber, "", validYear), ExpectCode: 1},
	{Card: NewCard(validNumber, validMonth, "0000"), ExpectCode: 2},
	{Card: NewCard(validNumber, validMonth, "ffff"), ExpectCode: 2},
	{Card: NewCard(validNumber, validMonth, "23"), ExpectCode: 2},
	{Card: NewCard(validNumber, validMonth, "rr"), ExpectCode: 2},
	{Card: NewCard(validNumber, validMonth, ""), ExpectCode: 2},
	{Card: NewCard(validNumber, validMonth, "2002"), ExpectCode: 3},
	{Card: NewCard(validNumber, validMonth, "2002"), ExpectCode: 3},
	{Card: NewCard(validNumber, validMonth, "2002"), ExpectCode: 3},
	{Card: NewCard(validNumber, validMonth, "2002"), ExpectCode: 3},
	{Card: NewCard("2415123", validMonth, validYear), ExpectCode: 4},
	{Card: NewCard("24151232415123241512324151232415123", validMonth, validYear), ExpectCode: 4},
	{Card: NewCard("41111111111111f1", validMonth, validYear), ExpectCode: 5},
	{Card: NewCard("4111111111111112", validMonth, validYear), ExpectCode: 6},

	{Card: NewCard("378282246310005", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("371449635398431", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("378734493671000", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("5610591081018250", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("30569309025904", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("38520000023237", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("6011111111111117", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("6011000990139424", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("3530111333300000", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("3566002020360505", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("5555555555554444", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("5105105105105100", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("4111111111111111", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("4012888888881881", validMonth, validYear), ExpectCode: 0},
	{Card: NewCard("4222222222222", validMonth, validYear), ExpectCode: 0},
}

// go test ./card --run TestCard_IsValid
func TestCard_IsValid(t *testing.T) {
	for _, tc := range isValidTestCases {
		err := tc.Card.IsValid()

		if tc.ExpectCode == 0 {
			require.Nil(t, err, tc.Card)

			continue
		}

		require.NotNil(t, err)

		var ewc errors.ErrorWithCode
		ok := stdErrors.As(err, &ewc)

		require.True(t, ok)
		require.Equal(t, ewc.Code(), tc.ExpectCode, tc.Card)
	}
}
