package ns

import (
	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/class"
	"github.com/noxworld-dev/noxscript/ns/v4/damage"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
	"github.com/noxworld-dev/noxscript/ns/v4/subclass"
)

// SELF object id.
//
// This constant can be used any place where an object id is used.
const SELF = ObjectID(-2)

// OTHER object id.
//
// This constant can be used any place where an object id is used.
const OTHER = ObjectID(-1)

type ObjectID int

func (id ObjectID) ScriptID() int {
	return int(id)
}

func (id ObjectID) ObjScriptID() int {
	return int(id)
}

type ObjectGroupID int

func (id ObjectGroupID) ScriptID() int {
	return int(id)
}

func (id ObjectGroupID) ObjGroupScriptID() int {
	return int(id)
}

// CreateObject creates an object at a location.
//
// This will create an object given a type and a starting location.
//
// Example usage:
//
//	spider := ns.CreateObject("SmallAlbinoSpider", ns.Waypoint("SpiderHole"))
func CreateObject(typ string, waypoint WaypointID) ObjectID {
	return asObj(ns4.CreateObject(typ, ns4.AsWaypoint(waypoint)))
}

// Object looks up an object by name.
//
// This will lookup an object by its script name and return the object id.
func Object(name string) ObjectID {
	return asObj(ns4.Object(name))
}

// GetTrigger gets SELF if valid.
func GetTrigger() ObjectID {
	return asObj(ns4.GetTrigger())

}

// GetCaller gets OTHER if valid.
func GetCaller() ObjectID {
	return asObj(ns4.GetCaller())
}

// GetObjectX gets object X coordinate.
func GetObjectX(id ObjectID) float32 {
	return ns4.AsObj(id).Pos().X
}

// GetObjectY gets object Y coordinate.
func GetObjectY(id ObjectID) float32 {
	return ns4.AsObj(id).Pos().Y
}

// GetObjectZ gets object Z coordinate.
func GetObjectZ(id ObjectID) float32 {
	return ns4.AsObj(id).Z()
}

// GetDirection gets object direction.
//
// See LookWithAngle for details.
func GetDirection(id ObjectID) Direction {
	return Direction(ns4.AsObj(id).Direction())
}

// MoveObject sets an object's location.
func MoveObject(id ObjectID, x, y float32) {
	ns4.AsObj(id).SetPos(ns4.Pointf{X: x, Y: y})
}

// Raise sets an object's Z coordinate.
//
// This will set an object's Z coordinate and then let the object fall down.
func Raise(id ObjectID, z float32) {
	ns4.AsObj(id).SetZ(z)
}

// IsTrigger gets whether object is SELF.
func IsTrigger(id ObjectID) bool {
	return ns4.IsTrigger(ns4.AsObj(id))
}

// IsCaller gets whether object is OTHER.
func IsCaller(id ObjectID) bool {
	return ns4.IsCaller(ns4.AsObj(id))
}

// IsGameBall gets whether object is a GameBall.
func IsGameBall(id ObjectID) bool {
	return ns4.IsGameBall(ns4.AsObj(id))
}

// IsCrown gets whether object is a Crown.
func IsCrown(id ObjectID) bool {
	return ns4.IsCrown(ns4.AsObj(id))
}

// IsSummoned gets whether object is a summoned creature.
func IsSummoned(id ObjectID) bool {
	return ns4.IsSummoned(ns4.AsObj(id))
}

// IsObjectOn gets whether object is enabled.
func IsObjectOn(id ObjectID) bool {
	return ns4.AsObj(id).IsEnabled()
}

// ObjectOn enables an object.
func ObjectOn(id ObjectID) {
	ns4.AsObj(id).Enable(true)
}

// ObjectOff disables an object.
func ObjectOff(id ObjectID) {
	ns4.AsObj(id).Enable(false)
}

// ObjectToggle toggles an object.
func ObjectToggle(id ObjectID) {
	ns4.AsObj(id).Toggle()
}

// Delete an object.
func Delete(id ObjectID) {
	ns4.AsObj(id).Delete()
}

// Wander causes an object to wander.
//
// This will cause a NPC or monster to wander.
func Wander(id ObjectID) {
	ns4.AsObj(id).Wander()
}

