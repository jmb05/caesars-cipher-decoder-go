package styling

import "fmt"

const BLACK = "\033[0;33m"
const RED = "\033[0;31m"
const GREEN = "\033[0;32m"
const YELLOW = "\033[0;33m"
const BLUE = "\033[0;34m"
const PURPLE = "\033[0;35m"
const CYAN = "\033[0;36m"
const WHITE = "\033[0;37m"

const B_BLACK = "\033[1;33m"
const B_RED = "\033[1;31m"
const B_GREEN = "\033[1;32m"
const B_YELLOW = "\033[1;33m"
const B_BLUE = "\033[1;34m"
const B_PURPLE = "\033[1;35m"
const B_CYAN = "\033[1;36m"
const B_WHITE = "\033[1;37m"

const DIM_BLACK = "\033[2;33m"
const DIM_RED = "\033[2;31m"
const DIM_GREEN = "\033[2;32m"
const DIM_YELLOW = "\033[2;33m"
const DIM_BLUE = "\033[2;34m"
const DIM_PURPLE = "\033[2;35m"
const DIM_CYAN = "\033[2;36m"
const DIM_WHITE = "\033[2;37m"

const I_BLACK = "\033[3;33m"
const I_RED = "\033[3;31m"
const I_GREEN = "\033[3;32m"
const I_YELLOW = "\033[3;33m"
const I_BLUE = "\033[3;34m"
const I_PURPLE = "\033[3;35m"
const I_CYAN = "\033[3;36m"
const I_WHITE = "\033[3;37m"

const UL_BLACK = "\033[4;33m"
const UL_RED = "\033[4;31m"
const UL_GREEN = "\033[4;32m"
const UL_YELLOW = "\033[4;33m"
const UL_BLUE = "\033[4;34m"
const UL_PURPLE = "\033[4;35m"
const UL_CYAN = "\033[4;36m"
const UL_WHITE = "\033[4;37m"

const BL_BLACK = "\033[5;33m"
const BL_RED = "\033[5;31m"
const BL_GREEN = "\033[5;32m"
const BL_YELLOW = "\033[5;33m"
const BL_BLUE = "\033[5;34m"
const BL_PURPLE = "\033[5;35m"
const BL_CYAN = "\033[5;36m"
const BL_WHITE = "\033[5;37m"

const REV_BLACK = "\033[7;33m"
const REV_RED = "\033[7;31m"
const REV_GREEN = "\033[7;32m"
const REV_YELLOW = "\033[7;33m"
const REV_BLUE = "\033[7;34m"
const REV_PURPLE = "\033[7;35m"
const REV_CYAN = "\033[7;36m"
const REV_WHITE = "\033[7;37m"

const BG_BLACK = "\033[40m"
const BG_RED = "\033[41m"
const BG_GREEN = "\033[42m"
const BG_YELLOW = "\033[43m"
const BG_BLUE = "\033[44m"
const BG_PURPLE = "\033[45m"
const BG_CYAN = "\033[46m"
const BG_WHITE = "\033[47m"

const HI_BLACK = "\033[0;90m"
const HI_RED = "\033[0;91m"
const HI_GREEN = "\033[0;92m"
const HI_YELLOW = "\033[0;93m"
const HI_BLUE = "\033[0;94m"
const HI_PURPLE = "\033[0;95m"
const HI_CYAN = "\033[0;96m"
const HI_WHITE = "\033[0;97m"

const BHI_BLACK = "\033[1;90m"
const BHI_RED = "\033[1;91m"
const BHI_GREEN = "\033[1;92m"
const BHI_YELLOW = "\033[1;93m"
const BHI_BLUE = "\033[1;94m"
const BHI_PURPLE = "\033[1;95m"
const BHI_CYAN = "\033[1;96m"
const BHI_WHITE = "\033[1;97m"

const HIBG_BLACK = "\033[0;100m"
const HIBG_RED = "\033[0;101m"
const HIBG_GREEN = "\033[0;102m"
const HIBG_YELLOW = "\033[0;103m"
const HIBG_BLUE = "\033[0;104m"
const HIBG_PURPLE = "\033[0;105m"
const HIBG_CYAN = "\033[0;106m"
const HIBG_WHITE = "\033[0;107m"

const RESET = "\033[0m"

func Color(color string, s string) string {
	return create(0, ColorStringToInt(color)) + s + RESET
}

func ColorBold(color string, s string) string {
	return create(1, ColorStringToInt(color)) + s + RESET
}

func ColorDim(color string, s string) string {
	return create(2, ColorStringToInt(color)) + s + RESET
}

func ColorItalic(color string, s string) string {
	return create(3, ColorStringToInt(color)) + s + RESET
}

func ColorUnderline(color string, s string) string {
	return create(4, ColorStringToInt(color)) + s + RESET
}

func ColorBlinking(color string, s string) string {
	return create(5, ColorStringToInt(color)) + s + RESET
}

//Does not work everywhere
func ColorReverse(color string, s string) string {
	return create(7, ColorStringToInt(color)) + s + RESET
}

func ColorBackground(color string, s string) string {
	return create(0, ColorStringToInt(color) + 10) + s + RESET
}

func ColorHighIntensity(color string, s string) string {
	return create(0, ColorStringToInt(color) + 60) + s + RESET
}

func ColorBoldHighIntensity(color string, s string) string {
	return create(1, ColorStringToInt(color) + 60) + s + RESET
}

func ColorBackgroundHighIntensity(color string, s string) string {
	return create(0, ColorStringToInt(color) + 70) + s + RESET
}

func create(style uint8, color uint8) string {
	return fmt.Sprintf("\033[%d;%dm",style, color)
}

func ColorStringToInt(color string) uint8 {
	if color == "black" {
		return 30
	} else if color == "red" {
		return 31
	} else if color == "green" {
		return 32
	} else if color == "yellow" {
		return 33
	} else if color == "blue" {
		return 34
	} else if color == "purple" {
		return 35
	} else if color == "cyan" {
		return 36
	} else {
		return 37
	}
}
