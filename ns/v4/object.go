package ns

import (
	"image/color"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/noxscript/ns/v4/class"
	"github.com/noxworld-dev/noxscript/ns/v4/damage"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
	"github.com/noxworld-dev/noxscript/ns/v4/subclass"
)

type Pointf = types.Pointf

func Ptf(x, y float32) Pointf {
	return types.Ptf(x, y)
}

// Positioner is an interface for objects that have position on the map.
type Positioner interface {
	// Pos returns current position of the object.
	Pos() Pointf
}

type Direction = int

const (
	NW Direction = 0
	N  Direction = 1
	NE Direction = 2
	W  Direction = 3
	E  Direction = 5
	SW Direction = 6
	S  Direction = 7
	SE Direction = 8
)

type ObjectEvent int

const (
	EventEnemySighted    = ObjectEvent(3)
	EventLookingForEnemy = ObjectEvent(4)
	EventDeath           = ObjectEvent(5)
	EventChangeFocus     = ObjectEvent(6)
	EventIsHit           = ObjectEvent(7)
	EventRetreat         = ObjectEvent(8)
	EventCollision       = ObjectEvent(9)
	EventEnemyHeard      = ObjectEvent(10)
	EventEndOfWaypoint   = ObjectEvent(11)
	EventLostEnemy       = ObjectEvent(13)
)

// ObjectType looks up an object type by name.
func ObjectType(name string) ObjType {
	if impl == nil {
		return nil
	}
	return impl.ObjectType(name)
}

// ObjectTypeByInd looks up an object type by index.
func ObjectTypeByInd(ind int) ObjType {
	if impl == nil {
		return nil
	}
	return impl.ObjectTypeByInd(ind)
}

// CreateObject creates an object given a type and a starting location.
//
//	Example:
//	  spider := CreateObject("SmallAlbinoSpider", Waypoint("SpiderHole"))
func CreateObject(typ string, pos Positioner) Obj {
	if impl == nil {
		return nil
	}
	return impl.CreateObject(typ, pos)
}

// Object looks up an object by name.
func Object(name string) Obj {
	if impl == nil {
		return nil
	}
	return impl.Object(name)
}

// ObjectGroup looks up object group by name.
func ObjectGroup(name string) ObjGroup {
	if impl == nil {
		return nil
	}
	return impl.ObjectGroup(name)
}

// GetTrigger returns object which triggered an event, if valid.
func GetTrigger() Obj {
	if impl == nil {
		return nil
	}
	return impl.GetTrigger()
}

// GetCaller returns object that was the target of an event, if valid.
func GetCaller() Obj {
	if impl == nil {
		return nil
	}
	return impl.GetCaller()
}

// IsTrigger checks whether object triggered an event.
func IsTrigger(obj Obj) bool {
	if impl == nil {
		return false
	}
	return impl.IsTrigger(obj)
}

// IsCaller checks whether object is a target of an event.
func IsCaller(obj Obj) bool {
	if impl == nil {
		return false
	}
	return impl.IsCaller(obj)
}

// IsGameBall gets whether object is a GameBall.
func IsGameBall(obj Obj) bool {
	if impl == nil {
		return false
	}
	return impl.IsGameBall(obj)
}

// IsCrown gets whether object is a Crown.
func IsCrown(obj Obj) bool {
	if impl == nil {
		return false
	}
	return impl.IsCrown(obj)
}

// IsSummoned gets whether object is a summoned creature.
func IsSummoned(obj Obj) bool {
	if impl == nil {
		return false
	}
	return impl.IsSummoned(obj)
}

type ObjType interface {
	// Name returns object type name.
	Name() string

	// Index returns an index of the type in the game's database.
	// It can be used to quickly compare object types, or as a map key for them.
	//
	// It is strongly advised against comparing it with hardcoded values, since mods may add/remove object types.
	// Use ObjectType to find object types instead.
	Index() int

	// Class returns object type class.
	Class() object.Class

	// HasClass checks whether object type has a class.
	// It uses string values instead of enum as Class does.
	HasClass(class class.Class) bool

	// HasSubclass tests whether an object has a specific subclass.
	// The subclass overlaps, so you should probably test for the class first (via HasClass).
	HasSubclass(subclass subclass.SubClass) bool

	// Flags returns object type flags.
	Flags() object.Flags

	// Create and object of this type.
	Create(pos Positioner) Obj
}

