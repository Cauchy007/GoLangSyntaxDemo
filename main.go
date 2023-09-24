package main

import (
	"flag"
	"fmt"
	cmd "github.com/Cauchy007/learn_go/hello"
	cmd_new "github.com/Cauchy007/learn_go/hello/hello_new"
	"os"
)

var name string

func init() {
	name = "cauchy"
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question1234567890")
		flag.PrintDefaults()
	}

	flag.StringVar(&name, "name", "cauchy", "help message for name")
	flag.Parse()

	fmt.Print(cmd.Hello(name))
	cmd_new.HelloNew()

}
