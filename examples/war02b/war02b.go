package war02b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	wp4    ns.WaypointID
	ivar5  int
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	obj11  ns.ObjectID
	obj12  ns.ObjectID
	obj13  ns.ObjectID
	obj14  ns.ObjectID
	obj15  ns.ObjectID
	wp16   ns.WaypointID
	wp17   ns.WaypointID
	wp18   ns.WaypointID
	wp19   ns.WaypointID
	wp20   ns.WaypointID
	wp21   ns.WaypointID
	wp22   ns.WaypointID
	wp23   ns.WaypointID
	wp24   ns.WaypointID
	wp25   ns.WaypointID
	wp26   ns.WaypointID
	wp27   ns.WaypointID
	wp28   ns.WaypointID
	wp29   ns.WaypointID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
	wp32   ns.WaypointID
	gvar33 int
	gvar34 int
	gvar35 int
	gvar36 int
	gvar37 int
	obj38  ns.ObjectID
	wp39   ns.WaypointID
	flag40 bool
	gvar41 int
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	gvar46 ns.ObjectGroupID
	gvar47 ns.ObjectGroupID
	gvar48 ns.ObjectGroupID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	wp53   ns.WaypointID
	flag54 bool
	flag55 bool
	flag56 bool
	flag57 bool
	flag58 bool
	obj59  ns.ObjectID
	obj60  ns.ObjectID
	obj61  ns.ObjectID
	obj62  ns.ObjectID
	obj63  ns.ObjectID
	obj64  ns.ObjectID
	obj65  ns.ObjectID
	obj66  ns.ObjectID
	obj67  ns.ObjectID
	obj68  ns.ObjectID
	obj69  ns.ObjectID
	obj70  ns.ObjectID
	obj71  ns.ObjectID
	obj72  ns.ObjectID
	obj73  ns.ObjectID
	wp74   ns.WaypointID
	wp75   ns.WaypointID
	wp76   ns.WaypointID
	wp77   ns.WaypointID
	wp78   ns.WaypointID
	wp79   ns.WaypointID
	wp80   ns.WaypointID
	wp81   ns.WaypointID
	wp82   ns.WaypointID
	wp83   ns.WaypointID
	wp84   ns.WaypointID
	wp85   ns.WaypointID
	wp86   ns.WaypointID
	wp87   ns.WaypointID
	wp88   ns.WaypointID
	wp89   ns.WaypointID
	wp90   ns.WaypointID
	wp91   ns.WaypointID
)

