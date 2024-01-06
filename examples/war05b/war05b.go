package war05b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	gvar5  ns.WallGroupID
	obj6   ns.ObjectID
	obj7   ns.ObjectID
	obj8   ns.ObjectID
	obj9   ns.ObjectID
	obj10  ns.ObjectID
	wp11   ns.WaypointID
	gvar12 int
	gvar13 int
	gvar14 int
	gvar15 int
	gvar16 int
	gvar17 int
	gvar18 int
	gvar19 int
	gvar20 int
	gvar21 int
	gvar22 int
	gvar23 int
	gvar24 int
	ivar25 int
	flag26 bool
	flag27 bool
	flag28 bool
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
	obj49  ns.ObjectID
	gvar50 ns.ObjectGroupID
	gvar51 ns.ObjectGroupID
	gvar52 ns.WallGroupID
	wp53   ns.WaypointID
	wp54   ns.WaypointID
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
	wp69   ns.WaypointID
	gvar70 ns.ObjectGroupID
)

func init() {
	gvar12 = 0
	gvar13 = 1
	gvar14 = 2
	gvar15 = 3
	gvar16 = 4
	gvar17 = 5
	gvar18 = 6
	gvar19 = 7
	gvar20 = 8
	gvar21 = gvar12
	gvar22 = gvar14
	gvar23 = gvar16
	gvar24 = gvar19
	ivar25 = 0
	flag26 = false
	flag27 = false
	flag28 = false
}
func StopElevator() {
	ns.ObjectOff(obj4)
	ns.FrameTimer(60, LowerKeyProtectWalls)
}
func LowerKeyProtectWalls() {
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar5)
	ns.NoWallSound(false)
}
func RaiseKey() {
	ns.ObjectOn(obj4)
	ns.FrameTimer(5, StopElevator)
}
func UnlockMainGates() {
	ns.ObjectOff(obj8)
	ns.UnlockDoor(obj6)
	ns.UnlockDoor(obj7)
}
func DefendCage() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Attack(ns.GetTrigger(), ns.GetHost())
LABEL1:
	return
}
func ReturnToPost() {
	if !(ns.IsCaller(obj9) || ns.IsCaller(obj10)) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func SecretFound() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp11, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp11)
	ns.GiveXp(ns.GetHost(), 40)
}
func CallForGuards() {
	ns.Frozen(obj45, true)
	ns.CreatureIdle(obj45)
	ns.SetDialog(obj45, ns.NEXT, KingStart, KingEnd)
	ns.StoryPic(obj45, "OgreWarlordPic")
	ns.StartDialog(obj45, ns.GetHost())
}
func ProtectTheKing() {
	LetterBoxOff()
	ns.ObjectGroupOn(gvar70)
	ns.Frozen(ns.GetHost(), false)
	ns.EnchantOff(obj45, ns.ENCHANT_INVULNERABLE)
	ns.ClearOwner(obj45)
	ns.Frozen(obj45, false)
	ns.ObjectOn(obj46)
	ns.ObjectOn(obj47)
	ns.ObjectOn(obj48)
	ns.AggressionLevel(obj45, 0.5)
	ns.MoveObject(obj46, ns.GetWaypointX(wp65), ns.GetWaypointY(wp65))
	ns.MoveObject(obj47, ns.GetWaypointX(wp66), ns.GetWaypointY(wp66))
	ns.MoveObject(obj48, ns.GetWaypointX(wp67), ns.GetWaypointY(wp67))
	ns.Wander(obj46)
	ns.Wander(obj47)
	ns.CreatureGuard(obj48, ns.GetWaypointX(wp68), ns.GetWaypointY(wp68), ns.GetWaypointX(wp64), ns.GetWaypointY(wp64), 0)
}
func KingStart() {
	var v0 int
	v0 = gvar24
	if v0 == gvar19 {
		goto LABEL1
	}
	if v0 == gvar20 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con05:OgreTalk07")
	ns.LookAtObject(obj45, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj45)
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con05:OgreTalk08")
	goto LABEL3
LABEL3:
	return
}
func KingEnd() {
	var v0 int
	v0 = gvar24
	if v0 == gvar19 {
		goto LABEL1
	}
	if v0 == gvar20 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	gvar24 = gvar20
	ns.StartDialog(obj45, ns.GetHost())
	goto LABEL3
LABEL2:
	ns.CancelDialog(obj45)
	ProtectTheKing()
	ns.CreatureGuard(obj45, ns.GetWaypointX(wp64), ns.GetWaypointY(wp64), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 500)
	goto LABEL3
LABEL3:
	return
}
func CaptainStart() {
	var v0 int
	v0 = gvar21
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj37, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj37)
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:CaptainWaiting")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj37, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj37)
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:CaptainSuccess3")
	goto LABEL3