type Obj interface {
	ObjHandle
	// Type returns object type.
	Type() ObjType
	// Pos returns current position of the object.
	Pos() Pointf
	// SetPos instantly moves object to a given position.
	SetPos(p Pointf)
	// Z returns current Z offset for the object.
	Z() float32
	// SetZ sets Z offset for the object.
	SetZ(z float32)
	// IsEnabled checks if object is currently enabled.
	IsEnabled() bool
	// Enable or disable the object.
	Enable(enable bool)
	// Toggle the object's enabled state.
	// It returns the new state after toggling the object.
	Toggle() bool
	// IsLocked checks if object is currently locked.
	IsLocked() bool
	// Lock or unlock the object.
	Lock(lock bool)

	// Class returns object class.
	Class() object.Class

	// HasClass checks whether object has a class.
	// It uses string values instead of enum as Class does.
	HasClass(class class.Class) bool

	// HasSubclass tests whether an object has a specific subclass.
	// The subclass overlaps, so you should probably test for the class first (via HasClass).
	HasSubclass(subclass subclass.SubClass) bool

	// Flags returns object flags.
	Flags() object.Flags

	// SetFlags sets object flags to the given value.
	// For switching individual flags see FlagsEnable and FlagsDisable.
	SetFlags(v object.Flags)

	// FlagsEnable enables specific object flags.
	FlagsEnable(v object.Flags)

	// FlagsDisable disables specific object flags.
	FlagsDisable(v object.Flags)

	// HasTeam checks if an object belongs to a team.
	HasTeam(t Team) bool

	// Team returns current team of an object, if any.
	Team() Team

	// SetTeam sets team of an object.
	SetTeam(t Team)

	// MonsterStatus returns monster/NPC status flags.
	// It returns 0 if object's class is not a monster/NPC.
	MonsterStatus() object.MonsterStatus

	// SetMonsterStatus sets monster status flags to the given value.
	// For switching individual statuses see MonsterStatusEnable and MonsterStatusDisable.
	// It does nothing if object's class is not a monster/NPC.
	SetMonsterStatus(v object.MonsterStatus)

	// MonsterStatusEnable enables specific monster status flags.
	// It does nothing if object's class is not a monster/NPC.
	MonsterStatusEnable(v object.MonsterStatus)

	// MonsterStatusDisable disables specific monster status flags.
	// It does nothing if object's class is not a monster/NPC.
	MonsterStatusDisable(v object.MonsterStatus)

	// HasEnchant gets whether object has an enchant.
	HasEnchant(enchant enchant.Enchant) bool

	// Direction gets object direction.
	//
	// See LookWithAngle.
	Direction() Direction

	// CurrentHealth gets object's health.
	CurrentHealth() int

	// MaxHealth gets object's maximum health.
	MaxHealth() int

	// RestoreHealth restores object's health.
	RestoreHealth(amount int)

	// SetHealth sets current object health.
	SetHealth(v int)

	// SetMaxHealth sets maximum object health.
	SetMaxHealth(v int)

	// CurrentMana gets object's mana. Only works on players.
	CurrentMana() int

	// MaxMana gets object's maximum mana. Only works on players.
	MaxMana() int

	// SetMana sets current object mana. Only works on players.
	SetMana(v int)

	// SetMaxMana sets maximum object mana. Only works on players.
	SetMaxMana(v int)

	// CurrentSpeed returns current speed of the object.
	CurrentSpeed() float32

	// BaseSpeed returns base speed of the object.
	BaseSpeed() float32

	// SetBaseSpeed sets base speed of the object.
	SetBaseSpeed(v float32)

	// Strength returns strength of the unit.
	Strength() int

	// SetStrength sets strength of the unit.
	SetStrength(v int)

	// SetColor sets NPC or player color, given its index.
	//
	// Example:
	//
	// 		obj.SetColor(0, color.NRGBA{R:80, A: 255}) // dark red, 100% opacity
	SetColor(ind int, cl color.Color)

	// GetGold gets amount of gold for player object.
	GetGold() int

	// ChangeGold changes amount of gold for player object.
	ChangeGold(delta int)

	// GiveXp grants experience to a player.
	GiveXp(xp float32)

	// GetLevel gets player's level.
	GetLevel() int

	// GetClass gets player's character class.
	GetClass() player.Class

	// Player returns player associated with the object.
	Player() Player

	// GetScore gets player's score.
	//
	// Deprecated: use Player.GetScore.
	GetScore() int

	// ChangeScore changes player's score.
	//
	// Deprecated: use Player.ChangeScore.
	ChangeScore(score int)

	// HasOwner checks whether target is owned by object.
	HasOwner(owner Obj) bool

	// HasOwnerIn checks whether target is owned by any object in the group.
	HasOwnerIn(owners ObjGroup) bool

	// SetOwner makes an object the owner of the target. This will make the target
	// friendly to the owner, and it will accredit the target's kills to the owner.
	//
	// Passing nil will clear the owner.
	//
	// For example, in a multiplayer map, you might have a switch that activates a
	// hazard. You can use this so that if the hazard kills anyone, the player who
	// activated the hazard gets the credit.
	SetOwner(owner Obj)

	// SetOwners is the same as SetOwner but with an object group as the owner.
	SetOwners(owners ObjGroup)

	// Freeze or unfreeze an object in place.
	Freeze(freeze bool)

	// Pause an object temporarily.
	Pause(dt Duration)

	// Move an object to a waypoint. The object must be movable or attached to a "Mover".
	//
	// If the waypoint is linked, the object will continue to move once it reaches the first waypoint.
	Move(wp WaypointObj)

	// WalkTo causes an object to walk to a location.
	WalkTo(p Pointf)

	// LookAtDirection causes object to look in a direction.
	LookAtDirection(dir Direction)

	// LookWithAngle sets an object's direction. The direction is an angle represented as an integer between 0 and 255.
	// Due east is 0, and the angle increases as the object turns clock-wise.
	LookWithAngle(angle int)

	// LookAtObject sets direction of object so it is looking at another object.
	LookAtObject(target Positioner)

	// CanSee first checks if the location of the objects are within 512 of each other coordinate-wise. It not, it returns false.
	//
	// It then checks whether the first object can see the second object.
	CanSee(obj Obj) bool

	// ApplyForce applies a force vector to an object.
	ApplyForce(force Pointf)

	// PushTo calculate a unit vector from the object's location to the specified
	// location, and multiply it by the specified magnitude. This vector will be
	// applied as a force via ApplyForce.
	PushTo(pos Positioner, force float32)

	// Damage the target with a given source object, amount, and damage type.
	Damage(source Obj, amount int, typ damage.Type)

	// Delete an object.
	Delete()

	// DeleteAfter delete object after a delay.
	DeleteAfter(dt Duration)

	// Idle causes creature to idle.
	Idle()

	// Wander causes an object to wander.
	Wander()

	// Hunt causes creature to hunt.
	Hunt()

	// Return cause object to move to its starting location.
	Return()

	// Follow causes a creature to follow target, and it won't attack anything unless disrupted or instructed to.
	Follow(target Positioner)

	// Guard causes a creature to move to a location, guard a nearby location,
	// and attack any enemies that move within range of the guarded location.
	Guard(p1, p2 Pointf, distance float32)

	// Attack a target.
	Attack(target Positioner)

	// IsAttackedBy gets whether object is being attacked by another object.
	IsAttackedBy(by Obj) bool

	// HitMelee causes object to melee attacks a location.
	HitMelee(p Pointf)

	// HitRanged causes object to ranged attacks a location.
	HitRanged(p Pointf)

	// Flee causes creature to run away from target.
	Flee(target Positioner, dt Duration)

	// HasItem gets whether this specific item is in the object's inventory.
	// For searching items by object type or other properties, see InItems().FindObjects(...).
	HasItem(item Obj) bool

	// HasEquipment gets whether this specific item is in the object's inventory and is equipped.
	// For searching items by object type or other properties, see InEquipment().FindObjects(...).
	HasEquipment(item Obj) bool

	// FindItems calls fnc for all items matching all the conditions.
	// It returns a number of items matched. If fnc returns false, the function stops the search.
	// If fnc is nil, the function only counts a number of objects matching a condition.
	//
	// Example:
	//
	//	// Check if unit has at least one apple
	//	if obj.FindItems(nil, ns.HasTypeName{"RedApple"}) != 0 {
	//		fmt.Println("Has an apple!")
	//	}
	//
	//	// Drop all armor
	//	obj.FindItems(
	//		func(it ns.Obj) {
	//			obj.Drop(it)
	//			return true
	//		},
	//		ns.EqualClass(object.ClassArmor),
	//	)
	//
	// Deprecated: use InItems().FindObjects(...)
	FindItems(fnc func(obj Obj) bool, conditions ...ObjCond) int

	// GetLastItem returns the object of the last item in the object's inventory. If the inventory is empty, it returns nil.
	//
	// This is used with GetPreviousItem to iterate through an object's inventory.
	//
	// Example:
	//
	//		for it := obj.GetLastItem(); it != nil; it = it.GetPreviousItem() {
	//			// ...
	//		}
	GetLastItem() Obj

	// GetPreviousItem returns the object of the previous item in the inventory from the given object.
	// If the specified object is not in an inventory, or there are no more items in the inventory, it returns nil.
	//
	// This is used with GetLastItem to iterate through an object's inventory.
	GetPreviousItem() Obj

	// Items returns all items in the object's inventory.
	// An optional list of conditions can be specified for filtering returned items.
	//
	// Resulting slice is read-only, changes to it won't be reflected on the object.
	// Use Pickup or Drop for adding or removing items.
	Items(conditions ...ObjCond) []Obj

	// Equipment returns all equipped items in the object's inventory.
	// An optional list of conditions can be specified for filtering returned items.
	//
	// Resulting slice is read-only, changes to it won't be reflected on the object.
	// Use Equip or Unequip for changing equipment.
	Equipment(conditions ...ObjCond) []Obj

	// InItems returns an ObjSearcher that can be used in FindObjectIn and FindAllObjectsIn or searched directly.
	//
	// Calling InItems().FindObjects() will call fnc for all items matching all the conditions.
	// It returns a number of items matched. If fnc returns false, the function stops the search.
	// If fnc is nil, the function only counts a number of objects matching a condition.
	//
	// Example:
	//
	//	// Check if unit has at least one apple
	//	if obj.InItems().FindObjects(nil, ns.HasTypeName{"RedApple"}) != 0 {
	//		fmt.Println("Has an apple!")
	//	}
	//
	//	// Drop all armor
	//	obj.InItems().FindObjects(
	//		func(it ns.Obj) {
	//			obj.Drop(it)
	//			return true
	//		},
	//		ns.EqualClass(object.ClassArmor),
	//	)
	InItems() ObjSearcher

	// InEquipment returns an ObjSearcher that can be used in FindObjectIn and FindAllObjectsIn or searched directly.
	//
	// Calling InEquipment().FindObjects() will call fnc for all equipped items matching all the conditions.
	// It returns a number of items matched. If fnc returns false, the function stops the search.
	// If fnc is nil, the function only counts a number of objects matching a condition.
	//
	// Example:
	//
	//	// Check if unit has sword equipped
	//	if obj.InEquipment().FindObjects(nil, ns.HasTypeName{"Sword"}) != 0 {
	//		fmt.Println("Has an apple!")
	//	}
	//
	//	// Drop all equipped armor
	//	obj.InEquipment().FindObjects(
	//		func(it ns.Obj) {
	//			obj.Drop(it)
	//			return true
	//		},
	//		ns.EqualClass(object.ClassArmor),
	//	)
	InEquipment() ObjSearcher

	// GetHolder returns the object that contains the item in its inventory.
	GetHolder() Obj

	// Pickup cause object to pick up an item.
	//
	// Returns false if object cannot pick up the item.
	Pickup(item Obj) bool

	// Drop cause object to drop an item.
	//
	// Returns false if object cannot drop the item.
	Drop(item Obj) bool

	// Equip cause object to equip an item.
	//
	// If an item is not in object's inventory, it will be transferred to it.
	//
	// Returns false if object cannot pick up or equip the item.
	Equip(item Obj) bool

	// Unequip cause object to unequip an item.
	//
	// If an item is not in object's inventory, it will do nothing.
	//
	// Returns false if object cannot unequip the item.
	Unequip(item Obj) bool

	// ZombieStayDown sets zombie to stay down.
	ZombieStayDown()

	// RaiseZombie raises a zombie. Also clears stay down state.
	RaiseZombie()

	// Chat displays a localized string in a speech bubble.
	//
	// If the string is not in the string database, it will instead print an error message with "MISSING:".
	Chat(message StringID)

	// ChatTimer displays a localized string in a speech bubble for a given duration (in seconds or frames).
	//
	// If the string is not in the string database, it will instead print an error message with "MISSING:".
	ChatTimer(message StringID, dt Duration)

	// ChatStr displays a string in a speech bubble.
	// It does not localize the string.
	ChatStr(message string)

	// ChatStrTimer displays a string in a speech bubble for a given duration (in seconds or frames).
	// It does not localize the string.
	ChatStrTimer(message string, dt Duration)

	// DestroyChat destroys object's speech bubble.
	DestroyChat()

	// CreateMover creates a Mover for an object.
	CreateMover(wp WaypointObj, speed float32) Obj

	// GetElevatorStatus gets elevator status.
	GetElevatorStatus() int

	// AggressionLevel sets a creature's aggression level. The most commonly used value is 0.83.
	AggressionLevel(level float32)

	// SetRoamFlag sets roaming flags for object. Default is 0x80.
	SetRoamFlag(flags int)

	// RetreatLevel causes the creature to retreat if its health falls below the specified percentage (0.0 - 1.0).
	RetreatLevel(percent float32)

	// ResumeLevel causes the creature to stop retreating if its health is above the specified percentage (0.0 - 1.0).
	ResumeLevel(percent float32)

	// AwardSpell awards spell level to object.
	//
	// This will raise the spell level of the object.
	// If the object can not cast this spell then it will have no effect.
	AwardSpell(spell spell.Spell) bool

	// Enchant grants object an enchantment of a specified duration.
	Enchant(enchant enchant.Enchant, dt Duration)

	// EnchantOff removes enchant from an object.
	EnchantOff(enchant enchant.Enchant)

	// TrapSpells sets spells on a bomber or a wizard trap.
	TrapSpells(spells ...spell.Spell)

	// TrapSpellsAdv sets spells on a bomber or a wizard trap. It's an advanced version of TrapSpells.
	TrapSpellsAdv(spells []TrapSpell)

	// OnEvent sets a function script to call for an event.
	OnEvent(event ObjectEvent, fnc Func)
}

