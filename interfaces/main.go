package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// 	fmt.Println((eb.getGreeting()))
// }

// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	// VERY vustom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
