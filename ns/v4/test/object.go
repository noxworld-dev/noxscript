package nstest

import (
	"image/color"
	"sync/atomic"
	"time"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/class"
	"github.com/noxworld-dev/noxscript/ns/v4/damage"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
	"github.com/noxworld-dev/noxscript/ns/v4/subclass"
)

type Objects struct {
	last   uint32
	ByID   map[uint32]*Object
	ByName map[string]*Object
	Host   *Object
}

func (r *Runtime) GetHost() ns.Obj {
	return r.Objects.Host.asObj()
}

func (r *Runtime) ObjectByHandle(h ns.ObjHandle) ns.Obj {
	if h == nil {
		return nil
	}
	if obj, ok := h.(ns.Obj); ok {
		return obj
	}
	obj := r.Objects.ByID[uint32(h.ObjScriptID())]
	return obj.asObj()
}

func (r *Runtime) Object(name string) ns.Obj {
	return r.Objects.ByName[name].asObj()
}

func (r *Runtime) FindObjects(fnc func(it ns.Obj) bool, conditions ...ns.ObjCond) int {
	var list ns.Objects
	for _, obj := range r.Objects.ByID {
		list = append(list, obj)
	}
	return list.FindObjects(fnc, conditions...)
}

func (r *Runtime) CreateObject(typ string, pos ns.Positioner) ns.Obj {
	t := r.ObjectType(typ)
	if t == nil {
		return nil
	}
	return t.Create(pos)
}

func (s *Objects) New(r *Runtime, typ *ObjectType, name string) *Object {
	id := atomic.AddUint32(&s.last, 1)
	obj := &Object{
		r:     r,
		id:    id,
		name:  name,
		typ:   typ,
		class: typ.class,
		flags: object.FlagEnabled,
	}
	if s.ByID == nil {
		s.ByID = make(map[uint32]*Object)
	}
	s.ByID[obj.id] = obj
	if name != "" {
		if s.ByName == nil {
			s.ByName = make(map[string]*Object)
		}
		s.ByName[name] = obj
	}
	return obj
}

func (r *Runtime) NewObject(typ *ObjectType, name string) *Object {
	return r.Objects.New(r, typ, name)
}

type Object struct {
	r         *Runtime
	id        uint32
	name      string
	typ       *ObjectType
	class     object.Class
	flags     object.Flags
	PosVec    types.Pointf
	VelVec    types.Pointf
	ZVal      float32
	PlayerPtr *Player
}

func (obj *Object) asObj() ns.Obj {
	if obj == nil {
		return nil
	}
	return obj
}

func (obj *Object) ScriptID() int {
	return int(obj.id)
}

func (obj *Object) ObjScriptID() int {
	return int(obj.id)
}

func (obj *Object) Type() ns.ObjType {
	return obj.typ.asObjType()
}

func (obj *Object) Pos() ns.Pointf {
	return obj.PosVec
}

func (obj *Object) SetPos(p ns.Pointf) {
	obj.PosVec = p
}

func (obj *Object) Z() float32 {
	return obj.ZVal
}

func (obj *Object) SetZ(z float32) {
	obj.ZVal = z
}

func (obj *Object) Vel() ns.Pointf {
	return obj.VelVec
}

func (obj *Object) IsEnabled() bool {
	return obj.flags.Has(object.FlagEnabled)
}

func (obj *Object) Enable(enable bool) {
	if enable {
		obj.FlagsEnable(object.FlagEnabled)
	} else {
		obj.FlagsDisable(object.FlagEnabled)
	}
}

func (obj *Object) Toggle() bool {
	v := !obj.IsEnabled()
	obj.Enable(v)
	return v
}

func (obj *Object) IsLocked() bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Lock(lock bool) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Class() object.Class {
	return obj.class
}

func (obj *Object) HasClass(class class.Class) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) HasSubclass(subclass subclass.SubClass) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Flags() object.Flags {
	return obj.flags
}

func (obj *Object) SetFlags(v object.Flags) {
	obj.flags = v
}

func (obj *Object) FlagsEnable(v object.Flags) {
	obj.flags |= v
}

func (obj *Object) FlagsDisable(v object.Flags) {
	obj.flags &^= v
}

func (obj *Object) HasTeam(t ns.Team) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Team() ns.Team {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetTeam(t ns.Team) {
	//TODO implement me
}

func (obj *Object) MonsterStatus() object.MonsterStatus {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetMonsterStatus(v object.MonsterStatus) {
	//TODO implement me
}

func (obj *Object) MonsterStatusEnable(v object.MonsterStatus) {
	//TODO implement me
}

func (obj *Object) MonsterStatusDisable(v object.MonsterStatus) {
	//TODO implement me
}

func (obj *Object) HasEnchant(enchant enchant.Enchant) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Direction() ns.Direction {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) CursorPos() types.Pointf {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) CursorObj() ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) CurrentHealth() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) MaxHealth() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) RestoreHealth(amount int) {
	//TODO implement me
}

func (obj *Object) SetHealth(v int) {
	//TODO implement me
}

func (obj *Object) SetMaxHealth(v int) {
	//TODO implement me
}

func (obj *Object) CurrentMana() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) MaxMana() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetMana(v int) {
	//TODO implement me
}

func (obj *Object) SetMaxMana(v int) {
	//TODO implement me
}

func (obj *Object) SetHealthRegenToMaxDur(t time.Duration) {
	//TODO implement me
}

