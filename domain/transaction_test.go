package domain

import (
	"testing"
)

func TestTransaction_Validate(t *testing.T) {
	type fields struct {
		Id           string
		Value        float64
		PaidAt       string
		IpLocation   string
		CardHoldName string
		Customer     Customer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		//test cases.
		{
			name: "Returns correct Boolean value when validation is unsuccessful.",
			fields: fields{
				Id:           "1",
				Value:        900.00,
				PaidAt:       "2020-01-10 09:00:00",
				IpLocation:   "RS/BR",
				CardHoldName: "Michael Scott",
				Customer: Customer{
					Id:        "1",
					Name:      "Michael Scott",
					BirthDate: "1998-06-20",
					State:     "RS/BR",
					Phone:     "48 98466-8473",
				},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				Id:           "Returns correct Boolean value when validation is successful.",
				Value:        0,
				PaidAt:       "",
				IpLocation:   "",
				CardHoldName: "",
				Customer:     Customer{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction := Transaction{
				Id:           tt.fields.Id,
				Value:        tt.fields.Value,
				PaidAt:       tt.fields.PaidAt,
				IpLocation:   tt.fields.IpLocation,
				CardHoldName: tt.fields.CardHoldName,
				Customer:     tt.fields.Customer,
			}
			if got := transaction.Validate(); got != tt.want {
				t.Errorf("Transaction.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
