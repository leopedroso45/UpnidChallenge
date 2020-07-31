package domain

import (
	"reflect"
	"testing"
)

func TestCheckFraud(t *testing.T) {
	transactions := []Transaction{
		{
			Id:           "1",
			Value:        800.00,
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
		{
			Id:           "2",
			Value:        900.00,
			PaidAt:       "2019-02-15 12:45:09",
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
	}
	transactions2 := append(transactions, Transaction{
		Id:           "3",
		Value:        -100.00,
		PaidAt:       "2022-01-10 09:00:00",
		IpLocation:   "RS/BR",
		CardHoldName: "Michael Scott",
		Customer: Customer{
			Id:        "4",
			Name:      "Michael",
			BirthDate: "2015-06-20",
			State:     "SC/BR",
			Phone:     "48 98466-8473",
		}})
	type args struct {
		transactions []Transaction
	}
	tests := []struct {
		name string
		args args
		want []TransactionResult
	}{
		//test cases.
		{
			name: "Return the correct result of checking transactions.",
			args: args{transactions: transactions},
			want: []TransactionResult{
				{Id: "1", Score: 12},
				{Id: "2", Score: 12},
			},
		},
		{
			name: "Return the correct result of checking transactions. //same case",
			args: args{transactions: transactions2},
			want: []TransactionResult{
				{Id: "1", Score: 12},
				{Id: "2", Score: 12},
				{Id: "3", Score: 100},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFraud(tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckFraud() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detectSB(t *testing.T) {
	transaction := Transaction{
		Id:           "1",
		Value:        800.00,
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
	}
	transaction2 := Transaction{
		Id:           "1",
		Value:        -100.00,
		PaidAt:       "2022-01-10 09:00:00",
		IpLocation:   "RS/BR",
		CardHoldName: "Michael Scott",
		Customer: Customer{
			Id:        "1",
			Name:      "Michael",
			BirthDate: "2015-06-20",
			State:     "SC/BR",
			Phone:     "48 98466-8473",
		},
	}
	SBFound := []SuspiciousBehavior{
		{Description: "The transaction location doesn't match the phone's DDD location.", Value: 12},
	}
	SBFound2 := []SuspiciousBehavior{
		{Description: "The transaction location doesn't match the customer's location.", Value: 12},
		{Description: "The cardholder's name doesn't match the customer's name.", Value: 10},
		{Description: "The transaction value is less than 0.", Value: 25},
		{Description: "The payment date is in the future.", Value: 16},
		{Description: "The Customer is a minor.", Value: 9},
	}

	type args struct {
		transaction Transaction
	}
	tests := []struct {
		name string
		args args
		want []SuspiciousBehavior
	}{
		//test cases.
		{
			name: "Return the correct object when there's no suspicious behavior.",
			args: args{
				transaction: transaction,
			},
			want: SBFound,
		},
		{
			name: "Return the correct object when there are several suspicious behavior.",
			args: args{
				transaction: transaction2,
			},
			want: SBFound2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectSB(tt.args.transaction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectSB() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"size", func(t *testing.T) {
			if got := DetectSB(tt.args.transaction); !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("detectSB() = %v, want %v", len(got), len(tt.want))
			}
		})
	}
}

func Test_pointCounter(t *testing.T) {
	supBehave := []SuspiciousBehavior{
		{
			Description: "Test value 1",
			Value:       6,
		},
		{
			Description: "Test value 2",
			Value:       6,
		},
	}
	type args struct {
		behaviorList []SuspiciousBehavior
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Test cases.
		{
			name: "Return correct score.",
			args: args{behaviorList: supBehave},
			want: 24,
		},
		{
			name: "Return correct score using negative numbers. out of scope.",
			args: args{behaviorList: []SuspiciousBehavior{
				{
					Description: "Test value 1",
					Value:       -10,
				},
				{
					Description: "Test value 2",
					Value:       6,
				},
			}},
			want: -8,
		},
		{
			name: "Returns 0 if the slice is empty.",
			args: args{behaviorList: []SuspiciousBehavior{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointCounter(tt.args.behaviorList); got != tt.want {
				t.Errorf("pointCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
