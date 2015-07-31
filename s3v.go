package main

import (
	"runtime"

	"github.com/kjmkznr/s3v/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
