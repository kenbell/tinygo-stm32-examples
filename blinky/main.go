package main

import (
	"machine"
	"time"

	"github.com/kenbell/tinygo-stm32-examples/common"
	"github.com/kenbell/tinygo-stm32/boards"
)

func main() {
	board := boards.NewBasicBoard(&common.BoardCfg)
	machine.InitializeBoard(board)

	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		time.Sleep(500 * time.Millisecond)
		machine.LED.High()
		time.Sleep(500 * time.Millisecond)
		machine.LED.Low()
	}
}
