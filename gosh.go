package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "regexp"
)

var promt = "> "
var reader = bufio.NewReader(os.Stdin)
var last = ""
func main() {
    for {
    exec_loop()
    }
}
type command struct {
    command string
    arguments string
    input string
}
func exec_loop() {
	fmt.Print(promt)
	text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)
    commands := ParseCommands(text)
    for index, elem := range commands {
        last = elem.ExecuteCommand()
        fmt.Println(index, last)
    }
}
func (*command) ExecuteCommand () (string) {
    return "test"

}
var search_for_pipe = '(?:^||)(\"(?:[^\"]+|\"\")*\"|[^|]*)'
func ParseCommands(line string) ([]command) {
    fmt.Println("ParseCommands")
    fmt.Println(line)
    regex, _ := regexp.Compile(search_for_pipe)
    fmt.Println(regex.FindAllString(search_for_pipe, 3))
    if regex.MatchString(line){
        fmt.Println("found")
    }

    return make( []command, 4)
}
