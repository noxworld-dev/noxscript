package eud

import (
	ns3 "github.com/noxworld-dev/noxscript/ns/v3"
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

func CheckGameRusLanguage() bool {
	panic("not implemented")
}

func CheckGameKorLanguage() bool {
	panic("not implemented")
}

// UniPrint displays a string on the screen of the player. It does not localize the string.
func UniPrint(unit ns3.ObjectID, msg string) {
	ns4.AsObj(unit).Player().PrintStr(msg)
}

// UniChatMessage displays a string in a speech bubble for a given duration (in frames).
// It does not localize the string.
func UniChatMessage(unit ns3.ObjectID, msg string, duration int) {
	ns4.AsObj(unit).ChatStrTimer(msg, ns4.Frames(duration))
}

// UniPrintToAll displays a string to everyone. It does not localize the string.
func UniPrintToAll(msg string) {
	ns4.PrintStrToAll(msg)
}
