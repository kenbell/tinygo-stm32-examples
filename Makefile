test:
	tinygo build -o ./uart/stm32l522generic.elf -size short -target stm32l552generic  ./uart
	tinygo build -o ./uart/stm32f103generic.elf -size short -target stm32f103generic  ./uart
	tinygo build -o ./uart/stm32f722generic.elf -size short -target stm32f722generic  ./uart
	tinygo build -o ./blinky/stm32l522generic.elf -size short -target stm32l552generic  ./blinky
	tinygo build -o ./blinky/stm32f103generic.elf -size short -target stm32f103generic  ./blinky
	tinygo build -o ./blinky/stm32f722generic.elf -size short -target stm32f722generic  ./blinky
