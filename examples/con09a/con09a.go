package con09a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   [2]ns.ObjectID
	wp5    [2]ns.WaypointID
	gvar6  ns.WallGroupID
	wp7    [5]ns.WaypointID
	wp8    [6]ns.WaypointID
	obj9   [6]ns.ObjectID
	gvar10 int
	ivar11 int
	ivar12 int
	gvar13 int
	flag14 bool
	flag15 bool
	flag16 bool
	fvar17 float32
	fvar18 float32
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
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
	gvar32 ns.ObjectGroupID
	gvar33 ns.ObjectGroupID
	gvar34 ns.ObjectGroupID
	gvar35 ns.ObjectGroupID
	gvar36 ns.ObjectGroupID
	gvar37 ns.ObjectGroupID
	gvar38 ns.ObjectGroupID
	gvar39 ns.WallGroupID
	gvar40 ns.WallGroupID
	gvar41 ns.WallGroupID
	gvar42 ns.WallGroupID
	gvar43 ns.WallGroupID
	gvar44 ns.WallGroupID
	gvar45 ns.WallGroupID
	wp46   ns.WaypointID
	wp47   ns.WaypointID
	wp48   ns.WaypointID
	gvar49 [9]int
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	flag52 bool
	flag53 bool
	gvar54 int
	gvar55 int
	gvar56 int
	gvar57 int
	gvar58 int
	gvar59 int
	gvar60 int
	gvar61 int
	gvar62 int
	gvar63 int
	flag64 bool
	gvar65 int
	gvar66 int
	gvar67 int
	gvar68 int
	gvar69 int
	gvar70 int
	ivar71 int
	ivar72 int
	obj73  [8]ns.ObjectID
	wp74   [8]ns.WaypointID
	gvar75 [8]int
	ivar76 [8]int
	gvar77 [8]ns.TimerID
)

func init() {
	gvar10 = 0
	gvar13 = 0
	flag14 = false
	flag15 = false
	flag16 = false
	flag52 = false
	flag53 = false
	gvar54 = 0
	gvar55 = 1
	gvar56 = 2
	gvar57 = 3
	gvar58 = 4
	gvar59 = 5
	gvar60 = 6
	gvar62 = gvar54
	gvar63 = gvar57
	flag64 = false
	gvar65 = 0
	gvar66 = 1
	gvar67 = 2
	gvar68 = 3
	gvar69 = 4
	gvar70 = 5
	ivar71 = 15
	ivar72 = 2
}
func PlantBarrier() {
	ivar11 = 0
	for {
		if !(ivar11 < 6) {
			goto LABEL1
		}
		CreatePlant(ivar11)
		ivar11 += 1
	}
LABEL1:
	return
}
func CreatePlant(a1 int) {
	obj9[a1] = ns.CreateObject("CarnivorousPlant", wp8[a1])
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj9[a1]), ns.GetObjectY(obj9[a1]), 0, 0)
	ns.LookAtObject(obj9[a1], ns.GetHost())
	ns.SetOwner(obj4[0], obj9[a1])
	ns.SetCallback(obj9[a1], 7, AttackBack)
}
func CheckPlants() {
	if !(flag15 == false || flag16 == false) {
		goto LABEL1
	}
	ivar11 = 0
	for {
		if !(ivar11 < 6) {
			goto LABEL2
		}
		if !(ns.CurrentHealth(obj9[ivar11]) <= 0) {
			goto LABEL3
		}
		ns.FrameTimerWithArg(14, ivar11, PlantDie)
	LABEL3:
		ivar11 += 1
	}
LABEL2:
	ns.FrameTimer(15, CheckPlants)
LABEL1:
	return
}
func PlantDie(a1 int) {
	if !(flag15 == false || flag16 == false) {
		goto LABEL1
	}
	ns.Delete(obj9[a1])
	CreatePlant(a1)
LABEL1:
	return
}
func AttackBack() {
	if ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Attack(ns.GetTrigger(), ns.GetCaller())
LABEL1:
	return
}
func PlayAction() {
	ns.Music(26, 100)
}
func PlaySwamp() {
	ns.Music(29, 100)
}
func DryadHuntPlayer() {
	ns.CreatureHunt(ns.GetTrigger())
}
func StartDryadSequence() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && flag14 == false) {
		goto LABEL1
	}
	flag14 = true
	ns.SetOwner(obj4[0], obj4[1])
	ns.MusicPushEvent()
	PlayAction()
	ns.NoWallSound(true)
	ns.WallGroupClose(gvar6)
	ns.NoWallSound(false)
	v0 = 0
	for {
		if !(v0 < 2) {
			goto LABEL2
		}
		ns.ObjectOn(obj4[v0])
		ns.Enchant(obj4[v0], ns.ENCHANT_INVISIBLE, 1)
		ns.MoveObject(obj4[v0], ns.GetWaypointX(wp5[v0]), ns.GetWaypointY(wp5[v0]))
		ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp5[v0]), ns.GetWaypointY(wp5[v0]), 0, 0)
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp5[v0]), ns.GetWaypointY(wp5[v0]), 0, 0)
		v0 += 1
	}
