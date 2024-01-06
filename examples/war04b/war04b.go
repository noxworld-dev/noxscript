package war04b

import (
	"strconv"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	gvar7  int
	gvar8  int
	gvar9  int
	gvar10 int
	gvar11 int
	gvar12 int
	gvar13 int
	gvar14 int
	gvar15 int
	gvar16 int
	gvar17 int
	gvar18 int
	flag19 bool
	flag20 bool
	flag21 bool
	flag22 bool
	flag23 bool
	ivar24 int
	ivar25 int
	gvar26 int
	gvar27 int
	gvar28 int
	obj29  [13]ns.ObjectID
	obj30  [13]ns.ObjectID
	ivar31 int
	obj32  [56]ns.ObjectID
	obj33  ns.ObjectID
	obj34  ns.ObjectID
	obj35  ns.ObjectID
	obj36  ns.ObjectID
	obj37  ns.ObjectID
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	gvar49 ns.ObjectGroupID
	gvar50 ns.WallGroupID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	wp53   ns.WaypointID
	wp54   ns.WaypointID
	wp55   ns.WaypointID
	wp56   [56]ns.WaypointID
	wp57   ns.WaypointID
	gvar58 ns.WaypointGroupID
	gvar59 ns.WaypointGroupID
	gvar60 ns.WaypointGroupID
)

func init() {
	gvar7 = 0
	gvar8 = 1
	gvar9 = 2
	gvar10 = 3
	gvar11 = 0
	gvar12 = 1
	gvar13 = 2
	gvar14 = 3
	gvar15 = 4
	gvar16 = 5
	gvar17 = 6
	gvar18 = 7
	flag19 = false
	flag20 = true
	flag21 = false
	flag22 = false
	flag23 = false
	ivar24 = 0
	ivar25 = 0
	gvar26 = gvar12
	gvar27 = 0
	gvar28 = gvar7
	ivar31 = 120
}
func InitializeSniperRoom() {
	obj4 = ns.Object("ExitGate")
	obj5 = ns.Object("SkeliGate01")
	obj6 = ns.Object("SkeliGate02")
	ns.LockDoor(obj4)
	ns.LockDoor(obj5)
	ns.LockDoor(obj6)
}
func UnlockExitGate() {
	ns.UnlockDoor(obj4)
	ns.UnlockDoor(obj5)
	ns.UnlockDoor(obj6)
}
func setUpFlameLever() {
	if !flag19 {
		goto LABEL1
	}
	return
LABEL1:
	flag19 = true
	gvar26 = gvar12
	ivar25 = 2
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL2
		}
		obj32[ivar24] = 0
		ivar24 += 1
	}
LABEL2:
	flameLever()
}
func flameLever() {
	var v0 int
	if !flag19 {
		goto LABEL1
	}
	v0 = gvar26
	if v0 == gvar12 {
		goto LABEL2
	}
	if v0 == gvar13 {
		goto LABEL3
	}
	if v0 == gvar14 {
		goto LABEL4
	}
	if v0 == gvar15 {
		goto LABEL5
	}
	if v0 == gvar16 {
		goto LABEL6
	}
	if v0 == gvar17 {
		goto LABEL7
	}
	if v0 == gvar18 {
		goto LABEL8
	}
	if v0 == gvar11 {
		goto LABEL9
	}
	goto LABEL10
LABEL2:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL11
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("SmallFlame", wp56[ivar24])
		ivar24 += 1
	}
LABEL11:
	gvar26 = gvar13
	ns.AudioEvent(ns.FireballCast, wp54)
	ns.AudioEvent(ns.FireballCast, wp55)
	ivar25 = 2
	goto LABEL10
LABEL3:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL13
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("MediumFlame", wp56[ivar24])
		ivar24 += 1
	}
LABEL13:
	gvar26 = gvar14
	ns.AudioEvent(ns.DemonBreath, wp54)
	ns.AudioEvent(ns.DemonBreath, wp55)
	ivar25 = 2
	goto LABEL10
LABEL4:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL15
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("Flame", wp56[ivar24])
		ivar24 += 1
	}
LABEL15:
	gvar26 = gvar15
	ivar25 = 2
	goto LABEL10
LABEL5:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL17
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("LargeFlame", wp56[ivar24])
		ivar24 += 1
	}
LABEL17:
	gvar26 = gvar16
	ivar25 = 150
	goto LABEL10
