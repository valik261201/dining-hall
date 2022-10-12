package main

import (
	"math/rand"
	"time"
)

type Order struct {
	Id         int   `json:"id"`
	Items      []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max-wait"`
	PickUpTime int   `json:"pick-up-time"`
}

// generate random number for order
func genRandomNum(min, max int) int {
	return min + rand.Intn(max-min)
}

// generate random order
func generateItems() []int {
	n := genRandomNum(1, 10)
	var items = make([]int, n)

	for i := 0; i < n; i++ {
		items[i] = genRandomNum(1, 10)
	}

	return items
}

// get max waitting time
func generateMaxWaitTime(menu Foods) int {
	maxTime := 0

	for i := 0; i < len(menu.Foods); i++ {
		if menu.Foods[i].Preparation_time > maxTime {
			maxTime = menu.Foods[i].Preparation_time
		}
	}
	maxWait := float32(maxTime) * 1.3
	return int(maxWait)
}

func getUnixTimestamp() int {
	now := time.Now()
	sec := now.Unix()

	return int(sec)
}

func generateOrder() Order {
	var menu Foods

	return Order{
		Id:         int(genRandomNum(1, 1000)),
		Items:      generateItems(),
		Priority:   int(genRandomNum(1, 5)),
		MaxWait:    generateMaxWaitTime(menu),
		PickUpTime: getUnixTimestamp(),
	}
}
