package ns

// GetQuestStatus gets quest status (int).
func GetQuestStatus(name string) int {
	if impl == nil {
		return 0
	}
	return impl.GetQuestStatus(name)
}

// GetQuestStatusFloat gets quest status (float).
func GetQuestStatusFloat(name string) float32 {
	if impl == nil {
		return 0
	}
	return impl.GetQuestStatusFloat(name)
}

// SetQuestStatus sets quest status (int).
func SetQuestStatus(status int, name string) {
	if impl == nil {
		return
	}
	impl.SetQuestStatus(status, name)
}

// SetQuestStatusFloat sets quest status (float).
func SetQuestStatusFloat(status float32, name string) {
	if impl == nil {
		return
	}
	impl.SetQuestStatusFloat(status, name)
}

// ResetQuestStatus deletes a quest status.
//
// The name can be a wildcard with an asterisk or with a map name.
func ResetQuestStatus(name string) {
	if impl == nil {
		return
	}
	impl.ResetQuestStatus(name)
}

type EntryType int

const (
	EntryGreen1 = EntryType(0) // green text, no sound
	EntryWhite  = EntryType(1) // white text
	EntryRed    = EntryType(2) // red text with quest label
	EntryGreen2 = EntryType(3) // green text
	EntryGrey   = EntryType(4) // grey text with completed label
	EntryYellow = EntryType(8) // yellow text with hint label
)

// JournalEntry adds entry to player's journal.
//
// If the player object is nil, then it will add the journal entry to all players.
func JournalEntry(obj Obj, message StringID, typ EntryType) {
	if impl == nil {
		return
	}
	impl.JournalEntry(obj, message, typ)
}

// JournalEdit edits entry in player's journal.
//
// If the player object is nil, then it will edit the journal entry for all players.
func JournalEdit(obj Obj, message StringID, typ EntryType) {
	if impl == nil {
		return
	}
	impl.JournalEdit(obj, message, typ)
}

// JournalDelete deletes entry from player's journal.
//
// If the player object is nil, then it will delete the journal entry for all players.
func JournalDelete(obj Obj, message StringID) {
	if impl == nil {
		return
	}
	impl.JournalDelete(obj, message)
}
