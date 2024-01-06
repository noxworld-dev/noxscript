package con04b

import (
	"strconv"

	"github.com/noxworld-dev/noxscript/ns/v3"
)

var (
	gvar4  int
	gvar5  int
	gvar6  int
	gvar7  int
	gvar8  int
	gvar9  int
	gvar10 int
	gvar11 int
	gvar12 int
	gvar13 int
	gvar14 int
	gvar15 int
	flag16 bool
	flag17 bool
	flag18 bool
	flag19 bool
	flag20 bool
	ivar21 int
	ivar22 int
	gvar23 int
	gvar24 int
	gvar25 int
	obj26  [13]ns.ObjectID
	obj27  [13]ns.ObjectID
	ivar28 int
	obj29  [56]ns.ObjectID
	obj30  ns.ObjectID
	obj31  ns.ObjectID
	obj32  ns.ObjectID
	obj33  ns.ObjectID
	obj34  ns.ObjectID
	obj35  ns.ObjectID
	obj36  ns.ObjectID
	obj37  ns.ObjectID
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	gvar41 int
	gvar42 int
	gvar43 int
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	gvar46 ns.ObjectGroupID
	gvar47 ns.WallGroupID
	wp48   ns.WaypointID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	wp53   [56]ns.WaypointID
	wp54   ns.WaypointID
	gvar55 ns.WaypointGroupID
	gvar56 ns.WaypointGroupID
	gvar57 ns.WaypointGroupID
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	gvar7 = 3
	gvar8 = 0
	gvar9 = 1
	gvar10 = 2
	gvar11 = 3
	gvar12 = 4
	gvar13 = 5
	gvar14 = 6
	gvar15 = 7
	flag16 = false
	flag17 = true
	flag18 = false
	flag19 = false
	flag20 = false
	ivar21 = 0
	ivar22 = 0
	gvar23 = gvar9
	gvar24 = 0
	gvar25 = gvar4
	ivar28 = 120
}
func setUpFlameLever() {
	if !flag16 {
		goto LABEL1
	}
	return
LABEL1:
	flag16 = true
	gvar23 = gvar9
	ivar22 = 2
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL2
		}
		obj29[ivar21] = 0
		ivar21 += 1
	}
LABEL2:
	ns.FrameTimer(3, flameLever)
}
func flameLever() {
	var v0 int
	if !flag16 {
		goto LABEL1
	}
	v0 = gvar23
	if v0 == gvar9 {
		goto LABEL2
	}
	if v0 == gvar10 {
		goto LABEL3
	}
	if v0 == gvar11 {
		goto LABEL4
	}
	if v0 == gvar12 {
		goto LABEL5
	}
	if v0 == gvar13 {
		goto LABEL6
	}
	if v0 == gvar14 {
		goto LABEL7
	}
	if v0 == gvar15 {
		goto LABEL8
	}
	if v0 == gvar8 {
		goto LABEL9
	}
	goto LABEL10
LABEL2:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL11
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("SmallFlame", wp53[ivar21])
		ivar21 += 1
	}
LABEL11:
	gvar23 = gvar10
	ns.AudioEvent(ns.FireballCast, wp51)
	ns.AudioEvent(ns.FireballCast, wp52)
	ivar22 = 2
	goto LABEL10
LABEL3:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL13
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("MediumFlame", wp53[ivar21])
		ivar21 += 1
	}
LABEL13:
	gvar23 = gvar11
	ns.AudioEvent(ns.DemonBreath, wp51)
	ns.AudioEvent(ns.DemonBreath, wp52)
	ivar22 = 2
	goto LABEL10
LABEL4:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL15
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("Flame", wp53[ivar21])
		ivar21 += 1
	}
LABEL15:
	gvar23 = gvar12
	ivar22 = 2
	goto LABEL10
LABEL5:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL17
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("LargeFlame", wp53[ivar21])
		ivar21 += 1
	}
LABEL17:
	gvar23 = gvar13
	ivar22 = 150
	goto LABEL10
