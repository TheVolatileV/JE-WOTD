package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

var senderID string
var pass string

type apiVals struct {
	Japanese string   `json:"japanese,omitempty"`
	Reading  string   `json:"reading,omitempty"`
	English  []string `json:"english"`
	POS      []string `json:"partOfSpeech"`
}

// Mail the struct for general email structure
type Mail struct {
	toIds   []string
	subject string
	body    string
}

// SMTPServer struct
type SMTPServer struct {
	host string
	port string
}

// EmailObj struct
type EmailObj struct {
	Email string `json:"email"`
}

// ServerName concats the host and port
func (s *SMTPServer) ServerName() string {
	return s.host + ":" + s.port
}

// BuildMessage puts all parts of mail together to create the complete "message"
func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", senderID)
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

func initMail() Mail {
	mail := Mail{}
	mail.toIds = []string{"jpn.eng.wotd@gmail.com"}
	resp, err := http.Get("https://jewotd-api.herokuapp.com/api/v1/list")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var emailList []EmailObj

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &emailList)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range emailList {
		mail.toIds = append(mail.toIds, v.Email)
	}

	mail.subject = "This is the email subject"

	vals := getData()
	mail.body = `<!DOCTYPE html>
	<html>
		<head>
			<style>
				.background {
					height: 100%;
				}
				.main__container {
					width: 70%;
					margin: 0 auto;
					margin-top: 5%;
					font-family: "Yu Gothic";
					border-radius: 30px;
					padding-top: 20px;
					padding-bottom: 30px;
				}
				.main__body {
					width: 70%;
					margin: 0 auto;
					margin-top: 5%;
					font-family: "Yu Gothic";
					border-radius: 30px;
					margin: 0 auto;
					font-size: 20px;
					padding-left:110px;
				}
				.title__container {
					width: 50%;
					margin: auto;
					text-align: center;
					border-bottom: 1px solid black;
					margin-bottom: 2%;
				}
				.title__text {
					font-weight: lighter;
				}
			</style>
		</head>
		<body class="background">
			<div class="main__container">
			<div class="title__container">
				<h1 class = "title__text">今日の言葉</h1>
			</div>
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
	return mail
}

func authAndWrite(mail Mail, messageBody string) {
	SMTPServer := SMTPServer{host: "smtp.gmail.com", port: "465"}
	//build an auth
	auth := smtp.PlainAuth("", senderID, pass, SMTPServer.host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTPServer.host,
	}

	conn, err := tls.Dial("tcp", SMTPServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, SMTPServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(senderID); err != nil {
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

func send() {

	mail := initMail()

	messageBody := mail.BuildMessage()

	authAndWrite(mail, messageBody)

}

func main() {
	senderID = os.Getenv("email")
	pass = os.Getenv("pass")
	send()
}