LABEL3:
	return
}
func CaptainEnd() {
	var v0 int
	v0 = gvar21
	if v0 == gvar12 {
		goto LABEL1
	}
	if v0 == gvar13 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	goto LABEL3
LABEL2:
	ns.JournalEdit(ns.GetHost(), "War05Quest4", 4)
	ns.CancelDialog(obj37)
	gvar23 = gvar18
	ns.SetDialog(obj38, ns.NORMAL, IngridStart, IngridEnd)
	ns.StartDialog(obj38, ns.GetHost())
	goto LABEL3
LABEL3:
	return
}
func LetterBoxOn() {
	ns.WideScreen(true)
}
func LetterBoxOff() {
	ns.WideScreen(false)
}
func PlayerDeath() {
	ns.DeathScreen(5)
}
func MaidenStart() {
	var v0 int
	v0 = gvar22
	if v0 == gvar14 {
		goto LABEL1
	}
	if v0 == gvar15 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:MaidenFollowing")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:MaidenRescued")
	goto LABEL3
LABEL3:
	return
}
func MaidenEnd() {
	var v0 int
	v0 = gvar22
	if v0 == gvar14 {
		goto LABEL1
	}
	if v0 == gvar15 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	goto LABEL3
LABEL2:
	goto LABEL3
LABEL3:
	return
}
func RescueMaidens() {
	ns.ObjectOff(obj34)
	ns.SetQuestStatus(1, "ChicksSaved")
	LetterBoxOn()
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureFollow(obj30, ns.GetHost())
	ns.CreatureFollow(obj31, ns.GetHost())
	ns.CreatureFollow(obj32, ns.GetHost())
	ns.CreatureFollow(obj33, ns.GetHost())
	ns.CreatureFollow(obj29, ns.GetHost())
	ns.SetDialog(obj29, ns.NORMAL, SisterStart, SisterEnd)
	ns.StoryPic(obj29, "GlyndaPic")
	ns.LookAtObject(obj29, ns.GetHost())
	ns.StartDialog(obj29, ns.GetHost())
	ns.StoryPic(obj30, "MaidenPic")
	ns.StoryPic(obj31, "MaidenPic2")
	ns.StoryPic(obj32, "MaidenPic3")
	ns.StoryPic(obj33, "MaidenPic4")
}
func SisterStart() {
	var v0 int
	v0 = gvar23
	if v0 == gvar16 {
		goto LABEL1
	}
	if v0 == gvar17 {
		goto LABEL2
	}
	if v0 == gvar18 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(obj29, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj29)
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:SisterFreed")
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj29, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj29)
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:SisterFollowing3")
	goto LABEL4
LABEL3:
	ns.LookAtObject(obj29, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj29)
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:SisterIdle")
	goto LABEL4