LABEL6:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL19
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("Flame", wp53[ivar21])
		ivar21 += 1
	}
LABEL19:
	gvar23 = gvar14
	ivar22 = 2
	goto LABEL10
LABEL7:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL21
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("MediumFlame", wp53[ivar21])
		ivar21 += 1
	}
LABEL21:
	gvar23 = gvar15
	ns.AudioEvent(ns.FireExtinguish, wp51)
	ns.AudioEvent(ns.FireExtinguish, wp52)
	ivar22 = 2
	goto LABEL10
LABEL8:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL23
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = ns.CreateObject("SmallFlame", wp53[ivar21])
		ivar21 += 1
	}
LABEL23:
	gvar23 = gvar8
	ivar22 = 2
	goto LABEL10
LABEL9:
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL25
		}
		ns.Delete(obj29[ivar21])
		obj29[ivar21] = 0
		ivar21 += 1
	}
LABEL25:
	gvar23 = gvar9
	flag16 = false
	goto LABEL10
LABEL10:
	ns.FrameTimer(ivar22, flameLever)
LABEL1:
	return
}
func skullHallOn() {
	if !flag17 {
		goto LABEL1
	}
	ns.CastSpellLocationLocation(ns.SPELL_FIREBALL, 3012, 2229, 3185, 2403)
	ns.CastSpellLocationLocation(ns.SPELL_FIREBALL, 2896, 2346, 3070, 2517)
	ns.FrameTimer(60, skullHallOn)
LABEL1:
	return
}
func skullHallOff() {
	flag17 = false
}
func Zombie1Rise() {
	var (
		v0 float32
		v1 float32
	)
	_ = v1
	_ = v0
	v0 = ns.GetObjectX(obj38)
	v1 = ns.GetObjectY(obj38)
	ns.RaiseZombie(obj38)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("ZombieAttackWP"))
	ns.LookAtObject(obj30, obj38)
}
func Zombie2Rise() {
	var (
		v0 float32
		v1 float32
	)
	_ = v1
	_ = v0
	v0 = ns.GetObjectX(obj39)
	v1 = ns.GetObjectY(obj39)
	ns.RaiseZombie(obj39)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("ZombieAttackWP"))
	ns.LookAtObject(obj30, obj39)
}
func Zombie3Rise() {
	var (
		v0 float32
		v1 float32
	)
	_ = v1
	_ = v0
	v0 = ns.GetObjectX(obj40)
	v1 = ns.GetObjectY(obj40)
	ns.RaiseZombie(obj40)
	ns.AudioEvent(ns.ZombieRecognize, ns.Waypoint("ZombieAttackWP"))
	ns.LookAtObject(obj30, obj40)
}
func ZombieAttackHecubah() {
	ns.WayPointGroupOn(gvar56)
	ns.WayPointGroupOff(gvar55)
	ns.HitLocation(obj38, ns.GetWaypointX(ns.Waypoint("ZombieAttackWP")), ns.GetWaypointY(ns.Waypoint("ZombieAttackWP")))
	ns.HitLocation(obj39, ns.GetWaypointX(ns.Waypoint("ZombieAttackWP")), ns.GetWaypointY(ns.Waypoint("ZombieAttackWP")))
	ns.HitLocation(obj40, ns.GetWaypointX(ns.Waypoint("ZombieAttackWP")), ns.GetWaypointY(ns.Waypoint("ZombieAttackWP")))
}
func HecubahShootZombie1() {
	ns.LookAtObject(obj30, obj38)
	ns.CastSpellObjectObject(ns.SPELL_FIREBALL, obj30, obj38)
	ns.FrameTimer(30, HecubahShootZombie2)
}
func HecubahShootZombie2() {
	ns.LookAtObject(obj30, obj39)
	ns.CastSpellObjectObject(ns.SPELL_FIREBALL, obj30, obj39)
	ns.FrameTimer(30, HecubahShootZombie3)
}
func HecubahShootZombie3() {
	ns.LookAtObject(obj30, obj40)
	ns.CastSpellObjectObject(ns.SPELL_FIREBALL, obj30, obj40)
	ns.FrameTimer(30, HecubahEncounterSEG3)
}
func HecubahEncounterSEG4() {
	ns.SetDialog(obj30, ns.NORMAL, HecubahDialogStart, HecubahDialogEnd)
	gvar25 = gvar7
	ns.StartDialog(obj30, ns.GetHost())
}
func HecubahEncounterSEG3() {
	ns.SetDialog(obj31, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	gvar25 = gvar6
	ns.StartDialog(obj31, ns.GetHost())
}
func hecGone() {
	if !ns.IsCaller(obj30) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj30, ns.GetWaypointX(wp50), ns.GetWaypointY(wp50))
	ns.ObjectOff(obj30)
LABEL1:
	return
}
func ReleasePlayer() {
	if !ns.IsCaller(obj30) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.LockDoor(obj36)
	ns.LockDoor(obj37)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar47)
	ns.NoWallSound(false)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj31, ns.NORMAL, necroTalkStart, necroTalkEnd)
	gvar24 = 5
	ns.StartDialog(obj31, ns.GetHost())
