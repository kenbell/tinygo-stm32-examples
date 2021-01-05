// +build stm32f10x

package common

import (
	"machine"

	"github.com/kenbell/tinygo-stm32/boards"
	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/pll"
	"github.com/kenbell/tinygo-stm32/timer"
	"github.com/kenbell/tinygo-stm32/uart"
)

var BoardCfg = boards.BasicBoardConfig{
	OscillatorConfig: &clock.OscillatorConfig{
		HSE: &hse.Config{State: hse.StateBypass},
		PLL: &pll.Config{
			Source: pll.SourceHSE,
			State:  pll.StateOn,
			Prediv: pll.HSEPreDiv1,
			Mul:    9,
		},
	},
	ClockConfig: &clock.Config{
		SYSCLKSource:    clock.SYSCLKSourcePLL,
		AHBCLKDivider:   clock.HPREDividerDiv1,
		APB1CLKDivider:  clock.PPREDividerDiv2,
		APB2CLKDivider:  clock.PPREDividerDiv1,
		FlashWaitStates: 2, // @ 72 MHz 2 wait states is required
	},
	SleepTimer: timer.TIM3,
	TickTimer:  timer.TIM1,
	UART:       uart.USART1,
	UARTConfig: &uart.Config{
		BaudRate: 115200,
		Clock:    clock.SourcePCLK2,
	},
	LED: machine.PC13,
}
