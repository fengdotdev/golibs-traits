package makerex

type ButtonWidget struct {
	content string
	action  Actions
}

func Button(content string, action Actions) *ButtonWidget {
	return &ButtonWidget{
		content: content,
		action:  action,
	}
}