// GoBackHome causes object to move to its starting location.
func GoBackHome(id ObjectID) {
	ns4.AsObj(id).Return()
}

// Chat causes an object to say a localized string.
//
// This will display a localized in string a speech bubble. If the string
// is not in the database string, it will instead print an error message with
// "MISSING:".
func Chat(id ObjectID, message string) {
	ns4.AsObj(id).Chat(ns4.StringID(message))
}

// IsLocked gets whether object is locked.
//
// This will return whether an object is locked. It works with any kind of lock.
func IsLocked(id ObjectID) bool {
	return ns4.AsObj(id).IsLocked()
}

// UnlockDoor unlocks a door.
//
// This will unlock a door. It has no effect if the object is not a door.
func UnlockDoor(id ObjectID) {
	ns4.AsObj(id).Lock(false)
}

// LockDoor locks a door.
//
// This will lock a door. It has no effect if the object is not a door.
func LockDoor(id ObjectID) {
	ns4.AsObj(id).Lock(true)
}

// Move an object to a waypoint.
//
// This moves an object to a waypoint. The object must be movable or attached
// to a "Mover". If the waypoint is linked, the object will continue to move
// once it reaches the first waypoint.
func Move(id ObjectID, waypoint WaypointID) {
	ns4.AsObj(id).Move(ns4.AsWaypoint(waypoint))
}

// LookAtDirection causes object to look in a direction.
//
// See Direction for details.
func LookAtDirection(id ObjectID, direction Direction) {
	ns4.AsObj(id).LookAtDirection(ns4.Direction(direction))
}

// LookWithAngle sets an object's direction.
//
// This will set an object's direction. The direction is an angle
// represented as an integer between 0 and 255. Due east is 0, and the angle
// increases as the object turns clock-wise.
//
// Angle is a number between 0 and 255.
func LookWithAngle(id ObjectID, angle int) {
	ns4.AsObj(id).LookWithAngle(angle)
}

// PushObjectTo pushes an object to a location.
func PushObjectTo(id ObjectID, x, y float32) {
	ns4.AsObj(id).ApplyForce(ns4.Ptf(x, y))
}

// PushObject pushes an object from a vector and magnitude.
//
// This will calculate a unit vector from the object's location to the specified
// location, and multiply it by the specified magnitude. This vector will be
// added to the object's location.
func PushObject(id ObjectID, magnitude float32, x, y float32) {
	ns4.AsObj(id).PushTo(ns4.Ptf(x, y), magnitude)
}

// GetLastItem gets object's last inventory item.
//
// This will return the object id of the last item in the object's inventory.
// If the inventory is empty, it will return 0.
//
// This is used with GetPreviousItem to iterate through an object's inventory.
func GetLastItem(id ObjectID) ObjectID {
	return asObj(ns4.AsObj(id).GetLastItem())
}

// GetPreviousItem gets previous inventory item.
//
// This will return the object id of the previous item in the inventory from
// the given object id. If the specified object id is not in an inventory, or
// there are no more items in the inventory, it will return 0.
//
// This is used with GetLastItem to iterate through an object's inventory.
func GetPreviousItem(id ObjectID) ObjectID {
	return asObj(ns4.AsObj(id).GetPreviousItem())
}

// HasItem gets whether the item is in the object's inventory.
func HasItem(holder ObjectID, item ObjectID) bool {
	return ns4.AsObj(holder).HasItem(ns4.AsObj(item))
}

// GetHolder gets the holder of an item.
//
// This will return the object id that contains the item in its inventory.
func GetHolder(item ObjectID) ObjectID {
	return asObj(ns4.AsObj(item).GetHolder())
}

// Pickup causes object to pickup an item.
func Pickup(id ObjectID, item ObjectID) bool {
	return ns4.AsObj(id).Pickup(ns4.AsObj(item))
}

// Drop causes object to drop an item.
func Drop(id ObjectID, item ObjectID) bool {
	return ns4.AsObj(id).Drop(ns4.AsObj(item))
}

// HasClass gets whether object has a class.
func HasClass(id ObjectID, className Class) bool {
	return ns4.AsObj(id).HasClass(class.Class(className))
}

// CurrentHealth gets object's health.
func CurrentHealth(id ObjectID) int {
	return ns4.AsObj(id).CurrentHealth()
}

