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

		cardSvc.GetNewCard("visa", 1000, "RUB", "5106 2100 0000 0007")
		cardSvc.GetNewCard("visa", 100, "RUB", "5106 2100 0000 0000 6")



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
		{
			name: "cardFrom yes, cardTo yes, balance ok",
			args: args{
				from:   "5106 2100 0000 0007",
				to:     "5106 2100 0000 0000 6",
				amount: 100,
			},
			wantErr: false,
		},
		{
			name: "cardFrom yes, cardTo yes, balance not ok",
			args: args{
				from:   "5106 2100 0000 0007",
				to:     "5106 2100 0000 0000 6",
				amount: 100000,
			},
			wantErr: true,
		},
		{
			name: "cardFrom yes, cardTo not found, balance ok",
			args: args{
				from:   "5106 2100 0000 0007",
				to:     "51106 2100 0000 0000 6",
				amount: 100,
			},
			wantErr: true,
		},
		{
			name: "cardFrom not found, cardTo yes, balance ok",
			args: args{
				from:   "511206 2100 0000 0007",
				to:     "5106 2100 0000 0000 6",
				amount: 100,
			},
			wantErr: true,
		},
		{
			name: "cardFrom not found, cardTo not found, balance ok",
			args: args{
				from:   "51106 2100 0000 0007",
				to:     "51106 2100 0000 0000 6",
				amount: 100,
			},
			wantErr: true,
		},
		{
			name: "cardFrom yes, cardTo not valid, balance ok",
			args: args{
				from:   "5106 2100 0000 0007",
				to:     "5106 2100 0000 01000 6",
				amount: 100,
			},
			wantErr: true,
		},
		{
			name: "cardFrom not valid, cardTo yes, balance ok",
			args: args{
				from:   "5106 2100 0000 0407",
				to:     "5106 2100 0000 0000 6",
				amount: 100,
			},
			wantErr: true,
		},
		{
			name: "cardFrom not valid, cardTo not valid, balance ok",
			args: args{
				from:   "5106 2100 0000 0207",
				to:     "5106 2100 0000 00030 6",
				amount: 100,
			},
			wantErr: true,
		},
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