package transfer

import (
	"github.com/i-hit/go-lesson2.2.git/pkg/card"
	"testing"
)



func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc               *card.Service
		CommissionToTinkoff   int64
		CommissionFromTinkoff int64
		MinimumFromTinkoff    int64
		CommissionOther       int64
		MinimumOther          int64
	}
	cardSvc := card.NewService("Tinkoff")

	cardSvc.GetNewCard("visa", 1000, "RUB", "0001")
	cardSvc.GetNewCard("visa", 100, "RUB", "0002")

	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "cardFrom yes, cardTo yes, balance ok",
			args: args{
				from:   "0001",
				to:     "0002",
				amount: 100,
			},
			want: nil,
		},
		{
			name: "cardFrom yes, cardTo yes, balance not ok",
			args: args{
				from:   "0001",
				to:     "0002",
				amount: 10000,
			},
			want: ErrFromCardBalance,
		},
		{
			name: "cardFrom yes, cardTo no, balance ok",
			args: args{
				from:   "0001",
				to:     "0003",
				amount: 100,
			},
			want: card.ErrCard,
		},
		{
			name: "cardFrom yes, cardTo no, balance not ok",
			args: args{
				from:   "0001",
				to:     "0003",
				amount: 10000,
			},
			want: ErrCard,
		},
		{
			name: "cardFrom no, cardTo yes, balance ok",
			args: args{
				from:   "0003",
				to:     "0001",
				amount: 100,
			},
			want: ErrCard,
		},
		{
			name: "cardFrom no, cardTo yes, balance not ok",
			args: args{
				from:   "0003",
				to:     "0001",
				amount: 10000,
			},
			want: ErrCardFrom,
		},
	}
	for _, tt := range tests {
			s := &Service{
				CardSvc: cardSvc,
				CommissionToTinkoff: 0,
				CommissionFromTinkoff: 5,
				MinimumFromTinkoff: 10,
				CommissionOther: 15,
				MinimumOther: 30,
			}
			got := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if got != tt.want {
				t.Errorf("Card2Card() got = %v, want %v", got, tt.want)
			}
	}
}

func TestService_Card2Card1(t *testing.T) {
	type fields struct {
		CardSvc               *card.Service
		CommissionToTinkoff   int64
		CommissionFromTinkoff int64
		MinimumFromTinkoff    int64
		CommissionOther       int64
		MinimumOther          int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:               tt.fields.CardSvc,
				CommissionToTinkoff:   tt.fields.CommissionToTinkoff,
				CommissionFromTinkoff: tt.fields.CommissionFromTinkoff,
				MinimumFromTinkoff:    tt.fields.MinimumFromTinkoff,
				CommissionOther:       tt.fields.CommissionOther,
				MinimumOther:          tt.fields.MinimumOther,
			}
			if err := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Card2Card() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}