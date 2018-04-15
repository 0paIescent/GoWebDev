package main

import (
	"log"

	"github.com/emersion/go-imap/client"
)

const server = "imap.gmail.com:993"

func main() {
	log.Println("Connecting to:", server)
	imap, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected")

	defer imap.Logout()

	if err := imap.Login("ticketing.practice@gmail.com", "18_@Dm1n~42"); err != nil {
		log.Fatalln(err)
	}
	log.Println("Logged In")

	mailboxes := make(chan *imap.MailboxInfo, 10)

}
