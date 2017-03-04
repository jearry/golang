package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"strings"
)

const MailContentTpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
    </head>
    <body>
		<h3>{{.TotalCount}}个文件被修改，清单如下：</h3>
		<table border="1">
		<tr>
		<th>序号</th>
		<th>文件</th>
		<th>操作</th>
		</tr>
		{{range $index, $value := .Files}}
		    <tr>
			<td>{{$index}}</td><td>{{$value.FileName}}</td><td>{{$value.Op}}</td>
			</tr>
		{{end}}
		</table>
        
    </body>
</html>
`

var (
	MailContentTemplate = template.Must(template.New("Web").Parse(MailContentTpl))
)

func sendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func sendAlarmMail(files []FileModifyDef) {
	user := "username@163.com"
	password := "passwd"
	host := "smtp.163.com:25"
	to := "username@163.com"

	subject := "服务器文件改动告警"

	v := ModifyDef{TotalCount: len(files), Files: files}

	b := bytes.NewBuffer(make([]byte, 0))
	if err := MailContentTemplate.Execute(b, v); err != nil {
		log.Println("template exectue err: ", err.Error())
		return
	}
	body := b.String()

	log.Println("send alarm email...")
	err := sendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		log.Println("send mail error", err)
	} else {
		log.Println("send mail success!")
	}

}
