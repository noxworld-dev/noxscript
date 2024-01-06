package wiz01a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4   ns.ObjectID
	obj5   ns.ObjectID
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
	obj16  ns.ObjectID
	obj17  ns.ObjectID
	obj18  ns.ObjectID
	obj19  ns.ObjectID
	obj20  ns.ObjectID
	obj21  ns.ObjectID
	obj22  ns.ObjectID
	gvar23 ns.ObjectGroupID
	gvar24 ns.ObjectGroupID
	gvar25 ns.ObjectGroupID
	gvar26 ns.ObjectGroupID
	gvar27 ns.ObjectGroupID
	gvar28 ns.ObjectGroupID
	gvar29 ns.ObjectGroupID
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
	wp40   ns.WaypointID
	wp41   ns.WaypointID
	wp42   ns.WaypointID
	wp43   ns.WaypointID
	gvar44 ns.WallGroupID
	wp45   ns.WaypointID
	wp46   ns.WaypointID
	flag47 bool
	flag48 bool
	flag49 bool
	flag50 bool
	flag51 bool
	obj52  ns.ObjectID
	gvar53 int
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
	ivar64 int
	ivar65 int
	obj66  ns.ObjectID
	obj67  ns.ObjectID
	obj68  ns.ObjectID
	obj69  ns.ObjectID
	obj70  ns.ObjectID
	obj71  ns.ObjectID
	obj72  ns.ObjectID
	obj73  ns.ObjectID
	obj74  ns.ObjectID
	obj75  ns.ObjectID
	obj76  ns.ObjectID
	obj77  ns.ObjectID
	flag78 bool
	wp79   ns.WaypointID
	wp80   ns.WaypointID
	wp81   ns.WaypointID
	wp82   ns.WaypointID
	wp83   ns.WaypointID
	wp84   ns.WaypointID
	wp85   ns.WaypointID
	flag86 bool
)

