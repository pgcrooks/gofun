package main

// Greeting structure
type Salutation struct {
	name     string
	greeting string
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "Hey " + name + "!"
}

func Greet(salutation Salutation) {
	message, alternate := (CreateMessage(salutation.name, salutation.greeting))
	println(message)
	println(alternate)
}

func main() {
	var s = Salutation{"Paul", "Hello"}

	Greet(s)
}
