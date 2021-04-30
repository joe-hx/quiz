package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the quiz! Enter \"q\" to quit")
	input := ""
	for !quit(input) {
		question := newQuestion()
		input = ask(question)
		answer(question, input)
	}
	fmt.Println("Goodbye!")
	wait := time.After(2 * time.Second); <-wait
}

func quit(input string) bool {
	return input == "q" || input == "quit" || input == "exit" || input == "stop"
}

type Question struct {
	valueA int
	valueB int
	operator string
}

func newQuestion() *Question {
	q := new(Question)
	q.valueA = rand.Intn(13)
	q.valueB = rand.Intn(13)
	ops := []string{"+","-","*"}
	q.operator = ops[rand.Intn(len(ops))]
	return q
}

func ask(q *Question) string {
	var input string
	question := "What is "+ strconv.Itoa(q.valueA) +" "+q.operator+" "+ strconv.Itoa(q.valueB) +"? "
	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}

func answer(q *Question, input string) {
	if !quit(input) {
		given, _ := strconv.Atoi(input)
		answer := 0
		switch q.operator {
		case "+":
			answer = q.valueA + q.valueB
		case "-":
			answer = q.valueA - q.valueB
		case "*":
			answer = q.valueA * q.valueB
		}
		if answer == given {
			fmt.Println("Correct! "+strconv.Itoa(q.valueA)+" "+q.operator+" "+strconv.Itoa(q.valueB)+" = "+strconv.Itoa(answer)+"!")
		} else {
			fmt.Println("Wrong! :( "+strconv.Itoa(q.valueA)+" "+q.operator+" "+strconv.Itoa(q.valueB)+" = "+strconv.Itoa(answer)+"")
		}
	}
}
