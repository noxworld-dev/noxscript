package nstest

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
)

func init() {
	ns.SetRuntime(NewRuntime())
}

const (
	FrameRate = 30
	FrameTime = time.Second/FrameRate + 1
)

func NewRuntime() *Runtime {
	r := &Runtime{}
	r.Types.Player = r.NewObjectType(object.ClassPlayer, "NewPlayer")
	r.Types.NPC = r.NewObjectType(object.ClassMonster, "NPC")

	u := r.NewObject(r.Types.Player, "", ns.Ptf(1000, 1000))
	pl := r.NewPlayer("Jack")

	u.PlayerPtr, pl.UnitPtr = pl, u
	r.Objects.Host = u
	r.PlayersData.Host = pl
	return r
}

func GetRuntime() *Runtime {
	return ns.Runtime().(*Runtime)
}

type Runtime struct {
	frame    int
	initDone bool

	Timers      Timers
	Storage     Storage
	Objects     Objects
	Types       ObjectTypes
	PlayersData Players
	Trigger     *Object
	Caller      *Object

	onFrame []ns.FrameFunc
	onInit  []ns.MapEventFunc
}

func (r *Runtime) WithArgs(caller, trigger *Object) func() {
	r.Caller, r.Trigger = caller, trigger
	return func() {
		r.Caller, r.Trigger = nil, nil
	}
}

func (r *Runtime) Update() {
	if r.frame == 0 && !r.initDone {
		r.initDone = true
		for _, fnc := range r.onInit {
			fnc()
		}
	}
	r.frame++
	r.Timers.Update()
	r.Objects.Update()
	for _, fnc := range r.onFrame {
		fnc()
	}
}

func (r *Runtime) UpdateN(frames int) {
	for i := 0; i < frames; i++ {
		r.Update()
	}
}

func (r *Runtime) UpdateFor(dt time.Duration) {
	t0 := r.Time()
	for r.Time()-t0 < dt {
		r.Update()
	}
}

func (r *Runtime) Frame() int {
	return r.frame
}

func (r *Runtime) Time() time.Duration {
	secs := float64(r.Frame()) / FrameRate
	return time.Duration(float64(time.Second) * secs)
}

func (r *Runtime) FrameRate() int {
	return FrameRate
}

func (r *Runtime) ResetFrame() {
	r.frame = 0
}

func (r *Runtime) Store(typ ns.StorageType) ns.Storage {
	switch typ := typ.(type) {
	default:
		panic(fmt.Errorf("unexpected storage type: %T", typ))
	case nil:
		return r.Storage.Session("default")
	case ns.Session:
		return r.Storage.Session(typ.Name)
	case ns.Persistent:
		return r.Storage.Persistent(typ.Name)
	}
}

func (r *Runtime) TimerByHandle(h ns.TimerHandle) ns.Timer {
	if h == nil {
		return nil
	}
	if t, ok := h.(ns.Timer); ok {
		return t
	}
	t := r.Timers.byID[uint32(h.TimerScriptID())]
	if t == nil {
		return nil
	}
	return t
}

func (r *Runtime) NewTimer(dt ns.Duration, fnc ns.Func, args ...any) ns.Timer {
	var call func()
	switch fnc := fnc.(type) {
	case func():
		call = fnc
	default:
		panic(fmt.Errorf("unsupported func type: %T", fnc))
	}
	if v, ok := dt.Frames(); ok {
		return r.Timers.New(uint32(v), call)
	}
	if t, ok := dt.Time(); ok {
		return r.Timers.NewDur(t, call)
	}
	panic("invalid interval")
}

func (r *Runtime) RandomFloat(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func (r *Runtime) Random(min int, max int) int {
	return min + rand.Intn(max+1)
}

func (r *Runtime) StopScript(value any) {
	panic("stop")
}

func (r *Runtime) AutoSave() {
	//TODO implement me
}

func (r *Runtime) StartupScreen(which int) {
	//TODO implement me
}

func (r *Runtime) DeathScreen(which int) {
	//TODO implement me
}

func (r *Runtime) ObjectGroupByHandle(h ns.ObjGroupHandle) ns.ObjGroup {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) ObjectGroup(name string) ns.ObjGroup {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) GetTrigger() ns.Obj {
	return r.Trigger.asObj()
}

func (r *Runtime) GetCaller() ns.Obj {
	return r.Caller.asObj()
}

func (r *Runtime) IsTrigger(obj ns.Obj) bool {
	if obj == nil {
		return false
	}
	return r.Trigger.asObj() == obj
}

func (r *Runtime) IsCaller(obj ns.Obj) bool {
	if obj == nil {
		return false
	}
	return r.Caller.asObj() == obj
}

func (r *Runtime) IsGameBall(obj ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) IsCrown(obj ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) IsSummoned(obj ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) MakeFriendly(obj ns.Obj) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) MakeEnemy(obj ns.Obj) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) BecomePet(obj ns.Obj) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) BecomeEnemy(obj ns.Obj) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) Teams() []ns.Team {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) GetCharacterData(field int) int {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) Print(message ns.StringID) {
	slog.Info("print", "id", message)
}