LABEL2:
	PlantBarrier()
	ns.FrameTimer(1, CheckPlants)
LABEL1:
	return
}
func DryadDie() {
	if !ns.IsTrigger(obj4[0]) {
		goto LABEL1
	}
	flag15 = true
LABEL1:
	if !ns.IsTrigger(obj4[1]) {
		goto LABEL2
	}
	flag16 = true
LABEL2:
	if !(flag15 == true && flag16 == true) {
		goto LABEL3
	}
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar6)
	ns.NoWallSound(false)
	ns.MusicPopEvent()
LABEL3:
	return
}
func DryadStay() {
	if !(ns.IsCaller(obj4[0]) || ns.IsCaller(obj4[1])) {
		goto LABEL1
	}
	ivar12 = ns.Random(0, 4)
	fvar17 = ns.GetWaypointX(wp7[ivar12])
	fvar18 = ns.GetWaypointY(wp7[ivar12])
	ns.Enchant(ns.GetCaller(), ns.ENCHANT_INVISIBLE, 1.5)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.MoveObject(ns.GetCaller(), fvar17, fvar18)
	ns.Effect(ns.BLUE_SPARKS, fvar17, fvar18, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, fvar17, fvar18, 0, 0)
LABEL1:
	return
}
func PlayerDeath() {
	ns.DeathScreen(9)
}
func CaptainDialogStart() {
	var v0 int
	v0 = gvar62
	if v0 == gvar54 {
		goto LABEL1
	}
	if v0 == gvar55 {
		goto LABEL2
	}
	if v0 == gvar56 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War09a:CaptainGreet")
	goto LABEL4
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War09a:CaptainLeave")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War09a:CaptainImpatient")
	goto LABEL4
