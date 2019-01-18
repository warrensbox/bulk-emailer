package lib

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/matcornic/hermes"
	gophermail "gopkg.in/jpoehls/gophermail.v0"
)

const (
	FALCONEMAIL   = "support@warrensbox.com"
	FALCONURL     = "https://warrensbox.github.io/falcon-form"
	FALCONAME     = "Falcon Form"
	FALCONCOPY    = "Ⓒ 2018 Warrensbox - Crafted with ❤ in Iowa"
	FALCONSUBJECT = "You've got a message from Falcon Form! "
	SEND_OK       = "{ \"message\": \"Message sent successfully\"}"
	SEND_NOT_OK   = "{ \"message\": \"Unble to send message\"}"
	IMGHEADER     = "https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/falcon_form/falcon-form_350.png"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func SendEmail(owner_email string, contact_email string, contact_phone string, contact_name string, msg_content string) string {

	session := session.Must(session.NewSession())
	svc := ssm.New(session)

	SMTPPASS := getEmailCredential(svc, "SMTP_PASS")

	SMTPUSER := getEmailCredential(svc, "SMTP_USER")

	SMTPEMAIL := getEmailCredential(svc, "SMTP_EMAIL")

	SMTPPORT := getEmailCredential(svc, "SMTP_PORT")

	emailContent := composeEmail(contact_email, contact_phone, contact_name, msg_content)
	emailHeadFoot := composeEmailFooterHeader(FALCONURL, FALCONAME, FALCONCOPY)

	emailBody, errBody := emailHeadFoot.GenerateHTML(emailContent)
	if errBody != nil {
		fmt.Println(errBody)
	}

	var msg gophermail.Message

	msg.SetFrom(FALCONEMAIL)
	msg.AddTo(owner_email)
	msg.SetReplyTo(contact_email)
	msg.Subject = FALCONSUBJECT
	msg.HTMLBody = emailBody

	auth := smtp.PlainAuth("", SMTPUSER, SMTPPASS, SMTPEMAIL)
	errEmail := gophermail.SendMail(SMTPPORT, auth, &msg)
	if errEmail != nil {
		fmt.Println(errEmail)
		return SEND_NOT_OK
	}
	return SEND_OK
}

func getEmailCredential(svc *ssm.SSM, val string) string {

	param := &ssm.GetParameterInput{
		Name:           aws.String(val),
		WithDecryption: aws.Bool(true),
	}

	paramVal, err := svc.GetParameter(param)
	ErrorExit("GetParameters", err)

	smtpInfo := *paramVal.Parameter.Value
	fmt.Println(smtpInfo)
	return smtpInfo
}

func composeEmail(contact_email string, phone_number string, contact_name string, msg_content string) hermes.Email {

	dictionary := []hermes.Entry{
		{Key: "Name", Value: contact_name},
		{Key: "Phone", Value: phone_number},
		{Key: "Email", Value: contact_email},
		{Key: "Message", Value: msg_content},
	}

	if phone_number == "" && contact_name == "" {
		dictionary = []hermes.Entry{
			{Key: "Email", Value: contact_email},
			{Key: "Message", Value: msg_content},
		}
	}

	if phone_number == "" && contact_name != "" {
		dictionary = []hermes.Entry{
			{Key: "Name", Value: contact_name},
			{Key: "Email", Value: contact_email},
			{Key: "Message", Value: msg_content},
		}
	}

	if phone_number != "" && contact_name == "" {
		dictionary = []hermes.Entry{
			{Key: "Phone", Value: phone_number},
			{Key: "Email", Value: contact_email},
			{Key: "Message", Value: msg_content},
		}
	}

	return hermes.Email{
		Body: hermes.Body{
			Title: "Hello",
			Intros: []string{
				"You got a message from a visitor..",
			},
			Dictionary: dictionary,
			Outros: []string{
				"Need help, or have questions? Shoot us an email at support@warrensbox.com.",
			},
		},
	}
}

func composeEmailFooterHeader(url string, name string, copyright string) hermes.Hermes {

	h := hermes.Hermes{
		Theme: new(hermes.Default),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: name,
			Link: url,
			// Optional product logo
			Logo:      IMGHEADER,
			Copyright: copyright,
		},
	}

	return h
}

func ErrorExit(msg string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s, %v\n", msg, e)
		os.Exit(1)
	}
}
