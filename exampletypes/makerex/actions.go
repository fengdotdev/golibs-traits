package makerex

import "fmt"

type Actions struct {
	action string
}

func OpenUrl(url string) *Actions {
	return &Actions{
		action: fmt.Sprintf("open url: %s", url),
	}
}

func SendEmail(email string) *Actions {
	return &Actions{
		action: fmt.Sprintf("send email: %s", email),
	}
}
