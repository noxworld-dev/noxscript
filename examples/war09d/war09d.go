package war09d

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
	gvar21 ns.ObjectGroupID
	gvar22 ns.ObjectGroupID
	gvar23 ns.ObjectGroupID
	gvar24 ns.ObjectGroupID
	gvar25 ns.ObjectGroupID
	gvar26 ns.ObjectGroupID
	gvar27 ns.WallGroupID
	gvar28 ns.WallGroupID
	gvar29 ns.WallGroupID
	wp30   ns.WaypointID
	wp31   ns.WaypointID
	wp32   ns.WaypointID
	wp33   ns.WaypointID
	wp34   ns.WaypointID
	wp35   ns.WaypointID
	wp36   ns.WaypointID
	wp37   ns.WaypointID
	wp38   ns.WaypointID
	wp39   ns.WaypointID
	wp40   ns.WaypointID
	wp41   [5]ns.WaypointID
	wp42   [5]ns.WaypointID
	wp43   ns.WaypointID
	gvar44 ns.WaypointGroupID
	gvar45 int
	gvar46 int
	gvar47 int
	gvar48 int
	gvar49 int
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	ivar52 int
	gvar53 int
	gvar54 int
	fvar55 float32
	fvar56 float32
	flag57 bool
	flag58 bool
	flag59 bool
	flag60 bool
	flag61 bool
	flag62 bool
	flag63 bool
	flag64 bool
	flag65 bool
	flag66 bool
	flag67 bool
	flag68 bool
	flag69 bool
	flag70 bool
	flag71 bool
	flag72 bool
	flag73 bool
)

