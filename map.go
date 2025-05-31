package main

//Challenge: Map the Days
var Days = map[int]string{
	0: "Sunday",
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
} // do initialization here

func findDay(n int) string {
	return Days[n]
}
