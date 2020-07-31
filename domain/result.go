package domain

import (
	"github.com/bearbin/go-age"
	"log"
	"time"
)

const (
	layoutISO   = "2006-01-02 15:04:05"
	layoutBIRTH = "2006-01-02"
)

type TransactionResult struct {
	Id    string `json:"id"`
	Score int    `json:"score,string"`
}

type SuspiciousBehavior struct {
	Description string `json:"description"`
	Value       int    `json:"value,string"`
}

var stateMap map[string]string

func CheckFraud(transactions []Transaction) []TransactionResult {
	stateMap = GetStateMap()
	var result []TransactionResult

	for _, transact := range transactions {
		if transact.Validate() {
			point := PointCounter(DetectSB(transact))
			result = append(result, TransactionResult{
				Id:    transact.Id,
				Score: point,
			})

		} else {
			log.Printf("The Transaction id: %v has issues...", transact)
			point := PointCounter(DetectSB(transact))
			result = append(result, TransactionResult{
				Id:    transact.Id,
				Score: point,
			})
		}
	}

	return result
}

func DetectSB(transaction Transaction) []SuspiciousBehavior {
	var SBFound []SuspiciousBehavior
	ddd := GetStateByDDD(transaction.Customer.Phone)
	currentState := stateMap[ddd] + "/BR"
	if currentState != transaction.IpLocation && transaction.IpLocation != transaction.Customer.State && currentState != transaction.Customer.State {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "All location information is conflicting.",
			Value:       25,
		})
	} else if transaction.IpLocation != transaction.Customer.State {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The transaction location doesn't match the customer's location.",
			Value:       12,
		})
	} else if transaction.IpLocation != currentState {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The transaction location doesn't match the phone's DDD location.",
			Value:       12,
		})
	} else if currentState != transaction.Customer.State {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The current DDD location doesn't match the customer's location.",
			Value:       15,
		})
	}
	if transaction.CardHoldName != transaction.Customer.Name {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The cardholder's name doesn't match the customer's name.",
			Value:       10,
		})
	}
	if transaction.Value < 0 {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The transaction value is less than 0.",
			Value:       25,
		})
	}
	paidAt, _ := time.Parse(layoutISO, transaction.PaidAt)
	now := time.Now()
	nowString := now.Format(layoutISO)
	now, _ = time.Parse(layoutISO, nowString)
	if paidAt.After(now) {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The payment date is in the future.",
			Value:       16,
		})
	}
	birth, _ := time.Parse(layoutBIRTH, transaction.Customer.BirthDate)
	ages := age.Age(birth)
	if ages < 18 {
		SBFound = append(SBFound, SuspiciousBehavior{
			Description: "The Customer is a minor.",
			Value:       9,
		})
	}

	return SBFound
}

func PointCounter(behaviorList []SuspiciousBehavior) int {
	count := 0
	behaviorPoints := 0

	for _, behavior := range behaviorList {
		count = count + 1
		behaviorPoints = behaviorPoints + behavior.Value
	}

	result := behaviorPoints * count
	if result > 100 {
		result = 100
	}
	return result
}
