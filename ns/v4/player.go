package ns

import "github.com/noxworld-dev/opennox-lib/player"

// GetHost gets host's player object.
func GetHost() Obj {
	if impl == nil {
		return nil
	}
	return impl.GetHost()
}

// GetCharacterData gets information about the loaded character.
func GetCharacterData(field int) int {
	if impl == nil {
		return 0
	}
	return impl.GetCharacterData(field)
}

// StringID is an ID of a localized string in the string database.
type StringID = string

// Print displays a localized string on the screen of Other.
//
// If the string is not in the string database, it will instead print an error message with "MISSING:".
func Print(message StringID) {
	if impl == nil {
		return
	}
	impl.Print(message)
}

// PrintToAll displays a localized string to everyone.
//
// If the string is not in the string database, it will instead print an error message with "MISSING:".
func PrintToAll(message StringID) {
	if impl == nil {
		return
	}
	impl.PrintToAll(message)
}

// ClearMessages clears messages on player's screen.
func ClearMessages(player Obj) {
	if impl == nil {
		return
	}
	impl.ClearMessages(player)
}

// UnBlind the host.
func UnBlind() {
	if impl == nil {
		return
	}
	impl.UnBlind()
}

// Blind the host.
func Blind() {
	if impl == nil {
		return
	}
	impl.Blind()
}

// WideScreen enables or disables cinematic wide-screen effect.
func WideScreen(enable bool) {
	if impl == nil {
		return
	}
	impl.WideScreen(enable)
}

// IsTalking gets whether host is talking.
func IsTalking() bool {
	if impl == nil {
		return false
	}
	return impl.IsTalking()
}

// IsTrading returns whether the host is currently talking to shopkeeper.
func IsTrading() bool {
	if impl == nil {
		return false
	}
	return impl.IsTrading()
}

type HalberdLevel int

const (
	OblivionHalberd   = HalberdLevel(0)
	OblivionHeart     = HalberdLevel(1)
	OblivionWierdling = HalberdLevel(2)
	OblivionOrb       = HalberdLevel(3)
)

// SetHalberd upgrades host's oblivion staff.
func SetHalberd(upgrade HalberdLevel) {
	if impl == nil {
		return
	}
	impl.SetHalberd(upgrade)
}

// ImmediateBlind immediately blinds the host.
func ImmediateBlind() {
	if impl == nil {
		return
	}
	impl.ImmediateBlind()
}

// EndGame end of game for a specific class.
func EndGame(class player.Class) {
	if impl == nil {
		return
	}
	impl.EndGame(class)
}
