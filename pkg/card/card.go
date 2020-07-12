package card

type Card struct {
	Issuer string
	Balance int64
	Currency string
	Number string
}

type Service struct {
	BankName string
	Cards []*Card
}

func NewService(bankname string) *Service  {
	return &Service {
		BankName : bankname,
	}
}

func (s *Service) GetNewCard(issuer string, balance int64, currency string, number string) *Card {
	card := &Card{
		Issuer: issuer,
		Balance: balance,
		Currency: currency,
		Number: number,
	}

	s.Cards = append(s.Cards, card)

	return card
}

func (s *Service) CheckNumber(number string) (*Card, bool) {
	for _, c := range s.Cards {
		if c.Number == number {
			return c, true
		}
	}
	return nil, false
}

func (s *Service) CheckBalance(card *Card, amount int64) (bool) {
	if card.Balance >= amount {
		return true
	}
	return false
}

