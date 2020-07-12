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
		want   int64
		want1  bool
	}{
		{
			name: "yes-yes-ok",
			args: args{
				from:   "0001",
				to:     "0002",
				amount: 100,
			},
			want: 100,
			want1: true,
		},
		{
			name: "yes-yes-not",
			args: args{
				from:   "0001",
				to:     "0002",
				amount: 10000,
			},
			want: 10000,
			want1: false,
		},
		{
			name: "yes-no-ok",
			args: args{
				from:   "0001",
				to:     "0003",
				amount: 100,
			},
			want: 110,
			want1: true,
		},
		{
			name: "yes-no-not",
			args: args{
				from:   "0001",
				to:     "0003",
				amount: 10000,
			},
			want: 10050,
			want1: false,
		},
		{
			name: "no-yes",
			args: args{
				from:   "0003",
				to:     "0002",
				amount: 100,
			},
			want: 100,
			want1: true,
		},
		{
			name: "no-no",
			args: args{
				from:   "0003",
				to:     "0004",
				amount: 1000,
			},
			want: 1030,
			want1: true,
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
			got, got1 := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if got != tt.want {
				t.Errorf("Card2Card() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Card2Card() got1 = %v, want %v", got1, tt.want1)
			}
	}
}
