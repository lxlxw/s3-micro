package main

import (
	"runtime"

	"github.com/lxlxw/s3-micro/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Execute()
}
