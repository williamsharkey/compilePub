package main

import (
	"github.com/williamsharkey/compilePub"
	"github.com/williamsharkey/compilePub/link"
	"os"
)

func main(){
	os.Args=append(os.Args, "hello/helloworld.go")
	//os.Args=append(os.Args, "C:/Users/william/go/src/github.com/williamsharkey/compilePub/hello/helloworld.go")
	compilePub.Compile()

	os.Args[1]="-ld_o"
	os.Args=append(os.Args, "hello/helloworld.exe","helloworld.o" )
	compilePub.Link()
	errorCount:=link.ErrorOrExit()
	os.Exit(errorCount)


}