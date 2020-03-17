package main

import (
	"fmt"
	"gitlab.intelligentb.com/devops/bplus/stm"
	"log"
	"os"
)

func main(){
	if len(os.Args) < 3  || (os.Args[2] != "events" && os.Args[2] != "autostates"){
		log.Printf("Usage: %s json-file events|autostates\n", os.Args[0])
		log.Println("events - emits the name of all events in the json file")
		log.Println("autostates - emits the name of all autostates in the json file")
		os.Exit(0)
	}
	parse()
}

func parse(){
	stm,err := stm.MakeStm(os.Args[1],nil)
	if err != nil {
		log.Println("Cannot construct an STM out of this. Check if file is valid!")
		os.Exit(1)
	}
	switch(os.Args[2]){
	case "events":
		printEvents(stm)
	case "autostates":
		printAutoStates(stm)
	}
}

func printAutoStates(stm1 *stm.Stm){
	for stateID,state := range stm1.States{
		if state.Automatic{
			fmt.Println(stateID)
		}
	}
}

func printEvents(stm1 *stm.Stm){
	for _,state := range stm1.States{
		for eventID,_ := range state.Events{
			fmt.Println(eventID)
		}
	}
}