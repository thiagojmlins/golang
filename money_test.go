package main

import (
	"testing"
)

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
func TestMultiplicationInDollars(t *testing.T){
	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := fiveDollars.Times(2)
	assertEqual(t, tenDollars, fiveDollars.Times(2))
}

func TestMultiplicationInEuros(t *testing.T){
	tenEuros := Money{amount: 10, currency: "EUR"}
	twentyEuros := tenEuros.Times(2)
	assertEqual(t, twentyEuros, tenEuros.Times(2))
}

func TestDivision(t *testing.T){
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}