func init() {
	ivar5 = 60
	gvar33 = 0
	gvar34 = 1
	gvar35 = 2
	gvar36 = 3
	gvar37 = 4
	flag40 = false
	gvar41 = gvar33
	flag54 = false
	flag55 = false
	flag56 = false
	flag57 = false
	flag58 = false
}
func PlayerDeath() {
	ns.DeathScreen(2)
}
func MapShutdown() {
	ns.SetQuestStatus(2, "War02A:EnteredGauntlet")
}
func BlockSecret() {
	ns.MoveWaypoint(wp4, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp4)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func OpenBlockSecretExit() {
	ns.NoWallSound(false)
	ns.WallOpen(ns.Wall(77, 157))
}
func ResetSquishTrapSEG1() {
	ns.Move(obj8, wp25)
	ns.Move(obj9, wp26)
	ns.Move(obj10, wp27)
	ns.Move(obj11, wp28)
	ns.Move(obj12, wp29)
	ns.Move(obj13, wp30)
	ns.Move(obj14, wp31)
	ns.Move(obj15, wp32)
	ns.AudioEvent(ns.BoulderMove, wp16)
	ns.FrameTimer(20, ResetSquishTrapSEG2)
}
func ResetSquishTrapSEG2() {
	ns.ObjectOn(obj6)
	ns.ObjectOn(obj7)
}
func InitializeSquishBlocks() {
	obj6 = ns.Object("SquishTrigger")
	obj7 = ns.Object("BoomTrigger")
	obj8 = ns.Object("SquishBlock01")
	obj9 = ns.Object("SquishBlock02")
	obj10 = ns.Object("SquishBlock03")
	obj11 = ns.Object("SquishBlock04")
	obj12 = ns.Object("SquishBlock05")
	obj13 = ns.Object("SquishBlock06")
	obj14 = ns.Object("SquishBlock07")
	obj15 = ns.Object("SquishBlock08")
	wp16 = ns.Waypoint("SquishAudioOrigin")
	wp17 = ns.Waypoint("SquishWP01")
	wp18 = ns.Waypoint("SquishWP02")
	wp19 = ns.Waypoint("SquishWP03")
	wp20 = ns.Waypoint("SquishWP04")
	wp21 = ns.Waypoint("SquishWP05")
	wp22 = ns.Waypoint("SquishWP06")
	wp23 = ns.Waypoint("SquishWP07")
	wp24 = ns.Waypoint("SquishWP08")
	wp25 = ns.Waypoint("SquishWPHome01")
	wp26 = ns.Waypoint("SquishWPHome02")
	wp27 = ns.Waypoint("SquishWPHome03")
	wp28 = ns.Waypoint("SquishWPHome04")
	wp29 = ns.Waypoint("SquishWPHome05")
	wp30 = ns.Waypoint("SquishWPHome06")
	wp31 = ns.Waypoint("SquishWPHome07")
	wp32 = ns.Waypoint("SquishWPHome08")
}
func Squish() {
	ns.ObjectOff(obj6)
	ns.Move(obj8, wp17)
	ns.Move(obj9, wp18)
	ns.Move(obj10, wp19)
	ns.Move(obj11, wp20)
	ns.Move(obj12, wp21)
	ns.Move(obj13, wp22)
	ns.Move(obj14, wp23)
	ns.Move(obj15, wp24)
	ns.AudioEvent(ns.BoulderMove, wp16)
	ns.FrameTimer(ivar5, ResetSquishTrapSEG1)
}
func SquishBoom() {
	if ns.IsCaller(obj10) {
		goto LABEL1
	}
	return
LABEL1:
	ns.ObjectOff(obj7)
	ns.AudioEvent(ns.HammerMissing, wp16)
	ns.Effect(ns.JIGGLE, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 10, 0)
}
func GearhartDialogStart() {
	var v0 int
	ns.LookAtObject(obj38, ns.GetHost())
	ns.ObjectGroupOff(gvar48)
	v0 = gvar41
	if v0 == gvar33 {
		goto LABEL1
	}
	if v0 == gvar34 {
		goto LABEL2
	}
	if v0 == gvar35 {
		goto LABEL3
	}
	if v0 == gvar36 {
		goto LABEL4
	}
	if v0 == gvar37 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "War02a:GearhartTalk02")
	goto LABEL6
LABEL2:
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "War02a:GearhartTalk03")
	goto LABEL6
LABEL3:
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "War02a:GearhartTalk04")
	goto LABEL6
LABEL4:
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "War02a:GearhartTalk05")
	goto LABEL6
LABEL5:
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "War02a:GearhartTalk06")
	goto LABEL6
