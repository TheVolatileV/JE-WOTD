package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

var pass = "rhinosarereallygray"

type apiVals struct {
	Japanese string   `json:"japanese,omitempty"`
	Reading  string   `json:"reading,omitempty"`
	English  []string `json:"english"`
	POS      []string `json:"partOfSpeech"`
}

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += fmt.Sprintf("Content-Type: text/html; charset=utf-8")
	message += "\r\n" + mail.body

	return message
}

func getData() apiVals {
	resp, err := http.Get("https://jewotd-api.herokuapp.com/api/v1")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var vals apiVals
	err = json.NewDecoder(resp.Body).Decode(&vals)
	if err != nil {
		log.Println(err)
		return apiVals{}
	}
	return vals
}

func isKatakana(vals apiVals) bool {
	if vals.Japanese == "" && vals.Reading != "" {
		return true
	}
	return false
}

func main() {
	mail := Mail{}
	mail.senderId = "jpn.eng.wotd@gmail.com"
	mail.toIds = []string{"jpn.eng.wotd@gmail.com"}
	mail.subject = "This is the email subject"

	vals := getData()
	mail.body = `<!DOCTYPE html>
	<html>
		<head>
			<style>
				.main__container {
					background-image: url(https://jewotd-spa.herokuapp.com/static/img/city-cropped.c18385d.jpeg);
					background-size: cover;
					filter: blur(5px);
					background-color: rgba(255,255,255,.65);
					width: 70%%;
					margin: 0 auto;
					margin-top: 5%%;
					font-family: "Yu Gothic";
					border-radius: 30px;
					padding-top: 20px;
					padding-bottom: 30px;
				}
				.main__body {
					background-color: rgba(255,255,255,.65);
					width: 70%%;
					margin: 0 auto;
					margin-top: 5%%;
					font-family: "Yu Gothic";
					border-radius: 30px;
					margin: 0 auto;
					font-size: 20px;
				}
			</style>
		</head>
		<body>
			<div class="main__container">
				<table class="main__body">
					<tbody>`
	if isKatakana(vals) {
		mail.body += fmt.Sprintf(`
						<tr>
							<td> 日本語: %s</td>
						</tr>
						<tr>
							<td>英語：%s</td>
						</tr>
						<tr>
							<td>品詞：%s</td>
						</tr>
					</tbody>
				</table>
			</div>
		</body>
	</html>`, vals.Reading, strings.Join(vals.English, ", "), strings.Join(vals.POS, ", "))
	} else {
		mail.body += fmt.Sprintf(`
						<tr>
							<td>日本語：%s</td>
						</tr>
						<tr>
							<td> 読み方: %s</td>
						</tr>
						<tr>
							<td>英語：%s</td>
						</tr>
						<tr>
							<td>品詞：%s</td>
						</tr>
					</tbody>
				</table>
			</div>
		</body>
	</html>`, vals.Japanese, vals.Reading, strings.Join(vals.English, ", "), strings.Join(vals.POS, ", "))
	}

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, pass, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
