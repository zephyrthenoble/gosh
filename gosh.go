package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"regexp"
	"os/exec"
)

var promt = "\n> "
var reader = bufio.NewReader(os.Stdin)
var last = ""
var PATH = "/usr/bin/"

func main() {
	for {
		exec_loop()
	}
}

type command struct {
	command   string
	arguments string
	input     string
}

func exec_loop() {
	fmt.Print(promt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	commands := SplitCommands(text)
	for _, elem := range commands {
		exe_cmd(elem)
	}
}

func SplitCommands(text string) []string {
	parts := strings.Fields(text)
	commands := []string{}
	last := 0
	for index, elem := range parts {
		if elem == "|" {
			commands = append(commands, strings.Join(parts[last:index], " "))
			last = index + 1
		}
	}
	commands = append(commands, strings.Join(parts[last:], " "))
	return commands
}

type Stack struct {
	stack []string
}

func (s Stack) len() int {
	return len(s.stack)
}
func (s Stack) Pop() (string, error) {
	if s.len() == 1 {
		val := s.stack[0]
		s.stack = s.stack[0:0]
		return val, nil
	}
	val := s.stack[s.len()-1]
	s.stack = s.stack[0 : s.len()-1]
	return val, nil
}
func (s Stack) Push(a string) {
	s.stack = append(s.stack, a)
}

func exe_cmd(cmd string) {
	// splitting head => g++ parts => rest of the command
	// assume whitespace in between all commands
	parts := strings.Fields(cmd)
	// if there is nothing, return
	if len(parts) == 0 {
		return
	}
	head := parts[0]
	if len(parts) > 1 {
		parts = parts[1:len(parts)]
		stack := Stack{}
		for _, elem := range parts {
			if strings.HasPrefix(elem, "\"") && strings.HasSuffix(elem, "\"") == false {
				fmt.Println("push", elem)
				stack.Push(elem)
			} else if strings.HasSuffix(elem, "\"") {
				fmt.Println("pop", elem)
				popped, _ := stack.Pop()
				fmt.Println(popped)
			}
		}
	}

	if head == parts[0] {
		command, err := exec.Command(PATH + head).CombinedOutput()
		fmt.Println(string(command), err)
	} else {
		command, err := exec.Command(PATH+head, parts...).CombinedOutput()
		fmt.Println(string(command), err)
	}
}
