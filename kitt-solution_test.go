package main

import (
	"testing"
)

func TestGetPrice1440(t *testing.T) {
	meetingTime := 1440
	expectedPrice := 60
	gotPrice := GetPrice(meetingTime)
	if expectedPrice != gotPrice {
		t.Fatalf("Expected a price of %d but got %d instead.", expectedPrice, gotPrice)
	}
}

func TestGetPrice2880(t *testing.T) {
	meetingTime := 2880
	expectedPrice := 105
	gotPrice := GetPrice(meetingTime)
	if expectedPrice != gotPrice {
		t.Fatalf("Expected a price of %d but got %d instead.", expectedPrice, gotPrice)
	}
}

func TestGetPrice20160(t *testing.T) {
	meetingTime := 20160
	expectedPrice := 210
	gotPrice := GetPrice(meetingTime)
	if expectedPrice != gotPrice {
		t.Fatalf("Expected a price of %d but got %d instead.", expectedPrice, gotPrice)
	}
}