func init() {
	gvar45 = 0
	gvar46 = 1
	gvar47 = 2
	gvar48 = 3
	gvar49 = 4
	ivar52 = 1000
	gvar53 = gvar45
	gvar54 = gvar48
	flag57 = false
	flag58 = false
	flag59 = false
	flag60 = false
	flag61 = false
	flag62 = false
	flag63 = false
	flag64 = false
	flag65 = false
	flag66 = false
	flag67 = false
	flag68 = false
	flag69 = false
	flag70 = false
	flag71 = false
	flag72 = false
	flag73 = false
}
func PlayerDeath() {
	ns.DeathScreen(9)
}
func MapInitialize() {
	obj4 = ns.Object("LoneWolf01")
	obj5 = ns.Object("Skel01a")
	obj6 = ns.Object("Skel02a")
	obj7 = ns.Object("Skel03a")
	obj8 = ns.Object("SquadLeader")
	obj9 = ns.Object("Cain")
	obj10 = ns.Object("Cain2")
	obj11 = ns.Object("CainWolf01")
	obj12 = ns.Object("CainWolf02")
	obj13 = ns.Object("CainWolf03")
	obj14 = ns.Object("Necromancer1")
	obj15 = ns.Object("Necromancer2")
	obj16 = ns.Object("Necromancer")
	obj17 = ns.Object("NecroStart")
	obj18 = ns.Object("Lich")
	obj19 = ns.Object("KeepOutTrigger")
	obj20 = ns.Object("TreasureLight")
	gvar21 = ns.ObjectGroup("GoldTrapSpikes")
	gvar22 = ns.ObjectGroup("WolfCaveTriggers")
	gvar23 = ns.ObjectGroup("SkeletonSquad")
	gvar25 = ns.ObjectGroup("NecroTriggers")
	gvar24 = ns.ObjectGroup("ElevatorGroup1")
	gvar26 = ns.ObjectGroup("SkeletonMarchTriggers")
	gvar27 = ns.WallGroup("SecretHallWalls01")
	gvar28 = ns.WallGroup("NecroWalls")
	gvar29 = ns.WallGroup("NecroSafetyWalls")
	wp30 = ns.Waypoint("NecroWP")
	wp31 = ns.Waypoint("GoldTrapWP1")
	wp32 = ns.Waypoint("GoldTrapWP2")
	wp33 = ns.Waypoint("GoldTrapWP3")
	wp34 = ns.Waypoint("GoldTrapWP4")
	wp37 = ns.Waypoint("WolfCreateWP01")
	wp38 = ns.Waypoint("WolfCreateWP02")
	wp39 = ns.Waypoint("LoneWolf01DestWP")
	wp40 = ns.Waypoint("Wolf01FacingWP")
	wp35 = ns.Waypoint("WolfCreate01Dest")
	wp36 = ns.Waypoint("WolfCreate02Dest")
	wp41[0] = ns.Waypoint("TeleWP0")
	wp41[1] = ns.Waypoint("TeleWP1")
	wp41[2] = ns.Waypoint("TeleWP2")
	wp41[3] = ns.Waypoint("TeleWP3")
	wp41[4] = ns.Waypoint("TeleWP4")
	wp42[0] = ns.Waypoint("WallSmokeWP0")
	wp42[1] = ns.Waypoint("WallSmokeWP1")
	wp42[2] = ns.Waypoint("WallSmokeWP2")
	wp42[3] = ns.Waypoint("WallSmokeWP3")
	wp42[4] = ns.Waypoint("WallSmokeWP4")
	wp43 = ns.Waypoint("PlayerSounds")
	gvar44 = ns.WaypointGroup("SearchNetwork1")
	ns.StoryPic(obj9, "WoundedConjurerPic")
	ns.StoryPic(obj10, "WoundedConjurerPic")
	ns.StoryPic(obj16, "NecromancerPic")
	ns.UnBlind()
}
func PlaySwamp() {
	ns.Music(25, 100)
}
func PlayAction() {
	ns.Music(26, 100)
}
func PlayAction2() {
	ns.Music(28, 100)
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
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(19, 100)
LABEL1:
	return
}
func Patrol() {
	ns.Wander(ns.GetTrigger())
}
func GoMedieval() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
LABEL1:
	return
}
func GoNormal() {
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
}
func StopAndListen() {
	var v0 int
	v0 = ns.Random(1, 3)
	if !(ns.IsCaller(ns.GetHost()) && v0 == 1) {
		goto LABEL1
	}
	ns.PauseObject(ns.GetTrigger(), 45)
	ns.LookAtObject(ns.GetTrigger(), ns.GetCaller())
	if !ns.IsVisibleTo(ns.GetTrigger(), ns.GetCaller()) {
		goto LABEL2
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
	goto LABEL1
LABEL2:
	ns.CreatureIdle(ns.GetTrigger())
LABEL1:
	return
}
func ReturnHome() {
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
	ns.GoBackHome(ns.GetTrigger())
}
func DisableGoldTrapSpikes() {
	ns.ObjectGroupOff(gvar21)
}
func TriggerGoldTrapPoison() {
	ns.CastSpellLocationLocation(ns.SPELL_TOXIC_CLOUD, ns.GetWaypointX(wp31), ns.GetWaypointY(wp31), 0, 0)
	ns.CastSpellLocationLocation(ns.SPELL_TOXIC_CLOUD, ns.GetWaypointX(wp32), ns.GetWaypointY(wp32), 0, 0)
	ns.CastSpellLocationLocation(ns.SPELL_TOXIC_CLOUD, ns.GetWaypointX(wp33), ns.GetWaypointY(wp33), 0, 0)
	ns.CastSpellLocationLocation(ns.SPELL_TOXIC_CLOUD, ns.GetWaypointX(wp34), ns.GetWaypointY(wp34), 0, 0)
}
func TriggerGoldTrap() {
	ns.ObjectGroupOn(gvar21)
	ns.FrameTimer(20, TriggerGoldTrapPoison)
}
func ToggleElevatorGroup1() {
	ns.ObjectGroupToggle(gvar24)
}
func DisableTL() {
	ns.ObjectOff(obj20)
}
func WolfOrders() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 float32
	)
	v0 = ns.GetWaypointX(wp35)
	v1 = ns.GetWaypointY(wp35)
	v2 = ns.GetWaypointX(wp36)
	v3 = ns.GetWaypointY(wp36)
	v4 = ns.GetWaypointX(wp39)
	v5 = ns.GetWaypointY(wp39)
	ns.CreatureGuard(obj50, v0, v1, v4, v5, 350)
	ns.CreatureGuard(obj51, v2, v3, v4, v5, 350)
	ns.AggressionLevel(obj50, 0.83)
	ns.AggressionLevel(obj51, 0.83)
}
func CainDialogStart() {
	var v0 int
	v0 = gvar53
	if v0 == gvar45 {
		goto LABEL1
	}
	if v0 == gvar46 {
		goto LABEL2
	}
	if v0 == gvar47 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "War09c:CainAsk")
	flag57 = true
	goto LABEL4
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War09c:CainHealed")
	flag57 = true
	goto LABEL4
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War09c:CainEscort")
	flag57 = true
	goto LABEL4
