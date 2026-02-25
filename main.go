package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/shivgitcode/sx/internals"
)

var myFigure = figure.NewFigure("easyssh","doom",true)

func main(){
	if len(os.Args)<=1 {
		myFigure.Print();
		color.Cyan("Welcome to easyssh")
		fmt.Println("What are you waiting for just write easyssh run! That's it!")
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