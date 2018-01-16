package main

// Greeting structure
type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey " + name + "!"
	return
}

func PrintLine(s string) {
	println(s)
}
func CreatePrinter(custom string) Printer {
	return func(s string) {
		println(s + custom)
	}
}
func Greet(salutation Salutation, do Printer) {
	message, alternate := (CreateMessage(salutation.name, salutation.greeting))
	do(message)
	do(alternate)
}

func main() {
	var s = Salutation{"Paul", "Hello"}

	Greet(s, CreatePrinter("!"))
}
