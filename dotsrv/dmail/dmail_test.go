package dmail_test

import (
	"github.com/decoz/go_lib/dot_tools/dmail"
	"github.com/decoz/go_lib/dot"
	"testing"
	"fmt"
)

func TestGomail(t *testing.T){

	mail := dmail.New()
	mail.Load("mail.acc")

	data := "to	zoced@naver.com \n" +
		"subject	test mail 한글\n" +
		"content	nothing special \n"

	d_msg := dot.TabParse( []byte(data) )
	fmt.Println(d_msg.ChildV("subject"))
	mail.Send(d_msg)

}
