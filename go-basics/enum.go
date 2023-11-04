package main

import "fmt"

type day int

const (
	mon day = iota
	tue
	wed
)

func (d day) string() string {
	switch d {
	case 0:
		return "MONDAY"
	case 1:
		return "TUESDAY"
	case 3:
		return "WEDNSDAY"

	}
	return ""
}

func getDayFromString(ip string) day {
	switch ip {
	case "MONDAY":
		return mon
	case "TUESDAY":
		return tue
	case "WEDNESDAY":
		return wed
	}
	return mon //default
}

func enum_test() {
	d := tue
	fmt.Println(d.string())
}

func realLifeExample() {
	httpInput := "MONDAY" //from an api call
	ipDay := getDayFromString(httpInput)
	_ = ipDay

	//perform business ops using ipDay now

	// storeInDb(ipDay) as int
}
