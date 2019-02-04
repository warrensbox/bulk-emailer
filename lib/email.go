package lib

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	gophermail "gopkg.in/jpoehls/gophermail.v0"
)

// const (
// 	SUPPORTEMAIL = "support@bittrove.com"
// 	APPURL       = "https://bittrove.org"
// 	APPNAME      = "BitTrove"
// 	APPCOPY      = "Ⓒ 2019 BitTrove - Crafted with ❤ in San Francisco"
// 	EMAILSUBJECT = "Wake Up! It's 2019"
// 	EMAILYOUTUBE = "https://youtu.be/5nrYOoPeUyA"
// 	EMAILWEBSITE = "https://bittrove.org"
// 	IMGHEADER = "https://s3.us-east-2.amazonaws.com/kepler-images/bittrove/bittrove_with_logo_email.png"
// )

// SendEmail : send email
func SendEmail(toEmail string, fromEmail string, msgSubject string, msgContent string) {

	session := session.Must(session.NewSession())
	svc := ssm.New(session)

	SMTPPASS := getEmailCredential(svc, "SMTP_PASS")

	SMTPUSER := getEmailCredential(svc, "SMTP_USER")

	SMTPEMAIL := getEmailCredential(svc, "SMTP_EMAIL")

	SMTPPORT := getEmailCredential(svc, "SMTP_PORT")

	emailContent := composeEmail(msgContent)

	var msg gophermail.Message

	msg.SetFrom(fromEmail)
	msg.AddTo(toEmail)
	msg.SetReplyTo(fromEmail)
	msg.Subject = msgSubject
	msg.HTMLBody = emailContent

	auth := smtp.PlainAuth("", SMTPUSER, SMTPPASS, SMTPEMAIL)
	errEmail := gophermail.SendMail(SMTPPORT, auth, &msg)
	if errEmail != nil {
		fmt.Println(errEmail)
		fmt.Printf("FAILED to send email to : %s\n", toEmail)
		return
	}

	fmt.Printf("SUCCESSFULLY to send email to : %s\n", toEmail)
}

func getEmailCredential(svc *ssm.SSM, val string) string {

	param := &ssm.GetParameterInput{
		Name:           aws.String(val),
		WithDecryption: aws.Bool(true),
	}

	paramVal, err := svc.GetParameter(param)
	ErrorExit("GetParameters", err)

	smtpInfo := *paramVal.Parameter.Value
	//fmt.Println(smtpInfo)
	return smtpInfo
}

func composeEmail(content string) string {

	return content
}

func ErrorExit(msg string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s, %v\n", msg, e)
		os.Exit(1)
	}
}
