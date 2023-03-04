package ns

import (
	"github.com/noxworld-dev/opennox-lib/player"

	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

// GetHost gets host's player object.
//
// This will return the host's player object id. The host is the player with
// a player id of 31.
func GetHost() ObjectID {
	return asObj(ns4.GetHost())
}

// IsTalking gets whether host is talking.
func IsTalking() bool {
	return ns4.IsTalking()
}

// IsTrading gets whether the host is currently trading.
//
// This will return whether the host is currently talking to shopkeeper.
func IsTrading() bool {
	return ns4.IsTrading()
}

// Print displays a localized string.
//
// This will display on a localized on string the screen of OTHER. If the string
// is not in the database string, it will instead print an error message with
// "MISSING:".
func Print(message string) {
	ns4.Print(ns4.StringID(message))
}

// PrintToAll displays a localized to string everyone.
//
// This will display on a localized on string everyone's screen. If the string
// is not in the database string, it will instead print an error message with
// "MISSING:".
func PrintToAll(message string) {
	ns4.PrintToAll(ns4.StringID(message))
}

// UnBlind the host.
func UnBlind() {
	ns4.UnBlind()
}

// Blind the host.
func Blind() {
	ns4.Blind()
}

func WideScreen(value bool) {
	ns4.WideScreen(value)
}

type EntryType = int

// JournalEntry adds entry to player's journal.
//
// This will add an entry to a player object's journal. The can string be any
// string, but should be less than 64 characters. The is string displayed
// according to the specified type:
//
//	0 = Green text, no sound
//	1 = White text
//	2 = Red text with quest label
//	3 = Green text
//	4 = Grey text with completed label
//	8 = Yellow text with hint label
//
// If the player object id is 0, then it will add the journal entry to all players.
func JournalEntry(id ObjectID, message string, typ EntryType) {
	ns4.JournalEntry(ns4.AsObj(id), message, ns4.EntryType(typ))
}

// JournalDelete deletes entry from player's journal.
//
// This will delete the first entry from a player object's journal whose message
// matches the specified message. If the player object id is 0, then all players
// are affected.
func JournalDelete(id ObjectID, message string) {
	ns4.JournalDelete(ns4.AsObj(id), message)
}

// JournalEdit edits entry in player's journal.
//
// This will modify the first entry in a player object's journal whose message
// matches the specified message. If the player object id is 0, then all players
// are affected.
func JournalEdit(id ObjectID, message string, typ EntryType) {
	ns4.JournalEdit(ns4.AsObj(id), message, ns4.EntryType(typ))
}

// GetGold gets amount of gold for player object.
func GetGold(id ObjectID) int {
	return ns4.AsObj(id).GetGold()
}

// ChangeGold changes amount of gold for player object. Delta can be negative.
func ChangeGold(id ObjectID, delta int) {
	ns4.AsObj(id).ChangeGold(delta)
}

// GiveXp grants experience to a player.
func GiveXp(id ObjectID, xp float32) {
	ns4.AsObj(id).GiveXp(xp)
}

// AutoSave triggers an autosave. Only solo games.
func AutoSave() {
	ns4.AutoSave()
}

// StartupScreen shows startup screen to host.
func StartupScreen(which int) {
	ns4.StartupScreen(which)
}

// SetHalberd upgrades host's oblivion staff.
//
// This will upgrade the oblivion staff. The upgrades are:
//
//	0 = OblivionHalberd
//	1 = OblivionHeart
//	2 = OblivionWierdling
//	3 = OblivionOrb
func SetHalberd(upgrade int) {
	ns4.SetHalberd(ns4.HalberdLevel(upgrade))
}

// DeathScreen shows death screen. Screen to show is in range 1-11.
func DeathScreen(which int) {
	ns4.DeathScreen(which)
}

// ClearMessages clears messages on player's screen.
func ClearMessages(id ObjectID) {
	ns4.ClearMessages(ns4.AsObj(id))
}

// EndGame ends of game.
//
// Type is one of:
//
//	0=Warrior
//	1=Wizard
//	2=Conjurer
func EndGame(typ int) {
	ns4.EndGame(player.Class(typ))
}

// ImmediateBlind immediately blinds the host.
func ImmediateBlind() {
	ns4.ImmediateBlind()
}

// GetScore gets player object's score.
func GetScore(id ObjectID) int {
	return ns4.AsObj(id).GetScore()
}

// ChangeScore changes player object's score.
func ChangeScore(id ObjectID, score int) {
	ns4.AsObj(id).ChangeScore(score)
}

// SetQuestStatus sets quest status (int).
func SetQuestStatus(status int, name string) {
	ns4.SetQuestStatus(status, name)
}

// SetQuestStatusFloat sets quest status (float).
func SetQuestStatusFloat(status float32, name string) {
	ns4.SetQuestStatusFloat(status, name)
}

// GetQuestStatus gets quest status (int).
func GetQuestStatus(name string) int {
	return ns4.GetQuestStatus(name)
}

// GetQuestStatusFloat get quest status (int).
func GetQuestStatusFloat(name string) float32 {
	return ns4.GetQuestStatusFloat(name)
}

// ResetQuestStatus deletes quest status.
//
// This will delete a quest status. The name can be a wildcard with an
// asterisk or with a map name.
//
// TODO: what is the wildcard syntax?
func ResetQuestStatus(name string) {
	ns4.ResetQuestStatus(name)
}

// GetCharacterData gets character data.
//
// Get information about the loaded character. Index is in range 0-5.
//
// Deprecated: return data is likely useless.
func GetCharacterData(idx int) int {
	return ns4.GetCharacterData(idx)
}