LABEL4:
	return
}
func SisterEnd() {
	var v0 int
	v0 = gvar23
	if v0 == gvar16 {
		goto LABEL1
	}
	if v0 == gvar17 {
		goto LABEL2
	}
	if v0 == gvar18 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	LetterBoxOff()
	ns.Frozen(ns.GetHost(), false)
	ns.BecomePet(obj30)
	ns.BecomePet(obj31)
	ns.BecomePet(obj32)
	ns.BecomePet(obj33)
	ns.BecomePet(obj29)
	MoveNPCs()
	gvar23 = gvar17
	ns.SetQuestStatus(1, "MaidensRescued")
	ns.CancelDialog(obj29)
	goto LABEL4
LABEL2:
	goto LABEL4
LABEL3:
	ns.SetDialog(obj38, ns.NORMAL, IngridStart, IngridEnd)
	ns.StartDialog(obj38, ns.GetHost())
	goto LABEL4
LABEL4:
	return
}
func MoveNPCs() {
	ns.MoveObject(obj37, ns.GetWaypointX(wp53), ns.GetWaypointY(wp53))
	ns.CreatureGuard(obj37, ns.GetWaypointX(wp53), ns.GetWaypointY(wp53), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.MoveObject(obj38, ns.GetWaypointX(wp63), ns.GetWaypointY(wp63))
	ns.CreatureGuard(obj38, ns.GetWaypointX(wp63), ns.GetWaypointY(wp63), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.MoveObject(obj41, ns.GetWaypointX(wp54), ns.GetWaypointY(wp54))
	ns.CreatureGuard(obj41, ns.GetWaypointX(wp54), ns.GetWaypointY(wp54), ns.GetWaypointX(wp56), ns.GetWaypointY(wp56), 0)
	ns.MoveObject(obj42, ns.GetWaypointX(wp55), ns.GetWaypointY(wp55))
	ns.CreatureGuard(obj42, ns.GetWaypointX(wp55), ns.GetWaypointY(wp55), ns.GetWaypointX(wp56), ns.GetWaypointY(wp56), 0)
}
func FinalScene() {
	LetterBoxOn()
	ns.Frozen(ns.GetHost(), true)
	gvar21 = gvar13
	ns.SetDialog(obj37, ns.NEXT, CaptainStart, CaptainEnd)
	ns.StartDialog(obj37, ns.GetHost())
}
func IngridStart() {
	ns.LookAtObject(obj38, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj38)
	ns.TellStory(ns.SwordsmanHurt, "War05B.scr:IngridReturned")
}
func IngridEnd() {
	ns.Pickup(ns.GetHost(), obj40)
	ns.Pickup(ns.GetHost(), obj39)
	ns.CancelDialog(obj38)
	MissionComplete()
}
func MissionComplete() {
	ns.GiveXp(ns.GetHost(), 1000)
	ns.Blind()
	ns.FrameTimer(45, Victory)
}
func Victory() {
	ns.Frozen(ns.GetHost(), false)
	LetterBoxOff()
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp69), ns.GetWaypointY(wp69))
}
func MaidenDeath() {
	ns.PrintToAll("War05B.scr:MaidenDeath")
	ns.DeathScreen(5)
}
func MapInitialize() {
	obj6 = ns.Object("MainGate1")
	obj7 = ns.Object("MainGate2")
	obj8 = ns.Object("MainGateSwitch")
	obj43 = ns.Object("TownGate1")
	obj44 = ns.Object("TownGate2")
	obj29 = ns.Object("W5Sister")
	obj30 = ns.Object("W5Maiden1")
	obj31 = ns.Object("W5Maiden2")
	obj32 = ns.Object("W5Maiden3")
	obj33 = ns.Object("W5Maiden4")
	obj34 = ns.Object("SetMaidenDialogTrigger")
	obj35 = ns.Object("CellDoor")
	obj36 = ns.Object("CellDoorSwitch")
	obj41 = ns.Object("Soldier1")
	obj42 = ns.Object("Soldier2")
	obj37 = ns.Object("W5Captain")
	obj38 = ns.Object("W5Ingrid")
	obj39 = ns.Object("PotionGift")
	obj40 = ns.Object("GoldGift")
	obj45 = ns.Object("OgreKing")
	obj49 = ns.Object("CallForGuardsTrigger")
	obj4 = ns.Object("KeyElevator")
	obj46 = ns.Object("Guard1")
	obj47 = ns.Object("Guard2")
	obj48 = ns.Object("Guard3")
	obj9 = ns.Object("CageGuard1")
	obj10 = ns.Object("CageGuard2")
	gvar50 = ns.ObjectGroup("KingTriggerGroup")
	gvar51 = ns.ObjectGroup("CaveEnemies")
	gvar70 = ns.ObjectGroup("WarlordOgres")
	gvar52 = ns.WallGroup("LordSafetyWalls")
	gvar5 = ns.WallGroup("KeyProtectWalls")
	wp64 = ns.Waypoint("OgreKingWaypoint")
	wp53 = ns.Waypoint("CaptainWP")
	wp54 = ns.Waypoint("Soldier1WP")
	wp55 = ns.Waypoint("Soldier2WP")
	wp56 = ns.Waypoint("SoldierLook")
	wp57 = ns.Waypoint("FaceWP")
	wp58 = ns.Waypoint("WP1")
	wp59 = ns.Waypoint("WP2")
	wp60 = ns.Waypoint("WP3")
	wp61 = ns.Waypoint("WP4")
	wp62 = ns.Waypoint("SisterWP")
	wp63 = ns.Waypoint("IngridWP")
	wp65 = ns.Waypoint("Guard1WP")
	wp66 = ns.Waypoint("Guard2WP")
	wp67 = ns.Waypoint("Guard3WP")
	wp68 = ns.Waypoint("Guard3SpotWP")
	wp69 = ns.Waypoint("MapExit")
	wp11 = ns.Waypoint("PlayerSounds")
	ns.StoryPic(obj37, "AirshipCaptainPic")
	ns.SetDialog(obj37, ns.NORMAL, CaptainStart, CaptainEnd)
	ns.StoryPic(obj38, "IngridPic")
	ns.LockDoor(obj35)
	ns.LockDoor(obj6)
	ns.LockDoor(obj7)
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
func Activate() {
	flag26 = true
	ns.ObjectGroupOff(gvar50)
	ns.ObjectOn(obj49)
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar70)
	ns.FrameTimer(1, LetterBoxOn)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar52)
	ns.NoWallSound(false)
	ns.Move(obj45, wp64)
}
func PummelPlayer() {
	ns.HitLocation(ns.GetTrigger(), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
}
func MapEntry() {
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.Music(7, 100)
	if ns.GetQuestStatus("MaidensRescued") != 1 {
		goto LABEL1
	}
	obj30 = ns.Object("War05B:W5Maiden1")
	obj31 = ns.Object("War05B:W5Maiden2")
	obj32 = ns.Object("War05B:W5Maiden3")
	obj33 = ns.Object("War05B:W5Maiden4")
	obj29 = ns.Object("War05B:W5Sister")
	ns.SetCallback(obj30, 5, MaidenDeath)
	ns.SetCallback(obj31, 5, MaidenDeath)
	ns.SetCallback(obj32, 5, MaidenDeath)
	ns.SetCallback(obj33, 5, MaidenDeath)
	ns.SetCallback(obj29, 5, MaidenDeath)
LABEL1:
	return
}
func FadeOut() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.Blind()
}
func UnlockCellDoor() {
	ns.ObjectOff(obj36)
	ns.UnlockDoor(obj35)
	ns.PrintToAll("War05B.scr:OgreCageUnlocked")
}
func PlayerLeft() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag28 = false
LABEL1:
	return
}
func CheckRescue() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	flag28 = true
	if ns.GetQuestStatus("MaidensRescued") != 1 {
		goto LABEL1
	}
	ns.CancelDialog(obj29)
	ns.Enchant(obj30, ns.ENCHANT_INVULNERABLE, 0)
	ns.Enchant(obj31, ns.ENCHANT_INVULNERABLE, 0)
	ns.Enchant(obj32, ns.ENCHANT_INVULNERABLE, 0)
	ns.Enchant(obj33, ns.ENCHANT_INVULNERABLE, 0)
	ns.Enchant(obj29, ns.ENCHANT_INVULNERABLE, 0)
	FinalScene()