LABEL6:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL19
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("Flame", wp56[ivar24])
		ivar24 += 1
	}
LABEL19:
	gvar26 = gvar17
	ivar25 = 2
	goto LABEL10
LABEL7:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL21
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("MediumFlame", wp56[ivar24])
		ivar24 += 1
	}
LABEL21:
	gvar26 = gvar18
	ns.AudioEvent(ns.FireExtinguish, wp54)
	ns.AudioEvent(ns.FireExtinguish, wp55)
	ivar25 = 2
	goto LABEL10
LABEL8:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL23
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = ns.CreateObject("SmallFlame", wp56[ivar24])
		ivar24 += 1
	}
LABEL23:
	gvar26 = gvar11
	ivar25 = 2
	goto LABEL10
LABEL9:
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL25
		}
		ns.Delete(obj32[ivar24])
		obj32[ivar24] = 0
		ivar24 += 1
	}
LABEL25:
	gvar26 = gvar12
	flag19 = false
	goto LABEL10
LABEL10:
	ns.FrameTimer(ivar25, flameLever)
LABEL1:
	return
}
func skullHallOn() {
	if !flag20 {
		goto LABEL1
	}
	ns.CastSpellLocationLocation(ns.SPELL_FIREBALL, 3012, 2229, 3185, 2403)
	ns.CastSpellLocationLocation(ns.SPELL_FIREBALL, 2896, 2346, 3070, 2517)
	ns.FrameTimer(60, skullHallOn)
LABEL1:
	return
}
func Zombie1Rise() {
	var (
		v0 float32
		v1 float32
	)
	_ = v1
	_ = v0
	v0 = ns.GetObjectX(obj41)
	v1 = ns.GetObjectY(obj41)
	ns.RaiseZombie(obj41)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("ZombieAttackWP"))
	ns.LookAtObject(obj33, obj41)
}
func Zombie2Rise() {
	var (
		v0 float32
		v1 float32
	)
	_ = v1
	_ = v0
	v0 = ns.GetObjectX(obj42)
	v1 = ns.GetObjectY(obj42)
	ns.RaiseZombie(obj42)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("ZombieAttackWP"))
	ns.LookAtObject(obj33, obj42)
}
func Zombie3Rise() {
	var (
		v0 float32
		v1 float32
	)
	_ = v1
	_ = v0
	v0 = ns.GetObjectX(obj43)
	v1 = ns.GetObjectY(obj43)
	ns.RaiseZombie(obj43)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("ZombieAttackWP"))
	ns.LookAtObject(obj33, obj43)
}
func ZombieAttackHecubah() {
	ns.WayPointGroupOn(gvar59)
	ns.WayPointGroupOff(gvar58)
	ns.HitLocation(obj41, ns.GetWaypointX(ns.Waypoint("ZombieAttackWP")), ns.GetWaypointY(ns.Waypoint("ZombieAttackWP")))
	ns.HitLocation(obj42, ns.GetWaypointX(ns.Waypoint("ZombieAttackWP")), ns.GetWaypointY(ns.Waypoint("ZombieAttackWP")))
	ns.HitLocation(obj43, ns.GetWaypointX(ns.Waypoint("ZombieAttackWP")), ns.GetWaypointY(ns.Waypoint("ZombieAttackWP")))
}
func HecubahEncounterSEG4() {
	ns.SetDialog(obj33, ns.NORMAL, HecubahDialogStart, HecubahDialogEnd)
	gvar28 = gvar10
	ns.StartDialog(obj33, ns.GetHost())
}
func HecubahEncounterSEG3() {
	ns.SetDialog(obj34, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	gvar28 = gvar9
	ns.StartDialog(obj34, ns.GetHost())
}
func hecGone() {
	if !ns.IsCaller(obj33) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj33, ns.GetWaypointX(wp53), ns.GetWaypointY(wp53))
	ns.ObjectOff(obj33)
LABEL1:
	return
}
func ReleasePlayer() {
	if !ns.IsCaller(obj33) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LockDoor(obj39)
	ns.LockDoor(obj40)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar50)
	ns.NoWallSound(false)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(ns.Object("War03b:Wolf1"), false)
	ns.Frozen(ns.Object("War03b:Wolf2"), false)
	ns.SetDialog(obj34, ns.NORMAL, necroTalkStart, necroTalkEnd)
	gvar27 = 5
	ns.StartDialog(obj34, ns.GetHost())
