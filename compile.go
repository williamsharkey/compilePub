// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compilePub

import (
	"fmt"
	"github.com/williamsharkey/compilePub/gotool"

	"github.com/williamsharkey/compilePub/internal/amd64"
	"github.com/williamsharkey/compilePub/internal/arm"
	"github.com/williamsharkey/compilePub/internal/arm64"
	"github.com/williamsharkey/compilePub/internal/gc"
	"github.com/williamsharkey/compilePub/internal/mips"
	"github.com/williamsharkey/compilePub/internal/mips64"
	"github.com/williamsharkey/compilePub/internal/objabi"
	"github.com/williamsharkey/compilePub/internal/ppc64"
	"github.com/williamsharkey/compilePub/internal/s390x"
	"github.com/williamsharkey/compilePub/internal/x86"
	"github.com/williamsharkey/compilePub/link"
	"log"
	"os"
)

var archInits = map[string]func(arch *gc.Arch){
	"386":      x86.Init,
	"amd64":    amd64.Init,
	"amd64p32": amd64.Init,
	"arm":      arm.Init,
	"arm64":    arm64.Init,
	"mips":     mips.Init,
	"mipsle":   mips.Init,
	"mips64":   mips64.Init,
	"mips64le": mips64.Init,
	"ppc64":    ppc64.Init,
	"ppc64le":  ppc64.Init,
	"s390x":    s390x.Init,
}

func Compile() {
	// disable timestamps for reproducible output
	log.SetFlags(0)
	log.SetPrefix("compile: ")

	archInit, ok := archInits[objabi.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", objabi.GOARCH)
		os.Exit(2)
	}

	gc.Main(archInit)
	//gc.Exit(0)
}

func Link() {
	link.Link()
}

func Get(cwd string) {
	z := os.Args[0:1]
	os.Args = append(z, "get", "-v")

	gotool.GoTool(cwd)
	//gotool.RunGet()
}

func Build(dir string) (err error) {
	err = os.Chdir(dir)

	if err != nil {
		return
	}

	x, err := os.Getwd()

	if err != nil {
		return
	}
	fmt.Println(x)

	z := os.Args[0:1]
	os.Args = append(z, "-u") //"hello/helloworld.go")

	Get(x)

	os.Args = append(z, "-o", "out.o", dir) //"hello/helloworld.go")
	//os.Args=append(os.Args, "C:/Users/william/go/src/github.com/williamsharkey/compilePub/hello/helloworld.go")
	Compile()

	os.Args = append(z, "-ld_o", "out.exe", "out.o")

	Link()
	errorCount := link.ErrorOrExit()
	os.Exit(errorCount)
	return
}
