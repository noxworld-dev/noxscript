//go:build !script

package ns

import (
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/script"

	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
)

var impl Implementation

// Runtime returns implementation of all the accessible functions as an interface.
func Runtime() Implementation {
	return impl
}

// Game is an optional interface for script.Game that exposes NoxScript runtime.
type Game interface {
	script.Game
	// NoxScript returns implementation of NoxScript runtime.
	NoxScript() Implementation
}

// Implementation of the script runtime.
type Implementation interface {
	TimerByHandle(h TimerHandle) Timer
	NewTimer(dt script.Duration, fnc Func, args ...any) Timer
	RandomFloat(min float32, max float32) float32
	Random(min int, max int) int
	StopScript(value any)
	AutoSave()
	StartupScreen(which int)
	DeathScreen(which int)

	ObjectByHandle(h ObjHandle) Obj
	Object(name string) Obj
	ObjectGroupByHandle(h ObjGroupHandle) ObjGroup
	ObjectGroup(name string) ObjGroup
	CreateObject(typ string, pos Positioner) Obj
	GetTrigger() Obj
	GetCaller() Obj
	GetHost() Obj
	IsTrigger(obj Obj) bool
	IsCaller(obj Obj) bool
	IsGameBall(obj Obj) bool
	IsCrown(obj Obj) bool
	IsSummoned(obj Obj) bool
	MakeFriendly(obj Obj)
	MakeEnemy(obj Obj)
	BecomePet(obj Obj)
	BecomeEnemy(obj Obj)

	GetCharacterData(field int) int
	Print(message StringID)
	PrintToAll(message StringID)
	ClearMessages(player Obj)
	UnBlind()
	Blind()
	ImmediateBlind()
	WideScreen(enable bool)
	IsTalking() bool
	IsTrading() bool
	SetHalberd(upgrade HalberdLevel)
	EndGame(class player.Class)

	DestroyEveryChat()
	SetShopkeeperText(obj Obj, text StringID)
	SetDialog(obj Obj, typ DialogType, start Func, end Func)
	CancelDialog(obj Obj)
	StoryPic(obj Obj, name string)
	TellStory(audio audio.Name, story StringID)
	StartDialog(obj Obj, other Obj)
	GetAnswer(obj Obj) DialogAnswer

	AudioEvent(audio audio.Name, p Positioner)
	Music(music int, volume int)
	MusicPushEvent()
	MusicPopEvent()
	MusicEvent()

	Effect(effect effect.Effect, p1, p2 Positioner)
	CastSpell(spell spell.Spell, source, target Positioner)

	GetQuestStatus(name string) int
	GetQuestStatusFloat(name string) float32
	SetQuestStatus(status int, name string)
	SetQuestStatusFloat(status float32, name string)
	ResetQuestStatus(name string)

	JournalEntry(obj Obj, message StringID, typ EntryType)
	JournalEdit(obj Obj, message StringID, typ EntryType)
	JournalDelete(obj Obj, message StringID)

	WaypointByHandle(h WaypointHandle) WaypointObj
	Waypoint(name string) WaypointObj
	WaypointGroupByHandle(h WaypointGroupHandle) WaypointGroupObj
	WaypointGroup(name string) WaypointGroupObj

	NoWallSound(noWallSound bool)
	WallByHandle(h WallHandle) WallObj
	Wall(x int, y int) WallObj
	WallGroupByHandle(h WallGroupHandle) WallGroupObj
	WallGroup(name string) WallGroupObj

	Unused1f(id int)
	Unused20(id int)
	Unused50()
	Unused58(arg1 int, arg2 int)
	Unused59(arg1 int, arg2 int)
	Unused5a(arg1 int, arg2 int)
	Unused5b(arg1 int, arg2 int)
	Unused5c(arg1 int, arg2 int)
	Unused5d(arg1 int, arg2 int)
	Unused5e(str string) int
	Unused74(arg1 int, arg2 int)
	Unknownb8(id int) bool
	Unknownb9(id int) bool
	Unknownc4()
}