LABEL1:
	return
}
func necroDies() {
	ns.UnlockDoor(obj36)
	ns.UnlockDoor(obj37)
	ns.CancelDialog(obj31)
	ns.SetDialog(obj31, ns.NORMAL, necroTalkStart, necroTalkEnd)
	gvar24 = 4
	ns.StartDialog(obj31, ns.GetHost())
}
func necroTalkStart() {
	var v0 int
	v0 = gvar24
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
	ns.Frozen(obj31, true)
	ns.Frozen(obj38, true)
	ns.Frozen(obj40, true)
	goto LABEL6
LABEL2:
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroGloats")
	ns.Frozen(obj31, true)
	ns.Frozen(obj38, true)
	ns.Frozen(obj40, true)
	goto LABEL6
LABEL3:
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroThreatens2")
	ns.Frozen(obj31, true)
	ns.Frozen(obj38, true)
	ns.Frozen(obj40, true)
	goto LABEL6
LABEL4:
	ns.WayPointGroupOff(gvar56)
	ns.WayPointGroupOn(gvar57)
	ns.LookAtObject(obj31, ns.GetHost())
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroDies")
	ns.Frozen(obj31, true)
	ns.Frozen(obj38, true)
	ns.Frozen(obj40, true)
	goto LABEL6
LABEL5:
	ns.AggressionLevel(obj31, 0.83)
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroThreatens")
	ns.LookAtObject(obj31, ns.GetHost())
	ns.Frozen(obj31, true)
	ns.Frozen(obj38, true)
	ns.Frozen(obj40, true)
	goto LABEL6
LABEL6:
	return
}
func necroTalkEnd() {
	ns.CancelDialog(obj31)
	ns.Frozen(obj31, false)
	ns.Frozen(obj38, false)
	ns.Frozen(obj40, false)
	if gvar24 != 4 {
		goto LABEL1
	}
	ns.JournalEdit(ns.GetHost(), "Chapter4SearchCrypts", 4)
	ns.JournalEntry(ns.GetHost(), "Chapter4Escape", 2)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL1:
	return
}
func KillZombie1() {
	ns.Damage(obj38, 0, 30, 0)
	ns.FrameTimer(5, KillZombie2)
}
func KillZombie2() {
	ns.Damage(obj39, 0, 30, 0)
	ns.FrameTimer(5, KillZombie3)
}
func KillZombie3() {
	ns.Damage(obj40, 0, 30, 0)
	ns.FrameTimer(5, KeepZombiesDead)
}
func KeepZombiesDead() {
	ns.ZombieStayDown(obj38)
	ns.ZombieStayDown(obj39)
	ns.ZombieStayDown(obj40)
}
func ZombieInjured() {
	ns.Damage(ns.GetTrigger(), 0, 100, 1)
}
func InitializeFistTraps() {
	obj27[0] = ns.Object("FistTrap01")
	obj27[1] = ns.Object("FistTrap02")
	obj27[2] = ns.Object("FistTrap03")
	obj27[3] = ns.Object("FistTrap04")
	obj27[4] = ns.Object("FistTrap05")
	obj27[5] = ns.Object("FistTrap06")
	obj27[6] = ns.Object("FistTrap07")
	obj27[7] = ns.Object("FistTrap08")
	obj27[8] = ns.Object("FistTrap09")
	obj27[9] = ns.Object("FistTrap10")
	obj27[10] = ns.Object("FistTrap11")
	obj27[11] = ns.Object("FistTrap12")
	obj27[12] = ns.Object("FistTrap13")
	obj26[0] = ns.Object("FistTrapLight01")
	obj26[1] = ns.Object("FistTrapLight02")
	obj26[2] = ns.Object("FistTrapLight03")
	obj26[3] = ns.Object("FistTrapLight04")
	obj26[4] = ns.Object("FistTrapLight05")
	obj26[5] = ns.Object("FistTrapLight06")
	obj26[6] = ns.Object("FistTrapLight07")
	obj26[7] = ns.Object("FistTrapLight08")
	obj26[8] = ns.Object("FistTrapLight09")
	obj26[9] = ns.Object("FistTrapLight10")
	obj26[10] = ns.Object("FistTrapLight11")
	obj26[11] = ns.Object("FistTrapLight12")
	obj26[12] = ns.Object("FistTrapLight13")
}
func MapInitialize() {
	ivar21 = 0
	for {
		if !(ivar21 < 56) {
			goto LABEL1
		}
		wp53[ivar21] = ns.Waypoint("FlameWay" + strconv.Itoa(ivar21+1))
		ivar21 += 1
	}
LABEL1:
	obj30 = ns.Object("Hecubah")
	obj31 = ns.Object("Necromancer")
	obj32 = ns.Object("EntranceGate1")
	obj33 = ns.Object("EntranceGate2")
	obj34 = ns.Object("EntranceGate3")
	obj35 = ns.Object("EntranceGate4")
	obj36 = ns.Object("PitGate1")
	obj37 = ns.Object("PitGate2")
	obj38 = ns.Object("DeadZombie1")
	obj39 = ns.Object("DeadZombie2")
	obj40 = ns.Object("DeadZombie3")
	obj44 = ns.Object("FlameLever1")
	obj45 = ns.Object("FlameLever2")
	gvar46 = ns.ObjectGroup("Zombies")
	gvar47 = ns.WallGroup("InvisibleBlockWall")
	wp48 = ns.Waypoint("SecretAudioOrigin")
	wp51 = ns.Waypoint("FlameSound1")
	wp52 = ns.Waypoint("FlameSound2")
	wp49 = ns.Waypoint("AudioOrigin")
	wp54 = ns.Waypoint("WP")
	wp50 = ns.Waypoint("HecubahStorage")
	gvar55 = ns.WaypointGroup("BeforeSetpieceWay")
	gvar56 = ns.WaypointGroup("SetpieceWay")
	gvar57 = ns.WaypointGroup("AfterSetpieceWay")
	ns.WayPointGroupOff(gvar56)
	ns.WayPointGroupOff(gvar57)
	ns.StoryPic(obj31, "NecromancerPic")
	InitializeFistTraps()
	ns.FrameTimer(1, KillZombie1)
}
func PlayerDeath() {
	ns.DeathScreen(4)
}
func ToggleSkullHall() {
	if !flag17 {
		goto LABEL1
	}
	flag17 = false
	goto LABEL2
LABEL1:
	flag17 = true
	skullHallOn()
LABEL2:
	return
}
func HecubahDialogStart() {
	var v0 int
	v0 = gvar25
	if v0 == gvar4 {
		goto LABEL1
	}
	if v0 == gvar5 {
		goto LABEL2
	}
	if v0 == gvar6 {
		goto LABEL3
	}
	if v0 == gvar7 {
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
	ns.LookAtObject(obj31, ns.GetHost())
	ns.TellStory(ns.DemonRecognize, "Con04a:NecroWarnsHec")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj30, ns.GetHost())
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
	v5 = gvar25
	if v5 == gvar4 {
		goto LABEL1
	}
	if v5 == gvar5 {
		goto LABEL2
	}
	if v5 == gvar6 {
		goto LABEL3
	}
	if v5 == gvar7 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.FrameTimer(20, Zombie2Rise)
	ns.FrameTimer(45, Zombie1Rise)
	ns.FrameTimer(60, Zombie3Rise)
	ns.FrameTimer(60, ZombieAttackHecubah)
	gvar25 = gvar5
	ns.CancelDialog(obj30)
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj30, obj39)
	gvar25 = gvar6
	ns.CancelDialog(obj30)
	ns.AggressionLevel(obj38, 0)
	ns.AggressionLevel(obj39, 0)
	ns.AggressionLevel(obj40, 0)
	ns.PushObject(obj38, 40, ns.GetObjectX(obj30), ns.GetObjectY(obj30))
	ns.PushObject(obj39, 40, ns.GetObjectX(obj30), ns.GetObjectY(obj30))
	ns.PushObject(obj40, 40, ns.GetObjectX(obj30), ns.GetObjectY(obj30))
	ns.SetOwner(obj36, obj38)
	ns.SetOwner(obj36, obj39)
	ns.SetOwner(obj36, obj40)
	ns.SetCallback(obj38, 7, ZombieInjured)
	ns.SetCallback(obj39, 7, ZombieInjured)
	ns.SetCallback(obj40, 7, ZombieInjured)
	ns.FrameTimer(20, HecubahShootZombie1)
	goto LABEL5
