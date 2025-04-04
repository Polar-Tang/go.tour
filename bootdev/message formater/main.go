package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}
type Bold struct {
	message string
}
type Code struct {
	message string
}

func (txt PlainText) Format() string {
	return txt.message
}

func (b Bold) Format() string {
	bolded := "**" + b.message + "**"
	return bolded
}

func (c Code) Format() string {
	coded := "`" + c.message + "`"
	return coded
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
