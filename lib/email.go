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
	SUPPORTEMAIL = "support@bittrove.com"
	APPURL       = "https://bittrove.org"
	APPNAME      = "BitTrove"
	APPCOPY      = "Ⓒ 2019 BitTrove - Crafted with ❤ in San Francisco"
	EMAILSUBJECT = "Wake Up! It's 2019"
	EMAILYOUTUBE = "https://youtu.be/5nrYOoPeUyA"
	EMAILWEBSITE = "https://bittrove.org"
	SEND_OK      = "{ \"message\": \"Message sent successfully\"}"
	SEND_NOT_OK  = "{ \"message\": \"Unble to send message\"}"
	IMGHEADER    = "https://s3.us-east-2.amazonaws.com/kepler-images/bittrove/bittrove_with_logo_email.png"
)

func SendEmail(contact_email string, orgname string, position string, name string) string {

	session := session.Must(session.NewSession())
	svc := ssm.New(session)

	SMTPPASS := getEmailCredential(svc, "SMTP_PASS")

	SMTPUSER := getEmailCredential(svc, "SMTP_USER")

	SMTPEMAIL := getEmailCredential(svc, "SMTP_EMAIL")

	SMTPPORT := getEmailCredential(svc, "SMTP_PORT")

	emailContent := composeEmail(orgname, position, name)
	emailHeadFoot := composeEmailFooterHeader(APPURL, APPNAME, APPCOPY)

	emailBody, errBody := emailHeadFoot.GenerateHTML(emailContent)
	if errBody != nil {
		fmt.Println(errBody)
	}

	var msg gophermail.Message

	msg.SetFrom(SUPPORTEMAIL)
	msg.AddTo(contact_email)
	msg.SetReplyTo(SUPPORTEMAIL)
	msg.Subject = orgname + " " + EMAILSUBJECT
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
	//fmt.Println(smtpInfo)
	return smtpInfo
}

func composeEmail(orgname string, position string, name string) hermes.Email {

	if orgname == "" {
		orgname = "your organization"
	}

	return hermes.Email{
		Body: hermes.Body{
			Title: "Hello " + name,
			Intros: []string{
				"Excel Spreadsheets and paper files are obsolete!\n",
				"Say hello to BitTrove! It’s an intuitive web app that allows organizations to keep cumulative records of members, meetings, events, and attendance in one convenient location. " +
					"Many organizations from different universities use this program. So what are you waiting for? 	" +
					"Leave a legacy - " + orgname + " will love you for it!\n",
				"Please share this video ad with the officers in your organization: " + EMAILYOUTUBE + "\n",
				"Register your organization for free at: " + EMAILWEBSITE + "\n",
				"We look forward to hearing from you.\n",
			},
			Outros: []string{
				"Have more questions? Shoot us an email at " + SUPPORTEMAIL + ".",
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
