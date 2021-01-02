package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/kenbell/tinygo-stm32/boards"
)

func main() {

	board := boards.NewBasicBoard(&boardCfg)
	machine.InitializeBoard(board)

	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go func() {
		buf := []byte{0}

		for {
			n, _ := machine.UART0.Read(buf)
			if n != 0 {
				fmt.Printf("Got 0x%x '%v'\r\n", buf[0], string(buf[0]))
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	i := 0
	for {
		time.Sleep(500 * time.Millisecond)
		machine.LED.High()
		time.Sleep(500 * time.Millisecond)
		machine.LED.Low()
		fmt.Printf("Counter: %v\r\n", i)
		i++
	}
}
