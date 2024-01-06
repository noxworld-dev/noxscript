package wiz06c

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	gvar11 ns.ObjectGroupID
	ivar12 int
	ivar13 int
	ivar14 int
	wp15   ns.WaypointID
	flag16 bool
	flag17 bool
	flag18 bool
	flag19 bool
	flag20 bool
	gvar21 int
	obj22  ns.ObjectID
	obj23  ns.ObjectID
	obj24  ns.ObjectID
	obj25  ns.ObjectID
	obj26  ns.ObjectID
	obj27  ns.ObjectID
	obj28  ns.ObjectID
	obj29  ns.ObjectID
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
	obj41  ns.ObjectID
	obj42  ns.ObjectID
	obj43  ns.ObjectID
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	gvar49 ns.ObjectGroupID
	gvar50 ns.ObjectGroupID
	gvar51 ns.ObjectGroupID
	gvar52 ns.WallGroupID
	gvar53 ns.WallGroupID
	gvar54 ns.WallGroupID
	wp55   ns.WaypointID
	wp56   ns.WaypointID
	wp57   ns.WaypointID
	wp58   ns.WaypointID
	wp59   ns.WaypointID
	wp60   ns.WaypointID
	wp61   ns.WaypointID
	wp62   ns.WaypointID
	wp63   ns.WaypointID
	wp64   ns.WaypointID
	wp65   ns.WaypointID
	wp66   ns.WaypointID
	wp67   ns.WaypointID
	wp68   ns.WaypointID
)

func init() {
	ivar12 = 60
	ivar13 = 0
	ivar14 = 90
	flag16 = true
	flag17 = true
	flag18 = true
	flag19 = true
	flag20 = true
	gvar21 = 0
}
func ReleaseMonsters() {
	ns.ObjectOff(ns.GetTrigger())
	ns.UnlockDoor(obj4)
	ns.UnlockDoor(obj5)
	ns.UnlockDoor(obj6)
	ns.UnlockDoor(obj7)
	ns.UnlockDoor(obj8)
	ns.UnlockDoor(obj9)
	ns.UnlockDoor(obj10)
	ns.SetOwner(ns.Object("OgreOwner"), ns.Object("Ogre01"))
	ns.SetOwner(ns.Object("OgreOwner"), ns.Object("Ogre02"))
	ns.SetOwner(ns.Object("OgreOwner"), ns.Object("Ogre03"))
	ns.SetOwner(ns.Object("TrollOwner"), ns.Object("Troll01"))
	ns.SetOwner(ns.Object("TrollOwner"), ns.Object("Troll02"))
	ns.GroupMove(gvar11, ns.Waypoint("CagedMonsterWP"))
}
func SayCountdown() {
	ivar13 -= 1
	if !(ivar13 <= 0) {
		goto LABEL1
	}
	ivar13 = 0
	goto LABEL2
LABEL1:
	ns.FrameTimer(30, SayCountdown)
LABEL2:
	return
}
func WarriorRecognize() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ivar13 == 0) {
		goto LABEL1
	}
	v0 = ns.Random(0, 3)
	ivar13 = ivar12
	SayCountdown()
	if v0 != 0 {
		goto LABEL2
	}
	ns.Chat(ns.GetTrigger(), "War01A.scr:SetpieceGuard1Talk01")
	goto LABEL1
LABEL2:
	ns.Chat(ns.GetTrigger(), "Wiz06a:Guard1Recognize")
LABEL1:
	return
}
func WarriorEngage() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ivar13 == 0) {
		goto LABEL1
	}
	v0 = ns.Random(0, 3)
	ivar13 = ivar12
	SayCountdown()
	if v0 != 0 {
		goto LABEL2
	}
	ns.Chat(ns.GetTrigger(), "War01A.scr:SetpieceGuard1Talk01")
	goto LABEL1
LABEL2:
	ns.Chat(ns.GetTrigger(), "Wiz06a:Guard1Recognize")
LABEL1:
	return
}
func WarriorListen() {
	if !(ns.IsCaller(ns.GetHost()) && ivar13 == 0) {
		goto LABEL1
	}
	if ns.IsVisibleTo(ns.GetTrigger(), ns.GetHost()) {
		goto LABEL1
	}
	ivar13 = ivar12
	SayCountdown()
	ns.ChatTimer(ns.GetTrigger(), "Wiz06a:Guard2Listen", ivar14)
LABEL1:
	return
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp15)
}
func heardYou() {
	if !flag19 {
		goto LABEL1
	}
	flag19 = false
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.Chat(ns.GetTrigger(), "Wiz06b:Guard1")
	ns.FrameTimer(90, hearingRestored)
LABEL1:
	return
}
func hearingRestored() {
	flag19 = true
}
func sightRestored() {
	flag18 = true
}
func flamesOn() {
	ns.ObjectGroupOn(gvar49)
}
func flamesOff() {
	ns.ObjectGroupOff(gvar49)
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
	ns.BecomePet(obj27)
	ns.BecomePet(obj28)
	ns.WallGroupOpen(gvar54)
}
func horrendousHurt() {
	if !flag17 {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj22) < 50) {
		goto LABEL1
	}
	ns.SetDialog(obj22, ns.NORMAL, horrendousTalk2Start, horrendousTalk2End)
	ns.StartDialog(obj22, ns.GetHost())