LABEL4:
	return
}
func CainDialogEnd() {
	var (
		v0 int
		v1 ns.ObjectID
		v2 int
	)
	v2 = gvar53
	if v2 == gvar45 {
		goto LABEL1
	}
	if v2 == gvar46 {
		goto LABEL2
	}
	if v2 == gvar47 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v0 = ns.GetAnswer(obj9)
	if v0 != 1 {
		goto LABEL5
	}
	v1 = ns.GetLastItem(ns.GetHost())
	for {
		if v1 == 0 {
			goto LABEL6
		}
		if !(ns.HasClass(v1, ns.FOOD) && ns.HasSubclass(v1, ns.HEALTH_POTION)) {
			goto LABEL7
		}
		ns.Drop(ns.GetHost(), v1)
		ns.Delete(v1)
		fvar55 = ns.GetObjectX(obj9)
		fvar56 = ns.GetObjectY(obj9)
		CainRescue()
		return
	LABEL7:
		v1 = ns.GetPreviousItem(v1)
	}
LABEL6:
	ns.PrintToAll("War09c:CainNoPotion")
	goto LABEL9
LABEL5:
	ns.FrameTimer(1, CainPlea2)
LABEL9:
	goto LABEL4
LABEL2:
	ns.CreatureFollow(obj10, ns.GetHost())
	ns.BecomePet(obj10)
	ns.PrintToAll("War09c:SavedCain")
	ns.GiveXp(ns.GetHost(), 2500)
	ns.CancelDialog(obj10)
	flag61 = true
	gvar53 = gvar47
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func LichFight() {
	ns.ObjectOn(obj18)
	ns.Enchant(obj18, ns.ENCHANT_INVISIBLE, 3)
	ns.MoveObject(obj18, ns.GetWaypointX(wp41[3]), ns.GetWaypointY(wp41[3]))
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(wp41[3]), ns.GetWaypointY(wp41[3]), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp41[3]), ns.GetWaypointY(wp41[3]), 0, 0)
}
func NecroTalk() {
	if !ns.IsCaller(obj16) {
		goto LABEL1
	}
	ns.ObjectOff(obj17)
	ns.SetDialog(obj16, ns.NORMAL, NecroDialogStart, NecroDialogEnd)
	ns.StartDialog(obj16, ns.GetHost())
LABEL1:
	return
}
func MoveBody() {
	ns.Delete(obj9)
	ns.MoveObject(obj10, fvar55, fvar56)
}
func StartNecroFight2() {
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.Walk(obj16, ns.GetWaypointX(wp30), ns.GetWaypointY(wp30))
}
func WolfHarass2() {
	var (
		v0 int
		v1 int
		v2 int
	)
	v0 = ns.Random(40, 65)
	v1 = ns.Random(0, 2)
	if ns.IsObjectOn(obj11) {
		goto LABEL1
	}
	ns.ObjectOn(obj11)
	ns.ObjectOn(obj12)
	ns.ObjectOn(obj13)
LABEL1:
	v2 = v1
	if v2 == 0 {
		goto LABEL2
	}
	if v2 == 1 {
		goto LABEL3
	}
	if v2 == 2 {
		goto LABEL4
	}
	goto LABEL5
LABEL2:
	if !(flag58 == false && flag63 == false) {
		goto LABEL6
	}
	ns.HitLocation(obj11, ns.GetObjectX(obj9), ns.GetObjectY(obj9))
	goto LABEL7
LABEL6:
	v0 = 2
LABEL7:
	goto LABEL5
LABEL3:
	if !(flag59 == false && flag64 == false) {
		goto LABEL8
	}
	ns.HitLocation(obj12, ns.GetObjectX(obj9), ns.GetObjectY(obj9))
	goto LABEL9
LABEL8:
	v0 = 2
LABEL9:
	goto LABEL5
LABEL4:
	if !(flag60 == false && flag65 == false) {
		goto LABEL10
	}
	ns.HitLocation(obj13, ns.GetObjectX(obj9), ns.GetObjectY(obj9))
	goto LABEL11
LABEL10:
	v0 = 2
LABEL11:
	goto LABEL5
LABEL5:
	if !(flag63 == false || flag64 == false || flag65 == false) {
		goto LABEL12
	}
	ns.FrameTimer(v0, WolfHarass2)
	goto LABEL13
LABEL12:
	return
LABEL13:
	return
}
func CainPlea2() {
	ns.Chat(obj9, "War09c:CainPlea2")
}
func MapEntry() {
	ns.NoWallSound(false)
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
func KeepOut2() {
	if flag69 {
		goto LABEL1
	}
	if !ns.IsAttackedBy(ns.GetCaller(), ns.GetHost()) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func Secret01Declare() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 75)
}
func Secret02Declare() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp43, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp43)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 150)
}
func SkeletonMarch() {
	ns.ObjectGroupOff(gvar26)
	ns.ObjectGroupOn(gvar23)
	ns.Wander(obj5)
	ns.Wander(obj6)
	ns.Wander(obj7)
	ns.Wander(obj8)
}
func LoneWolf01Retreat() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetWaypointX(wp39)
	v1 = ns.GetWaypointY(wp39)
	v2 = ns.GetWaypointX(wp40)
	v3 = ns.GetWaypointY(wp40)
	ns.RetreatLevel(obj4, 0)
	ns.CreatureGuard(obj4, v0, v1, v2, v3, 150)
}
func WolfAmbush() {
	ns.ObjectGroupOff(gvar22)
	obj50 = ns.CreateObject("BlackWolf", wp37)
	obj51 = ns.CreateObject("BlackWolf", wp38)
	ns.FrameTimer(1, WolfOrders)
}
func HomePatrol() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.GoBackHome(ns.GetTrigger())
	ns.PauseObject(ns.GetTrigger(), 120)
	ns.Wander(ns.GetTrigger())
}
func WolfGone() {
	if !ns.IsTrigger(obj11) {
		goto LABEL1
	}
	flag58 = true
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL1
	}
	ns.AggressionLevel(obj11, 0.5)