LABEL1:
	return
}
func necroDies() {
	ns.UnlockDoor(obj39)
	ns.UnlockDoor(obj40)
	ns.CancelDialog(obj34)
	ns.SetDialog(obj34, ns.NORMAL, necroTalkStart, necroTalkEnd)
	gvar27 = 4
	ns.StartDialog(obj34, ns.GetHost())
}
func necroTalkStart() {
	var v0 int
	v0 = gvar27
	if v0 == 1 {
		goto LABEL1
	}
	if v0 == 2 {
		goto LABEL2
	}
	if v0 == 3 {
		goto LABEL3
	}
	if v0 == 4 {
		goto LABEL4
	}
	if v0 == 5 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroTaunts")
	ns.Frozen(obj34, true)
	ns.Frozen(obj41, true)
	ns.Frozen(obj43, true)
	goto LABEL6
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroGloats")
	ns.Frozen(obj34, true)
	ns.Frozen(obj41, true)
	ns.Frozen(obj43, true)
	goto LABEL6
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroThreatens2")
	ns.Frozen(obj34, true)
	ns.Frozen(obj41, true)
	ns.Frozen(obj43, true)
	goto LABEL6
LABEL4:
	ns.WayPointGroupOff(gvar59)
	ns.WayPointGroupOn(gvar60)
	ns.LookAtObject(obj34, ns.GetHost())
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroDies")
	ns.Frozen(obj34, true)
	ns.Frozen(obj41, true)
	ns.Frozen(obj43, true)
	goto LABEL6
LABEL5:
	ns.AggressionLevel(obj34, 0.83)
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroThreatens")
	ns.LookAtObject(obj34, ns.GetHost())
	ns.Frozen(obj34, true)
	ns.Frozen(obj41, true)
	ns.Frozen(obj43, true)
	goto LABEL6
LABEL6:
	return
}
func necroTalkEnd() {
	ns.CancelDialog(obj34)
	ns.Frozen(obj34, false)
	ns.Frozen(obj41, false)
	ns.Frozen(obj43, false)
	if gvar27 != 4 {
		goto LABEL1
	}
	ns.JournalEdit(ns.GetHost(), "Chapter4SearchCrypts", 4)
	ns.JournalEntry(ns.GetHost(), "Chapter4Escape", 2)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL1:
	return
}
func HecubahShootZombie1() {
	ns.LookAtObject(obj33, obj41)
	ns.CastSpellObjectObject(ns.SPELL_FIREBALL, obj33, obj41)
	ns.FrameTimer(30, HecubahShootZombie2)
}
func HecubahShootZombie2() {
	ns.LookAtObject(obj33, obj42)
	ns.CastSpellObjectObject(ns.SPELL_FIREBALL, obj33, obj42)
	ns.FrameTimer(30, HecubahShootZombie3)
}
func HecubahShootZombie3() {
	ns.LookAtObject(obj33, obj43)
	ns.CastSpellObjectObject(ns.SPELL_FIREBALL, obj33, obj43)
	ns.FrameTimer(30, HecubahEncounterSEG3)
}
func ZombieInjured() {
	ns.Damage(ns.GetTrigger(), 0, 100, 1)
}
func InitializeFistTraps() {
	obj30[0] = ns.Object("FistTrap01")
	obj30[1] = ns.Object("FistTrap02")
	obj30[2] = ns.Object("FistTrap03")
	obj30[3] = ns.Object("FistTrap04")
	obj30[4] = ns.Object("FistTrap05")
	obj30[5] = ns.Object("FistTrap06")
	obj30[6] = ns.Object("FistTrap07")
	obj30[7] = ns.Object("FistTrap08")
	obj30[8] = ns.Object("FistTrap09")
	obj30[9] = ns.Object("FistTrap10")
	obj30[10] = ns.Object("FistTrap11")
	obj30[11] = ns.Object("FistTrap12")
	obj30[12] = ns.Object("FistTrap13")
	obj29[0] = ns.Object("FistTrapLight01")
	obj29[1] = ns.Object("FistTrapLight02")
	obj29[2] = ns.Object("FistTrapLight03")
	obj29[3] = ns.Object("FistTrapLight04")
	obj29[4] = ns.Object("FistTrapLight05")
	obj29[5] = ns.Object("FistTrapLight06")
	obj29[6] = ns.Object("FistTrapLight07")
	obj29[7] = ns.Object("FistTrapLight08")
	obj29[8] = ns.Object("FistTrapLight09")
	obj29[9] = ns.Object("FistTrapLight10")
	obj29[10] = ns.Object("FistTrapLight11")
	obj29[11] = ns.Object("FistTrapLight12")
	obj29[12] = ns.Object("FistTrapLight13")
}
func MapInitialize() {
	ivar24 = 0
	for {
		if !(ivar24 < 56) {
			goto LABEL1
		}
		wp56[ivar24] = ns.Waypoint("FlameWay" + strconv.Itoa(ivar24+1))
		ivar24 += 1
	}
LABEL1:
	obj33 = ns.Object("Hecubah")
	obj34 = ns.Object("Necromancer")
	obj35 = ns.Object("EntranceGate1")
	obj36 = ns.Object("EntranceGate2")
	obj37 = ns.Object("EntranceGate3")
	obj38 = ns.Object("EntranceGate4")
	obj39 = ns.Object("PitGate1")
	obj40 = ns.Object("PitGate2")
	obj41 = ns.Object("DeadZombie1")
	obj42 = ns.Object("DeadZombie2")
	obj43 = ns.Object("DeadZombie3")
	obj44 = ns.Object("Zombie01")
	obj45 = ns.Object("Zombie02")
	obj46 = ns.Object("Zombie03")
	obj47 = ns.Object("FlameLever1")
	obj48 = ns.Object("FlameLever2")
	gvar49 = ns.ObjectGroup("Zombies")
	gvar50 = ns.WallGroup("InvisibleBlockWall")
	wp52 = ns.Waypoint("SecretAudioOrigin")
	wp54 = ns.Waypoint("FlameSound1")
	wp55 = ns.Waypoint("FlameSound2")
	wp51 = ns.Waypoint("AudioOrigin")
	wp57 = ns.Waypoint("WP")
	wp53 = ns.Waypoint("HecubahStorage")
	gvar58 = ns.WaypointGroup("BeforeSetpieceWay")
	gvar59 = ns.WaypointGroup("SetpieceWay")
	gvar60 = ns.WaypointGroup("AfterSetpieceWay")
	ns.WayPointGroupOff(gvar59)
	ns.WayPointGroupOff(gvar60)
	ns.Damage(obj41, 0, 100, 0)
	ns.Damage(obj42, 0, 100, 0)
	ns.Damage(obj43, 0, 100, 0)
	ns.ZombieStayDown(obj41)
	ns.ZombieStayDown(obj42)
	ns.ZombieStayDown(obj43)
	ns.StoryPic(obj34, "NecromancerPic")
	InitializeFistTraps()
	InitializeSniperRoom()
}
func PlayerDeath() {
	ns.DeathScreen(4)
}
func ToggleSkullHall() {
	if !flag20 {
		goto LABEL1
	}
	flag20 = false
	goto LABEL2
LABEL1:
	flag20 = true
	skullHallOn()
LABEL2:
	return
}
func HecubahDialogStart() {
	var v0 int
	v0 = gvar28
	if v0 == gvar7 {
		goto LABEL1
	}
	if v0 == gvar8 {
		goto LABEL2
	}
	if v0 == gvar9 {
		goto LABEL3
	}
	if v0 == gvar10 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.GhostRecognize, "War04a:Hecubah")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.GhostRecognize, "War04a:HecubahSEG2")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj34, ns.GetHost())
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroWarnsHec")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj33, ns.GetHost())
	ns.TellStory(ns.GhostRecognize, "War04a:HecubahSEG3")
	goto LABEL5
