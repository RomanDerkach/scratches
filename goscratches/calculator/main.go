package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calc(sign, a, b string) (result string) {
	var calcRes float64
	first, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Printf("Sorry one of the el was not number: %v\n", a)
		os.Exit(1)
	}
	second, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Printf("Sorry one of the el was not number: %v\n", a)
		os.Exit(1)
	}
	switch sign {
	case "+":
		calcRes = first + second
	case "-":
		calcRes = first - second
	case "*":
		calcRes = first * second
	case "/":
		calcRes = first / second
	default:
		fmt.Printf("Sorry we do not support that operation yet %v\n", sign)
		os.Exit(1)
	}
	result = fmt.Sprintf("%f", calcRes)
	return
}

func reverse(data []string) {
	last := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		data[i], data[last-i] = data[last-i], data[i]
	}
}

func prioritisation(opStack, postfixExp []string, sign string) ([]string, []string) {
	var priorityMap = map[string]int{"/": 2, "*": 2, "+": 1, "-": 1}
	if len(opStack) == 0 {
		opStack = append(opStack, sign)
	} else if lastEl := opStack[len(opStack)-1]; priorityMap[sign] > priorityMap[lastEl] {
		opStack = append(opStack, sign)
	} else {
		reverse(opStack)
		postfixExp = append(postfixExp, opStack...)
		opStack = []string{sign}
	}
	return opStack, postfixExp
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your data: ")
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("something went wrong")
		os.Exit(1)
	}

	notvalidRe, _ := regexp.Compile(`[^+\-*\/\s\d\.]`)
	valuesRe, _ := regexp.Compile(`\d+(\.\d+)?`)
	signRe, _ := regexp.Compile(`[+\-*/]+`)

	// get all the operand and signs using regex
	valueArr := valuesRe.FindAllString(data, -1)
	signArr := signRe.FindAllString(data, -1)
	if notvalidRe.MatchString(data) || len(valueArr) == 0 {
		fmt.Println("Your data is not valid, you cannot enter letters and operators different from {+ - / *}")
		os.Exit(1)
	}

	var exprArr []string

	// put everythin in right order into array
	for i, val := range signArr {
		exprArr = append(exprArr, valueArr[i])
		exprArr = append(exprArr, val)
	}
	exprArr = append(exprArr, valueArr[len(valueArr)-1])

	var postfixExp []string
	var opStack []string

	for _, el := range exprArr {
		switch el {
		case "*", "/", "+", "-":
			// why here we cant just send list and get them updated like with reverse
			opStack, postfixExp = prioritisation(opStack, postfixExp, el)
		default:
			postfixExp = append(postfixExp, el)
		}
	}
	reverse(opStack)
	postfixExp = append(postfixExp, opStack...)

	var calcArr []string
	for _, el := range postfixExp {
		switch el {
		case "*", "/", "+", "-":
			firstOp := calcArr[len(calcArr)-2]
			secondOp := calcArr[len(calcArr)-1]
			calcArr = append(calcArr[:len(calcArr)-2], calc(el, firstOp, secondOp))
		default:
			calcArr = append(calcArr, el)
		}
	}
	fmt.Printf("Here is a result : %s\n", calcArr[0])
}
