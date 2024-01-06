package wiz03b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	gvar5  ns.ObjectGroupID
	wp6    ns.WaypointID
	gvar7  ns.ObjectGroupID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	gvar14 ns.ObjectGroupID
	gvar15 ns.ObjectGroupID
	gvar16 ns.WallGroupID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	gvar27 ns.ObjectGroupID
	wp28   ns.WaypointID
	wp29   ns.WaypointID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
	wp32   ns.WaypointID
	wp33   ns.WaypointID
	wp34   ns.WaypointID
	wp35   ns.WaypointID
	wp36   ns.WaypointID
	wp37   ns.WaypointID
	wp38   ns.WaypointID
	obj39  ns.ObjectID
	obj40  ns.ObjectID
	gvar41 ns.ObjectGroupID
	gvar42 ns.ObjectGroupID
	gvar43 ns.ObjectGroupID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	gvar46 ns.ObjectGroupID
	obj47  ns.ObjectID
	gvar48 ns.WallGroupID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	gvar52 ns.WallGroupID
	obj53  ns.ObjectID
	wp54   ns.WaypointID
	wp55   ns.WaypointID
	wp56   ns.WaypointID
	gvar57 ns.ObjectGroupID
	ivar58 int
)

func init() {
	ivar58 = 0
}
func PlayerDeath() {
	ns.DeathScreen(3)
}
func TauntBegin() {
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar16)
	ns.NoWallSound(false)
	ns.ObjectGroupOn(gvar5)
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar15)
	ns.WideScreen(true)
	ns.Move(obj4, wp6)
	ns.ObjectOn(obj12)
	ns.ObjectGroupOff(gvar7)
}
func Taunt() {
	ns.StartDialog(obj4, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Wiz03b:RogueLeader")
	ns.ObjectOff(obj12)
}
func Unfreeze() {
	ns.Music(27, 100)
	ns.UnlockDoor(obj13)
	ns.CancelDialog(obj4)
	ns.ClearOwner(obj4)
	ns.ClearOwner(obj8)
	ns.ClearOwner(obj9)
	ns.ClearOwner(obj10)
	ns.ClearOwner(obj11)
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar15)
	ns.GroupAggressionLevel(gvar14, 0.83)
}
func GilgoreTalkEnd() {
	ns.WideScreen(false)
	ns.FrameTimer(60, Unfreeze)
}
func GilgoreTalkStart() {
}
func DisableElevator() {
	if ns.GetElevatorStatus(obj39) != 2 {
		goto LABEL1
	}
	ns.ObjectOff(obj39)
	goto LABEL2
LABEL1:
	ns.FrameTimer(1, DisableElevator)
LABEL2:
	return
}
func BanditDeath() {
	ivar58 += 1
	if !(ivar58 >= 8) {
		goto LABEL1
	}
	ns.Music(16, 100)
LABEL1:
	return
}
func PlaySubMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(16, 100)
LABEL1:
	return
}
func PlayActionMusic() {
	ns.Music(27, 100)
}
func MapInitialize() {
	obj17 = ns.Object("Swordsman01")
	obj18 = ns.Object("Swordsman02")
	obj19 = ns.Object("Swordsman03")
	obj20 = ns.Object("Swordsman04")
	obj21 = ns.Object("Swordsman05")
	obj22 = ns.Object("Archer01")
	obj23 = ns.Object("Archer02")
	obj24 = ns.Object("Archer03")
	obj25 = ns.Object("Archer04")
	obj26 = ns.Object("Archer05")
	wp28 = ns.Waypoint("AmbushWP01")
	wp29 = ns.Waypoint("AmbushWP02")
	wp30 = ns.Waypoint("AmbushWP03")
	wp31 = ns.Waypoint("AmbushWP04")
	wp32 = ns.Waypoint("AmbushWP05")
	wp33 = ns.Waypoint("AmbushWP06")
	wp34 = ns.Waypoint("AmbushWP07")
	wp35 = ns.Waypoint("AmbushWP08")
	wp36 = ns.Waypoint("AmbushWP09")
	wp37 = ns.Waypoint("AmbushWP10")
	wp38 = ns.Waypoint("AmbushLook")
	obj39 = ns.Object("AmbushElevator")
	obj40 = ns.Object("AmbushElevator2")
	gvar41 = ns.ObjectGroup("AmbushTrigGroup")
	gvar27 = ns.ObjectGroup("AmbushMonsterGroup")
	gvar42 = ns.ObjectGroup("WaspGroup")
	ns.GroupWander(gvar42)
	gvar43 = ns.ObjectGroup("Wasp2Group")
	ns.GroupWander(gvar43)
	gvar46 = ns.ObjectGroup("BoulderTrigGroup")
	gvar48 = ns.WallGroup("RogueWallGroup")
	wp49 = ns.Waypoint("BoulderWP")
	obj45 = ns.Object("Boulder")
	obj47 = ns.Object("RogueDoor")
	wp50 = ns.Waypoint("IronBlockWP")
	wp51 = ns.Waypoint("SoundWP")
	gvar52 = ns.WallGroup("RogueSecretWallGroup")
	DisableElevator()
	obj53 = ns.Object("Wiz03b:AmuletofTeleportation")
	obj4 = ns.Object("Gilgore")
	gvar5 = ns.ObjectGroup("GilgoreGroup")
	ns.StoryPic(obj4, "GilgorePic")
	ns.SetDialog(obj4, ns.NORMAL, GilgoreTalkStart, GilgoreTalkEnd)
	ns.GroupSetOwner(ns.GetHost(), gvar5)
	ns.SetOwner(ns.GetHost(), obj44)
	wp6 = ns.Waypoint("TauntWP")
	gvar7 = ns.ObjectGroup("GilgoreTrigGroup")
	obj8 = ns.Object("LeadRogue1")
	obj9 = ns.Object("LeadRogue2")
	obj10 = ns.Object("LeadRogue3")
	obj11 = ns.Object("LeadRogue4")
	obj12 = ns.Object("TauntTrig")
	gvar15 = ns.ObjectGroup("BackDoorGroup")
	gvar14 = ns.ObjectGroup("RogueGroup1")
	obj13 = ns.Object("RogueDoor2")
	ns.LockDoor(obj13)
	wp54 = ns.Waypoint("Wiz03aExitWP")
	wp55 = ns.Waypoint("Wiz03cExitWP")
	gvar57 = ns.ObjectGroup("ExitTrigGroup")
	wp56 = ns.Waypoint("Secret01WP")
	gvar16 = ns.WallGroup("InvisibleWallGroup")
}
func MapEntry() {
	ns.UnBlind()
	ns.ObjectGroupOn(gvar57)
}
func Ambush() {
	var (
		v0  float32
		v1  float32
		v2  float32
		v3  float32
		v4  float32
		v5  float32
		v6  float32
		v7  float32
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
		v16 float32
		v17 float32
		v18 float32
		v19 float32
		v20 float32
		v21 float32
	)
	v0 = ns.GetWaypointX(wp28)
	v1 = ns.GetWaypointY(wp28)
	v2 = ns.GetWaypointX(wp29)
	v3 = ns.GetWaypointY(wp29)
	v4 = ns.GetWaypointX(wp30)
	v5 = ns.GetWaypointY(wp30)
	v6 = ns.GetWaypointX(wp31)
	v7 = ns.GetWaypointY(wp31)
	v8 = ns.GetWaypointX(wp32)
	v9 = ns.GetWaypointY(wp32)
	v10 = ns.GetWaypointX(wp33)
	v11 = ns.GetWaypointY(wp33)
	v12 = ns.GetWaypointX(wp34)
	v13 = ns.GetWaypointY(wp34)
	v14 = ns.GetWaypointX(wp35)
	v15 = ns.GetWaypointY(wp35)
	v16 = ns.GetWaypointX(wp36)
	v17 = ns.GetWaypointY(wp36)
	v18 = ns.GetWaypointX(wp37)
	v19 = ns.GetWaypointY(wp37)
	v20 = ns.GetWaypointX(wp38)
	v21 = ns.GetWaypointY(wp38)
	ns.ObjectGroupOn(gvar27)
	ns.MoveObject(obj17, v0, v1)
	ns.MoveObject(obj18, v2, v3)
	ns.MoveObject(obj19, v4, v5)
	ns.MoveObject(obj20, v16, v17)
	ns.MoveObject(obj21, v18, v19)
	ns.MoveObject(obj22, v6, v7)
	ns.CreatureGuard(obj22, v6, v7, v20, v21, 500)
	ns.MoveObject(obj23, v8, v9)
	ns.CreatureGuard(obj23, v8, v9, v20, v21, 500)
	ns.MoveObject(obj24, v10, v11)
	ns.CreatureGuard(obj24, v10, v11, v20, v21, 500)
	ns.MoveObject(obj25, v12, v13)
	ns.CreatureGuard(obj25, v12, v13, v20, v21, 500)
	ns.MoveObject(obj26, v14, v15)
	ns.CreatureGuard(obj26, v14, v15, v20, v21, 500)
	ns.ObjectGroupOff(gvar41)
	ns.ObjectOn(obj39)
	ns.ObjectGroupOn(gvar15)
}
func BanditAmbush() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Chat(obj17, "Wiz03b:BanditAmbush")
	PlayActionMusic()
