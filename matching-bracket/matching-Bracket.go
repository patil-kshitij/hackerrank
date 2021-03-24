package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Arr []string
	Tos int
}

func (s *Stack) Push(bracket string) {
	if s.Arr == nil {
		s.Arr = make([]string, 0, 1)
		s.Tos = -1
	}
	s.Tos = s.Tos + 1
	s.Arr = append(s.Arr, bracket)

}

func (s *Stack) Pop() error {
	if len(s.Arr) == 0 || s.Tos < 0 {
		return errors.New("Popping from nil stack")
	}
	s.Arr = s.Arr[:s.Tos]
	s.Tos = s.Tos - 1
	return nil
}

func (s *Stack) Peek() string {
	if s.Tos >= 0 {
		return s.Arr[s.Tos]
	}
	return "nil"
}

// Complete the isBalanced function below.
func isBalanced(s string) string {
	stackObj := &Stack{Tos: -1}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stackObj.Push(string(s[i]))
		} else if string(s[i]) == ")" || string(s[i]) == "]" || string(s[i]) == "}" {

			if string(s[i]) == ")" && stackObj.Peek() == "(" {
				err := stackObj.Pop()
				if err != nil {
					return "NO"
				}
			} else if string(s[i]) == "]" && stackObj.Peek() == "[" {
				err := stackObj.Pop()
				if err != nil {
					return "NO"
				}
			} else if string(s[i]) == "}" && stackObj.Peek() == "{" {
				err := stackObj.Pop()
				if err != nil {
					return "NO"
				}
			} else {
				return "NO"
			}

		}
	}
	if stackObj.Tos == -1 {
		return "YES"
	}
	return "NO"
}

func main() {
	inputFile, _ := os.Open("trial.txt")
	reader := bufio.NewReaderSize(inputFile, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
