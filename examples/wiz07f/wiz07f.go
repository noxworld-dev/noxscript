package wiz07f

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
	gvar38 ns.WallGroupID
	gvar39 ns.WallGroupID
	gvar40 ns.ObjectGroupID
	gvar41 ns.ObjectGroupID
	wp42   ns.WaypointID
	gvar43 int
	gvar44 int
	gvar45 int
	gvar46 int
	gvar47 int
	gvar48 int
	gvar49 int
	gvar50 int
	flag51 bool
)

func init() {
	gvar43 = 100
	gvar44 = 0
	gvar45 = 1
	gvar46 = 2
	gvar47 = 3
	gvar48 = 4
	gvar49 = 5
	gvar50 = gvar44
	flag51 = false
}
func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj18)
	ns.SetOwner(ns.GetHost(), obj16)
	ns.SetOwner(ns.GetHost(), obj17)
	ns.SetOwner(ns.GetHost(), obj19)
	ns.SetOwner(ns.GetHost(), obj20)
	ns.SetOwner(ns.GetHost(), obj21)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj11)
	ns.SetOwner(ns.GetHost(), obj7)
	ns.SetOwner(ns.GetHost(), obj14)
	ns.SetOwner(ns.GetHost(), obj13)
	ns.SetOwner(ns.GetHost(), obj12)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.SetOwner(ns.GetHost(), obj24)
	ns.SetOwner(ns.GetHost(), obj25)
	ns.SetOwner(ns.GetHost(), obj26)
	ns.SetOwner(ns.GetHost(), obj27)
	if ns.GetQuestStatus("OwnMax") != 1 {
		goto LABEL1
	}
	ns.UnlockDoor(obj28)
	ns.UnlockDoor(obj29)
LABEL1:
	return
}
func PriestStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj11)
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:PriestTalk01")
}
func PriestEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj11, ns.NORMAL, PriestStart, PriestEnd)
}
func GrillfStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:GrillfTalk01")
}
func GrillfEnd() {
	ns.SetDialog(obj13, ns.NORMAL, GrillfStart, GrillfEnd)
}
func UndertakerStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj11)
	ns.TellStory(ns.HumanMaleEatFood, "Wiz02A.scr:UndertakerTalk01")
}
func UndertakerEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj12, ns.NORMAL, UndertakerStart, UndertakerEnd)
}
func MaxDialogStart() {
	var v0 int
	v0 = gvar50
	if v0 == gvar44 {
		goto LABEL1
	}
	if v0 == gvar45 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War07a.scr:MaxOffer01")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MaxWaiting")
	goto LABEL3
LABEL3:
	return
}
func MaxDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v0 = ns.GetAnswer(obj7)
	v5 = v0
	if v5 == v4 {
		goto LABEL1
	}
	if v5 == v3 {
		goto LABEL2
	}
	if v5 == v2 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	gvar50 = gvar45
	ResetMaxDialog()
	goto LABEL4
