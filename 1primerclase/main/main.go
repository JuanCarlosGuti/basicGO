package main

import "fmt"

func main() {
	monthMap := make(map[int]string)
	monthMap[1] = "jenuary"
	monthMap[2] = "february"
	monthMap[3] = "march"
	monthMap[4] = "april"
	monthMap[5] = "may"
	monthMap[6] = "june"
	monthMap[7] = "july"
	monthMap[8] = "august"
	monthMap[9] = "september"
	monthMap[10] = "october"
	monthMap[11] = "november"
	monthMap[12] = "december"

	mont := 15
	value, exists := monthMap[mont]
	if exists {
		fmt.Printf("The month is %s\n", value)
	} else {
		fmt.Printf("There is no month\n")
	}
}