func (r *Runtime) PrintStr(message string) {
	slog.Info(message)
}

func (r *Runtime) PrintToAll(message ns.StringID) {
	slog.Info("print", "id", message, "all", true)
}

func (r *Runtime) PrintStrToAll(message string) {
	slog.Info(message, "all", true)
}

func (r *Runtime) ClearMessages(player ns.Obj) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) UnBlind() {
	//TODO implement me
}

func (r *Runtime) Blind() {
	//TODO implement me
}

func (r *Runtime) ImmediateBlind() {
	//TODO implement me
}

func (r *Runtime) WideScreen(enable bool) {
	//TODO implement me
}

func (r *Runtime) IsTalking() bool {
	//TODO implement me
	return false
}

func (r *Runtime) IsTrading() bool {
	//TODO implement me
	return false
}

func (r *Runtime) SetHalberd(upgrade ns.HalberdLevel) {
	//TODO implement me
}

func (r *Runtime) EndGame(class player.Class) {
	//TODO implement me
}

func (r *Runtime) DestroyEveryChat() {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) SetShopkeeperText(obj ns.Obj, text ns.StringID) {
	//TODO implement me
}

func (r *Runtime) SetShopkeeperTextStr(obj ns.Obj, text string) {
	//TODO implement me
}

func (r *Runtime) SetDialog(obj ns.Obj, typ ns.DialogType, start ns.Func, end ns.Func) {
	//TODO implement me
}

func (r *Runtime) CancelDialog(obj ns.Obj) {
	//TODO implement me
}

func (r *Runtime) StoryPic(obj ns.Obj, name string) {
	//TODO implement me
}

func (r *Runtime) TellStory(audio audio.Name, story ns.StringID) {
	//TODO implement me
}

func (r *Runtime) TellStoryStr(audio audio.Name, story string) {
	//TODO implement me
}

func (r *Runtime) StartDialog(obj ns.Obj, other ns.Obj) {
	//TODO implement me
}

func (r *Runtime) GetAnswer(obj ns.Obj) ns.DialogAnswer {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) AudioEvent(audio audio.Name, p ns.Positioner) {
	//TODO implement me
}

func (r *Runtime) Music(music int, volume int) {
	//TODO implement me
}

func (r *Runtime) MusicPushEvent() {
	//TODO implement me
}

func (r *Runtime) MusicPopEvent() {
	//TODO implement me
}

func (r *Runtime) MusicEvent() {
	//TODO implement me
}

func (r *Runtime) Effect(effect effect.Effect, p1, p2 ns.Positioner) {
	//TODO implement me
}

func (r *Runtime) CastSpell(spell spell.Spell, source, target ns.Positioner) {
	//TODO implement me
}

func (r *Runtime) CastSpellLvl(spell spell.Spell, lvl int, source, target ns.Positioner) {
	//TODO implement me
}

func (r *Runtime) NewTrap(pos ns.Positioner, spells []ns.TrapSpell) ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) GetQuestStatus(name string) int {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) GetQuestStatusFloat(name string) float32 {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) SetQuestStatus(status int, name string) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) SetQuestStatusFloat(status float32, name string) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) ResetQuestStatus(name string) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) JournalEntry(obj ns.Obj, message ns.StringID, typ ns.EntryType) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) JournalEdit(obj ns.Obj, message ns.StringID, typ ns.EntryType) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) JournalDelete(obj ns.Obj, message ns.StringID) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) JournalEntryStr(obj ns.Obj, message string, typ ns.EntryType) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) JournalEditStr(obj ns.Obj, message string, typ ns.EntryType) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) JournalDeleteStr(obj ns.Obj, message string) {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) Waypoints() []ns.WaypointObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) WaypointByHandle(h ns.WaypointHandle) ns.WaypointObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) Waypoint(name string) ns.WaypointObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) NewWaypoint(name string, pos ns.Pointf) ns.WaypointObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) WaypointGroupByHandle(h ns.WaypointGroupHandle) ns.WaypointGroupObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) WaypointGroup(name string) ns.WaypointGroupObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) NoWallSound(noWallSound bool) {
	//TODO implement me
}

func (r *Runtime) WallByHandle(h ns.WallHandle) ns.WallObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) Wall(x int, y int) ns.WallObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) WallGroupByHandle(h ns.WallGroupHandle) ns.WallGroupObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) WallGroup(name string) ns.WallGroupObj {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) FindWalls(fnc func(it ns.WallObj) bool, conditions ...ns.WallCond) int {
	//TODO implement me
	panic("implement me")
}

func (r *Runtime) OnFrame(fnc ns.FrameFunc) {
	r.onFrame = append(r.onFrame, fnc)
}

func (r *Runtime) OnMapEvent(typ ns.MapEvent, fnc ns.MapEventFunc) {
	//TODO implement other events
	switch typ {
	case ns.MapInitialize:
		r.onInit = append(r.onInit, fnc)
	}
}

func (r *Runtime) OnChat(fnc ns.ChatFunc) {
	//TODO implement me
}