LABEL1:
	if !ns.IsTrigger(obj12) {
		goto LABEL2
	}
	flag59 = true
	if !(ns.CurrentHealth(obj12) > 0) {
		goto LABEL2
	}
	ns.AggressionLevel(obj12, 0.5)
LABEL2:
	if !ns.IsTrigger(obj13) {
		goto LABEL3
	}
	flag60 = true
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL3
	}
	ns.AggressionLevel(obj13, 0.5)
LABEL3:
	if !(flag58 == true && flag59 == true && flag60 == true) {
		goto LABEL4
	}
	ns.DestroyChat(obj9)
	flag57 = true
	gvar54 = gvar49
	ns.SetDialog(obj9, ns.YESNO, CainDialogStart, CainDialogEnd)
LABEL4:
	return
}
func CainPlea() {
	var v0 int
	if !(ns.IsCaller(ns.GetHost()) && ns.CurrentHealth(obj9) > 0) {
		goto LABEL1
	}
	v0 = gvar54
	if v0 == gvar48 {
		goto LABEL2
	}
	if v0 == gvar49 {
		goto LABEL3
	}
	goto LABEL1
LABEL2:
	ns.Chat(obj9, "War09c:CainPlea")
	goto LABEL1
LABEL3:
	ns.Chat(obj9, "War09c:CainPlea2")
	goto LABEL1
LABEL1:
	return
}
func WolfHarass() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !(r1 && flag61 == false) {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	if !(flag61 == false && flag62 == false) {
		goto LABEL2
	}
	WolfHarass2()
LABEL2:
	return
}
func WolvesTargetPlayer() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	if flag58 {
		goto LABEL2
	}
	flag63 = true
	ns.AggressionLevel(obj11, 0.5)
	ns.Attack(obj11, ns.GetHost())
