package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ladies and Gentlemen, the Weeknd")
	default:
		fmt.Println("Work work")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("dont know what type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}