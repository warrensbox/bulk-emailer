package main

/*
* Version 0.0.1
* Compatible on command line - tested on mac
 */

/*** WORKFLOW ***/
/*
* 1- Read command line args - pass csv file
* 2- Loop through file, collect properties on each line (currently app is hardcoded to accept four columns)
* 3- Check if each email address is valid
* 4- Send information to email library
* 5- Connect to SES using AWS Credentials
* 6- Sends message using AWS SES
* 7- Returns OK message if sent was successful
 */

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/pborman/getopt"
	"github.com/warrensbox/bulk-emailer/lib"
)

func main() {

	helpFlag := getopt.BoolLong("help", 'h', "displays help message")
	filePath := getopt.StringLong("file", 'f', "", "--file", "Path to csv file")

	getopt.Parse()

	if *helpFlag {
		usageMessage()
	} else if *filePath == "" {
		fmt.Println("You must provide a csv file")
		usageMessage()
	}

	csvFile, errorFile := os.Open(*filePath)

	if errorFile != nil {
		log.Fatal("Unable to open file")
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var wg sync.WaitGroup
	fmt.Println("Attempting to send messages!")
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		var orgname = line[0]
		var position = line[1]
		var name = line[2]
		var email = line[3]

		wg.Add(1)
		go func(email string, orgname string, position string, name string) {
			lib.SendEmail(email, orgname, position, name)
			wg.Done()
		}(email, orgname, position, name)

	}

	wg.Wait()
}

func usageMessage() {
	fmt.Print("\n\n")
	getopt.PrintUsage(os.Stderr)
	fmt.Println("Pass path to csv file to -f or --file; example ./main -f test.csv")
}