LABEL2:
	ns.SetDialog(obj7, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:TurnMaxDown")
	gvar50 = gvar45
	goto LABEL4
LABEL3:
	v1 = ns.GetGold(ns.GetHost())
	if !(v1 >= 50000) {
		goto LABEL5
	}
	ns.ObjectGroupOff(gvar40)
	ns.SetDialog(obj7, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:HaveGoldForMax")
	ns.ChangeGold(ns.GetHost(), -50000)
	gvar50 = gvar49
	ns.UnlockDoor(obj28)
	ns.UnlockDoor(obj29)
	ns.SetQuestStatus(1, "OwnMax")
	goto LABEL4
	goto LABEL4
LABEL5:
	ns.SetDialog(obj7, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:NoGoldForMax")
	gvar50 = gvar45
	goto LABEL4
LABEL4:
	return
}
func AlbiStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj18, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj18)
	ns.LookAtObject(obj18, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj18)
	ns.TellStory(ns.HumanMaleEatFood, "Con07:HecRaisesDead")
}
func AlbiEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj18, false)
	ns.Wander(obj18)
}
func DorianStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj19, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj19)
	ns.LookAtObject(obj19, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj19)
	ns.TellStory(ns.HumanMaleEatFood, "Con07:LeavingGalava2")
}
func DorianEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj19, false)
	ns.Wander(obj19)
}
func GrunbarStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj17, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj17)
	ns.LookAtObject(obj17, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj17)
	ns.TellStory(ns.HumanMaleEatFood, "Con07:LeavingGalava1")
}
func GrunbarEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj17, false)
	ns.Wander(obj17)
}
func EowynnStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj25, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj25)
	ns.LookAtObject(obj25, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj17)
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:MaidenTalk03")
}
func EowynnEnd() {
	ns.Frozen(obj25, false)
	ns.Frozen(ns.GetHost(), false)
	ns.Wander(obj25)
}
func ExitTrigger() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.UnlockDoor(obj28)
	ns.UnlockDoor(obj29)
LABEL1:
	return
}
func BoothStart() {
	ns.Frozen(obj26, true)
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(ns.GetHost(), obj26)
	ns.LookAtObject(obj26, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War03b:T4Post")
}
func BoothEnd() {
	ns.Frozen(obj26, false)
	ns.Frozen(ns.GetHost(), false)
}
func Jump() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj36) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj37) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar38)
	ns.Chat(obj34, "Con02a:BarMaiden2")
	ns.SecondTimer(3, Upset)
