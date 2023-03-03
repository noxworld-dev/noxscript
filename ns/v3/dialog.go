package ns

import (
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
)

// SetShopkeeperText sets shopkeeper text.
func SetShopkeeperText(id ObjectID, text string) {
	ns4.SetShopkeeperText(ns4.AsObj(id), ns4.StringID(text))
}

type DialogType = string

const (
	NORMAL    = "NORMAL"
	NEXT      = "NEXT"
	YESNO     = "YESNO"
	YESNONEXT = "YESNONEXT"
)

// SetDialog sets up a conversation with object.
//
// This will cause a conversation with the object. The type of conversation is
// one of: NORMAL, NEXT, YESNO, YESNONEXT. The start and end are script
// functions that are called at the start and end of the conversation.
//
// If using a YESNO conversation, the end script function should use ::GetAnswer
// to retrieve the result.
func SetDialog(id ObjectID, typ DialogType, start Func, end Func) {
	ns4.SetDialog(ns4.AsObj(id), ns4.DialogType(typ), start, end)
}

// CancelDialog cancels a conversation with object.
func CancelDialog(id ObjectID) {
	ns4.CancelDialog(ns4.AsObj(id))
}

// StoryPic assigns a picture to a conversation.
func StoryPic(id ObjectID, name string) {
	ns4.StoryPic(ns4.AsObj(id), name)
}

// TellStory causes the telling of a story.
//
// This will cause a story to be told. It relies on SELF and OTHER to be
// particular values, which limits this to being used in the SetDialog
// callbacks.
//
// Example usage:
//
//	ns.TellStory("SwordsmanHurt", "Con05:OgreTalk07")
func TellStory(audioID AudioName, story string) {
	ns4.TellStory(audio.Name(audioID), ns4.StringID(story))
}

// StartDialog starts a conversation between two objects.
//
// This requires that SetDialog has already been used to setup the conversation
// on the NPC object.
func StartDialog(npc ObjectID, other ObjectID) {
	ns4.StartDialog(ns4.AsObj(npc), ns4.AsObj(other))
}

type Answer = int

const (
	Goodbye = 0
	Yes     = 1
	No      = 2
)

// GetAnswer gets answer from conversation.
//
// Result is one of:
//
//	0=Goodbye
//	1=Yes
//	2=No
func GetAnswer(id ObjectID) Answer {
	return Answer(ns4.GetAnswer(ns4.AsObj(id)))
}