// MaxHealth gets object's maximum health.
func MaxHealth(id ObjectID) int {
	return ns4.AsObj(id).MaxHealth()
}

// RestoreHealth restores object's health.
func RestoreHealth(id ObjectID, amount int) {
	ns4.AsObj(id).RestoreHealth(amount)
}

// Damage an object.
//
// This will damage the target with a given source object, amount, and damage type.
//
// See DamageType for details.
func Damage(target ObjectID, source ObjectID, amount int, typ DamageType) {
	ns4.AsObj(target).Damage(ns4.AsObj(source), amount, damage.Type(typ))
}

// CreateMover creates a Mover for an object.
func CreateMover(id ObjectID, waypoint WaypointID, speed float32) ObjectID {
	return asObj(ns4.AsObj(id).CreateMover(ns4.AsWaypoint(waypoint), speed))
}

// MakeFriendly sets object friendly with host.
func MakeFriendly(id ObjectID) {
	ns4.MakeFriendly(ns4.AsObj(id))
}

// MakeEnemy unsets object as friendly.
func MakeEnemy(id ObjectID) {
	ns4.MakeEnemy(ns4.AsObj(id))
}

// BecomePet sets object as pet of host.
func BecomePet(id ObjectID) {
	ns4.BecomePet(ns4.AsObj(id))
}

// BecomeEnemy unsets object as pet of host.
func BecomeEnemy(id ObjectID) {
	ns4.BecomeEnemy(ns4.AsObj(id))
}

// Frozen sets frozen status of an object.
func Frozen(id ObjectID, frozen bool) {
	ns4.AsObj(id).Freeze(frozen)
}

type ObjEvent = int

// SetCallback sets a callback on an object.
//
// This will set a function script to call for an event. The callback index is
// one of the following:
//
//	3  = Enemy sighted
//	4  = Looking for enemy
//	5  = Death
//	6  = Change focus
//	7  = Is hit
//	8  = Retreat
//	9  = Collision
//	10 = Enemy heard
//	11 = End of waypoint
//	13 = Lost sight of enemy
//
// No other indexes are defined.
func SetCallback(id ObjectID, idx ObjEvent, callback Func) {
	ns4.AsObj(id).OnEvent(ns4.ObjectEvent(idx), callback)
}

// DeleteObjectTimer deletes object after a delay.
func DeleteObjectTimer(id ObjectID, delay int) {
	ns4.AsObj(id).DeleteAfter(ns4.Frames(delay))
}

// TrapSpells sets spells on a bomber.
func TrapSpells(id ObjectID, spell1, spell2, spell3 string) {
	ns4.AsObj(id).TrapSpells(spell.Spell(spell1), spell.Spell(spell2), spell.Spell(spell3))
}

// IsVisibleTo gets whether an object can see another object.
//
// This will first check if the location of the objects are within 512 of each
// other coordinate-wise. It not, it returns false.
//
// It then checks whether the first object can see the second object.
func IsVisibleTo(object1 ObjectID, object2 ObjectID) bool {
	return ns4.AsObj(object1).CanSee(ns4.AsObj(object2))
}

// LookAtObject sets direction of object so it is looking at another object.
func LookAtObject(id ObjectID, target ObjectID) {
	ns4.AsObj(id).LookAtObject(ns4.AsObj(target))
}

// Walk causes an object to walk to a location.
func Walk(id ObjectID, x, y float32) {
	ns4.AsObj(id).WalkTo(ns4.Ptf(x, y))
}

// SetOwner sets the owner of an object.
//
// This will make an object the owner of the target. This will make the target
// friendly to the owner, and it will accredit the target's kills to the owner.
//
// For example, in a multiplayer map, you might have a switch that activates a
// hazard. You can use this so that if the hazard kills anyone, the player who
// activated the hazard gets the credit.
func SetOwner(owner ObjectID, target ObjectID) {
	ns4.AsObj(target).SetOwner(ns4.AsObj(owner))
}

// SetOwners sets the objects in a group as owners of target.
//
// This is the same as SetOwner but with a object group as the owner.
func SetOwners(owners ObjectGroupID, target ObjectID) {
	ns4.AsObj(target).SetOwners(ns4.AsObjGroup(owners))
}

