package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"gitlab.intelligentb.com/devops/code-gen/util"
)



func main() { 
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s interface-file url template-file [start error code]", os.Args[0])
		os.Exit(0)
	}
	templateFile := os.Args[3]
	s := "100000"
	if len(os.Args) > 4 {
		s = os.Args[4]
	}
	processAndPrint(util.ParseService(os.Args[1],os.Args[2],s),templateFile)
}

func processAndPrint(sdet util.Servicedetail,templateFile string) {
	tpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Printf("uh oh problem with template.err = %s\n", err.Error())
		return
	}

	tpl.Execute(os.Stdout, sdet)
}


