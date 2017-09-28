package main

import (
	"os"
	"github.com/hpcloud/tail"
	"fmt"
	"strings"
	"github.com/fatih/color"
)

var colors []func(a ...interface{}) string

func main() {
	args := os.Args[1:]

	colors = []func(a ...interface{}) string{
		color.New(color.FgGreen).SprintFunc(),
		color.New(color.FgYellow).SprintFunc(),
		color.New(color.FgBlue).SprintFunc(),
		color.New(color.FgMagenta).SprintFunc(),
		color.New(color.FgCyan).SprintFunc(),
	}

	for i, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			c := colors[i%len(colors)]
			if strings.Contains(strings.ToLower(arg), "error") {
				c = color.New(color.FgRed).SprintFunc()
			}

			go tailFile(arg, c)
		}
	}

	select {}
}

func tailFile(file string, colorFunc func(a ...interface{}) string) {
	t, err := tail.TailFile(file, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		fmt.Printf("[%s] %s\n", colorFunc(fmt.Sprintf(" %s ", file)), line.Text)
	}
}
