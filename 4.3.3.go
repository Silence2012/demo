package main

import "fmt"

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

func main() {
	fmt.Println(LevelEmergency)
	fmt.Println(LevelAlert)
	fmt.Println(LevelCritical)
	fmt.Println(LevelError)
	fmt.Println(LevelWarning)
	fmt.Println(LevelNotice)
	fmt.Println(LevelInformational)
	fmt.Println(LevelDebug)


}
