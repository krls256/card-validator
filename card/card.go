package card

func NewCard(number string, month string, year string) Card {
	return Card{
		Number: number,
		Month:  month,
		Year:   year,
	}
}

type Card struct {
	Number string `json:"number"`
	Month  string `json:"month"`
	Year   string `json:"year"`
}

func (c Card) IsValid() error {
	for _, rule := range Rules {
		if err := rule(c); err != nil {
			return err
		}
	}

	return nil
}
