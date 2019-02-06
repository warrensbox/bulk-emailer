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
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/warrensbox/bulk-emailer/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	versionFlag   *bool
	fromEmail     *string
	msgContent    *string
	msgSubject    *string
	contactEmails *string
)

func init() {

	const (
		cmdDesc           = "Lets you send bulk emails"
		versionFlagDesc   = "Displays the version of bulk-emailer"
		fromEmailgDesc    = "Provide sender's email"
		msgContentDesc    = "Provide message content"
		msgSubjectDesc    = "Provide email subject"
		contactEmailsDecs = "Provide contact cvs file"
	)

	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	fromEmail = kingpin.Flag("from", fromEmailgDesc).Short('f').String()
	msgContent = kingpin.Flag("message", msgContentDesc).Short('c').String()
	msgSubject = kingpin.Flag("subject", msgSubjectDesc).Short('s').String()
	contactEmails = kingpin.Flag("contacts", contactEmailsDecs).Short('e').String()

}

func main() {

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()

	if *contactEmails == "" {
		fmt.Println("You must provide a csv file")
	}

	if *msgContent == "" {
		fmt.Println("You must provide a message content file")
	}

	csvFile, errorFile := os.Open(*contactEmails)
	defer csvFile.Close()
	contentFile, errorContentFile := os.Open(*msgContent)
	defer contentFile.Close()

	if errorFile != nil {
		log.Fatal("Unable to open csv file")
		os.Exit(1)
	}

	if errorContentFile != nil {
		log.Fatal("Unable to open content file")
		os.Exit(1)
	}

	baseStr, _ := ioutil.ReadAll(contentFile)

	strContent := string(baseStr)

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

		var email = line[0]

		var FromEmail = ""
		FromEmail = *fromEmail

		var MsgSubject = ""
		MsgSubject = *msgSubject

		wg.Add(1)
		go func(email string, FromEmail string, msgSubject string, msgStr string) {
			lib.SendEmail(email, FromEmail, MsgSubject, strContent)
			wg.Done()
		}(email, FromEmail, MsgSubject, strContent)

	}

	fmt.Println(strContent)
	wg.Wait()
}