func init() {
	flag47 = false
	flag48 = false
	flag49 = false
	flag50 = false
	flag51 = false
	gvar53 = 0
	gvar54 = 1
	gvar55 = 2
	gvar56 = 3
	gvar57 = 4
	gvar58 = 5
	gvar59 = 6
	gvar60 = gvar53
	gvar61 = 0
	gvar62 = 0
	gvar63 = 0
	flag78 = false
	ivar64 = 0
	ivar65 = 12
	flag86 = false
}
func InitializeDialogs() {
	ns.SetDialog(obj4, ns.NORMAL, JandorIntroStart, JandorIntroEnd)
}
func OwnObjects() {
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj6)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(ns.GetHost(), obj9)
	ns.JournalEntry(ns.GetHost(), "FindHorvath", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.Music(21, 100)
	ns.SetDialog(obj4, ns.NORMAL, JandorIntroStart, JandorIntroEnd)
	ns.StartDialog(obj4, ns.GetHost())
}
func OpenSecretWalls() {
	ns.ObjectOff(ns.GetTrigger())
	ns.WallGroupOpen(gvar44)
}
func MirrorText() {
	ns.PrintToAll("Wiz01A.scr:Mirror")
}
func MovingToLab() {
	if !ns.IsCaller(obj5) {
		goto LABEL1
	}
	if !flag48 {
		goto LABEL1
	}
	ns.UnlockDoor(obj17)
	ns.UnlockDoor(obj18)
	ns.Walk(obj5, 2467, 2815)
LABEL1:
	return
}
func FirstSecret() {
	ns.ObjectGroupOff(gvar27)
	ns.MoveWaypoint(wp41, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp41)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func BeachSecret2() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp41, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp41)
	ns.GiveXp(ns.GetHost(), 50)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func CaveSecret() {
	ns.ObjectGroupOff(gvar28)
	ns.MoveWaypoint(wp41, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp41)
	ns.GiveXp(ns.GetHost(), 75)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func RiverSecret() {
	ns.ObjectGroupOff(gvar29)
	ns.MoveWaypoint(wp41, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp41)
	ns.GiveXp(ns.GetHost(), 30)
	ns.PrintToAll("GeneralPrint:SecretFound")
}
func JandorIntroStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:JandorTalk01")
}
func JandorIntroEnd() {
	ns.SetDialog(obj4, ns.NORMAL, JandorDoneStart, JandorDoneEnd)
}
func ApprenticeStart() {
	ns.Damage(obj7, 0, 500, 0)
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01:ApprenticeTalk01")
}
func ApprenticeEnd() {
	ns.CancelDialog(obj6)
	ns.FrameTimer(5, KillApprentice)
}
func KillApprentice() {
	ns.MoveWaypoint(wp40, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.HumanMaleDie, wp40)
	ns.MoveObject(obj7, ns.GetObjectX(obj6), ns.GetObjectY(obj6))
	ns.MoveObject(obj6, 4279, 5285)
	ns.MoveObject(obj36, 3760, 4868)
	ns.MoveObject(obj34, 3691, 4887)
}
func PlayerInRobeRoom() {
	if !(!ns.IsLocked(obj19) || !ns.IsLocked(obj20)) {
		goto LABEL1
	}
	if flag49 {
		goto LABEL1
	}
	flag49 = true
	ns.SetDialog(obj6, ns.NORMAL, ApprenticeStart, ApprenticeEnd)
LABEL1:
	return
}
func StopSpawning() {
	ns.CreatureGuard(obj5, ns.GetObjectX(obj5), ns.GetObjectY(obj5), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	gvar60 = gvar55
	SetHorvathDialog()
}
func ScriptInit() {
	wp79 = ns.Waypoint("SP1WP")
	wp80 = ns.Waypoint("SP2WP")
	wp81 = ns.Waypoint("SP3WP")
	wp82 = ns.Waypoint("SP4WP")
	wp83 = ns.Waypoint("SP5WP")
	wp84 = ns.Waypoint("SP6WP")
	wp85 = ns.Waypoint("HorvathAttackWP")
}
func GoAway() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MapInitialize() {
	obj4 = ns.Object("Airship_Captain")
	obj5 = ns.Object("Horvath")
	obj6 = ns.Object("HorvathApprentice")
	obj7 = ns.Object("Apprentice2")
	obj8 = ns.Object("Lance")
	obj9 = ns.Object("TowerNPC")
	obj10 = ns.Object("UrchinA")
	obj11 = ns.Object("UrchinB")
	obj12 = ns.Object("FinalUrchin1")
	obj13 = ns.Object("FinalUrchin2")
	obj14 = ns.Object("HorvathHomeDoor")
	obj15 = ns.Object("HorvathGate1")
	obj16 = ns.Object("HorvathGate2")
	obj17 = ns.Object("LabDoor1")
	obj18 = ns.Object("LabDoor2")
	obj19 = ns.Object("RobeDoor1")
	obj20 = ns.Object("RobeDoor2")
	obj21 = ns.Object("FacadeGate")
	obj22 = ns.Object("HorvathGiveStuffTrigger")
	gvar23 = ns.ObjectGroup("UrchinCreateTriggers")
	gvar24 = ns.ObjectGroup("FreezeTriggers")
	gvar25 = ns.ObjectGroup("FinalTriggers")
	gvar26 = ns.ObjectGroup("UrchinTriggers")
	gvar27 = ns.ObjectGroup("FirstSecretTriggers")
	gvar28 = ns.ObjectGroup("CaveTriggers")
	gvar29 = ns.ObjectGroup("RiverTriggers")
	obj30 = ns.Object("Crystal1")
	obj31 = ns.Object("Heal1")
	obj32 = ns.Object("Heal2")
	obj33 = ns.Object("LesserHealBook")
	obj34 = ns.Object("MagicMissileBook")
	obj35 = ns.Object("HorvathStaff")
	obj36 = ns.Object("HorvathRobe")
	obj37 = ns.Object("Basket")
	obj38 = ns.Object("BasketShadow")
	obj39 = ns.Object("BriefTeleporter")
	wp40 = ns.Waypoint("PlayerSounds")
	wp41 = ns.Waypoint("SecretSounds")
	wp42 = ns.Waypoint("HorvathMoveWP")
	gvar44 = ns.WallGroup("SecretWalls")
	wp45 = ns.Waypoint("HorvathTeleporterWP")
	wp46 = ns.Waypoint("HorvathStorage")
	wp43 = ns.Waypoint("BriefWP")
	ns.LockDoor(obj15)
	ns.LockDoor(obj16)
	ns.LockDoor(obj17)
	ns.LockDoor(obj18)
	ns.LockDoor(obj21)
	ns.StoryPic(obj4, "AirshipCaptainPic")
	ns.StoryPic(obj5, "HorvathPic")
	ns.StoryPic(obj6, "WoundedApprenticePic")
	ns.StoryPic(obj8, "WizardGuard1Pic")
	ns.StartupScreen(1)
	ScriptInit()
	ns.FrameTimer(1, OwnObjects)
	ns.FrameTimer(1, StopSpawning)
}
func PlayerDeath() {
	ns.DeathScreen(1)
}
func TeleportHorvath() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp43), ns.GetWaypointY(wp43))
}
func HorvathGateDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "Wiz02A.scr:Horvath1")
}
func HorvathGateDialogEnd() {
	ns.AudioEvent(ns.TeleportOut, ns.Waypoint("FacadeAudioOrigin"))
	ns.Effect(ns.TELEPORT, ns.GetObjectX(obj5), ns.GetObjectY(obj5), 0, 0)
	ns.MoveObject(obj5, ns.GetWaypointX(wp46), ns.GetWaypointY(wp46))
	ns.Frozen(ns.GetHost(), false)
	ns.JournalEntry(ns.GetHost(), "MeetHorvath", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
}
func LanceDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "War01A.scr:Guard1Talk01")
}
func LanceDialogEnd() {
}
func HorvathSpeak() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp40, ns.GetWaypointX(wp43), ns.GetWaypointY(wp43))
	ns.AudioEvent(ns.TeleportIn, wp40)
	ns.Effect(ns.TELEPORT, ns.GetWaypointX(wp43), ns.GetWaypointY(wp43), ns.GetWaypointX(wp43), ns.GetWaypointY(wp43))
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj9)
	ns.MoveObject(obj5, ns.GetWaypointX(wp45), ns.GetWaypointY(wp45))
}
func PlayerArrivesAtGate() {
	ns.ObjectOff(ns.GetTrigger())
	ns.StartupScreen(2)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(ns.Waypoint("FacadeAudioOrigin"))) // FIXME
	ns.FrameTimer(1, TeleportHorvath)
	ns.SetDialog(obj8, ns.NORMAL, LanceDialogStart, LanceDialogEnd)
}
func HorvathArrivesAtGate() {
	ns.ObjectOff(ns.GetTrigger())
	ns.CreatureIdle(obj5)
	ns.LookAtObject(obj5, ns.GetHost())
	ns.SetDialog(obj5, ns.NORMAL, HorvathGateDialogStart, HorvathGateDialogEnd)
	ns.StartDialog(obj5, ns.GetHost())
}
func SetHorvathDialog() {
	var v0 int
	v0 = gvar60
	if v0 == gvar55 {
		goto LABEL1
	}
	if v0 == gvar56 {
		goto LABEL2
	}
	if v0 == gvar57 {
		goto LABEL3
	}
	if v0 == gvar58 {
		goto LABEL4
	}
	if v0 == gvar59 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.SetDialog(obj5, ns.NORMAL, HorvathIntroStart, HorvathIntroEnd)
	goto LABEL6
LABEL2:
	ns.SetDialog(obj5, ns.NEXT, HorvathBriefStart, HorvathBriefEnd)
	goto LABEL6
LABEL3:
	ns.SetDialog(obj5, ns.NORMAL, HorvathWaitStart, HorvathWaitEnd)
	goto LABEL6
LABEL4:
	ns.SetDialog(obj5, ns.NORMAL, HorvathWait2Start, HorvathWait2End)
	goto LABEL6
LABEL5:
	ns.SetDialog(obj5, ns.NORMAL, HorvathLabStart, HorvathLabEnd)
	goto LABEL6
LABEL6:
	return
}
func HorvathRecognize() {
	if gvar60 != gvar53 {
		goto LABEL1
	}
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	gvar60 = gvar54
	ns.ClearOwner(obj10)
	ns.ClearOwner(obj11)
	ns.ObjectOn(obj5)
LABEL1:
	return
}
func HorvathIntroStart() {
	ns.Frozen(obj5, true)
	ns.CreatureIdle(obj5)
	ns.LookAtObject(obj5, ns.GetHost())
	ns.AggressionLevel(obj5, 0.5)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk01")
}
func HorvathIntroEnd() {
	ns.Frozen(obj5, false)
	ns.CancelDialog(obj5)
	ns.JournalEdit(ns.GetHost(), "FindHorvath", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.Move(obj5, wp42)
}
func HorvathBriefStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk02")
	ns.MoveWaypoint(wp40, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.LookAtObject(obj5, ns.GetHost())
}
func HorvathBriefEnd() {
	ns.Pickup(ns.GetHost(), obj30)
	ns.Pickup(ns.GetHost(), obj31)
	ns.Pickup(ns.GetHost(), obj32)
	ns.Pickup(ns.GetHost(), obj33)
	gvar60 = gvar58
	SetHorvathDialog()
	ns.StartDialog(obj5, ns.GetHost())
}
func HorvathWaitStart() {
	if !ns.HasItem(ns.GetHost(), obj36) {
		goto LABEL1
	}
	ns.LookAtObject(obj5, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk04")
	flag48 = true
	goto LABEL2
LABEL1:
	ns.LookAtObject(obj5, ns.GetHost())
	flag48 = false
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk03")
LABEL2:
	return
}
func HorvathWaitEnd() {
	if !flag48 {
		goto LABEL1
	}
	ns.Walk(obj5, 2633, 2934)
	ns.GiveXp(ns.GetHost(), 500)
	ns.JournalEdit(ns.GetHost(), "FindHorvathApprentice", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.CancelDialog(obj5)
	goto LABEL2
LABEL1:
	ns.SetDialog(obj5, ns.NORMAL, HorvathWaitStart, HorvathWaitEnd)
LABEL2:
	return
}
func HorvathWait2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk02a")
}
func HorvathWait2End() {
	ns.UnlockDoor(obj15)
	ns.UnlockDoor(obj16)
	gvar60 = gvar57
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "FindHorvathApprentice", 2)
	ns.LookAtObject(obj5, obj15)
	ns.CreatureGuard(obj5, ns.GetObjectX(obj5), ns.GetObjectY(obj5), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	SetHorvathDialog()
}
func HorvathLabStart() {
	ns.LookAtObject(obj5, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:HorvathTalk05")
}
func HorvathLabEnd() {
	ns.SetDialog(obj5, ns.NORMAL, HorvathLabStart, HorvathLabEnd)
}
func HorvathGoHome() {
	if !ns.IsCaller(obj5) {
		goto LABEL1
	}
	ns.Wander(obj5)
LABEL1:
	return
}
func HorvathOpenHomeDoor() {
	ns.UnlockDoor(obj14)
}
func HorvathAtGiveStuff() {
	ns.CreatureGuard(obj5, ns.GetObjectX(obj5), ns.GetObjectY(obj5), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	gvar60 = gvar56
	ns.ObjectOff(obj22)
	ns.LookAtObject(obj5, ns.GetHost())
	SetHorvathDialog()
}
func SayTeleportStuff() {
	ns.CreatureGuard(obj5, 2501, 2865, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	gvar60 = gvar59
	SetHorvathDialog()
}
func HorvathFacePlayer() {
	if flag48 {
		goto LABEL1
	}
	ns.LookAtObject(obj5, ns.GetHost())
	ns.FrameTimer(15, HorvathFacePlayer)
LABEL1:
	return
}
func SPDie() {
	ivar64 += 1
	if !(ivar64 >= ivar65) {
		goto LABEL1
	}
	ns.WideScreen(false)
	ns.Frozen(obj5, true)
	ns.FrameTimer(1, HorvathStop)
LABEL1:
	if flag51 {
		goto LABEL2
	}
	flag51 = true
	ns.SecondTimer(10, DoubleCheck)
LABEL2:
	return
}
func HorvathStop() {
	ns.Frozen(obj5, true)
	ns.CreatureIdle(obj5)
	ns.LookAtObject(obj5, ns.GetHost())
	ns.AggressionLevel(obj5, 0)
	ns.FrameTimer(1, HorvathGuard)
}
func HorvathGuard() {
	ns.Frozen(obj5, false)
	ns.CreatureGuard(obj5, ns.GetObjectX(obj5), ns.GetObjectY(obj5), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	SetHorvathDialog()
	ns.FrameTimer(15, HorvathHi)
}
func HorvathHi() {
	ns.Frozen(ns.GetHost(), false)
	ns.StartDialog(obj5, ns.GetHost())
}
func SetHorvathLook() {
	ns.CreatureGuard(obj5, ns.GetObjectX(obj5), ns.GetObjectY(obj5), ns.GetObjectX(obj15), ns.GetObjectY(obj15), 0)
}
func DoubleCheck() {
	ns.Damage(obj66, 0, 50, 9)
	ns.Damage(obj67, 0, 50, 9)
	ns.Damage(obj68, 0, 50, 9)
	ns.Damage(obj69, 0, 50, 9)
	ns.Damage(obj70, 0, 50, 9)
	ns.Damage(obj71, 0, 50, 9)
	ns.Damage(obj72, 0, 50, 9)
	ns.Damage(obj73, 0, 50, 9)
	ns.Damage(obj74, 0, 50, 9)
	ns.Damage(obj75, 0, 50, 9)
	ns.Damage(obj76, 0, 50, 9)
	ns.Damage(obj77, 0, 50, 9)
}
func JandorDoneStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz01A.scr:JandorTalk02")
}
func JandorDoneEnd() {
	ns.SetDialog(obj4, ns.NORMAL, JandorDoneStart, JandorDoneEnd)
}
func DestroyCaptain() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar23)
	ns.Delete(obj4)
	ns.Delete(obj37)
	ns.Delete(obj38)
LABEL1:
	return
}
func FinalUrchins() {
	obj52 = ns.GetLastItem(ns.GetHost())
	for {
		if obj52 == 0 {
			goto LABEL1
		}
		if obj52 != obj36 {
			goto LABEL2
		}
		flag48 = true
		ns.ObjectGroupOff(gvar25)
		ns.MoveObject(obj12, 3449, 2378)
		ns.MoveObject(obj13, 3516, 2425)
		ns.ObjectOn(obj12)
		ns.ObjectOn(obj13)
		goto LABEL1
	LABEL2:
		obj52 = ns.GetPreviousItem(obj52)
	}
LABEL1:
	return
}
func HorrendousPainting() {
	ns.PrintToAll("Wiz01A.scr:UrchinPainting01")
}
func UrchinPainting1() {
	ns.PrintToAll("Wiz01A.scr:UrchinPainting2")
}
func UrchinPainting2() {
	ns.PrintToAll("Wiz01A.scr:UrchinPainting3")
}
func UrchinAEnable() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(obj10, 0.83)
LABEL1:
	return
}
func UrchinBEnable() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(obj11, 0.83)
LABEL1:
	return
}
func UrchinAway() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func UrchinReport() {
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
}
func DenReport() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func UrchinSetup() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if flag50 {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar26)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj68)
	ns.WideScreen(true)
	ns.ClearOwner(obj10)
	ns.ClearOwner(obj11)
	obj66 = ns.CreateObject("Urchin", wp79)
	obj67 = ns.CreateObject("Urchin", wp80)
	obj68 = ns.CreateObject("Urchin", wp81)
	obj69 = ns.CreateObject("Urchin", wp82)
	obj70 = ns.CreateObject("Urchin", wp83)
	obj71 = ns.CreateObject("Urchin", wp84)
	flag50 = true
	ns.FrameTimer(1, UrchinBegin)