// IsOwnedBy get whether target is owned by object.
func IsOwnedBy(owner ObjectID, target ObjectID) bool {
	return ns4.AsObj(target).HasOwner(ns4.AsObj(owner))
}

// IsOwnedByAny gets whether target is owned by any object in the group.
func IsOwnedByAny(owners ObjectGroupID, target ObjectID) bool {
	return ns4.AsObj(target).HasOwnerIn(ns4.AsObjGroup(owners))
}

// ClearOwner clears the owner of an object.
func ClearOwner(id ObjectID) {
	ns4.AsObj(id).SetOwner(nil)
}

// ChatTimerSeconds causes an object to say a localized for string duration in seconds.
//
// This will display a localized in string a speech bubble. If the string
// is not in the database string, it will instead print an error message with
// "MISSING:".
func ChatTimerSeconds(id ObjectID, message string, duration int) {
	ns4.AsObj(id).ChatTimer(ns4.StringID(message), ns4.Seconds(float64(duration)))
}

// ChatTimer causes an object to say a localized for string duration in frames.
//
// This will display a localized in string a speech bubble. If the string
// is not in the database string, it will instead print an error message with
// "MISSING:".
func ChatTimer(id ObjectID, message string, duration int) {
	ns4.AsObj(id).ChatTimer(ns4.StringID(message), ns4.Frames(duration))
}

// DestroyChat destroys object's speech bubble.
func DestroyChat(id ObjectID) {
	ns4.AsObj(id).DestroyChat()
}

// DestroyEveryChat destroys all speech bubbles.
func DestroyEveryChat() {
	ns4.DestroyEveryChat()
}

// GetElevatorStatus gets elevator status.
func GetElevatorStatus(id ObjectID) int {
	return ns4.AsObj(id).GetElevatorStatus()
}

// CreatureGuard causes a creature to guard a location.
//
// This will cause a creature to move to a location, guard a nearby location,
// and attack any enemies that move within range of the guarded location.
func CreatureGuard(id ObjectID, x1, y1, x2, y2 float32, distance float32) {
	ns4.AsObj(id).Guard(ns4.Ptf(x1, y1), ns4.Ptf(x2, y2), distance)
}

// CreatureHunt causes creature to hunt.
func CreatureHunt(id ObjectID) {
	ns4.AsObj(id).Hunt()
}

// CreatureIdle causes creature to idle.
func CreatureIdle(id ObjectID) {
	ns4.AsObj(id).Idle()
}

// CreatureFollow causes creature to follow an object.
//
// This will cause a creature to follow target, and it won't attack anything
// unless disrupted or instructed to.
func CreatureFollow(id ObjectID, target ObjectID) {
	ns4.AsObj(id).Follow(ns4.AsObj(target))
}

// AggressionLevel sets creature's aggression level.
//
// This will set a creature's aggression level. The most commonly used value
// is 0.83.
func AggressionLevel(id ObjectID, level float32) {
	ns4.AsObj(id).AggressionLevel(level)
}

// HitLocation melee attacks a location.
func HitLocation(id ObjectID, x, y float32) {
	ns4.AsObj(id).HitMelee(ns4.Ptf(x, y))
}

// HitFarLocation ranged attacks a location.
func HitFarLocation(id ObjectID, x, y float32) {
	ns4.AsObj(id).HitRanged(ns4.Ptf(x, y))
}

// SetRoamFlag sets roaming flags.
//
// Flags a 8-bit value (default is 0x80).
func SetRoamFlag(id ObjectID, flags int) {
	ns4.AsObj(id).SetRoamFlag(flags)
}

// Attack an object.
func Attack(id ObjectID, target ObjectID) {
	ns4.AsObj(id).Attack(ns4.AsObj(target))
}

// RetreatLevel sets when creature retreats due to low health.
//
// This will cause the creature to retreat if its health falls below the
// specified percentage.
//
// Percent low health ratio is in range 0.0-1.0.
func RetreatLevel(id ObjectID, percent float32) {
	ns4.AsObj(id).RetreatLevel(percent)
}

// ResumeLevel sets when creature resumes due to health.
//
// This will cause the creature to stop retreating if its health is above the
// specified percentage.
//
// Percent health ratio is in range 0.0-1.0.
func ResumeLevel(id ObjectID, percent float32) {
	ns4.AsObj(id).ResumeLevel(percent)
}