LABEL1:
	return
}
func BoulderGo() {
	ns.AudioEvent(ns.BoulderRoll, wp51)
	ns.WallGroupOpen(gvar48)
	ns.Move(obj45, wp49)
	ns.ObjectGroupOff(gvar46)
	ns.LockDoor(obj47)
}
func IronBlock() {
	ns.CreateObject("ExtentBoxLarge", wp50)
}
func BlockFilter() int {
	if !ns.IsCaller(obj45) {
		goto LABEL1
	}
	return 1
	goto LABEL2
LABEL1:
	return 0
LABEL2:
	return 0
}
func OpenRogueSecretWallGroup() {
	ns.WallGroupOpen(gvar52)
}
func MonsterGoHome() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MonsterGoHomeShop() {
	if !ns.HasClass(ns.GetCaller(), ns.MONSTER) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func PlayerWiz03aExit() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp54)
	v1 = ns.GetWaypointY(wp54)
	ns.MoveObject(ns.GetHost(), v0, v1)
}
func FreezePlayer() {
	ns.Frozen(ns.GetHost(), true)
}
func GotoWiz03a() {
	ns.ObjectGroupOff(gvar57)
	ns.Blind()
	ns.FrameTimer(30, FreezePlayer)
	ns.FrameTimer(60, PlayerWiz03aExit)
}
func PlayerWiz03cExit() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp55)
	v1 = ns.GetWaypointY(wp55)
	ns.MoveObject(ns.GetHost(), v0, v1)
}
func GotoWiz03c() {
	ns.ObjectGroupOff(gvar57)
	ns.Blind()
	ns.FrameTimer(30, FreezePlayer)
	ns.FrameTimer(60, PlayerWiz03cExit)
}
func Secret01XP() {
	ns.AudioEvent(ns.SecretFound, wp56)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func FoundAmulet() {
	ns.GiveXp(ns.GetHost(), 100)
	ns.PrintToAll("Con05B.scr:FoundAmulet")
}
func PlayWanderMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(22, 100)
LABEL1:
	return
}
func PlayDunMirMusic() {
	ns.Music(16, 100)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	}
}