LABEL1:
	return
}
func UrchinBegin() {
	if !(ns.CurrentHealth(obj10) > 0) {
		goto LABEL1
	}
	ns.CreatureFollow(obj10, ns.GetHost())
LABEL1:
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL2
	}
	ns.CreatureFollow(obj11, ns.GetHost())
LABEL2:
	ns.SetCallback(obj66, 5, SPDie)
	ns.SetCallback(obj67, 5, SPDie)
	ns.SetCallback(obj68, 5, SPDie)
	ns.SetCallback(obj69, 5, SPDie)
	ns.SetCallback(obj70, 5, SPDie)
	ns.SetCallback(obj71, 5, SPDie)
	ns.SetCallback(obj66, 3, AttackHorvath)
	ns.SetCallback(obj67, 3, AttackHorvath)
	ns.SetCallback(obj68, 3, AttackHorvath)
	ns.SetCallback(obj69, 3, AttackHorvath)
	ns.SetCallback(obj70, 3, AttackHorvath)
	ns.SetCallback(obj71, 3, AttackHorvath)
	ns.AggressionLevel(obj66, 0)
	ns.AggressionLevel(obj67, 0)
	ns.AggressionLevel(obj68, 0)
	ns.AggressionLevel(obj69, 0)
	ns.AggressionLevel(obj70, 0)
	ns.AggressionLevel(obj71, 0)
	ns.CreatureFollow(obj66, ns.GetHost())
	ns.CreatureFollow(obj67, ns.GetHost())
	ns.CreatureFollow(obj68, ns.GetHost())
	ns.CreatureFollow(obj69, ns.GetHost())
	ns.CreatureFollow(obj70, ns.GetHost())
	ns.CreatureFollow(obj71, ns.GetHost())
	ns.RetreatLevel(obj66, 0)
	ns.RetreatLevel(obj67, 0)
	ns.RetreatLevel(obj68, 0)
	ns.RetreatLevel(obj69, 0)
	ns.RetreatLevel(obj70, 0)
	ns.RetreatLevel(obj71, 0)
	ns.SecondTimer(1, UrchinBegin2)
}
func UrchinBegin2() {
	obj72 = ns.CreateObject("Urchin", wp79)
	obj73 = ns.CreateObject("Urchin", wp80)
	obj74 = ns.CreateObject("Urchin", wp81)
	obj75 = ns.CreateObject("Urchin", wp82)
	obj76 = ns.CreateObject("Urchin", wp83)
	obj77 = ns.CreateObject("Urchin", wp84)
	ns.FrameTimer(1, UrchinGroup2Go)
}
func UrchinGroup2Go() {
	ns.SetCallback(obj72, 5, SPDie)
	ns.SetCallback(obj73, 5, SPDie)
	ns.SetCallback(obj74, 5, SPDie)
	ns.SetCallback(obj75, 5, SPDie)
	ns.SetCallback(obj76, 5, SPDie)
	ns.SetCallback(obj77, 5, SPDie)
	ns.SetCallback(obj72, 3, AttackHorvath)
	ns.SetCallback(obj73, 3, AttackHorvath)
	ns.SetCallback(obj74, 3, AttackHorvath)
	ns.SetCallback(obj75, 3, AttackHorvath)
	ns.SetCallback(obj76, 3, AttackHorvath)
	ns.SetCallback(obj77, 3, AttackHorvath)
	ns.AggressionLevel(obj72, 0)
	ns.AggressionLevel(obj73, 0)
	ns.AggressionLevel(obj74, 0)
	ns.AggressionLevel(obj75, 0)
	ns.AggressionLevel(obj76, 0)
	ns.AggressionLevel(obj77, 0)
	ns.CreatureFollow(obj72, obj66)
	ns.CreatureFollow(obj73, obj67)
	ns.CreatureFollow(obj74, obj68)
	ns.CreatureFollow(obj75, obj69)
	ns.CreatureFollow(obj76, obj70)
	ns.CreatureFollow(obj77, obj71)
	ns.RetreatLevel(obj72, 0)
	ns.RetreatLevel(obj73, 0)
	ns.RetreatLevel(obj74, 0)
	ns.RetreatLevel(obj75, 0)
	ns.RetreatLevel(obj76, 0)
	ns.RetreatLevel(obj77, 0)
	ns.SecondTimer(2, HorvathApproach)
}
func RenewAttack() {
	ns.HitFarLocation(ns.GetTrigger(), ns.GetObjectX(obj5), ns.GetObjectY(obj5))
}
func HorvathApproach() {
	ns.Move(obj5, wp85)
	if !(ns.CurrentHealth(obj10) > 0) {
		goto LABEL1
	}
	ns.AggressionLevel(obj10, 0)
	ns.GoBackHome(obj10)
LABEL1:
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL2
	}
	ns.AggressionLevel(obj11, 0)
	ns.GoBackHome(obj11)
