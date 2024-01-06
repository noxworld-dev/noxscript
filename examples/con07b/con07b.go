package con07b

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
	obj38  ns.ObjectID
	obj39  ns.ObjectID
	gvar40 ns.WallGroupID
	gvar41 ns.WallGroupID
	wp42   ns.WaypointID
	wp43   ns.WaypointID
	ivar44 int
	gvar45 int
	flag46 bool
	gvar47 int
	gvar48 int
	gvar49 int
	gvar50 int
	gvar51 int
	gvar52 int
	gvar53 int
	flag54 bool
	flag55 bool
	flag56 bool
	flag57 bool
	flag58 bool
	flag59 bool
)

func init() {
	ivar44 = 0
	gvar45 = 25
	flag46 = true
	gvar47 = 0
	gvar48 = 1
	gvar49 = 2
	gvar50 = 3
	gvar51 = 4
	gvar52 = 5
	gvar53 = gvar47
	flag54 = false
	flag55 = false
	flag56 = false
	flag57 = false
	flag58 = false
	flag59 = false
}
func DialogInit() {
	obj4 = ns.Object("Albi")
	obj5 = ns.Object("Dorian")
	obj6 = ns.Object("Grunbar")
	obj7 = ns.Object("Jorgan")
	obj8 = ns.Object("Civvy3")
	obj9 = ns.Object("Civvy4")
	obj10 = ns.Object("Eowynn")
	obj12 = ns.Object("Shari")
	obj11 = ns.Object("Kayla")
	ns.StoryPic(obj5, "Townsman2Pic")
	ns.StoryPic(obj6, "MalePic7")
	ns.StoryPic(obj4, "MalePic1")
	ns.StoryPic(obj7, "Townsman3Pic")
	ns.StoryPic(obj10, "MaidenPic")
	ns.StoryPic(obj11, "MaidenPic2")
	ns.StoryPic(obj12, "MaidenPic3")
	ns.SetDialog(obj10, ns.NORMAL, EowynnStart, EowynnEnd)
	ns.SetDialog(obj11, ns.NORMAL, KaylaStart, KaylaEnd)
	ns.SetDialog(obj12, ns.NORMAL, ShariStart, ShariEnd)
	ns.SetDialog(obj7, ns.NORMAL, JorganStart, JorganEnd)
	ns.SetDialog(obj4, ns.NORMAL, AlbiStart, AlbiEnd)
	ns.SetDialog(obj5, ns.NORMAL, DorianStart, DorianEnd)
	ns.SetDialog(obj6, ns.NORMAL, GrunbarStart, GrunbarEnd)
	ns.FrameTimer(1, OwnTownies)
}
func OwnTownies() {
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj6)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.SetOwner(ns.GetHost(), obj7)
	ns.SetOwner(ns.GetHost(), obj8)
	ns.SetOwner(ns.GetHost(), obj9)
	ns.SetOwner(ns.GetHost(), obj10)
	ns.SetOwner(ns.GetHost(), obj12)
	ns.SetOwner(ns.GetHost(), obj11)
}
func EowynnStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj10, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj10)
	ns.LookAtObject(obj10, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj10)
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:MaidenTalk03")
}
func EowynnEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj10, false)
	ns.Wander(obj10)
	ns.SetDialog(obj10, ns.NORMAL, EowynnStart, EowynnEnd)
}
func ShariStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj12, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj12)
	ns.LookAtObject(obj12, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj12)
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:MaidenTalk02")
}
func ShariEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj12, false)
	ns.Wander(obj12)
	ns.SetDialog(obj12, ns.NORMAL, ShariStart, ShariEnd)
}
func KaylaStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj11, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj11)
	ns.LookAtObject(ns.GetHost(), obj11)
	ns.LookAtObject(obj11, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:MaidenTalk01")
}
func KaylaEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj11, false)
	ns.Wander(obj11)
	ns.SetDialog(obj11, ns.NORMAL, KaylaStart, KaylaEnd)
}
func JorganStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj7, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj7)
	ns.LookAtObject(ns.GetHost(), obj7)
	ns.LookAtObject(obj7, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con07:HecRaisesDead")
}
func JorganEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj7, false)
	ns.Wander(obj7)
	ns.SetDialog(obj7, ns.NORMAL, JorganStart, JorganEnd)
}
func AlbiStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj4, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj4)
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.LookAtObject(obj4, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con07:LeavingGalava2")
}
func AlbiEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj4, false)
	ns.Wander(obj4)
	ns.SetDialog(obj4, ns.NORMAL, AlbiStart, AlbiEnd)
}
func DorianStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj5, true)
	ns.CreatureIdle(ns.GetHost())
	ns.CreatureIdle(obj5)
	ns.LookAtObject(ns.GetHost(), obj5)
	ns.LookAtObject(obj5, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con07:LeavingGalava1")
}
func DorianEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj5, false)
	ns.Wander(obj5)
	ns.SetDialog(obj5, ns.NORMAL, DorianStart, DorianEnd)
}
func GrunbarStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War03b:T4Post")
}
func GrunbarEnd() {
	ns.Frozen(ns.GetHost(), false)
}
func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj20)
	ns.SetOwner(ns.GetHost(), obj19)
	ns.SetOwner(ns.GetHost(), obj22)
	ns.SetOwner(ns.GetHost(), obj21)
	ns.SetOwner(ns.GetHost(), obj15)
	ns.SetOwner(ns.GetHost(), obj28)
	ns.SetOwner(ns.GetHost(), obj25)
	ns.SetOwner(ns.GetHost(), obj26)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.SetOwner(ns.GetHost(), obj27)
	ns.SetOwner(ns.GetHost(), obj35)
	ns.SetOwner(ns.GetHost(), obj36)
	ns.Music(17, 100)
}
func CivvyBump() {
	var v0 int
	if !flag46 {
		goto LABEL1
	}
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ivar44 = ns.Random(1, 4)
	v0 = ivar44
	if v0 == 1 {
		goto LABEL2
	}
	if v0 == 2 {
		goto LABEL3
	}
	if v0 == 3 {
		goto LABEL4
	}
	if v0 == 4 {
		goto LABEL5
	}
	goto LABEL6
LABEL2:
	ns.Chat(ns.GetTrigger(), "Wiz02A.scr:CivvyTalk01")
	goto LABEL6
LABEL3:
	ns.Chat(ns.GetTrigger(), "Wiz02A.scr:CivvyTalk02")
	goto LABEL6
LABEL4:
	ns.Chat(ns.GetTrigger(), "Wiz02A.scr:CivvyTalk03")
	goto LABEL6
LABEL5:
	ns.Chat(ns.GetTrigger(), "Wiz02A.scr:CivvyTalk04")
	goto LABEL6
LABEL6:
	flag46 = false
	ns.SecondTimer(3, ResetCivvyReady)
LABEL1:
	return
}
func ResetCivvyReady() {
	flag46 = true
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.MoveWaypoint(wp42, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.AudioEvent(ns.RestoreHealthName, wp42)
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, ns.Waypoint("WellWP"))
}
func MaxDialogStart() {
	var v0 int
	v0 = gvar53
	if v0 == gvar47 {
		goto LABEL1
	}
	if v0 == gvar48 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MaxOffer01")
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
	v0 = ns.GetAnswer(obj28)
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
	gvar53 = gvar48
	ResetMaxDialog()
	goto LABEL4
LABEL2:
	ns.SetDialog(obj28, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:TurnMaxDown")
	gvar53 = gvar48
	goto LABEL4
LABEL3:
	v1 = ns.GetGold(ns.GetHost())
	if !(v1 >= 50000) {
		goto LABEL5
	}
	ns.SetDialog(obj28, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:HaveGoldForMax")
	ns.ChangeGold(ns.GetHost(), -50000)
	gvar53 = gvar52
	ns.UnlockDoor(obj29)
	ns.UnlockDoor(obj30)
	goto LABEL4
	goto LABEL4
LABEL5:
	ns.SetDialog(obj28, ns.NORMAL, NullFunction, ResetMaxDialog)
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:NoGoldForMax")
	gvar53 = gvar48
	goto LABEL4
LABEL4:
	return
}
func PriestTalkStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj22)
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:PriestTalk01")
}
func PriestTalkEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj22, ns.NORMAL, PriestTalkStart, PriestTalkEnd)
}
func UndertakerTalkStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj21)
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:UndertakerTalk01")
}
func UndertakerTalkEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj21, ns.NORMAL, UndertakerTalkStart, UndertakerTalkEnd)
}
func GrillfTalkStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:GrillfTalk01")
}
func GrillfTalkEnd() {
	ns.SetDialog(obj19, ns.NORMAL, GrillfTalkStart, GrillfTalkEnd)
}
func MlurghTalkStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:MlurghTalk01")
}
func MlurghTalkEnd() {
	ns.SetDialog(obj20, ns.NORMAL, MlurghTalkStart, MlurghTalkEnd)
}
func WardenTalkStart() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj23)
	ns.TellStory(ns.HumanMaleEatFood, "Con07B.scr:WardenTalk01")
}
func WardenTalkEnd() {
	ns.Frozen(ns.GetHost(), false)
	ns.SetDialog(obj23, ns.NORMAL, WardenTalkStart, WardenTalkEnd)
}
func GoAway() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func SpiderTrigger() {
	if !ns.IsCaller(obj24) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.UnlockDoor(obj29)
	ns.UnlockDoor(obj30)
LABEL1:
	return
}
func BoothStart() {
	ns.LookAtObject(obj25, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "War03b:T4Post")
}
func BoothEnd() {
}
func Jump() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj38) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj39) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar40)
	ns.Chat(obj36, "Con02a:BarMaiden2")
	ns.SecondTimer(3, Upset)
