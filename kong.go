package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"path"
	"path/filepath"
)

type Context struct {
	Debug bool
}

//rm --prune --pull --config="key1=valu1;key2=value2" a.txt b.txt
type RmCmd struct {
	Force     bool              `help:"Force removal."`
	Recursive bool              `help:"Recursively remove files."`
	Pull      bool              `help:"Pull Image" `
	Prune     bool              `help:"Prune" `
	Config    map[string]string ``
	Paths     []string          `arg:"" name:"path" help:"Paths to remove." type:"path"`
}

func (r *RmCmd) Run(ctx *Context) error {
	fmt.Println("rm", r.Pull)
	fmt.Println("rm", r.Prune)
	fmt.Println("rm", r.Paths)
	fmt.Println("rm", r.Config)
	env := make([]string, 0)
	if r.Config != nil {
		for k, v := range r.Config {
			env = append(env, k+"="+v)
		}
	}
	str := "abcde=1;abc=3;cde=3;"
	str = str[0 : len(str)-1]
	fmt.Println(str)
	return nil
}

type LsCmd struct {
	Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
}

func (l *LsCmd) Run(ctx *Context) error {
	fmt.Println("ls", l.Paths)
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Rm RmCmd `cmd:"" help:"Remove files."`
	Ls LsCmd `cmd:"" help:"List paths."`
}

func show() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)

	root := "/"
	fmt.Println(path.Join(root, "good.exe"))
	fmt.Println(filepath.Join("/root/var", "stacks", "123"))

}
