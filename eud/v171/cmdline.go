package eud

import (
	"github.com/noxworld-dev/opennox-lib/script"
)

const (
	CONSOLE_COLOR_BLACK       = 1
	CONSOLE_COLOR_GREY        = 2
	CONSOLE_COLOR_WHITE       = 3
	CONSOLE_COLOR_WHITE_LIGHT = 4
	CONSOLE_COLOR_BROWN       = 5
	CONSOLE_COLOR_RED         = 6
	CONSOLE_COLOR_PINK        = 7
	CONSOLE_COLOR_GREEN       = 8
	CONSOLE_COLOR_LIME        = 10
	CONSOLE_COLOR_DARKBLUE    = 11
	CONSOLE_COLOR_BLUE        = 12
	CONSOLE_COLOR_SKYBLUE     = 13
	CONSOLE_COLOR_ORANGE      = 14
	CONSOLE_COLOR_YELLOW      = 15
	CONSOLE_COLOR_PALE_YELLOW = 16
)

func CmdLine(commandMessage string, show bool) {
	panic("not implemented")
}

func Loadmap(mapFileName string) {
	panic("not implemented")
}

// NoxConsolePrint prints a message
func NoxConsolePrint(msg string, color int) {
	// TODO: expose colored console
	script.Runtime().Console(false).Print(msg)
}
