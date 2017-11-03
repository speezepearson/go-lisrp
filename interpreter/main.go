package main

import (
	"bufio"
	"fmt"
	"lisrp"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("waiting for input...")
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("input is %s; tokenizing...\n", strconv.Quote(line))
	tokens, err := lisrp.Tokenize(line)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("tokens are %s; parsing...\n", tokens)
	expr, err := lisrp.ParseTokens(tokens)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("parsed is %s; evaluating...\n", expr)
	val, lerr := expr.Eval(lisrp.MakeDefaultEnv())
	if lerr != nil {
		fmt.Printf("lisrp exception: %s\n", lerr)
	} else {
		fmt.Printf("lisrp value: %s\n", val)
	}
}