LABEL1:
	return
}
func Upset() {
	ns.DestroyEveryChat()
	ns.Chat(obj35, "Wiz02B.scr:LewisTalk06")
	ns.SecondTimer(3, StartJump)
}
func StartJump() {
	ns.WallGroupOpen(gvar41)
	ns.Wander(obj35)
}
func MoveJumper() {
	if !ns.IsCaller(obj35) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.CreatureIdle(obj35)
	ns.ObjectOn(obj37)
	ns.Chat(obj35, "Wiz02B.scr:LewisTalk03")
LABEL1:
	return
}
func KillJumper() {
	if !ns.IsCaller(obj35) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.DestroyEveryChat()
	ns.Damage(obj35, 0, 500, 0)
	ns.WallGroupClose(gvar41)
	ns.WallGroupClose(gvar40)
	ns.NoWallSound(false)
LABEL1:
	return
}
func NewMaxStart() {
	ns.TellStory(ns.HumanMaleEatFood, "War07A.scr:MaxOffer01")
}
func NewMaxEnd() {
}
func MaxSecret() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 500)
LABEL1:
	return
}
func FireballSecret() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
LABEL1:
	return
}
func MapInitialize() {
	obj35 = ns.Object("TowerNPC")
	obj36 = ns.Object("TowerMaiden")
	obj37 = ns.Object("NPCMover")
	gvar40 = ns.WallGroup("JumpWalls")
	gvar41 = ns.WallGroup("FallWalls")
	obj38 = ns.Object("SecretShirt")
	obj39 = ns.Object("SecretApple")
	obj13 = ns.Object("MageVendor")
	obj15 = ns.Object("BookVendor")
	obj14 = ns.Object("Shopkeeper1")
	obj16 = ns.Object("Kincaid")
	obj17 = ns.Object("BrightBlades")
	obj18 = ns.Object("AppleMan")
	obj19 = ns.Object("Grillf")
	obj20 = ns.Object("Mlurgh")
	obj21 = ns.Object("Undertaker")
	obj22 = ns.Object("Priest")
	obj23 = ns.Object("Warden")
	obj24 = ns.Object("GoodSpider")
	obj25 = ns.Object("F6Elite4")
	obj26 = ns.Object("Maiden3")
	obj27 = ns.Object("TowerNPC2")
	obj28 = ns.Object("Max")
	obj29 = ns.Object("MaxHomeDoor1")
	obj30 = ns.Object("MaxHomeDoor2")
	obj31 = ns.Object("GearRoomDoor")
	obj32 = ns.Object("MorganCellDoor")
	obj33 = ns.Object("FirstJailDoor")
	obj34 = ns.Object("GuardDoor1")
	wp42 = ns.Waypoint("PlayerSounds")
	wp43 = ns.Waypoint("SecretSounds")
	ns.SetDialog(obj28, ns.NORMAL, NewMaxStart, NewMaxEnd)
	ns.SetDialog(obj22, ns.NORMAL, PriestTalkStart, PriestTalkEnd)
	ns.SetDialog(obj21, ns.NORMAL, UndertakerTalkStart, UndertakerTalkEnd)
	ns.SetDialog(obj23, ns.NORMAL, WardenTalkStart, WardenTalkEnd)
	ns.SetDialog(obj25, ns.NORMAL, BoothStart, BoothEnd)
	ns.StoryPic(obj23, "WardenPic")
	ns.StoryPic(obj28, "MaxPic")
	ns.StoryPic(obj21, "UndertakerPic")
	ns.StoryPic(obj22, "GalavaPriestPic")
	ns.StoryPic(obj25, "MalePic4")
	ns.LockDoor(obj29)
	ns.LockDoor(obj30)
	ns.LockDoor(obj31)
	ns.LockDoor(obj32)
	ns.LockDoor(obj33)
	ns.LockDoor(obj34)
	DialogInit()
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func MaxEnter() {
	if flag54 {
		goto LABEL1
	}
	flag54 = true
	ns.Chat(obj28, "War07A.scr:MaxWelcome01")
LABEL1:
	return
}
func MaxExit() {
	flag54 = false
}
func ResetMaxDialog() {
	if gvar53 != gvar52 {
		goto LABEL1
	}
	ns.CancelDialog(obj28)
	goto LABEL2
LABEL1:
	ns.SetDialog(obj28, ns.YESNO, MaxDialogStart, MaxDialogEnd)
LABEL2:
	return
}
func NullFunction() {
}
func Shopkeeper1Talk() {
	if flag55 {
		goto LABEL1
	}
	flag55 = true
	ns.Chat(obj14, "Con07B.scr:ShopkeeperTalk01")
LABEL1:
	return
}
func ResetShopkeeper() {
	flag55 = false
}
func BookShopEnter() {
	if flag57 {
		goto LABEL1
	}
	flag57 = true
	ns.Chat(obj15, "Con07B.scr:BookVendorTalk01")
LABEL1:
	return
}
func BookShopExit() {
	flag57 = false
}
func MagicShopEnter() {
	if flag56 {
		goto LABEL1
	}
	flag56 = true
	ns.Chat(obj13, "Con07B.scr:MageVendorTalk01")
LABEL1:
	return
}
func MagicShopExit() {
	flag56 = false
}
func ArmorShopEnter() {
	if flag58 {
		goto LABEL1
	}
	flag58 = true
	ns.Chat(obj16, "Con07B.scr:LoremanTalk01")
LABEL1:
	return
}
func ArmorShopExit() {
	flag58 = false
}
func WeaponShopEnter() {
	if flag59 {
		goto LABEL1
	}
	flag59 = true
	ns.Chat(obj17, "Con07B.scr:BrightBladesTalk01")
LABEL1:
	return
}
func WeaponShopExit() {
	flag59 = false
}
func AppleManRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Chat(obj18, "Con07B.scr:AppleManTalk01")
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
