package controllers

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

func GenerateSender(nama string, email string) string {
	var hasil string
	hasil = nama + " " + "<" + email + ">"
	return hasil
}

func SendMail(to []string, cc []string, subject, message string, email string, password string, sender string) error {
	body := "From: " + sender + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", email, password, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, email, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func GenerateEmail(jurusan string, angkatan int, absen int) string {
	email := jurusan + strconv.Itoa(angkatan)
	if absen/10 < 0 {
		email = email + "00"
	} else {
		email = email + "0"
	}
	email = email + strconv.Itoa(absen) + "@students.ithb.ac.id"
	return email
}
