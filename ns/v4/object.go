package ns

import (
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/script"
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
	Pause(dt script.Duration)

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
	DeleteAfter(dt script.Duration)

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
	Flee(target Positioner, dt script.Duration)

	// HasItem gets whether the item is in the object's inventory.
	HasItem(item Obj) bool

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

	// GetHolder returns the object that contains the item in its inventory.
	GetHolder() Obj

	// Pickup cause object to pickup an item.
	Pickup(item Obj) bool

	// Drop cause object to drop an item.
	Drop(item Obj) bool

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
	ChatTimer(message StringID, dt script.Duration)

	// ChatStr displays a string in a speech bubble.
	// It does not localize the string.
	ChatStr(message string)

	// ChatStrTimer displays a string in a speech bubble for a given duration (in seconds or frames).
	// It does not localize the string.
	ChatStrTimer(message string)

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
	Enchant(enchant enchant.Enchant, dt script.Duration)

	// EnchantOff removes enchant from an object.
	EnchantOff(enchant enchant.Enchant)

	// TrapSpells sets spells on a bomber.
	TrapSpells(spell1 spell.Spell, spell2 spell.Spell, spell3 spell.Spell)

	// OnEvent sets a function script to call for an event.
	OnEvent(event ObjectEvent, fnc Func)
}

type ObjGroup interface {
	ObjGroupHandle
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
	Pause(dt script.Duration)

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
	Flee(target Positioner, dt script.Duration)

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
	Enchant(enchant enchant.Enchant, dt script.Duration)

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
