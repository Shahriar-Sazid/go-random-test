package email

import (
	"log"

	clients "magic.pathao.com/intsvc/comm-clients"
	"magic.pathao.com/intsvc/comm-clients/email"
)

func TestMail() {
	c := clients.New("http://mailpusher.p-stageenv.xyz",
		"pathao_username",
		"pathao_password")

	emailClient := email.New(c)
	err := emailClient.Send(&email.Request{
		From: map[string]string{
			"email": "no-reply@pathao.com",
			"name":  "Pathao Ride",
		},
		To: map[string]string{
			"email": "saddam@pathao.com",
			"name":  "Saddam H",
		},
		Subject:  "Hello from Pathao Ride",
		BodyType: email.BodyTypeHTML,
		Body:     "<h4>This is a test email from Pathao ride service.</h4>",
		Attachments: []map[string]string{
			{
				"type":        "text/text",
				"filename":    "hello.txt",
				"disposition": "attachment",
				"content":     "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdlLCBleGNlZWRzIHRoZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4=",
			},
		},
	})
	if err != nil {
		log.Println("Failed to send email:", err)
		return
	}
}
