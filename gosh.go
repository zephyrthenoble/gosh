package gosh

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
}
var search_for_pipe = "(?:^||)(\"(?:[^\"]+|\"\")*\"|[^|]*)"
func ParseCommands(line string) ([]command) {
    regex, err := regexp.Compile(search_for_pipe)
    if regex.MatchString(line){
        fmt.Println("found")
    }

    return make( []command, 4)
}