// RunAway causes creature to run away from target.
func RunAway(id ObjectID, target ObjectID, duration int) {
	ns4.AsObj(id).Flee(ns4.AsObj(target), ns4.Frames(duration))
}

// PauseObject pauses an object temporarily.
func PauseObject(id ObjectID, duration int) {
	ns4.AsObj(id).Pause(ns4.Frames(duration))
}

// IsAttackedBy gets whether object1 is being attacked by object2.
func IsAttackedBy(id1 ObjectID, id2 ObjectID) bool {
	return ns4.AsObj(id1).IsAttackedBy(ns4.AsObj(id2))
}

// HasSubclass gets whether object has subclass.
//
// This will test whether an item has a specific subclass. The subclass
// overlaps, so you should probably test for the class first (HasClass).
func HasSubclass(id ObjectID, subclassName SubClass) bool {
	return ns4.AsObj(id).HasSubclass(subclass.SubClass(subclassName))
}

// ZombieStayDown sets zombie to stay down.
func ZombieStayDown(id ObjectID) {
	ns4.AsObj(id).ZombieStayDown()
}

// RaiseZombie raises a zombie. Also clears stay down state.
func RaiseZombie(id ObjectID) {
	ns4.AsObj(id).RaiseZombie()
}

// ObjectGroup looks up object group by name.
func ObjectGroup(name string) ObjectGroupID {
	return asObjGroup(ns4.ObjectGroup(name))
}

// ObjectGroupOn enables objects in a group.
func ObjectGroupOn(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Enable(true)
}

// ObjectGroupOff disables objects in a group.
func ObjectGroupOff(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Enable(false)
}

// ObjectGroupToggle toggles objects in group.
func ObjectGroupToggle(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Toggle()
}

// GroupDelete deletes objects in a group.
func GroupDelete(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Delete()
}

// GroupWander cause objects in a group to wander.
func GroupWander(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Wander()
}

// GroupMove moves objects in a group to a waypoint.
//
// This moves the objects in a group to a waypoint. The objects must be movable
// or attached to a "Mover". If the waypoint is linked, the objects will
// continue to move once they reach the first waypoint.
func GroupMove(objectGroup ObjectGroupID, waypoint WaypointID) {
	ns4.AsObjGroup(objectGroup).Move(ns4.AsWaypoint(waypoint))
}

// GroupLookAtDirection causes objects in a group to look in a direction.
//
// See Direction for details.
func GroupLookAtDirection(objectGroup ObjectGroupID, direction Direction) {
	ns4.AsObjGroup(objectGroup).LookAtDirection(ns4.Direction(direction))
}

// GroupDamage damages objects in a group.
//
// This will damage the target objects with a given source object, amount, and
// damage type.
//
// See DamageType for details.
func GroupDamage(targetGroup ObjectGroupID, source ObjectID, amount int, typ DamageType) {
	ns4.AsObjGroup(targetGroup).Damage(ns4.AsObj(source), amount, damage.Type(typ))
}

// GroupCreateMover creates a Mover for every object in a group.
func GroupCreateMover(objectGroup ObjectGroupID, waypoint WaypointID, speed float32) {
	ns4.AsObjGroup(objectGroup).CreateMover(ns4.AsWaypoint(waypoint), speed)
}

// ZombieGroupStayDown sets group of zombies to stay down.
func ZombieGroupStayDown(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).ZombieStayDown()
}

// RaiseZombieGroup raises a zombie. Also clears stay down state.
func RaiseZombieGroup(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).RaiseZombie()
}

// GroupSetRoamFlag sets roaming flags.
//
// Flags a 8-bit value (default is 0x80).
func GroupSetRoamFlag(objectGroup ObjectGroupID, flags int) {
	ns4.AsObjGroup(objectGroup).SetRoamFlag(flags)
}

// GroupAttack attacks an object.
func GroupAttack(objectGroup ObjectGroupID, target ObjectID) {
	ns4.AsObjGroup(objectGroup).Attack(ns4.AsObj(target))
}

