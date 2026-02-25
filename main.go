package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/shivgitcode/sx/internals"
)

var myFigure = figure.NewFigure("sx","doom",true)

func main(){
	if len(os.Args)<=1 {
		myFigure.Print();
		color.Cyan("Welcome to sx (sshexpress)")
		fmt.Println("List of available commands")
		fmt.Print("\tinit\n\trun")
		return
	}
	cmd := os.Args[1]

	switch cmd{
	case "run":
		internals.RunCmd()
	case "init":
		internals.InitCmd()
	}

}