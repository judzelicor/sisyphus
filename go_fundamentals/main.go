package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var points int = 3

	question, answer := bioQuestion()

	fmt.Println(question)
	fmt.Println(answer)
	fmt.Println("Correct! You now have", points, "points!")
}

func bioQuestion() (string, string) {
	return "What is the enzyme responsible for starting and controling pyrimidine synthesis?", "CAD protein"
}
