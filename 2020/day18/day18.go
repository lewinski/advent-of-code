package main

import (
	"fmt"
	"strconv"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	sum := 0
	for _, line := range input {
		sum += eval1(line)
	}
	fmt.Println("part1:", sum)

	sum = 0
	for _, line := range input {
		sum += eval2(line)
	}
	fmt.Println("part2:", sum)
}

type stack struct {
	vals []int
}

func (s *stack) len() int {
	return len(s.vals)
}

func (s *stack) push(x int) {
	s.vals = append(s.vals, x)
}

func (s *stack) pop() int {
	x := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return x
}

func (s *stack) peek() int {
	return s.vals[len(s.vals)-1]
}

const (
	add = iota
	times
	lparen
	rparen
)

func eval1(expr string) int {
	tokens := tokenize(expr)

	values := stack{}
	ops := stack{}
	mins := stack{}
	mins.push(2)

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			ops.push(add)
		case "*":
			ops.push(times)
		case "(":
			ops.push(lparen)
			mins.push(values.len() + 2)
		case ")":
			mins.pop()
			ops.pop()
		default:
			values.push(util.MustAtoi(tokens[i]))
		}

		if values.len() >= mins.peek() {
			if ops.peek() == add {
				y := values.pop()
				x := values.pop()
				values.push(x + y)
				ops.pop()
			} else if ops.peek() == times {
				y := values.pop()
				x := values.pop()
				values.push(x * y)
				ops.pop()
			}
		}
	}

	return values.pop()
}

func eval2(expr string) int {
	return eval2tokens(tokenize(expr))
}

func eval2tokens(tokens []string) int {
	// evaluate parens recursively
	simpleTokens := []string{}
	pos := 0
	depth := 0
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(" {
			if depth == 0 {
				pos = i + 1
			}
			depth++
		} else if tokens[i] == ")" {
			depth--
			if depth == 0 {
				num := eval2tokens(tokens[pos:i])
				simpleTokens = append(simpleTokens, strconv.Itoa(num))
			}
		} else if depth == 0 {
			simpleTokens = append(simpleTokens, tokens[i])
		}
	}

	// everything should be simple now
	// perform adds
adds:
	for {
		for i := 1; i < len(simpleTokens); i++ {
			if simpleTokens[i] == "+" {
				sum := strconv.Itoa(util.MustAtoi(simpleTokens[i-1]) + util.MustAtoi(simpleTokens[i+1]))
				newTokens := append(simpleTokens[:i-1], sum)
				newTokens = append(newTokens, simpleTokens[i+2:]...)
				simpleTokens = newTokens
				continue adds
			}
		}
		break
	}

	// perform products
mults:
	for {
		for i := 1; i < len(simpleTokens); i++ {
			if simpleTokens[i] == "*" {
				sum := strconv.Itoa(util.MustAtoi(simpleTokens[i-1]) * util.MustAtoi(simpleTokens[i+1]))
				newTokens := append(simpleTokens[:i-1], sum)
				newTokens = append(newTokens, simpleTokens[i+2:]...)
				simpleTokens = newTokens
				continue mults
			}
		}
		break
	}
	return util.MustAtoi(simpleTokens[0])
}

func tokenize(expr string) []string {
	tokens := []string{}
	for i := 0; i < len(expr); i++ {
		ch := expr[i]
		switch ch {
		case ' ':
			continue
		case '(':
			tokens = append(tokens, "(")
		case ')':
			tokens = append(tokens, ")")
		case '+':
			tokens = append(tokens, "+")
		case '*':
			tokens = append(tokens, "*")
		default:
			start := i
			for isDigit(ch) {
				i++
				if i == len(expr) {
					break
				}
				ch = expr[i]
			}
			tokens = append(tokens, expr[start:i])
			i--
		}
	}
	return tokens
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