LABEL5:
	return
}
func HecubahDialogEnd() {
	var (
		v0 float32
		v1 float32
		v4 ns.WaypointID
		v5 int
	)
	_ = v1
	_ = v0
	v5 = gvar28
	if v5 == gvar7 {
		goto LABEL1
	}
	if v5 == gvar8 {
		goto LABEL2
	}
	if v5 == gvar9 {
		goto LABEL3
	}
	if v5 == gvar10 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.FrameTimer(20, Zombie2Rise)
	ns.FrameTimer(45, Zombie1Rise)
	ns.FrameTimer(60, Zombie3Rise)
	ns.FrameTimer(60, ZombieAttackHecubah)
	gvar28 = gvar8
	ns.CancelDialog(obj33)
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj33, obj42)
	gvar28 = gvar9
	ns.CancelDialog(obj33)
	ns.AggressionLevel(obj41, 0)
	ns.AggressionLevel(obj42, 0)
	ns.AggressionLevel(obj43, 0)
	ns.PushObject(obj41, 40, ns.GetObjectX(obj33), ns.GetObjectY(obj33))
	ns.PushObject(obj42, 40, ns.GetObjectX(obj33), ns.GetObjectY(obj33))
	ns.PushObject(obj43, 40, ns.GetObjectX(obj33), ns.GetObjectY(obj33))
	ns.SetOwner(obj39, obj41)
	ns.SetOwner(obj39, obj42)
	ns.SetOwner(obj39, obj43)
	ns.SetCallback(obj41, 7, ZombieInjured)
	ns.SetCallback(obj42, 7, ZombieInjured)
	ns.SetCallback(obj43, 7, ZombieInjured)
	ns.FrameTimer(20, HecubahShootZombie1)
	goto LABEL5
