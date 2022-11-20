package main

import "fmt"

func main() {
	var question string = "What is the region of repetitive nucleotides at the end of a chromosome?"
	var answer string = "Telomerase!"

	fmt.Println(question)
	fmt.Println(answer)

	answerPointer(&answer)

	fmt.Println("Wrong! The correct answer is", answer)
}

func answerPointer(correctAnswer *string) {
	var answer string = "Telomere"
	*correctAnswer = answer
}
