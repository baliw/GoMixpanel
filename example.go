package main

import (
	"fmt"
	"mixpanel"
)

func main() {
	m := mixpanel.NewMixpanel()
	m.SetApiToken("yourApiToken")
	
	//granular api
	e := m.NewEvent()
	e.SetName("Granular API Event")
	e.SetProperties(map[string] interface{} {
		"lib": "GoMixpanel",
		"author": "Rich Collins"})
	e.SetProperty("version", "20121127")
	
	if success, err := e.Send(); success {
		fmt.Println("Granular Success")
	} else if err == nil {
		fmt.Println("Granular Failure")
	} else {
		fmt.Println(err)
	}
	
	//terse api
	
	success, err := m.SendEvent("Terse API Event", map[string] interface{} {
		"lib": "GoMixpanel",
		"author": "Rich Collins",
		"version": "20121127"})
		
	if success {
		fmt.Println("Terse Success")
	} else if err == nil {
		fmt.Println("Terse Failure")
	} else {
		fmt.Println(err)
	}
}