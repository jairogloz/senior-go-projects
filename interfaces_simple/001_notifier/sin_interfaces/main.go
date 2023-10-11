package main

import "fmt"

type User struct {
	Name  string
	Email string
	Phone string
}

type EmailSender struct {
	//...
}

func (e *EmailSender) SendEmail(user User, subject, body string) error {
	fmt.Printf("Sending email to %s at %s, with subject %s and body %s\n", user.Name, user.Email, subject, body)
	return nil
}

func NotifyUsers(users []User, subject string, body string, sender *EmailSender) error {
	for _, user := range users {
		err := sender.SendEmail(user, subject, body)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	users := []User{
		{
			Name:  "Juan",
			Email: "juan@gmail.com",
			Phone: "123456789",
		},
		{
			Name:  "Monica",
			Email: "monica@gmail.com",
			Phone: "987654321",
		},
	}

	sender := &EmailSender{}

	err := NotifyUsers(users, "Hello", "Hello world", sender)
	if err != nil {
		fmt.Println(err)
	}
}