LABEL4:
	return
}
func CaptainDialogEnd() {
	var v0 int
	v0 = gvar62
	if v0 == gvar54 {
		goto LABEL1
	}
	if v0 == gvar55 {
		goto LABEL2
	}
	if v0 == gvar56 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.JournalEntry(ns.GetHost(), "FindMordwyn", 2)
	ns.JournalEntry(ns.GetHost(), "FindOutpost", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	PlaySwamp()
	gvar62 = gvar55
	goto LABEL4
LABEL2:
	gvar62 = gvar56
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func MordwynDialogStart() {
	var v0 int
	v0 = gvar63
	if v0 == gvar57 {
		goto LABEL1
	}
	if v0 == gvar58 {
		goto LABEL2
	}
	if v0 == gvar59 {
		goto LABEL3
	}
	if v0 == gvar60 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.LookAtObject(obj19, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War09a:MordwynGreet")
	goto LABEL5
LABEL2:
	ns.LookAtObject(obj19, ns.GetHost())
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.TellStory(ns.SwordsmanHurt, "Con09a:MordwynHome")
	goto LABEL5
LABEL3:
	ns.LookAtObject(obj19, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con09a:MordwynHome2")
	goto LABEL5
LABEL4:
	ns.LookAtObject(obj19, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War09a:MordwynDone")
	goto LABEL5
LABEL5:
	return
}
func MordwynDialogEnd() {
	var v0 int
	v0 = gvar63
	if v0 == gvar57 {
		goto LABEL1
	}
	if v0 == gvar58 {
		goto LABEL2
	}
	if v0 == gvar59 {
		goto LABEL3
	}
	if v0 == gvar60 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar63 = gvar58
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.JournalEdit(ns.GetHost(), "FindMordwyn", 4)
	ns.CancelDialog(obj19)
	ns.SetCallback(obj19, 11, AllowDialog)
	ns.Move(obj19, wp47)
	goto LABEL5
LABEL2:
	gvar63 = gvar59
	ns.SetCallback(obj19, 11, GetArmorOfJandor)
	ns.Move(obj19, wp48)
	goto LABEL5
LABEL3:
	gvar63 = gvar60
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.CreatureGuard(obj19, ns.GetWaypointX(wp47), ns.GetWaypointY(wp47), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetCallback(obj19, 11, NullFunction)
	ns.Pickup(ns.GetHost(), obj23)
	ns.Pickup(ns.GetHost(), obj24)
	ns.Pickup(ns.GetHost(), obj25)
	ns.PrintToAll("Con09a:GainedBowAlt")
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func StartCaptainConversation() {
	ns.SetDialog(obj20, ns.NORMAL, CaptainDialogStart, CaptainDialogEnd)
	ns.StartDialog(obj20, ns.GetHost())
}
func WispInitialize() {
	obj73[0] = ns.Object("Wisp0")
	gvar75[0] = gvar65
	wp74[0] = ns.Waypoint("Wisp0Dest")
	obj73[1] = ns.Object("Wisp1")
	gvar75[1] = gvar65
	wp74[1] = ns.Waypoint("Wisp1Dest")
}
func MapInitialize() {
	obj20 = ns.Object("Jandor")
	obj21 = ns.Object("Basket")
	obj22 = ns.Object("BasketShadow")
	obj19 = ns.Object("Mordwyn")
	obj23 = ns.Object("JBow")
	obj24 = ns.Object("JQuiver1")
	obj25 = ns.Object("JQuiver2")
	obj26 = ns.Object("PitElev01")
	obj27 = ns.Object("PitElev02")
	obj28 = ns.Object("PitElev03")
	obj29 = ns.Object("TreasureTrap01")
	obj30 = ns.Object("TreasureTrap01Light")
	obj31 = ns.Object("Secret02Trigger")
	obj4[0] = ns.Object("Dryad0")
	obj4[1] = ns.Object("Dryad1")
	gvar39 = ns.WallGroup("Surprise00Walls")
	gvar40 = ns.WallGroup("Surprise00DestWalls")
	gvar41 = ns.WallGroup("Surprise01Walls")
	gvar42 = ns.WallGroup("Surprise01DWalls")
	gvar43 = ns.WallGroup("Surprise02Walls")
	gvar44 = ns.WallGroup("Surprise02DestWalls")
	gvar45 = ns.WallGroup("TreasureTrap01Walls")
	gvar6 = ns.WallGroup("DryadWalls")
	gvar34 = ns.ObjectGroup("Surprise00Triggers")
	gvar33 = ns.ObjectGroup("Surprise00Creatures")
	gvar35 = ns.ObjectGroup("Surprise01Creatures")
	gvar36 = ns.ObjectGroup("Surprise01Triggers")
	gvar32 = ns.ObjectGroup("MordwynTriggers")
	gvar37 = ns.ObjectGroup("TreasureTrap01Zombies")
	gvar38 = ns.ObjectGroup("Secret01Triggers")
	wp46 = ns.Waypoint("TalkToPlayer")
	wp47 = ns.Waypoint("MordwynHome")
	wp48 = ns.Waypoint("MordwynBedroom")
	wp5[0] = ns.Waypoint("DryadHome0")
	wp5[1] = ns.Waypoint("DryadHome1")
	wp7[0] = ns.Waypoint("DryadWarp0")
	wp7[1] = ns.Waypoint("DryadWarp1")
	wp7[2] = ns.Waypoint("DryadWarp2")
	wp7[3] = ns.Waypoint("DryadWarp3")
	wp7[4] = ns.Waypoint("DryadWarp4")
	wp8[0] = ns.Waypoint("PlantLoc0")
	wp8[1] = ns.Waypoint("PlantLoc1")
	wp8[2] = ns.Waypoint("PlantLoc2")
	wp8[3] = ns.Waypoint("PlantLoc3")
	wp8[4] = ns.Waypoint("PlantLoc4")
	wp8[5] = ns.Waypoint("PlantLoc5")
	wp50 = ns.Waypoint("WellWP")
	wp51 = ns.Waypoint("PlayerSounds")
	ns.SetOwner(ns.GetHost(), obj20)
	ns.UnBlind()
	WispInitialize()
	ns.StoryPic(obj20, "AirshipCaptainPic")
	ns.StoryPic(obj19, "MordwynPic")
	ns.SetDialog(obj19, ns.NORMAL, MordwynDialogStart, MordwynDialogEnd)
	ns.StartupScreen(9)
	ns.Music(0, 100)
	ns.FrameTimer(5, StartCaptainConversation)
	ns.ObjectOn(ns.Object("StayHere1"))
	ns.Move(ns.Object("StayHere1"), ns.Waypoint("Here1"))
	ns.ObjectOn(ns.Object("StayHere2"))
	ns.Move(ns.Object("StayHere2"), ns.Waypoint("Here2"))
	ns.ObjectOn(ns.Object("StayHere3"))
	ns.Move(ns.Object("StayHere3"), ns.Waypoint("Here3"))
	ns.ObjectOn(ns.Object("StayHere4"))
	ns.Move(ns.Object("StayHere4"), ns.Waypoint("Here4"))
	ns.ObjectOn(ns.Object("StayHere5"))
	ns.Move(ns.Object("StayHere5"), ns.Waypoint("Here5"))
}
func MordwynGreet() {
	ns.LookAtObject(obj19, ns.GetHost())
	ns.StartDialog(obj19, ns.GetHost())
	ns.SetCallback(obj19, 11, NullFunction)
}
func GetArmorOfJandor() {
	ns.PauseObject(obj19, 75)
	ns.SetCallback(obj19, 11, GiveArmorToPlayer)
	ns.Move(obj19, wp47)
}
func GiveArmorToPlayer() {
	ns.LookAtObject(obj19, ns.GetHost())
	ns.StartDialog(obj19, ns.GetHost())
}
func KeepOut() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func CapTakeoff() {
	if !(ns.IsCaller(ns.GetHost()) && flag53 == false) {
		goto LABEL1
	}
	flag53 = true
	ns.Delete(obj21)
	ns.Delete(obj22)
	ns.Delete(obj20)
LABEL1:
	return
}
func NullFunction() {
}
func MordwynMove() {
	ns.ObjectGroupOff(gvar32)
	ns.ObjectOn(obj19)
	ns.SetOwner(ns.GetHost(), obj19)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.Move(obj19, wp46)
}
func AllowDialog() {
	ns.CreatureGuard(obj19, ns.GetWaypointX(wp47), ns.GetWaypointY(wp47), ns.GetWaypointX(wp46), ns.GetWaypointY(wp46), 0)
	ns.SetDialog(obj19, ns.NORMAL, MordwynDialogStart, MordwynDialogEnd)
}
func Surprise00() {
	ns.ObjectGroupOff(gvar34)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar39)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar40)
	ns.ObjectGroupOn(gvar33)
	ns.GroupWander(gvar33)
}
func Surprise01() {
	ns.ObjectGroupOff(gvar36)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar41)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar42)
	ns.ObjectGroupOn(gvar35)
}
func Secret01Declare() {
	ns.ObjectGroupOff(gvar38)
	ns.MoveWaypoint(wp51, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp51)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret02Declare() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp51, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp51)
	ns.GiveXp(ns.GetHost(), 100)
}
func EnablePitElev01() {
	ns.ObjectOn(obj26)
}
func EnablePitElev02() {
	ns.ObjectOn(obj27)
}
func EnablePitElev03() {
	ns.ObjectOn(obj28)
}
func TriggerTreasureTrap() {
	ns.WallGroupBreak(gvar45)
	ns.ObjectOff(obj30)
	ns.GroupAggressionLevel(gvar37, 0.83)
	ns.GroupWander(gvar37)
}
func MapEntry() {
	ns.NoWallSound(false)
}
func whichWisp(a1 ns.ObjectID) int {
	var v0 int
	v0 = 0
	for {
		if !(v0 < ivar72) {
			goto LABEL1
		}
		if obj73[v0] != a1 {
			goto LABEL2
		}
		return v0
	LABEL2:
		v0 += 1
	}
LABEL1:
	return -1
	return 1
}
func WispAction(a1 int) {
	var (
		v0 float32
		v1 int
	)
	v1 = gvar75[a1]
	if v1 == gvar66 {
		goto LABEL1
	}
	if v1 == gvar68 {
		goto LABEL2
	}
	if v1 == gvar67 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v0 = ns.Distance(ns.GetObjectX(obj73[a1]), ns.GetObjectY(obj73[a1]), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	if !(v0 < 70) {
		goto LABEL5
	}
	gvar75[a1] = gvar67
	ns.Move(obj73[a1], wp74[a1])
LABEL5:
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj73[a1], ns.GetHost())
	ivar76[a1] += 1
	if !(ivar76[a1] > 4) {
		goto LABEL6
	}
	gvar75[a1] = gvar66
	ns.CreatureFollow(obj73[a1], ns.GetHost())
LABEL6:
	goto LABEL4
LABEL3:
	v0 = ns.Distance(ns.GetObjectX(obj73[a1]), ns.GetObjectY(obj73[a1]), ns.GetWaypointX(wp74[a1]), ns.GetWaypointY(wp74[a1]))
	if !(v0 < 30) {
		goto LABEL7
	}
	gvar75[a1] = gvar69
	ns.CreatureIdle(obj73[a1])
	return
LABEL7:
	goto LABEL4
LABEL4:
	ns.FrameTimerWithArg(ivar71, a1, WispAction)
}
func WispRecognize() {
	var (
		v0 int
		v1 int
	)
	if ns.HasClass(ns.GetCaller(), ns.PLAYER) {
		goto LABEL1
	}
	return
LABEL1:
	v0 = whichWisp(ns.GetTrigger())
	if !(v0 < 0) {
		goto LABEL2
	}
	return
LABEL2:
	v1 = gvar75[v0]
	if v1 == gvar65 {
		goto LABEL3
	}
	if v1 == gvar68 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	gvar75[v0] = gvar66
	ns.CreatureFollow(ns.GetTrigger(), ns.GetCaller())
	gvar77[v0] = ns.FrameTimerWithArg(ivar71, v0, WispAction)
	goto LABEL5
LABEL4:
	gvar75[v0] = gvar67
	ns.Move(ns.GetTrigger(), wp74[v0])
	goto LABEL5
LABEL5:
	return
}
func WispLoseSight() {
	var (
		v0 int
		v1 int
	)
	if ns.HasClass(ns.GetCaller(), ns.PLAYER) {
		goto LABEL1
	}
	return
LABEL1:
	v0 = whichWisp(ns.GetTrigger())
	if !(v0 < 0) {
		goto LABEL2
	}
	return
LABEL2:
	v1 = gvar75[v0]
	if v1 == gvar67 {
		goto LABEL3
	}
	goto LABEL4
LABEL3:
	gvar75[v0] = gvar68
	ivar76[v0] = 0
	ns.CreatureIdle(ns.GetTrigger())
	goto LABEL4
LABEL4:
	return
}
func WispGetMad() {
	var v0 int
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.CreatureIdle(ns.GetTrigger())
	v0 = whichWisp(ns.GetTrigger())
	if !(v0 >= 0) {
		goto LABEL1
	}
	gvar75[v0] = gvar70
	ns.CancelTimer(gvar77[v0])
LABEL1:
	return
}
func WispMedieval() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.CreatureIdle(ns.GetTrigger())
LABEL1:
	return
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp50)
}
func PlayMordwyn() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(21, 100)
LABEL1:
	return
}
func PlayOgre() {
	ns.Music(7, 100)
}
func PlayCaves() {
	ns.Music(18, 100)
}
func PlayWasteland() {
	ns.Music(19, 100)
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
