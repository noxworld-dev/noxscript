package ns

// ChatFunc is a callback type for OnChat.
// It returns a modified message. If message is empty, it will be ignored (muted).
//
// A nil player will be passed to the callback for non-player messages (e.g. ones sent via Obj.Chat).
//
// For team chat a team will be passed to the callback. It will be nil for regular chat.
type ChatFunc func(t Team, p Player, obj Obj, msg string) string

// OnChat registers a callback function that will be called when each chat message is received.
// See ChatFunc for details.
func OnChat(fnc ChatFunc) {
	if impl == nil {
		return
	}
	impl.OnChat(fnc)
}
