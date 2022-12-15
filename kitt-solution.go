package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const pPM = 2         //Price per minute
const pPH = 22        //Price per hour
const pPD = 60        //Price per day
const pPW = 105       //Price per week
const minsInHour = 60 //per hour
const minsInDay = 1440
const minsInWeek = 10080

func main() {
	meetingTime := os.Args[1] //set meeting time in minutes
	meetingTimeInt, err := strconv.Atoi(meetingTime)
	if err != nil {
		log.Fatalf("Could not convert %s to int. Please supply a valid integer", meetingTime)
	}
	price := GetPrice(meetingTimeInt)
	fmt.Printf("%d minutes = £%d\n", meetingTimeInt, price)
}

func GetPrice(meetingTime int) int {
	var prices []int

	//If the Meeting Time is less than a particular rate, that rate can be paid.
	//eg: If
	if meetingTime <= minsInWeek {
		prices = append(prices, pPW)
	}
	if meetingTime <= minsInDay {
		prices = append(prices, pPD)
	}
	if meetingTime <= minsInHour {
		prices = append(prices, pPH)
	}

	prices = append(prices, pPM*meetingTime)

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