LABEL6:
	return
}
func GearhartDialogEnd() {
	var v0 int
	ns.ObjectGroupOn(gvar48)
	v0 = gvar41
	if v0 == gvar33 {
		goto LABEL1
	}
	if v0 == gvar34 {
		goto LABEL2
	}
	if v0 == gvar35 {
		goto LABEL3
	}
	if v0 == gvar36 {
		goto LABEL4
	}
	if v0 == gvar37 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.CancelDialog(obj38)
	ns.CreatureGuard(obj38, ns.GetWaypointX(wp50), ns.GetWaypointY(wp50), ns.GetWaypointX(wp52), ns.GetWaypointY(wp52), 0)
	ns.Pickup(ns.GetHost(), ns.Object("SilverKey01"))
	ns.PrintToAll("GeneralPrint:GainedKey")
	ns.AudioEvent(ns.KeyDrop, wp39)
	gvar41 = gvar34
	goto LABEL6
LABEL2:
	flag55 = true
	ns.CreatureGuard(obj38, ns.GetWaypointX(wp53), ns.GetWaypointY(wp53), ns.GetWaypointX(wp52), ns.GetWaypointY(wp52), 0)
	gvar41 = gvar35
	goto LABEL6
LABEL3:
	goto LABEL6
LABEL4:
	goto LABEL6
LABEL5:
	flag56 = true
	ns.CreatureFollow(obj38, ns.GetHost())
	ns.CancelDialog(obj38)
	ns.ObjectOn(obj45)
	goto LABEL6
LABEL6:
	return
}
func GearRoomMusic() {
	ns.Music(26, 100)
}
func InitializeSwitchPuzzle() {
	obj42 = ns.Object("ExitElevator")
	obj43 = ns.Object("PlayerSwitch")
	obj44 = ns.Object("GearhartSwitch")
	obj45 = ns.Object("ExitTrigger")
	gvar46 = ns.ObjectGroup("ElevatorSwitches")
	gvar47 = ns.ObjectGroup("ElevatorGears")
	wp49 = ns.Waypoint("FenceWP")
	wp50 = ns.Waypoint("GateWP")
	wp51 = ns.Waypoint("GearRoomEntranceWP")
	wp52 = ns.Waypoint("PlayerSwitchWP")
	wp53 = ns.Waypoint("GearhartSwitchWP")
	wp39 = ns.Waypoint("SilverKeyWay")
}
func SetGearhartDialog() {
	if !ns.IsCaller(obj38) {
		goto LABEL1
	}
	ns.StoryPic(obj38, "GearhartPic")
	ns.SetDialog(obj38, "FALSE", GearhartDialogStart, GearhartDialogEnd)
	ns.ObjectOff(ns.GetTrigger())
LABEL1:
	return
}
func GearhartCallForHelp() {
	if !flag54 {
		goto LABEL1
	}
	return
LABEL1:
	flag54 = true
	ns.Chat(obj38, "War02a:GearhartTalk01")
	ns.Move(obj38, wp49)
	ns.StoryPic(obj38, "GearhartPic")
	ns.SetDialog(obj38, "FALSE", GearhartDialogStart, GearhartDialogEnd)
}
func CheckForBothButtons() {
	if flag55 {
		goto LABEL1
	}
	return
LABEL1:
	if !(flag58 && flag57) {
		goto LABEL2
	}
	ns.Music(16, 100)
	ns.ObjectOn(obj42)
	ns.ObjectGroupOn(gvar47)
	ns.ObjectGroupOff(gvar46)
	gvar41 = gvar37
	ns.StartDialog(obj38, ns.GetHost())
LABEL2:
	return
}
func PlayerSwitchActivate() {
	flag58 = true
	CheckForBothButtons()
}
func PlayerSwitchDeactivate() {
	flag58 = false
}
func GearhartSwitchActivate() {
	if !ns.IsCaller(obj38) {
		goto LABEL1
	}
	ns.Chat(obj38, "War02a:GearhartTalk05")
LABEL1:
	flag57 = true
	CheckForBothButtons()
}
func GearhartSwitchDeactivate() {
	flag57 = false
}
func LeaveGearRoom() {
	if !ns.IsCaller(obj38) {
		goto LABEL1
	}
	ns.CreatureGuard(obj38, ns.GetWaypointX(wp51), ns.GetWaypointY(wp51), 0, 0, 0)
LABEL1:
	return
}
func EnterGearRoom() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !flag56 {
		goto LABEL1
	}
	ns.CreatureFollow(obj38, ns.GetHost())