// GroupRetreatLevel sets when creature retreats due to low health.
//
// This will cause the creatures to retreat if its health falls below the
// specified percentage.
//
// Percent low health ratio is in range 0.0-1.0.
func GroupRetreatLevel(objectGroup ObjectGroupID, percent float32) {
	ns4.AsObjGroup(objectGroup).RetreatLevel(percent)
}

// GroupResumeLevel sets when creature resumes due to health.
//
// This will cause the creatures to stop retreating if its health is above the
// specified percentage.
//
// Percent health ratio is in range 0.0-1.0.
func GroupResumeLevel(objectGroup ObjectGroupID, percent float32) {
	ns4.AsObjGroup(objectGroup).ResumeLevel(percent)
}

// GroupRunAway causes creatures to run away from target.
func GroupRunAway(objectGroup ObjectGroupID, target ObjectID, duration int) {
	ns4.AsObjGroup(objectGroup).Flee(ns4.AsObj(target), ns4.Frames(duration))
}

// GroupPauseObject pauses objects of a group temporarily.
func GroupPauseObject(objectGroup ObjectGroupID, duration int) {
	ns4.AsObjGroup(objectGroup).Pause(ns4.Frames(duration))
}

// GroupWalk causes objects in a group to walk to a location.
func GroupWalk(objectGroup ObjectGroupID, x, y float32) {
	ns4.AsObjGroup(objectGroup).WalkTo(ns4.Ptf(x, y))
}

// GroupSetOwner sets the owner of objects in a group.
//
// This is the same as SetOwner but with a object group as the target.
func GroupSetOwner(owner ObjectID, targets ObjectGroupID) {
	ns4.AsObjGroup(targets).SetOwner(ns4.AsObj(owner))
}

// GroupSetOwners sets the objects in a group as owners of target.
//
// This is the same as SetOwners but with a object group as the target.
func GroupSetOwners(owners ObjectGroupID, targets ObjectGroupID) {
	ns4.AsObjGroup(targets).SetOwners(ns4.AsObjGroup(owners))
}

// GroupIsOwnedBy gets whether any object in target group is owned by object.
func GroupIsOwnedBy(owner ObjectID, target ObjectGroupID) bool {
	return ns4.AsObjGroup(target).HasOwner(ns4.AsObj(owner))
}

// GroupIsOwnedByAny gets whether any object in target is owned by any object in the group.
func GroupIsOwnedByAny(owners ObjectGroupID, target ObjectGroupID) bool {
	return ns4.AsObjGroup(target).HasOwnerIn(ns4.AsObjGroup(owners))
}

// CreatureGroupGuard causes creatures in a group to guard a location.
//
// This is the same as CreatureGuard but applies to creatures in a group.
func CreatureGroupGuard(objectGroup ObjectGroupID, x1, y1, x2, y2 float32, distance float32) {
	ns4.AsObjGroup(objectGroup).Guard(ns4.Ptf(x1, y1), ns4.Ptf(x2, y2), distance)
}

// CreatureGroupHunt causes creatures in a group to hunt.
func CreatureGroupHunt(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Hunt()
}

// CreatureGroupIdle causes creatures in a group to idle.
func CreatureGroupIdle(objectGroup ObjectGroupID) {
	ns4.AsObjGroup(objectGroup).Idle()
}

// CreatureGroupFollow causes creatures in a group to follow an object.
//
// This will cause the creatures to follow target, and they won't attack
// anything unless disrupted or instructed to.
func CreatureGroupFollow(objectGroup ObjectGroupID, target ObjectID) {
	ns4.AsObjGroup(objectGroup).Follow(ns4.AsObj(target))
}

// GroupAggressionLevel sets group of creature's aggression level.
//
// This will set a group of creature's aggression level. The most commonly used
// value is 0.83.
func GroupAggressionLevel(objectGroup ObjectGroupID, level float32) {
	ns4.AsObjGroup(objectGroup).AggressionLevel(level)
}

// GroupHitLocation melee attacks a location.
func GroupHitLocation(objectGroup ObjectGroupID, x, y float32) {
	ns4.AsObjGroup(objectGroup).HitMelee(ns4.Ptf(x, y))
}

// GroupHitFarLocation ranged attacks a location.
func GroupHitFarLocation(objectGroup ObjectGroupID, x, y float32) {
	ns4.AsObjGroup(objectGroup).HitRanged(ns4.Ptf(x, y))
}
