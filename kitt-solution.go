package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const pPM = 2            //Price per minute
const pPH = 22           //Price per hour
const pPD = 60           //Price per day
const pPW = 105          //Price per week
const minsInHour = 60    //Minutes per hour
const minsInDay = 1440   //Minutes per day
const minsInWeek = 10080 //Minutes per Week

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Expected a single argument to be supplied for the meeting time eg. 1400")
	}

	meetingTime := os.Args[1] //set meeting time in minutes

	meetingTimeInt, err := strconv.Atoi(meetingTime)
	if err != nil {
		log.Fatalf("Could not convert %s to int. Please supply a valid integer", meetingTime)
	}

	price := GetPrice(meetingTimeInt)

	fmt.Printf("%d minutes = £%d\n", meetingTimeInt, price)
}

func GetPrice(meetingTime int) int {

	//There are 6 conditions we need to check for. If the meeting time is less than a full unit's rate (eg. 23 hours is less than a day rate),
	//the rentee can pay for a full day or week, but not hour or minute. These are the first 3 logical checks. The next 3 checks check that
	//a meeting time is greater than a full unit's rate (eg. 3 hours is greater than 1 hour), the rentee can use the hourly rate for the 3
	//hour duration.

	//If a check is passed, the price will be added to an int slice. This int slice contains all the possible prices, which is then orded.
	//The lowest price in the slice is then returned by the function.

	var prices []int

	if meetingTime <= minsInWeek {
		prices = append(prices, pPW)
	}
	if meetingTime <= minsInDay {
		prices = append(prices, pPD)
	}
	if meetingTime <= minsInHour {
		prices = append(prices, pPH)
	}

	prices = append(prices, pPM*meetingTime) //always add this price, as you can always pay by the minute rate

	if meetingTime > minsInHour {
		prices = append(prices, pPH*meetingTime/minsInHour)
	}
	if meetingTime > minsInDay {
		prices = append(prices, pPD*meetingTime/minsInDay)
	}
	if meetingTime > minsInWeek {
		prices = append(prices, pPW*meetingTime/minsInWeek)
	}

	sort.Ints(prices)

	return prices[0]
}
