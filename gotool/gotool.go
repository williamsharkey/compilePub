// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate ./mkalldocs.sh

package gotool

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/williamsharkey/compilePub/gotool/base"
	"github.com/williamsharkey/compilePub/gotool/bug"
	"github.com/williamsharkey/compilePub/gotool/cfg"
	"github.com/williamsharkey/compilePub/gotool/clean"
	"github.com/williamsharkey/compilePub/gotool/doc"
	"github.com/williamsharkey/compilePub/gotool/envcmd"
	"github.com/williamsharkey/compilePub/gotool/fix"
	"github.com/williamsharkey/compilePub/gotool/fmtcmd"
	"github.com/williamsharkey/compilePub/gotool/generate"
	"github.com/williamsharkey/compilePub/gotool/get"
	"github.com/williamsharkey/compilePub/gotool/help"
	"github.com/williamsharkey/compilePub/gotool/list"
	"github.com/williamsharkey/compilePub/gotool/run"
	"github.com/williamsharkey/compilePub/gotool/test"
	"github.com/williamsharkey/compilePub/gotool/tool"
	"github.com/williamsharkey/compilePub/gotool/version"
	"github.com/williamsharkey/compilePub/gotool/vet"
	"github.com/williamsharkey/compilePub/gotool/work"
)

func init() {
	base.Commands = []*base.Command{
		work.CmdBuild,
		clean.CmdClean,
		doc.CmdDoc,
		envcmd.CmdEnv,
		bug.CmdBug,
		fix.CmdFix,
		fmtcmd.CmdFmt,
		generate.CmdGenerate,
		get.CmdGet,
		work.CmdInstall,
		list.CmdList,
		run.CmdRun,
		test.CmdTest,
		tool.CmdTool,
		version.CmdVersion,
		vet.CmdVet,

		help.HelpC,
		help.HelpBuildmode,
		help.HelpCache,
		help.HelpFileType,
		help.HelpGopath,
		help.HelpEnvironment,
		help.HelpImportPath,
		help.HelpPackages,
		test.HelpTestflag,
		test.HelpTestfunc,
	}
}

func RunGet(cwd string) {
	cmd := base.Command{}

	get.RunGet(&cmd, []string{}, cwd)
}

func GoTool(cwd string) {
	_ = go11tag
	flag.Usage = base.Usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		base.Usage()
	}

	cfg.CmdName = args[0] // for error messages
	if args[0] == "help" {
		help.Help(args[1:])
		return
	}

	// Diagnose common mistake: GOPATH==GOROOT.
	// This setting is equivalent to not setting GOPATH at all,
	// which is not what most people want when they do it.
	if gopath := cfg.BuildContext.GOPATH; filepath.Clean(gopath) == filepath.Clean(runtime.GOROOT()) {
		fmt.Fprintf(os.Stderr, "warning: GOPATH set to GOROOT (%s) has no effect\n", gopath)
	} else {
		for _, p := range filepath.SplitList(gopath) {
			// Some GOPATHs have empty directory elements - ignore them.
			// See issue 21928 for details.
			if p == "" {
				continue
			}
			// Note: using HasPrefix instead of Contains because a ~ can appear
			// in the middle of directory elements, such as /tmp/git-1.8.2~rc3
			// or C:\PROGRA~1. Only ~ as a path prefix has meaning to the shell.
			if strings.HasPrefix(p, "~") {
				fmt.Fprintf(os.Stderr, "go: GOPATH entry cannot start with shell metacharacter '~': %q\n", p)
				os.Exit(2)
			}
			if !filepath.IsAbs(p) {
				fmt.Fprintf(os.Stderr, "go: GOPATH entry is relative; must be absolute path: %q.\nFor more details see: 'go help gopath'\n", p)
				os.Exit(2)
			}
		}
	}

	if fi, err := os.Stat(cfg.GOROOT); err != nil || !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "go: cannot find GOROOT directory: %v\n", cfg.GOROOT)
		os.Exit(2)
	}

	// Set environment (GOOS, GOARCH, etc) explicitly.
	// In theory all the commands we invoke should have
	// the same default computation of these as we do,
	// but in practice there might be skew
	// This makes sure we all agree.
	cfg.OrigEnv = os.Environ()
	cfg.CmdEnv = envcmd.MkEnv()
	for _, env := range cfg.CmdEnv {
		if os.Getenv(env.Name) != env.Value {
			os.Setenv(env.Name, env.Value)
		}
	}

	for _, cmd := range base.Commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args, cwd)
			base.Exit()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "go: unknown subcommand %q\nRun 'go help' for usage.\n", args[0])
	base.SetExitStatus(2)
	base.Exit()
}

func init() {
	base.Usage = mainUsage
}

func mainUsage() {
	// special case "go test -h"
	if len(os.Args) > 1 && os.Args[1] == "test" {
		test.Usage()
	}
	help.PrintUsage(os.Stderr)
	os.Exit(2)
}
