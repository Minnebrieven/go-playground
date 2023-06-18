package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	recipient := "libr.libr1711@gmail.com"
	m := gomail.NewMessage()
	m.SetHeader("From", "gofitapi@gmail.com")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "GoFit OTP Forgot Password")
	m.SetBody("text/plain", "Here's your OTP : 8965. Don't share it with anybody.")

	d := gomail.NewDialer("email-smtp.ap-southeast-1.amazonaws.com", 587, "AKIARZAQ63N5HOFJGVML", "BAiES959EDzwBfRPqzXexQokZVJ3vnS370et48HBvPwz")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Printf("email sended to %s", recipient)
	}

}