LABEL1:
	return
}
func horrendousDie() {
	ns.SetDialog(obj22, ns.NORMAL, horrendousTalk3Start, horrendousTalk3End)
	ns.StartDialog(obj22, ns.GetHost())
}
func horrendousEngage() {
	if !flag16 {
		goto LABEL1
	}
	lockDoors()
	ns.Music(28, 100)
	ns.WideScreen(true)
	ns.SetDialog(obj22, ns.NORMAL, horrendousTalk1Start, horrendousTalk1End)
	ns.StartDialog(obj22, ns.GetHost())
LABEL1:
	return
}
func horrendousTalk1Start() {
	flag16 = false
	ns.CreatureIdle(obj22)
	ns.ObjectGroupOff(gvar51)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	ns.Frozen(ns.GetHost(), true)
	ns.TellStory(ns.DemonRecognize, "Wiz06c:Horrendous2")
}
func horrendousTalk1End() {
	ns.ObjectGroupOn(gvar51)
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.CancelDialog(obj22)
}
func horrendousTalk2Start() {
	flag17 = false
	ns.ObjectGroupOff(gvar51)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	ns.Damage(ns.GetHost(), 0, 1, 0)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.TellStory(ns.DemonRecognize, "Wiz06c:Horrendous3")
}
func horrendousTalk2End() {
	ns.ObjectGroupOn(gvar51)
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
}
func horrendousTalk3Start() {
	ns.ObjectGroupOff(gvar51)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.TellStory(ns.DemonRecognize, "Wiz06c:Horrendous5")
}
func horrendousTalk3End() {
	ns.ObjectGroupOn(gvar51)
	ns.ObjectOn(obj27)
	ns.ObjectOn(obj28)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(60, getHalberd)
}
func getHalberd() {
	ns.Pickup(ns.GetHost(), obj23)
	ns.AudioEvent(ns.StaffOblivionAchieve1, wp60)
	ns.ObjectOn(obj48)
	ns.UnlockDoor(obj33)
	ns.UnlockDoor(obj34)
}
func lockDoors() {
	if !flag20 {
		goto LABEL1
	}
	flag20 = false
	ns.ObjectOff(obj41)
	ns.ObjectOn(obj40)
	ns.LockDoor(obj33)
	ns.LockDoor(obj34)
	ns.LockDoor(obj35)
	ns.LockDoor(obj36)
	ns.LockDoor(obj37)
	ns.LockDoor(obj38)
	ns.LockDoor(obj39)
LABEL1:
	return
}
func unlockCells() {
	ns.UnlockDoor(obj4)
	ns.UnlockDoor(obj5)
	ns.UnlockDoor(obj6)
	ns.UnlockDoor(obj7)
	ns.UnlockDoor(obj8)
	ns.UnlockDoor(obj9)
	ns.UnlockDoor(obj10)
}
func lockCells() {
	ns.LockDoor(obj4)
	ns.LockDoor(obj5)
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
	ns.LockDoor(obj8)
	ns.LockDoor(obj9)
	ns.LockDoor(obj10)
}
func openMainGates() {
	ns.Move(obj44, wp65)
	ns.Move(obj45, wp66)
	ns.Move(obj46, wp67)
	ns.Move(obj47, wp68)
}
func closeMainGates() {
	ns.Move(obj44, wp61)
	ns.Move(obj45, wp62)
	ns.Move(obj46, wp63)
	ns.Move(obj47, wp64)
}
func toggleMainGates() {
	if gvar21 != 0 {
		goto LABEL1
	}
	gvar21 = 1
	openMainGates()
	goto LABEL2
LABEL1:
	gvar21 = 0
	closeMainGates()
LABEL2:
	ns.AudioEvent(ns.BoulderMove, wp64)
	ns.AudioEvent(ns.SecretWallStoneClose, wp64)
}
func endLevel() {
	ns.Blind()
	ns.DestroyEveryChat()
	ns.ObjectGroupOff(gvar51)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(60, exitLevel)
}
func exitLevel() {
	ns.DestroyEveryChat()
	ns.Frozen(ns.GetHost(), false)
	ns.UnBlind()
	ns.MoveObject(ns.GetHost(), 655, 2726)
}
func MapInitialize() {
	obj22 = ns.Object("Horrendous")
	obj23 = ns.Object("OblivionHalberd")
	obj24 = ns.Object("ManaDrainFocus")
	obj25 = ns.Object("EliteGuard1")
	obj26 = ns.Object("EliteGuard2")
	obj27 = ns.Object("Wizard1")
	obj28 = ns.Object("Wizard2")
	obj29 = ns.Object("Ogre1")
	obj30 = ns.Object("Ogre2")
	obj31 = ns.Object("Ogre3")
	obj32 = ns.Object("Ogre4")
	obj33 = ns.Object("ArenaDoor1")
	obj34 = ns.Object("ArenaDoor2")
	obj35 = ns.Object("ArenaDoor3")
	obj36 = ns.Object("ArenaDoor4")
	obj37 = ns.Object("ArenaDoor5")
	obj38 = ns.Object("ArenaDoor6")
	obj39 = ns.Object("ArenaDoor7")
	obj40 = ns.Object("OpenArenaTrigger")
	obj41 = ns.Object("CloseArenaTrigger")
	ns.ObjectOff(obj40)
	ns.ObjectOff(obj41)
	obj42 = ns.Object("RSideLocked")
	obj43 = ns.Object("LSideLocked")
	obj4 = ns.Object("CellDoor1")
	obj5 = ns.Object("CellDoor2")
	obj6 = ns.Object("CellDoor3")
	obj7 = ns.Object("CellDoor4")
	obj8 = ns.Object("CellDoor5")
	obj9 = ns.Object("CellDoor6")
	obj10 = ns.Object("CellDoor7")
	obj48 = ns.Object("ExitTrigger1")
	ns.ObjectOff(obj48)
	obj44 = ns.Object("FG1Mover")
	obj45 = ns.Object("FG2Mover")
	obj46 = ns.Object("FG3Mover")
	obj47 = ns.Object("FG4Mover")
	gvar49 = ns.ObjectGroup("BlueFlames")
	ns.ObjectGroupOn(gvar49)
	gvar11 = ns.ObjectGroup("CagedMonsters")
	gvar50 = ns.ObjectGroup("RoamingWarriors")
	gvar51 = ns.ObjectGroup("EveryMonsterOnMap")
	gvar52 = ns.WallGroup("OgreCages")
	gvar53 = ns.WallGroup("UndeadBreakIn")
	gvar54 = ns.WallGroup("WizardCages")
	wp55 = ns.Waypoint("PlayerSounds")
	wp15 = ns.Waypoint("WellWP")
	wp56 = ns.Waypoint("PlayerNorth")
	wp57 = ns.Waypoint("PlayerSouth")
	wp58 = ns.Waypoint("PlayerEast")
	wp59 = ns.Waypoint("PlayerWest")
	wp60 = ns.Waypoint("HorrendousNorth")
	wp61 = ns.Waypoint("FG1Closed")
	wp62 = ns.Waypoint("FG2Closed")
	wp63 = ns.Waypoint("FG3Closed")
	wp64 = ns.Waypoint("FG4Closed")
	wp65 = ns.Waypoint("FG1Open")
	wp66 = ns.Waypoint("FG2Open")
	wp67 = ns.Waypoint("FG3Open")
	wp68 = ns.Waypoint("FG4Open")
	ns.LockDoor(obj42)
	ns.LockDoor(obj43)
	ns.LockDoor(obj4)
	ns.LockDoor(obj5)
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
	ns.LockDoor(obj8)
	ns.LockDoor(obj9)
	ns.LockDoor(obj10)
	ns.ObjectOff(obj27)
	ns.ObjectOff(obj28)
	ns.StoryPic(obj22, "HorrendousPic")
	ns.StoryPic(obj27, "WizardGuard1Pic")
	ns.StoryPic(obj28, "WizardGuard2Pic")
	ns.GroupWander(gvar50)
	ns.Music(16, 100)
}
func PlayerDeath() {
	ns.DeathScreen(6)
}
func unlockDoors() {
	if flag20 {
		goto LABEL1
	}
	flag20 = true
	ns.ObjectOff(obj40)
	ns.ObjectOn(obj41)
	ns.UnlockDoor(obj33)
	ns.UnlockDoor(obj34)
LABEL1:
	return
}
func Secret01Found() {
	ns.MoveWaypoint(wp55, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp55)
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