LABEL3:
	gvar28 = gvar10
	ns.CancelDialog(obj34)
	ns.FrameTimer(5, HecubahEncounterSEG4)
	goto LABEL5
LABEL4:
	v4 = ns.Waypoint("WP")
	v0 = ns.GetObjectX(obj33)
	v1 = ns.GetObjectY(obj33)
	ns.CreatureIdle(obj34)
	ns.UnlockDoor(obj35)
	ns.UnlockDoor(obj36)
	ns.LockDoor(obj37)
	ns.LockDoor(obj38)
	ns.ObjectOn(obj33)
	ns.Move(obj33, v4)
	flag22 = true
	ns.CancelDialog(obj33)
	goto LABEL5
LABEL5:
	return
}
func AttackSuccessful() {
	ns.SetDialog(obj33, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj33, ns.GetHost())
}
func HecubahEncounterSEG2() {
	ns.StoryPic(obj33, "HecubahPic")
	ns.SetDialog(obj33, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj33, ns.GetHost())
}
func HecubahEncounterSEG1() {
	if !(!flag21 && ns.IsCaller(ns.GetHost())) {
		goto LABEL1
	}
	flag21 = true
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(ns.Object("War03b:Wolf1"), true)
	ns.Frozen(ns.Object("War03b:Wolf2"), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, HecubahEncounterSEG2)
LABEL1:
	return
}
func ResetFistTrap(a1 int) {
	ns.ObjectOn(obj29[a1])
	ns.ObjectOn(obj30[a1])
	ns.MoveWaypoint(wp51, ns.GetObjectX(obj30[a1]), ns.GetObjectY(obj30[a1]))
	ns.AudioEvent(ns.TriggerReleased, wp51)
}
func DropFist(a1 int) {
	ns.CastSpellObjectLocation(ns.SPELL_FIST, obj30[a1], ns.GetObjectX(obj30[a1]), ns.GetObjectY(obj30[a1]))
	ns.ObjectOff(obj30[a1])
	ns.ObjectOff(obj29[a1])
	ns.FrameTimerWithArg(ivar31, a1, ResetFistTrap)
}
func ActivateFistTrap01() {
	DropFist(0)
}
func ActivateFistTrap02() {
	DropFist(1)
}
func ActivateFistTrap03() {
	DropFist(2)
}
func ActivateFistTrap04() {
	DropFist(3)
}
func ActivateFistTrap05() {
	DropFist(4)
}
func ActivateFistTrap06() {
	DropFist(5)
}
func ActivateFistTrap07() {
	DropFist(6)
}
func ActivateFistTrap08() {
	DropFist(7)
}
func ActivateFistTrap09() {
	DropFist(8)
}
func ActivateFistTrap10() {
	DropFist(9)
}
func ActivateFistTrap11() {
	DropFist(10)
}
func ActivateFistTrap12() {
	DropFist(11)
}
func ActivateFistTrap13() {
	DropFist(12)
}
func NoEnemys() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func SecretSFX() {
	ns.MoveWaypoint(wp52, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp52)
}
func FoundSecret01() {
	ns.ObjectOff(ns.GetTrigger())
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	SecretSFX()
}
func FoundSecret02() {
	ns.ObjectOff(ns.GetTrigger())
	ns.GiveXp(ns.GetHost(), 75)
	ns.PrintToAll("GeneralPrint:SecretFound")
	SecretSFX()
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
