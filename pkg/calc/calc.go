package calc

import (
	"container/list"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	rpn := toRPN(splitExpression(expression))
	stack := list.New()
	var tmp float64

	for _, token := range rpn {
		switch token {
		case "+", "-", "*", "/":
			if stack.Len() < 2 {
				return 0, ErrInvalidExpr
			}
			b, a := stack.Remove(stack.Back()).(float64), stack.Remove(stack.Back()).(float64)

			switch token {
			case "+":
				tmp = a + b
			case "-":
				tmp = a - b
			case "*":
				tmp = a * b
			case "/":
				if b == 0 {
					return 0, ErrDivisionByZero
				}
				tmp = a / b
			}
			stack.PushBack(tmp)
		default:
			value, err := strconv.ParseFloat(strings.ReplaceAll(token, ",", "."), 64)
			if err != nil {
				return 0, ErrInvalidExpr
			}
			stack.PushBack(value)
		}
	}

	if stack.Len() != 1 {
		return 0, ErrInvalidExpr
	}

	return stack.Back().Value.(float64), nil
}

func splitExpression(expression string) []string {
	answer := make([]string, 0, len(expression))
	subStr := make([]string, 0, len(expression))
	for _, v := range strings.Split(expression, "") {
		if ((v >= "0") && (v <= "9")) || (v == ".") || (v == ",") {
			subStr = append(subStr, v)
		} else if v != " " {
			if len(subStr) > 0 {
				answer = append(answer, strings.Join(subStr, ""))
				subStr = make([]string, 0, len(expression))
			}
			answer = append(answer, v)
		}
	}
	if len(subStr) > 0 {
		answer = append(answer, strings.Join(subStr, ""))
	}
	return answer
}

func toRPN(expression []string) []string {
	stack := list.New()
	answer := make([]string, 0, len(expression))
	for _, v := range expression {
		switch {
		case v >= "0" && v <= "9" || strings.ContainsAny(v, ".,"):
			answer = append(answer, v)
		case v == "(":
			stack.PushBack(v)
		case v == ")":
			for stack.Len() > 0 && stack.Back().Value.(string) != "(" {
				answer = append(answer, stack.Remove(stack.Back()).(string))
			}
			if stack.Len() > 0 {
				stack.Remove(stack.Back())
			}
		default:
			for stack.Len() > 0 && precedence(stack.Back().Value.(string)) >= precedence(v) {
				answer = append(answer, stack.Remove(stack.Back()).(string))
			}
			stack.PushBack(v)
		}
	}
	for stack.Len() > 0 {
		answer = append(answer, stack.Remove(stack.Back()).(string))
	}
	return answer
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}
