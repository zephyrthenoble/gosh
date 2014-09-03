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

func SplitCommands(text string) ([]string) {
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
func exe_cmd(cmd string) {
	// splitting head => g++ parts => rest of the command
    // assume whitespace in between all commands
    parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return
	}
	head := parts[0]
	if len(parts) > 1 {
		parts = parts[1:len(parts)]
	}

    if head == parts[0] {
        command, err:= exec.Command(PATH+head).CombinedOutput()
    fmt.Println(string(command), err)
    } else {
        command, err:= exec.Command(PATH+head, parts...).CombinedOutput()
    fmt.Println(string(command), err)
    }
}