LABEL2:
	if flag59 {
		goto LABEL3
	}
	flag64 = true
	ns.AggressionLevel(obj12, 0.5)
	ns.Attack(obj12, ns.GetHost())
LABEL3:
	if flag60 {
		goto LABEL1
	}
	flag65 = true
	ns.AggressionLevel(obj13, 0.5)
	ns.Attack(obj13, ns.GetHost())
LABEL1:
	return
}
func CainDie() {
	if !(ns.CurrentHealth(obj11) > 0) {
		goto LABEL1
	}
	ns.AggressionLevel(obj11, 0.83)
LABEL1:
	if !(ns.CurrentHealth(obj12) > 0) {
		goto LABEL2
	}
	ns.AggressionLevel(obj12, 0.83)
LABEL2:
	if !(ns.CurrentHealth(obj13) > 0) {
		goto LABEL3
	}
	ns.AggressionLevel(obj13, 0.83)
LABEL3:
	fvar55 = ns.GetObjectX(obj9)
	fvar56 = ns.GetObjectY(obj9)
	flag62 = true
	ns.ObjectOn(obj10)
	ns.Damage(obj10, 0, 200000, 0)
	ns.FrameTimer(1, MoveBody)
}
func CainRescue() {
	ns.Delete(obj9)
	ns.MoveObject(obj10, fvar55, fvar56)
	ns.ObjectOn(obj10)
	ns.CreatureGuard(obj10, fvar55, fvar56, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	gvar53 = gvar46
	ns.SetDialog(obj10, ns.NORMAL, CainDialogStart, CainDialogEnd)
	ns.LookAtObject(obj10, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj10)
	ns.StartDialog(obj10, ns.GetHost())
}
func NecroDialogStart() {
	var (
		v0 int
		v1 int
	)
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL1
	}
	if v1 == 2 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.LookAtObject(obj16, ns.GetHost())
	ns.LookAtObject(obj15, ns.GetHost())
	ns.LookAtObject(obj14, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj16)
	ns.TellStory(ns.SwordsmanHurt, "Con09a:NecroThreat1")
	goto LABEL3
LABEL2:
	ns.LookAtObject(obj16, ns.GetHost())
	ns.LookAtObject(obj15, ns.GetHost())
	ns.LookAtObject(obj14, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj16)
	ns.TellStory(ns.SwordsmanHurt, "Con09a:NecroThreat4")
	goto LABEL3
LABEL3:
	return
}
func NecroDialogEnd() {
	ns.CancelDialog(obj16)
	ns.AggressionLevel(obj14, 0.83)
	ns.AggressionLevel(obj15, 0.83)
	ns.AggressionLevel(obj16, 0.83)
	ns.CreatureIdle(obj14)
	ns.CreatureIdle(obj15)
	ns.Frozen(ns.GetHost(), false)
	ns.WideScreen(false)
	ns.FrameTimer(ivar52, LichFight)
}
func StartNecroFight() {
	var v0 int
	ns.WayPointGroupOff(gvar44)
	ns.ObjectOff(obj19)
	ns.ObjectGroupOff(gvar25)
	ns.ObjectOn(obj17)
	ns.ObjectOn(obj14)
	ns.ObjectOn(obj15)
	ns.ObjectOn(obj16)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar29)
	ns.NoWallSound(false)
	ns.MusicPushEvent()
	PlayAction2()
	ns.WallGroupClose(gvar28)
	v0 = 0
	for {
		if !(v0 < 5) {
			goto LABEL1
		}
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp42[v0]), ns.GetWaypointY(wp42[v0]), 0, 0)
		v0 += 1
	}
