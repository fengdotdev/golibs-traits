package makerex

type TextWidget struct {
	content string
}

func Text(content string) *TextWidget {
	return &TextWidget{
		content: content,
	}
}