type ObjGroup interface {
	ObjGroupHandle
	ObjSearcher

	// Name returns object group name.
	Name() string
	// Enable or disable the object.
	Enable(enable bool)
	// Toggle the object's enabled state.
	// It returns the new state after toggling the object.
	Toggle() bool

	// HasOwner checks whether any object in target group is owned by object.
	HasOwner(owner Obj) bool

	// HasOwnerIn checks whether any object in target is owned by any object in the group.
	HasOwnerIn(owners ObjGroup) bool

	// SetOwner sets the owner for each object in a group. See Obj.SetOwner for details.
	SetOwner(owner Obj)

	// SetOwners is the same as SetOwner but with an object group as the owner.
	SetOwners(owners ObjGroup)

	// Pause objects of a group temporarily.
	Pause(dt Duration)

	// Move moves the objects in a group to a waypoint. The objects must be movable or attached to a "Mover".
	//
	// If the waypoint is linked, the objects will continue to move once they reach the first waypoint.
	Move(wp WaypointObj)

	// WalkTo causes objects in a group to walk to a location.
	WalkTo(p Pointf)

	// LookAtDirection causes objects in a group to look in a direction.
	LookAtDirection(dir Direction)

	// Damage damages the target objects with a given source object, amount, and damage type.
	Damage(source Obj, amount int, typ damage.Type)

	// Delete deletes objects in a group.
	Delete()

	// Idle causes creatures in a group to idle.
	Idle()

	// Wander cause objects in a group to wander.
	Wander()

	// Hunt causes creatures in a group to hunt.
	Hunt()

	// Follow causes the creatures to follow target, and they won't attack anything unless disrupted or instructed to.
	Follow(target Positioner)

	// Guard is the same as CreatureGuard but applies to creatures in a group.
	Guard(p1, p2 Pointf, distance float32)

	// Flee causes creatures to run away from target.
	Flee(target Positioner, dt Duration)

	// HitMelee causes objects in a group to melee attacks a location.
	HitMelee(p Pointf)

	// HitRanged causes objects in a group to ranged attacks a location.
	HitRanged(p Pointf)

	// Attack causes objects in a group to attack a target.
	Attack(target Positioner)

	// ZombieStayDown sets group of zombies to stay down.
	ZombieStayDown()

	// RaiseZombie raises a zombie group. Also clears stay down state.
	RaiseZombie()

	// CreateMover creates a Mover for every object in a group.
	CreateMover(wp WaypointObj, speed float32)

	// AggressionLevel sets a group of creature's aggression level. The most commonly used value is 0.83.
	AggressionLevel(level float32)

	// SetRoamFlag sets roaming flags for objects in a group. Default is 0x80.
	SetRoamFlag(flags int)

	// RetreatLevel causes the creatures to retreat if its health falls below the specified percentage (0.0 - 1.0).
	RetreatLevel(percent float32)

	// ResumeLevel causes the creatures to stop retreating if its health is above the specified percentage (0.0 - 1.0).
	ResumeLevel(percent float32)

	// AwardSpell awards spell level to objects in a group.
	//
	// This will raise the spell level of the objects in the group.
	// If an object can not cast this spell then it will have no effect on that object.
	AwardSpell(spell spell.Spell)

	// Enchant grants objects in a group an enchantment of a specified duration.
	Enchant(enchant enchant.Enchant, dt Duration)

	// AllObjects returns all objects as a list.
	AllObjects() Objects

	// EachObject calls fnc for all objects in the group.
	// If fnc returns false, the iteration stops.
	// If recursive is true, iteration will include items from nested groups.
	EachObject(recursive bool, fnc func(obj Obj) bool)
}

