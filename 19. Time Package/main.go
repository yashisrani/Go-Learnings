package main

import (
	"fmt"
	"time"
)

func main()  {
	currentTime:= time.Now()
	fmt.Println(currentTime)							// for current time

	formatted:= currentTime.Format("02-01-2006 , Monday, 15:04:05, 3:04 PM" )		// for today's date, day, time in 24-hour format,
	fmt.Println(formatted)

	// add 1 day in current time
	new_date:= currentTime.Add(24 * time.Hour) // 24 * hour
	fmt.Println(new_date)
	formattednewdate:= new_date.Format("2006-01-02 Monday")
	fmt.Println(formattednewdate)
}