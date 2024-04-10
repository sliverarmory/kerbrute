package main

import "C"
import (
	"github.com/sliverarmory/kerbrute/cmd"
	"github.com/sliverarmory/kerbrute/pkg/parser"
	"github.com/sliverarmory/kerbrute/pkg/stdredir"
	"github.com/sliverarmory/kerbrute/util"
)

const (
	Success = 0
	Error   = 1
)

//export Run
func Run(data uintptr, dataLen uintptr, callback uintptr) uintptr {
	// Prepare the output buffer used to send data back to the implant
	outBuff := parser.NewOutBuffer(callback)

	// Simulate command-line argument passing
	err := parser.GetCommandLineArgs(data, dataLen)
	if err != nil {
		outBuff.SendError(err)
		outBuff.Flush()
		return Error
	}

	// Start capturing stdout and stderr
	stdredir.StartCapture(outBuff)

	main()

	// Stop capturing stdout and stderr
	stdredir.StopCapture()

	outBuff.Flush()
	return Success
}

func main() {
	util.PrintBanner()
	cmd.Execute()
}