// DestroyEveryChat destroys all speech bubbles.
func DestroyEveryChat() {
	if impl == nil {
		return
	}
	impl.DestroyEveryChat()
}

// MakeFriendly sets object friendly with host.
func MakeFriendly(obj Obj) {
	if impl == nil {
		return
	}
	impl.MakeFriendly(obj)
}

// MakeEnemy unsets object as friendly.
func MakeEnemy(obj Obj) {
	if impl == nil {
		return
	}
	impl.MakeEnemy(obj)
}

// BecomePet sets object as pet of host.
func BecomePet(obj Obj) {
	if impl == nil {
		return
	}
	impl.BecomePet(obj)
}

// BecomeEnemy unsets object as pet of host.
func BecomeEnemy(obj Obj) {
	if impl == nil {
		return
	}
	impl.BecomeEnemy(obj)
}

var _ ObjSearcher = Objects{}

// Objects is a list of objects. It can be searched.
type Objects []Obj

// Enable or disable the objects.
func (arr Objects) Enable(enable bool) {
	for _, obj := range arr {
		obj.Enable(enable)
	}
}

// Toggle the objects enabled state.
// It returns the new state after toggling the objects. If at least one object is enabled, it returns true.
func (arr Objects) Toggle() bool {
	one := false
	for _, obj := range arr {
		if obj.Toggle() {
			one = true
		}
	}
	return one
}

// Delete deletes objects in a list.
func (arr Objects) Delete() {
	for _, obj := range arr {
		obj.Delete()
	}
}

// SetOwner sets the owner for each object in a list. See Obj.SetOwner for details.
func (arr Objects) SetOwner(owner Obj) {
	for _, obj := range arr {
		obj.SetOwner(owner)
	}
}

// SetOwners is the same as SetOwner but with an object group as the owner.
func (arr Objects) SetOwners(owners ObjGroup) {
	for _, obj := range arr {
		obj.SetOwners(owners)
	}
}

// FindObjects calls fnc for all objects in the list. See ObjSearcher.
func (arr Objects) FindObjects(fnc func(it Obj) bool, conditions ...ObjCond) int {
	if len(arr) == 0 {
		return 0
	}
	filter := AND(conditions)
	cnt := 0
	for _, obj := range arr {
		if !filter.Matches(obj) {
			continue
		}
		cnt++
		if fnc != nil && !fnc(obj) {
			break
		}
	}
	return cnt
}
