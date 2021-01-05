// +build stm32l552

package common

import (
	"machine"

	"github.com/kenbell/tinygo-stm32/boards"
	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/clock/hsi"
	"github.com/kenbell/tinygo-stm32/clock/hsi48"
	"github.com/kenbell/tinygo-stm32/clock/lse"
	"github.com/kenbell/tinygo-stm32/clock/msi"
	"github.com/kenbell/tinygo-stm32/clock/pll"
	"github.com/kenbell/tinygo-stm32/timer"
	"github.com/kenbell/tinygo-stm32/uart"
)

var BoardCfg = boards.BasicBoardConfig{
	OscillatorConfig: &clock.OscillatorConfig{
		LSE: &lse.Config{
			State:     lse.StateOn,
			Frequency: 32768}, // 32.768 KHz
		HSI: &hsi.Config{
			State:            hsi.StateOn,
			CalibrationValue: 0x40}, // default
		HSI48: &hsi48.Config{
			State: hsi48.StateOn},
		MSI: &msi.Config{
			State:            msi.StateOn,
			CalibrationValue: 0x0, // default
			ClockRange:       6},  // 4 MHz
		PLL: &pll.Config{
			Source: pll.SourceMSI,
			State:  pll.StateOn,
			M:      1,
			N:      55,
			P:      7,
			Q:      2,
			R:      2,
		},
	},
	ClockConfig: &clock.Config{
		SYSCLKSource:    clock.SYSCLKSourcePLL,
		AHBCLKDivider:   clock.HPREDividerDiv1,
		APB1CLKDivider:  clock.PPREDividerDiv1,
		APB2CLKDivider:  clock.PPREDividerDiv1,
		FlashWaitStates: 5, // @100 MHz, 5 wait states required
	},
	SleepTimer: timer.TIM15,
	TickTimer:  timer.TIM16,
	UART:       uart.LPUART1,
	UARTConfig: &uart.Config{
		BaudRate: 115200,
		Clock:    clock.SourcePCLK1,
	},
	LED: machine.PC7,
}