func (obj *Object) SetHealthRegenPerFrame(v float32) {
	//TODO implement me
}

func (obj *Object) Mass() float32 {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetMass(v float32) {
	//TODO implement me
}

func (obj *Object) CurrentSpeed() float32 {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) BaseSpeed() float32 {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetBaseSpeed(v float32) {
	//TODO implement me
}

func (obj *Object) Strength() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetStrength(v int) {
	//TODO implement me
}

func (obj *Object) SetColor(ind int, cl color.Color) {
	//TODO implement me
}

func (obj *Object) GetGold() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) ChangeGold(delta int) {
	//TODO implement me
}

func (obj *Object) GiveXp(xp float32) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) GetLevel() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) GetClass() player.Class {
	return obj.PlayerPtr.ClassVal
}

func (obj *Object) Player() ns.Player {
	return obj.PlayerPtr.asPlayer()
}

func (obj *Object) GetScore() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) ChangeScore(score int) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) HasOwner(owner ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) HasOwnerIn(owners ns.ObjGroup) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) SetOwner(owner ns.Obj) {
	//TODO implement me
}

func (obj *Object) SetOwners(owners ns.ObjGroup) {
	//TODO implement me
}

func (obj *Object) Freeze(freeze bool) {
	//TODO implement me
}

func (obj *Object) Pause(dt ns.Duration) {
	//TODO implement me
}

func (obj *Object) Move(wp ns.WaypointObj) {
	//TODO implement me
}

func (obj *Object) WalkTo(p ns.Pointf) {
	//TODO implement me
}

func (obj *Object) LookAtDirection(dir ns.Direction) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) LookWithAngle(angle int) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) LookAtObject(target ns.Positioner) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) CanSee(obj2 ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) ApplyForce(force ns.Pointf) {
	//TODO implement me
	obj.VelVec = obj.VelVec.Add(force)
}

func (obj *Object) PushTo(pos ns.Positioner, force float32) {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Damage(source ns.Obj, amount int, typ damage.Type) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Delete() {
	_, ok := obj.r.Objects.ByID[obj.id]
	if !ok {
		panic("trying to delete an already deleted object")
	}
	delete(obj.r.Objects.ByID, obj.id)
	if obj.name != "" {
		delete(obj.r.Objects.ByName, obj.name)
	}
}

func (obj *Object) DeleteAfter(dt ns.Duration) {
	obj.r.NewTimer(dt, func() {
		obj.Delete()
	})
}

func (obj *Object) Idle() {
	//TODO implement me
}

func (obj *Object) Wander() {
	//TODO implement me
}

func (obj *Object) Hunt() {
	//TODO implement me
}

func (obj *Object) Return() {
	//TODO implement me
}

func (obj *Object) Follow(target ns.Positioner) {
	//TODO implement me
}

func (obj *Object) Guard(p1, p2 ns.Pointf, distance float32) {
	//TODO implement me
}

func (obj *Object) Attack(target ns.Positioner) {
	//TODO implement me
}

func (obj *Object) IsAttackedBy(by ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) HitMelee(p ns.Pointf) {
	//TODO implement me
}

func (obj *Object) HitRanged(p ns.Pointf) {
	//TODO implement me
}

func (obj *Object) Flee(target ns.Positioner, dt ns.Duration) {
	//TODO implement me
}

func (obj *Object) HasItem(item ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) HasEquipment(item ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) FindItems(fnc func(obj ns.Obj) bool, conditions ...ns.ObjCond) int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) GetLastItem() ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) GetPreviousItem() ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Items(conditions ...ns.ObjCond) []ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Equipment(conditions ...ns.ObjCond) []ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) InItems() ns.ObjSearcher {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) InEquipment() ns.ObjSearcher {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) GetHolder() ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Pickup(item ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Drop(item ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Equip(item ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Unequip(item ns.Obj) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) ZombieStayDown() {
	//TODO implement me
}

func (obj *Object) RaiseZombie() {
	//TODO implement me
}

func (obj *Object) Chat(message ns.StringID) {
	//TODO implement me
}

func (obj *Object) ChatTimer(message ns.StringID, dt ns.Duration) {
	//TODO implement me
}

func (obj *Object) ChatStr(message string) {
	//TODO implement me
}

func (obj *Object) ChatStrTimer(message string, dt ns.Duration) {
	//TODO implement me
}

func (obj *Object) DestroyChat() {
	//TODO implement me
}

func (obj *Object) CreateMover(wp ns.WaypointObj, speed float32) ns.Obj {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) GetElevatorStatus() int {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) AggressionLevel(level float32) {
	//TODO implement me
}

func (obj *Object) SetRoamFlag(flags int) {
	//TODO implement me
}

func (obj *Object) RetreatLevel(percent float32) {
	//TODO implement me
}

func (obj *Object) ResumeLevel(percent float32) {
	//TODO implement me
}

func (obj *Object) AwardSpell(spell spell.Spell) bool {
	//TODO implement me
	panic("implement me")
}

func (obj *Object) Enchant(enchant enchant.Enchant, dt ns.Duration) {
	//TODO implement me
}

func (obj *Object) EnchantOff(enchant enchant.Enchant) {
	//TODO implement me
}

func (obj *Object) TrapSpells(spells ...spell.Spell) {
	//TODO implement me
}

func (obj *Object) TrapSpellsAdv(spells []ns.TrapSpell) {
	//TODO implement me
}

func (obj *Object) OnEvent(event ns.ObjectEvent, fnc ns.Func) {
	//TODO implement me
}
