package main

import "fmt"

func main()  {
	
	// Basic switch statement
	day:=3

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thersday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	default:
		fmt.Println("sunday")
	}

	today:= "monday"

	switch today {
	case "sunday", "saturday":
		fmt.Println("holiday")
	default:
		fmt.Println("working day")
	}

	month:= "January"

	switch month {
	case "December", "January", "February":
		fmt.Println("Winter")
	case "March", "April","May":
		fmt.Println("Spring")
	case "June", "July", "August":
		fmt.Println("Summer")
	case "September", "October", "November":
		fmt.Println("Fall")
	default:
		fmt.Println("Invalid month")
	}
	

	var temperature int = 25

	switch temperature {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("cool")
    case temperature >= 20 && temperature < 30:
		fmt.Println("warm")
	default: 
	    fmt.Println("hot")
	}

	
}