LABEL1:
	ns.FrameTimer(1, StartNecroFight2)
}
func NecroTeleport() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = ns.Random(0, 4)
	v0 = ns.GetWaypointX(wp41[v2])
	v1 = ns.GetWaypointY(wp41[v2])
	if !(ns.IsCaller(obj14) || ns.IsCaller(obj15) || ns.IsCaller(obj16) || ns.IsCaller(obj18)) {
		goto LABEL1
	}
	ns.Enchant(ns.GetCaller(), ns.ENCHANT_INVISIBLE, 2)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.MoveObject(ns.GetCaller(), v0, v1)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
LABEL1:
	return
}
func NecroTeleportInjured() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = ns.Random(0, 4)
	v0 = ns.GetWaypointX(wp41[v2])
	v1 = ns.GetWaypointY(wp41[v2])
	if !(ns.IsTrigger(obj14) || ns.IsTrigger(obj15) || ns.IsTrigger(obj16)) {
		goto LABEL1
	}
	ns.PauseObject(ns.GetTrigger(), 15)
LABEL1:
	ns.Enchant(ns.GetTrigger(), ns.ENCHANT_INVISIBLE, 2)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), 0, 0)
	ns.MoveObject(ns.GetTrigger(), v0, v1)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
}
func HuntPlayer() {
	ns.CreatureHunt(ns.GetTrigger())
}
func OpenMagicWalls() {
	var v0 int
	if !ns.IsTrigger(obj14) {
		goto LABEL1
	}
	flag66 = true
LABEL1:
	if !ns.IsTrigger(obj15) {
		goto LABEL2
	}
	flag67 = true
LABEL2:
	if !ns.IsTrigger(obj16) {
		goto LABEL3
	}
	flag68 = true
LABEL3:
	if !ns.IsTrigger(obj18) {
		goto LABEL4
	}
	flag69 = true
LABEL4:
	if !(flag66 == true && flag67 == true && flag68 == true && flag69 == true) {
		goto LABEL5
	}
	ns.WallGroupOpen(gvar28)
	v0 = 0
	for {
		if !(v0 < 5) {
			goto LABEL6
		}
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp42[v0]), ns.GetWaypointY(wp42[v0]), 0, 0)
		v0 += 1
	}
LABEL6:
	ns.WayPointGroupOn(gvar44)
	ns.MusicPopEvent()
LABEL5:
	return
}
func TeleportToPlayer() {
	var (
		v0 int
		v1 float32
		v2 float32
		v3 ns.WaypointID
	)
	v2 = 1e+06
	v0 = 0
	for {
		if !(v0 < 5) {
			goto LABEL1
		}
		v1 = ns.Distance(ns.GetWaypointX(wp41[v0]), ns.GetWaypointY(wp41[v0]), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
		if !(v1 < v2) {
			goto LABEL2
		}
		v2 = v1
		v3 = wp41[v0]
	LABEL2:
		v0 += 1
	}
LABEL1:
	ns.Enchant(obj18, ns.ENCHANT_INVISIBLE, 1)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(obj18), ns.GetObjectY(obj18), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(obj18), ns.GetObjectY(obj18), 0, 0)
	ns.MoveObject(obj18, ns.GetWaypointX(v3), ns.GetWaypointY(v3))
	ns.Effect(ns.BLUE_SPARKS, ns.GetWaypointX(v3), ns.GetWaypointY(v3), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(v3), ns.GetWaypointY(v3), 0, 0)
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
