package dmail
import (
	"github.com/go-gomail/gomail"
	"github.com/decoz/go_lib/dot"
	"fmt"
	"io/ioutil"
)


type DMail struct{
	dialer *gomail.Dialer
	mqueue chan *dot.Dot

	smtp	string
	server 	string
	id 		string
	pass	string
}

func New() *DMail {

	d := new( DMail )
	d.mqueue = make(chan *dot.Dot)

	go d.sendmail()
	return d
}

func (dmail *DMail) sendmail() {

	for {
		mail := <-dmail.mqueue
		//fmt.Println("--- got req for mail --------------------------------------------")

		dmail.Send(mail)
	}
}

func (dmail *DMail) Put(dreq *dot.Dot) *dot.Dot{
	switch( dreq.Val()){
	case "setsmtp" : dmail.SetSmtp(dreq)
	case "send"		:  dmail.mqueue <- dreq
	}
	return nil
}

func (dmail *DMail) Load(fname string){

	buff, err := ioutil.ReadFile(fname)
	if err == nil {
		d_mail := dot.TabParse(buff)
		dmail.SetSmtp(d_mail)
	}
}


func (dmail *DMail) SetSmtp(d_mail *dot.Dot){


	dmail.server = d_mail.ChildV("server").CVal()
	dmail.smtp = d_mail.ChildV("smtp").CVal()
	dmail.id = d_mail.ChildV("id").CVal()
	dmail.pass = d_mail.ChildV("pass").CVal()

	fmt.Println("dmail:",dmail.server,dmail.id,dmail.pass,"\n")
	dmail.dialer = gomail.NewDialer(dmail.smtp, 587,dmail.id, dmail.pass)

}

func (dmail *DMail) Send(d_msg *dot.Dot){

	//fmt.Println(d_msg)

	who := d_msg.ChildV("to").CVal()
	subject := d_msg.ChildV("subject").CVal()
	content := d_msg.ChildV("content").CVal()
	from := dmail.id + "@" + dmail.server

	fmt.Println(who, subject, content, from)


	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", who )
	m.SetAddressHeader("Cc", from, dmail.id)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	if err:= dmail.dialer.DialAndSend(m); err !=  nil {
		panic(err)
	}




}