LABEL1:
	return
}
func CheckRescue2() {
	if !ns.IsCaller(obj30) {
		goto LABEL1
	}
	ns.BecomeEnemy(ns.GetCaller())
	ns.CreatureGuard(ns.GetCaller(), ns.GetWaypointX(wp58), ns.GetWaypointY(wp58), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.CancelDialog(obj30)
	ivar25 += 1
LABEL1:
	if !ns.IsCaller(obj31) {
		goto LABEL2
	}
	ns.BecomeEnemy(ns.GetCaller())
	ns.CreatureGuard(ns.GetCaller(), ns.GetWaypointX(wp59), ns.GetWaypointY(wp59), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.CancelDialog(obj31)
	ivar25 += 1
LABEL2:
	if !ns.IsCaller(obj32) {
		goto LABEL3
	}
	ns.BecomeEnemy(ns.GetCaller())
	ns.CreatureGuard(ns.GetCaller(), ns.GetWaypointX(wp60), ns.GetWaypointY(wp60), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.CancelDialog(obj32)
	ivar25 += 1
LABEL3:
	if !ns.IsCaller(obj33) {
		goto LABEL4
	}
	ns.BecomeEnemy(ns.GetCaller())
	ns.CreatureGuard(ns.GetCaller(), ns.GetWaypointX(wp61), ns.GetWaypointY(wp61), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.CancelDialog(obj33)
	ivar25 += 1
LABEL4:
	if !ns.IsCaller(obj29) {
		goto LABEL5
	}
	ns.BecomeEnemy(ns.GetCaller())
	ns.CreatureGuard(ns.GetCaller(), ns.GetWaypointX(wp62), ns.GetWaypointY(wp62), ns.GetWaypointX(wp57), ns.GetWaypointY(wp57), 0)
	ns.CancelDialog(obj29)
	ivar25 += 1
LABEL5:
	r26 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r26 {
		goto LABEL6
	}
	ns.GoBackHome(ns.GetCaller())
LABEL6:
	return
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
