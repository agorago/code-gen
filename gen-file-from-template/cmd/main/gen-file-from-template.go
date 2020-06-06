package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	util "github.com/agorago/wego-gen/util"
)

func main() { 
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s interface-file apiURL serviceURL wegoURL template-file [start error code]", os.Args[0])
		os.Exit(0)
	}
	templateFile := os.Args[5]
	s := "100000"
	if len(os.Args) > 4 {
		s = os.Args[6]
	}
	processAndPrint(util.ParseService(os.Args[1],os.Args[2],os.Args[3],os.Args[4],s),templateFile)
}

// processAndPrint - Accept a template file and make sure that it ends with .gohtml
// string .gohtml from the end of the file and create a new file by that name
// Write to the created file  with  template
func processAndPrint(sdet util.Servicedetail,templateFile string) {
	if !strings.HasSuffix(templateFile,".gohtml") {
		fmt.Printf("Template file %s does not end with .gohtml\n",templateFile)
		return
	}

	targetFile := strings.TrimSuffix(templateFile,".gohtml")
	f, err := os.Create(targetFile)
	if err != nil {
		fmt.Printf("Could not open file %s for writing.Err = %s\n",targetFile,err.Error())
		return
	}

	defer f.Close()

	tpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Printf("uh oh problem with template.err = %s\n", err.Error())
		return
	}

	tpl.Execute(f, sdet)
}


