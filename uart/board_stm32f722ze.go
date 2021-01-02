// +build stm32f722

package main

import (
	"machine"

	"github.com/kenbell/tinygo-stm32/boards"
	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/pll"
	"github.com/kenbell/tinygo-stm32/timer"
	"github.com/kenbell/tinygo-stm32/uart"
)

var boardCfg = boards.BasicBoardConfig{
	OscillatorConfig: &clock.OscillatorConfig{
		HSE: &hse.Config{State: hse.StateBypass},
		PLL: &pll.Config{
			Source: pll.SourceHSE,
			State:  pll.StateOn,
			M:      4,
			N:      216,
			P:      2,
			Q:      9,
		},
	},
	ClockConfig: &clock.Config{
		SYSCLKSource:   clock.SYSCLKSourcePLL,
		AHBCLKDivider:  clock.HPREDividerDiv1,
		APB1CLKDivider: clock.PPREDividerDiv4,
		APB2CLKDivider: clock.PPREDividerDiv2,
	},
	SleepTimer: timer.TIM3,
	TickTimer:  timer.TIM5,
	UART:       uart.USART3,
	UARTConfig: &uart.Config{
		BaudRate: 115200,
		Clock:    clock.SourcePCLK1,
	},
	LED: machine.PB0,
}