LABEL2:
	ns.SetCallback(obj66, 11, RenewAttack)
	ns.SetCallback(obj67, 11, RenewAttack)
	ns.SetCallback(obj68, 11, RenewAttack)
	ns.SetCallback(obj69, 11, RenewAttack)
	ns.SetCallback(obj70, 11, RenewAttack)
	ns.SetCallback(obj71, 11, RenewAttack)
	ns.SetCallback(obj72, 11, RenewAttack)
	ns.SetCallback(obj73, 11, RenewAttack)
	ns.SetCallback(obj74, 11, RenewAttack)
	ns.SetCallback(obj75, 11, RenewAttack)
	ns.SetCallback(obj76, 11, RenewAttack)
	ns.SetCallback(obj77, 11, RenewAttack)
	ns.SecondTimer(3, StopTheSpawning)
}
func AttackHorvath() {
	if !ns.IsCaller(obj5) {
		goto LABEL1
	}
	ns.HitFarLocation(ns.GetTrigger(), ns.GetObjectX(obj5), ns.GetObjectY(obj5))
LABEL1:
	return
}
func HorvathArrived() {
	ns.AggressionLevel(obj5, 0.83)
}
func StopTheSpawning() {
	flag86 = true
}
func UrchinSecret1() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp41, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp41)
	ns.GiveXp(ns.GetHost(), 25)
	ns.PrintToAll("GeneralPrint:SecretFound")
LABEL1:
	return
}
func UrchinSecret2() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp41, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp41)
	ns.GiveXp(ns.GetHost(), 25)
	ns.PrintToAll("GeneralPrint:SecretFound")
LABEL1:
	return
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
