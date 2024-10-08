package card

import (
	"github.com/krls256/card-validator/errors"
	"strconv"
	"time"
)

var (
	ErrWrongMonth              = errors.NewErrorWithCode("month must be in [01, ... ,12]", 1)
	ErrWrongYear               = errors.NewErrorWithCode("year must be 4 digit number", 2)
	ErrPastExpiryDate          = errors.NewErrorWithCode("expiry date must not be in the past", 3)
	ErrWrongNumberLength       = errors.NewErrorWithCode("number length must be between 8 and 19", 4)
	ErrNumberContainNotNumbers = errors.NewErrorWithCode("number must contain only numbers", 5)
	ErrNumberCantPassLuhn      = errors.NewErrorWithCode("number can not pass Luhn test", 6)
)

type Rule func(card Card) error

var Rules = []Rule{
	MonthRule,
	YearRule,
	NotExpiredRule,
	CardNumberLenRule,
	CardNumberLuhnRule,
}

func MonthRule(card Card) error {
	if len(card.Month) != 2 {
		return ErrWrongMonth
	}

	iMonth, err := strconv.Atoi(card.Month)
	if err != nil {
		return ErrWrongMonth
	}

	if iMonth < 1 || iMonth > 12 {
		return ErrWrongMonth
	}

	return nil
}

func YearRule(card Card) error {
	if len(card.Year) != 4 {
		return ErrWrongYear
	}

	iYear, err := strconv.Atoi(card.Year)
	if err != nil {
		return ErrWrongYear
	}

	if iYear < 1000 || iYear > 9999 {
		return ErrWrongYear
	}

	return nil
}

func NotExpiredRule(card Card) error {
	iYear, err := strconv.Atoi(card.Year)
	if err != nil {
		return ErrWrongYear
	}

	iMonth, err := strconv.Atoi(card.Month)
	if err != nil {
		return ErrWrongMonth
	}

	now := time.Now()

	if iYear < now.Year() || (iYear == now.Year() && iMonth < int(now.Month())) {
		return ErrPastExpiryDate
	}

	return nil
}

const MinCardNumberLength = 8
const MaxCardNumberLength = 19

func CardNumberLenRule(card Card) error {
	if len(card.Number) < MinCardNumberLength || len(card.Number) > MaxCardNumberLength {
		return ErrWrongNumberLength
	}

	return nil
}

func CardNumberLuhnRule(card Card) error {
	checksum := 0
	isSecond := false
	numbers := make([]int, 0, len(card.Number))

	for _, symbol := range card.Number {
		n, err := strconv.Atoi(string(symbol))
		if err != nil {
			return ErrNumberContainNotNumbers
		}

		numbers = append(numbers, n)
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		item := numbers[i]

		if isSecond {
			item *= 2

			if item > 9 {
				item -= 9
			}
		}

		isSecond = !isSecond
		checksum += item
	}

	if checksum%10 != 0 {
		return ErrNumberCantPassLuhn
	}

	return nil
}
