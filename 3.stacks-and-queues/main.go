package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Problem Statement:
// Given a postfix expression, evaluate it.

// Example:
// Input: ["2", "1", "+", "3", "*"]
// Output: 9 (Explanation: (2 + 1) * 3)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("end of stack")
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, nil
}

func evaluate(list []string) int {
	res := 0
	stack := make(Stack, 0)
	for _, s := range list {
		v, err := strconv.Atoi(s)
		if err == nil {
			stack.Push(v)
			continue
		}

		switch s {
		case "+":
			for i, err := stack.Pop(); err == nil; i, err = stack.Pop() {
				res += i
			}

		case "-":
			for i, err := stack.Pop(); err == nil; i, err = stack.Pop() {
				res -= i
			}
		case "*":
			for i, err := stack.Pop(); err == nil; i, err = stack.Pop() {
				res *= i
			}
		case "/":
			for i, err := stack.Pop(); err == nil; i, err = stack.Pop() {
				res /= i
			}
		}

	}

	return res

}

func main() {
	list := []string{"2", "1", "+", "3", "*"}

	fmt.Println(evaluate(list))
}