LABEL1:
	return
}
func gotoExit() {
	if flag40 {
		goto LABEL1
	}
	ns.MoveObject(ns.GetHost(), 494, 610)
LABEL1:
	return
}
func ExitSewers() {
	ns.Frozen(ns.GetHost(), true)
	ns.Blind()
	ns.FrameTimer(60, gotoExit)
}
func Exit() {
	flag40 = true
	ns.Frozen(ns.GetHost(), false)
}
func sweepsOn123() {
	if ns.IsObjectOn(obj59) {
		goto LABEL1
	}
	ns.ObjectOn(obj59)
LABEL1:
	ns.MoveWaypoint(wp4, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BoulderMove, wp4)
	ns.ObjectOn(obj60)
	ns.ObjectOn(obj61)
	ns.ObjectOn(obj62)
	ns.Move(obj60, wp75)
	ns.Move(obj61, wp77)
	ns.Move(obj62, wp79)
}
func sweepsOn456() {
	ns.MoveWaypoint(wp4, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BoulderMove, wp4)
	ns.ObjectOn(obj63)
	ns.ObjectOn(obj64)
	ns.ObjectOn(obj65)
	ns.Move(obj63, wp80)
	ns.Move(obj64, wp82)
	ns.Move(obj65, wp84)
}
func sweepsOn789() {
	ns.MoveWaypoint(wp4, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.BoulderMove, wp4)
	ns.ObjectOn(obj66)
	ns.ObjectOn(obj67)
	ns.ObjectOn(obj68)
	ns.Move(obj66, wp87)
	ns.Move(obj67, wp89)
	ns.Move(obj68, wp91)
}
func sweepsOn10thru14() {
	ns.ObjectOn(obj69)
	ns.ObjectOn(obj70)
	ns.ObjectOn(obj71)
	ns.ObjectOn(obj72)
	ns.ObjectOn(obj73)
}
func sweepsOff123() {
	ns.Move(obj60, wp74)
	ns.Move(obj61, wp76)
	ns.Move(obj62, wp78)
}
func sweepsOff456() {
	ns.Move(obj63, wp81)
	ns.Move(obj64, wp83)
	ns.Move(obj65, wp85)
}
func sweepsOff789() {
	ns.Move(obj66, wp86)
	ns.Move(obj67, wp88)
	ns.Move(obj68, wp90)
}
func sweepsOff10thru14() {
	ns.ObjectOff(obj69)
	ns.ObjectOff(obj70)
	ns.ObjectOff(obj71)
	ns.ObjectOff(obj72)
	ns.ObjectOff(obj73)
}
func MapInitialize() {
	obj38 = ns.Object("F2Gearhart")
	obj60 = ns.Object("SweeperMover1")
	obj61 = ns.Object("SweeperMover2")
	obj62 = ns.Object("SweeperMover3")
	obj63 = ns.Object("SweeperMover4")
	obj64 = ns.Object("SweeperMover5")
	obj65 = ns.Object("SweeperMover6")
	obj66 = ns.Object("SweeperMover7")
	obj67 = ns.Object("SweeperMover8")
	obj68 = ns.Object("SweeperMover9")
	obj69 = ns.Object("SweeperMover10")
	obj70 = ns.Object("SweeperMover11")
	obj71 = ns.Object("SweeperMover12")
	obj72 = ns.Object("SweeperMover13")
	obj73 = ns.Object("SweeperMover14")
	gvar48 = ns.ObjectGroup("EveryMonsterOnMap")
	wp4 = ns.Waypoint("AudioOrigin")
	wp74 = ns.Waypoint("SweeperWay1a")
	wp75 = ns.Waypoint("SweeperWay1b")
	wp76 = ns.Waypoint("SweeperWay2a")
	wp77 = ns.Waypoint("SweeperWay2b")
	wp78 = ns.Waypoint("SweeperWay3a")
	wp79 = ns.Waypoint("SweeperWay3b")
	wp80 = ns.Waypoint("SweeperWay4")
	wp81 = ns.Waypoint("SweeperWay4a")
	wp82 = ns.Waypoint("SweeperWay5")
	wp83 = ns.Waypoint("SweeperWay5a")
	wp84 = ns.Waypoint("SweeperWay6")
	wp85 = ns.Waypoint("SweeperWay6a")
	wp86 = ns.Waypoint("SweeperWay7")
	wp87 = ns.Waypoint("SweeperWay7a")
	wp88 = ns.Waypoint("SweeperWay8")
	wp89 = ns.Waypoint("SweeperWay8a")
	wp90 = ns.Waypoint("SweeperWay9")
	wp91 = ns.Waypoint("SweeperWay9a")
	obj59 = ns.Object("Spider01")
	InitializeSwitchPuzzle()
	InitializeSquishBlocks()
}
func MapEntry() {
	ns.UnBlind()
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapShutdown":
		MapShutdown()
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	}
}