LABEL3:
	gvar25 = gvar7
	ns.CancelDialog(obj31)
	ns.FrameTimer(5, HecubahEncounterSEG4)
	goto LABEL5
LABEL4:
	v4 = ns.Waypoint("WP")
	v0 = ns.GetObjectX(obj30)
	v1 = ns.GetObjectY(obj30)
	ns.CreatureIdle(obj31)
	ns.UnlockDoor(obj32)
	ns.UnlockDoor(obj33)
	ns.LockDoor(obj34)
	ns.LockDoor(obj35)
	ns.ObjectOn(obj30)
	ns.Move(obj30, v4)
	flag19 = true
	ns.CancelDialog(obj30)
	goto LABEL5
LABEL5:
	return
}
func AttackSuccessful() {
	ns.SetDialog(obj30, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj30, ns.GetHost())
}
func HecubahEncounterSEG2() {
	ns.StoryPic(obj30, "HecubahPic")
	ns.SetDialog(obj30, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj30, ns.GetHost())
}
func HecubahEncounterSEG1() {
	if !(!flag18 && ns.IsCaller(ns.GetHost())) {
		goto LABEL1
	}
	flag18 = true
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.FrameTimer(60, HecubahEncounterSEG2)
LABEL1:
	return
}
func ResetFistTrap(a1 int) {
	ns.ObjectOn(obj26[a1])
	ns.ObjectOn(obj27[a1])
	ns.MoveWaypoint(wp49, ns.GetObjectX(obj27[a1]), ns.GetObjectY(obj27[a1]))
	ns.AudioEvent(ns.TriggerReleased, wp49)
}
func DropFist(a1 int) {
	ns.CastSpellObjectLocation(ns.SPELL_FIST, obj27[a1], ns.GetObjectX(obj27[a1]), ns.GetObjectY(obj27[a1]))
	ns.ObjectOff(obj27[a1])
	ns.ObjectOff(obj26[a1])
	ns.FrameTimerWithArg(ivar28, a1, ResetFistTrap)
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
	ns.MoveWaypoint(wp48, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp48)
}
func Secret01Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	SecretSFX()
}
func Secret02Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.ObjectOn(ns.Object("SecretLord"))
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