LABEL1:
	return
}
func Upset() {
	ns.DestroyEveryChat()
	ns.Chat(obj33, "Wiz02B.scr:LewisTalk06")
	ns.SecondTimer(3, StartJump)
}
func StartJump() {
	ns.WallGroupOpen(gvar39)
	ns.Wander(obj33)
}
func MoveJumper() {
	if !ns.IsCaller(obj33) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.CreatureIdle(obj33)
	ns.ObjectOn(obj35)
	ns.Chat(obj33, "Wiz02B.scr:LewisTalk03")
LABEL1:
	return
}
func KillJumper() {
	if !ns.IsCaller(obj33) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.DestroyEveryChat()
	ns.Damage(obj33, 0, 500, 0)
	ns.WallGroupClose(gvar39)
	ns.WallGroupClose(gvar38)
	ns.NoWallSound(false)
LABEL1:
	return
}
func NewMaxStart() {
	ns.Frozen(obj7, true)
	ns.Frozen(ns.GetHost(), true)
	ns.LookAtObject(ns.GetHost(), obj7)
	ns.LookAtObject(obj7, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MaxOffer01")
}
func NewMaxEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj7, false)
}
func DemonAway() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func MapInitialize() {
	obj33 = ns.Object("TowerNPC")
	obj34 = ns.Object("TowerMaiden")
	obj35 = ns.Object("NPCMover")
	obj36 = ns.Object("SecretShirt")
	obj37 = ns.Object("SecretApple")
	gvar38 = ns.WallGroup("JumpWalls")
	gvar39 = ns.WallGroup("FallWalls")
	obj4 = ns.Object("BrightBlades")
	obj5 = ns.Object("Loreman")
	obj6 = ns.Object("MageVendor")
	obj7 = ns.Object("Max")
	obj8 = ns.Object("MiscGuy")
	obj9 = ns.Object("BookVendor")
	obj10 = ns.Object("AppleMan")
	obj11 = ns.Object("Priest")
	obj12 = ns.Object("Undertaker")
	obj13 = ns.Object("Grillf")
	obj14 = ns.Object("Mlurgh")
	obj15 = ns.Object("Warden")
	obj16 = ns.Object("Jorgan")
	obj17 = ns.Object("Grunbar")
	obj18 = ns.Object("Albi")
	obj19 = ns.Object("Dorian")
	obj20 = ns.Object("Civvy3")
	obj21 = ns.Object("Civvy4")
	obj22 = ns.Object("Morgan")
	obj25 = ns.Object("Eowynn")
	obj24 = ns.Object("Shari")
	obj23 = ns.Object("Kayla")
	obj26 = ns.Object("F6Elite4")
	obj27 = ns.Object("Maiden3")
	obj28 = ns.Object("MaxHomeDoor1")
	obj29 = ns.Object("MaxHomeDoor2")
	obj30 = ns.Object("MorganCellDoor")
	obj31 = ns.Object("GuardDoor1")
	obj32 = ns.Object("GearRoomDoor")
	gvar40 = ns.ObjectGroup("MaxWelcomeTriggers")
	gvar41 = ns.ObjectGroup("ResetMaxTriggers")
	wp42 = ns.Waypoint("PlayerSounds")
	ns.LockDoor(obj28)
	ns.LockDoor(obj29)
	ns.LockDoor(obj30)
	ns.LockDoor(obj31)
	ns.LockDoor(obj32)
	ns.Wander(obj19)
	ns.Wander(obj18)
	ns.Wander(obj17)
	ns.Wander(obj16)
	ns.Wander(obj20)
	ns.Wander(obj21)
	ns.SetDialog(obj11, ns.NORMAL, PriestStart, PriestEnd)
	ns.SetDialog(obj12, ns.NORMAL, UndertakerStart, UndertakerEnd)
	ns.SetDialog(obj7, ns.NORMAL, NewMaxStart, NewMaxEnd)
	ns.SetDialog(obj18, ns.NORMAL, AlbiStart, AlbiEnd)
	ns.SetDialog(obj19, ns.NORMAL, DorianStart, DorianEnd)
	ns.SetDialog(obj17, ns.NORMAL, GrunbarStart, GrunbarEnd)
	ns.SetDialog(obj25, ns.NORMAL, EowynnStart, EowynnEnd)
	ns.SetDialog(obj26, ns.NORMAL, BoothStart, BoothEnd)
	ns.StoryPic(obj26, "MalePic4")
	ns.StoryPic(obj15, "WardenPic")
	ns.StoryPic(obj7, "MaxPic")
	ns.StoryPic(obj14, "IxGuard2Pic")
	ns.StoryPic(obj12, "UndertakerPic")
	ns.StoryPic(obj19, "Townsman2Pic")
	ns.StoryPic(obj17, "MalePic7")
	ns.StoryPic(obj18, "MalePic1")
	ns.StoryPic(obj16, "Townsman3Pic")
	ns.StoryPic(obj11, "GalavaPriestPic")
	ns.StoryPic(obj25, "MaidenPic")
}
func MapEntry() {
	ns.FrameTimer(1, CreatureSetup)
	ns.Music(17, 100)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func ResetMaxDialog() {
	if gvar50 != gvar49 {
		goto LABEL1
	}
	ns.CancelDialog(obj7)
	goto LABEL2
LABEL1:
	ns.SetDialog(obj7, ns.YESNO, MaxDialogStart, MaxDialogEnd)
LABEL2:
	return
}
func NullFunction() {
}
func SetMaxWelcome() {
	flag51 = false
}
func MaxSayWelcome() {
	if flag51 {
		goto LABEL1
	}
	ns.Chat(obj7, "Wiz02A.scr:MaxTalk01")
	flag51 = true
	if ns.GetQuestStatus("OwnMax") != 1 {
		goto LABEL1
	}
	ns.ObjectGroupOff(gvar41)
	ns.MoveObject(obj7, 1514, 4800)
	ns.FrameTimer(1, DisableMax)
LABEL1:
	return
}
func DisableMax() {
	ns.ObjectOff(obj7)
}
func MakeAWish() {
	var v0 int
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp42)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "MapEntry":
		MapEntry()
	case "PlayerDeath":
		PlayerDeath()
	}
}
