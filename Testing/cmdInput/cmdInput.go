package cmdInput

import (
	"fmt"
	"log"
)

// Kunne ikke henvise til import av "log" pakke som ligger i samme mappe
// ./log henviser til search for pakker i mappen log
// Se mappestruktur for klarifisering eller ask Tor
import "LogCalc/log"

func input()	{


	fmt.Println("Initiating")

	// InputReaders from terminal on number and base
	// Calls the corresponding funcs for arithmics
	finish := false
	for finish == false {

		var logVal int
		fmt.Println("Enter the numeric number you wish to calculate")
		fmt.Scan(&logVal)
		floatVal := float64(logVal)
		fmt.Println("Number to be calculated on is: ", floatVal)

		var inputVal int
		fmt.Println("Enter the base you wish to calculate with: Either 2 or 10")
		fmt.Scan(&inputVal)
		fmt.Println("Input was read as", inputVal, " executing the corresponding func")

		if inputVal == 2 {
			fmt.Println("The answer from log2 from value:", logVal, "is", log.Log2(floatVal))
		}
		if inputVal == 10 {
			fmt.Println("The answer from log10 from value:", logVal, "is", log.Log10(floatVal))
			// test
		}
		var exit string
		fmt.Println("Exit the app? Type yes to exit, no to keep caculating ")
		fmt.Scan(&exit)
		if exit == "yes" {
			finish = true
		}
		if exit == "no" {
			finish = false
		}

	}
}
