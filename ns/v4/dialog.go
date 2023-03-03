package ns

import "github.com/noxworld-dev/noxscript/ns/v4/audio"

// SetShopkeeperText sets shopkeeper text.
func SetShopkeeperText(obj Obj, text StringID) {
	if impl == nil {
		return
	}
	impl.SetShopkeeperText(obj, text)
}

type DialogType string

const (
	DialogNormal    = DialogType("NORMAL")
	DialogNext      = DialogType("NEXT")
	DialogYesNo     = DialogType("YESNO")
	DialogYesNoNext = DialogType("YESNONEXT")
	DialogFalse     = DialogType("FALSE")
)

// SetDialog sets up a conversation with object.
//
// The type of conversation is one of: DialogNormal, DialogNext, DialogYesNo, DialogYesNoNext.
// The start and end are script functions that are called at the start and end of the conversation.
//
// If using a DialogYesNo conversation, the end script function should use GetAnswer to retrieve the result.
func SetDialog(obj Obj, typ DialogType, start Func, end Func) {
	if impl == nil {
		return
	}
	impl.SetDialog(obj, typ, start, end)
}

// CancelDialog cancels a conversation with object.
func CancelDialog(obj Obj) {
	if impl == nil {
		return
	}
	impl.CancelDialog(obj)
}

// StoryPic assigns a picture to a conversation.
func StoryPic(obj Obj, name string) {
	if impl == nil {
		return
	}
	impl.StoryPic(obj, name)
}

// TellStory causes the telling of a story.
//
// This will cause a story to be told. It relies on Self and Other to be
// particular values, which limits this to being used in the SetDialog callbacks.
//
// Example:
//
//	TellStory(audio.SwordsmanHurt, "Con05:OgreTalk07")
func TellStory(audio audio.Name, story StringID) {
	if impl == nil {
		return
	}
	impl.TellStory(audio, story)
}

// StartDialog starts a conversation between two objects.
//
// This requires that SetDialog has already been used to set up the conversation on the object.
func StartDialog(obj Obj, other Obj) {
	if impl == nil {
		return
	}
	impl.StartDialog(obj, other)
}

type DialogAnswer int

const (
	AnswerGoodbye = DialogAnswer(0)
	AnswerYes     = DialogAnswer(1)
	AnswerNo      = DialogAnswer(2)
)

// GetAnswer gets answer from conversation.
func GetAnswer(obj Obj) DialogAnswer {
	if impl == nil {
		return 0
	}
	return impl.GetAnswer(obj)
}
