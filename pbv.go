package main

func main() {
	// pass by reference
	msg := "Hello"
	say(msg)
}
func say(msg string) {
	println(msg)
}
