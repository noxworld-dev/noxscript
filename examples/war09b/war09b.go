package war09b

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
	gvar25 ns.ObjectGroupID
	gvar26 ns.ObjectGroupID
	gvar27 ns.ObjectGroupID
	gvar28 ns.ObjectGroupID
	gvar29 ns.ObjectGroupID
	gvar30 ns.ObjectGroupID
	gvar31 ns.ObjectGroupID
	gvar32 ns.ObjectGroupID
	gvar33 ns.ObjectGroupID
	gvar34 ns.ObjectGroupID
	gvar35 ns.ObjectGroupID
	gvar36 ns.ObjectGroupID
	gvar37 ns.WallGroupID
	gvar38 ns.WallGroupID
	gvar39 ns.WallGroupID
	gvar40 ns.WallGroupID
	gvar41 ns.WallGroupID
	gvar42 ns.WallGroupID
	wp43   ns.WaypointID
	wp44   ns.WaypointID
	wp45   [9]ns.WaypointID
	wp46   [3]ns.WaypointID
	wp47   ns.WaypointID
	wp48   ns.WaypointID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	flag53 bool
	gvar54 int
	gvar55 int
	gvar56 int
	gvar57 int
	gvar58 int
	obj59  [9]ns.ObjectID
	obj60  ns.ObjectID
	obj61  ns.ObjectID
	gvar62 int
	flag63 bool
	flag64 bool
	flag65 bool
	flag66 bool
	flag67 bool
	gvar68 int
	gvar69 int
	gvar70 int
	gvar71 int
	gvar72 int
	gvar73 int
	ivar74 int
	ivar75 int
	obj76  [8]ns.ObjectID
	wp77   [8]ns.WaypointID
	gvar78 [8]int
	ivar79 [8]int
	gvar80 [8]ns.TimerID
)

func init() {
	flag53 = false
	gvar54 = 0
	gvar55 = 1
	gvar56 = 2
	gvar57 = 3
	gvar58 = 4
	gvar62 = gvar54
	flag63 = false
	flag64 = false
	flag65 = false
	flag66 = false
	flag67 = false
	gvar68 = 0
	gvar69 = 1
	gvar70 = 2
	gvar71 = 3
	gvar72 = 4
	gvar73 = 5
	ivar74 = 15
	ivar75 = 2
}
func PlaySwamp() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(29, 100)
LABEL1:
	return
}
func PlaySwamp2() {
	ns.Music(29, 100)
}
func PlayAction() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(26, 100)
LABEL1:
	return
}
func PlayOgre() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(7, 100)
LABEL1:
	return
}
func PlayerDeath() {
	ns.DeathScreen(9)
}
func AidanDialogStart() {
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
	if v0 == gvar57 {
		goto LABEL4
	}
	if v0 == gvar58 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.DestroyChat(obj5)
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanGreet1")
	flag53 = true
	goto LABEL6
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanGreet2")
	goto LABEL6
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanWaiting")
	goto LABEL6
LABEL4:
	ns.DestroyChat(obj5)
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanThankful")
	goto LABEL6
LABEL5:
	ns.DestroyChat(obj5)
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanLeaving")
	goto LABEL6
LABEL6:
	return
}
func AidanDialogEnd() {
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
	if v0 == gvar57 {
		goto LABEL4
	}
	if v0 == gvar58 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	gvar62 = gvar55
	goto LABEL6
LABEL2:
	gvar62 = gvar56
	goto LABEL6
LABEL3:
	goto LABEL6
LABEL4:
	flag53 = true
	flag66 = true
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.CancelDialog(obj5)
	ns.GiveXp(ns.GetHost(), 1000)
	ns.Pickup(ns.GetHost(), obj24)
	ns.PrintToAll("GeneralPrint:FindScroll")
	ns.CreatureFollow(obj5, ns.GetHost())
	ns.AggressionLevel(obj5, 0.5)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.BecomePet(obj5)
	goto LABEL6
LABEL5:
	goto LABEL6
LABEL6:
	return
}
func WispInitialize() {
	obj76[0] = ns.Object("Wisp2")
	gvar78[0] = gvar68
	wp77[0] = ns.Waypoint("Wisp2Dest")
	obj76[1] = ns.Object("Wisp3")
	gvar78[1] = gvar68
	wp77[1] = ns.Waypoint("Wisp3Dest")
}
func MapInitialize() {
	obj5 = ns.Object("Aidan")
	obj6 = ns.Object("DeadPrisoner")
	obj7 = ns.Object("Mystic_Brotherhood")
	obj8 = ns.Object("OgreCageDoor01")
	obj9 = ns.Object("CheckRescueTrigger")
	obj10 = ns.Object("PitElev02")
	obj11 = ns.Object("PitElev03")
	obj12 = ns.Object("CryptTrigger01")
	obj13 = ns.Object("CryptEnd1")
	obj14 = ns.Object("CryptDoor3")
	obj15 = ns.Object("CryptDoor4")
	obj16 = ns.Object("CryptGuardian")
	obj17 = ns.Object("UrchinCageDoor")
	obj18 = ns.Object("UrchinShaman")
	obj19 = ns.Object("Urchin1")
	obj20 = ns.Object("Urchin2")
	obj21 = ns.Object("Urchin3")
	obj22 = ns.Object("Urchin4")
	obj4 = ns.Object("Necromancer")
	obj23 = ns.Object("OgreDude")
	obj24 = ns.Object("WarlordScroll")
	gvar37 = ns.WallGroup("Surprise02Walls")
	gvar38 = ns.WallGroup("Surprise02DestWalls")
	gvar39 = ns.WallGroup("GraveyardTrapWalls")
	gvar40 = ns.WallGroup("GraveyardTrapWalls2")
	gvar41 = ns.WallGroup("CryptWall01")
	gvar42 = ns.WallGroup("SubGraveyardWalls")
	gvar25 = ns.ObjectGroup("GraveyardTrapEnemies")
	gvar26 = ns.ObjectGroup("TriggerGroup01")
	gvar27 = ns.ObjectGroup("TriggerGroup02")
	gvar28 = ns.ObjectGroup("SubGraveyardTriggers")
	gvar29 = ns.ObjectGroup("SubGraveyardCreatures")
	gvar30 = ns.ObjectGroup("CryptGuardianTriggers")
	gvar31 = ns.ObjectGroup("CryptWarriorTriggers")
	gvar32 = ns.ObjectGroup("UrchinPrisoners")
	gvar33 = ns.ObjectGroup("Secret03Enemies")
	gvar34 = ns.ObjectGroup("Secret03Triggers")
	gvar35 = ns.ObjectGroup("Secret04Triggers")
	gvar36 = ns.ObjectGroup("Secret06Triggers")
	wp45[0] = ns.Waypoint("CreateWP1")
	wp45[1] = ns.Waypoint("CreateWP2")
	wp45[2] = ns.Waypoint("CreateWP3")
	wp45[3] = ns.Waypoint("CreateWP4")
	wp45[4] = ns.Waypoint("CreateWP5")
	wp45[5] = ns.Waypoint("CreateWP6")
	wp45[6] = ns.Waypoint("CreateWP7")
	wp45[7] = ns.Waypoint("CreateWP8")
	wp45[8] = ns.Waypoint("CreateWP9")
	wp43 = ns.Waypoint("VileZombieCreateWP01")
	wp44 = ns.Waypoint("GhostCreateWP01")
	wp49 = ns.Waypoint("FaceHereWP")
	wp46[0] = ns.Waypoint("TeleWP1")
	wp46[1] = ns.Waypoint("TeleWP2")
	wp46[2] = ns.Waypoint("TeleWP3")
	wp47 = ns.Waypoint("CryptGuardianEntrance")
	wp48 = ns.Waypoint("CryptWarriorEntrance")
	wp50 = ns.Waypoint("AidanWait")
	wp51 = ns.Waypoint("AidanLook")
	wp52 = ns.Waypoint("PlayerSounds")
	ns.LockDoor(obj17)
	ns.UnBlind()
	WispInitialize()
	ns.StoryPic(obj5, "AidanPic")
	ns.SetDialog(obj5, ns.NORMAL, AidanDialogStart, AidanDialogEnd)
	ns.SetOwner(obj23, obj5)
	ns.Damage(obj6, 0, 50, 8)
}
func AidanRescued() {
	ns.ObjectOff(obj9)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	gvar62 = gvar57
	ns.StartDialog(obj5, ns.GetHost())
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
func AidanNearPlayer() {
	if !(flag53 == false && ns.IsCaller(ns.GetHost())) {
		goto LABEL1
	}
	ns.Chat(obj5, "War09a:AidanPlea")
	if flag65 {
		goto LABEL1
	}
	ns.ClearOwner(obj5)
	ns.SetOwner(ns.GetHost(), obj5)
	flag65 = true
LABEL1:
	return
}
func CheckRescue() {
	if ns.IsLocked(obj8) {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj5) > 0) {
		goto LABEL1
	}
	AidanRescued()
LABEL1:
	return
}
func AidanMustGo() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	PlayOgre()
LABEL1:
	if !(ns.IsCaller(ns.GetHost()) && flag64 == false) {
		goto LABEL2
	}
	flag64 = true
	ns.GiveXp(ns.GetHost(), 500)
	ns.JournalEdit(ns.GetHost(), "FindOutpost", 4)
	ns.JournalEntry(ns.GetHost(), "FindLOTD", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL2:
	if !(ns.IsCaller(ns.GetHost()) && flag66 == true && ns.CurrentHealth(obj5) > 0 && flag67 == false) {
		goto LABEL3
	}
	flag67 = true
	ns.RestoreHealth(obj5, 120)
	ns.SetDialog(obj5, ns.NORMAL, AidanDialogStart, AidanDialogEnd)
	gvar62 = gvar58
	ns.CreatureGuard(obj5, ns.GetWaypointX(wp50), ns.GetWaypointY(wp50), ns.GetWaypointX(wp51), ns.GetWaypointY(wp51), 0)
	ns.BecomeEnemy(obj5)
	ns.SetOwner(ns.GetHost(), obj5)
LABEL3:
	return
}
func DestroyAidan() {
	if !flag67 {
		goto LABEL1
	}
	ns.Delete(obj5)
LABEL1:
	return
}
func Secret03Declare() {
	ns.ObjectGroupOff(gvar34)
	ns.ObjectGroupOn(gvar33)
	ns.MoveWaypoint(wp52, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp52)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 5)
}
func Secret04Declare() {
	ns.ObjectGroupOff(gvar35)
	ns.MoveWaypoint(wp52, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp52)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 150)
}
func Secret05Declare() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp52, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp52)
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret06Declare() {
	ns.ObjectGroupOff(gvar36)
	ns.MoveWaypoint(wp52, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp52)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
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
func EnablePitElev02() {
	ns.ObjectOn(obj10)
}
func EnablePitElev03() {
	ns.ObjectOn(obj11)
}
func GraveyardCreateOrders() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 float32
		v5 float32
	)
	v0 = ns.GetWaypointX(wp43)
	v1 = ns.GetWaypointY(wp43)
	v2 = ns.GetWaypointX(wp44)
	v3 = ns.GetWaypointY(wp44)
	v4 = ns.GetWaypointX(wp49)
	v5 = ns.GetWaypointY(wp49)
	ns.LookWithAngle(obj60, 224)
	ns.LookWithAngle(obj61, 192)
	ns.CreatureGuard(obj60, v0, v1, v4, v5, 500)
	ns.CreatureGuard(obj61, v2, v3, v4, v5, 500)
}
func GraveyardTrapTrigger() {
	ns.ObjectGroupOff(gvar26)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar39)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar40)
	ns.ObjectGroupOn(gvar25)
	obj60 = ns.CreateObject("VileZombie", wp43)
	obj61 = ns.CreateObject("Ghost", wp44)
	ns.FrameTimer(1, GraveyardCreateOrders)
}
func SubGraveSurprise() {
	ns.ObjectGroupOff(gvar28)
	ns.ObjectGroupOn(gvar29)
	ns.WallGroupBreak(gvar42)
}
func CryptGuardian() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp47)
	v1 = ns.GetWaypointY(wp47)
	ns.ObjectGroupOff(gvar30)
	ns.LockDoor(obj14)
	ns.LockDoor(obj15)
	ns.ObjectOn(obj16)
	ns.Enchant(obj16, ns.ENCHANT_INVISIBLE, 2)
	ns.MoveObject(obj16, v0, v1)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
}
func TurnInvisible() {
	ns.Enchant(obj16, ns.ENCHANT_INVISIBLE, 1.5)
}
func UnlockCryptDoors2() {
	ns.UnlockDoor(obj14)
	ns.UnlockDoor(obj15)
}
func CreatureOrders() {
	var v0 int
	v0 = 0
	for {
		if !(v0 < 9) {
			goto LABEL1
		}
		ns.SetRoamFlag(obj59[v0], 1)
		ns.Wander(obj59[v0])
		ns.AggressionLevel(obj59[v0], 0.83)
		v0 += 1
	}
LABEL1:
	return
}
func CreateZombies01() {
	var v0 int
	ns.ObjectGroupOff(gvar27)
	ns.ObjectOn(obj4)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar37)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar38)
	v0 = 0
	for {
		if !(v0 < 5) {
			goto LABEL1
		}
		obj59[v0] = ns.CreateObject("Zombie", wp45[v0])
		v0 += 1
	}
LABEL1:
	v0 = 5
	for {
		if !(v0 < 7) {
			goto LABEL3
		}
		obj59[v0] = ns.CreateObject("Ghost", wp45[v0])
		v0 += 1
	}
LABEL3:
	v0 = 7
	for {
		if !(v0 < 9) {
			goto LABEL5
		}
		obj59[v0] = ns.CreateObject("VileZombie", wp45[v0])
		v0 += 1
	}
LABEL5:
	ns.FrameTimer(1, CreatureOrders)
}
func NecroTeleport() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = ns.Random(0, 2)
	v0 = ns.GetWaypointX(wp46[v2])
	v1 = ns.GetWaypointY(wp46[v2])
	if !ns.IsCaller(obj4) {
		goto LABEL1
	}
	ns.Enchant(ns.GetCaller(), ns.ENCHANT_INVISIBLE, 1)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.MoveObject(obj4, v0, v1)
	ns.Effect(ns.BLUE_SPARKS, v0, v1, 0, 0)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
LABEL1:
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL2
	}
	PlaySwamp()
LABEL2:
	return
}
func NecroTeleportInjured() {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = ns.Random(0, 2)
	v0 = ns.GetWaypointX(wp46[v2])
	v1 = ns.GetWaypointY(wp46[v2])
	ns.PauseObject(ns.GetTrigger(), 15)
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
func FreeUrchins01() {
	ns.UnlockDoor(obj17)
	if flag63 {
		goto LABEL1
	}
	flag63 = true
	ns.SetOwner(ns.GetHost(), obj19)
	ns.SetOwner(ns.GetHost(), obj20)
	ns.SetOwner(ns.GetHost(), obj21)
	ns.SetOwner(ns.GetHost(), obj22)
	ns.SetOwner(ns.GetHost(), obj18)
	ns.SetRoamFlag(obj18, 1)
	ns.SetRoamFlag(obj19, 1)
	ns.SetRoamFlag(obj20, 1)
	ns.SetRoamFlag(obj21, 1)
	ns.SetRoamFlag(obj22, 1)
	ns.AggressionLevel(obj18, 0.83)
	ns.AggressionLevel(obj19, 0.83)
	ns.AggressionLevel(obj20, 0.83)
	ns.AggressionLevel(obj21, 0.83)
	ns.AggressionLevel(obj22, 0.83)
LABEL1:
	return
}
func LockUrchinCage() {
	ns.LockDoor(obj17)
}
func HurtUrchins01() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.AggressionLevel(ns.GetTrigger(), 0.5)
	ns.RunAway(ns.GetTrigger(), ns.GetCaller(), 120)
LABEL1:
	return
}
func MapEntry() {
	ns.NoWallSound(false)
}
func whichWisp(a1 ns.ObjectID) int {
	var v0 int
	v0 = 0
	for {
		if !(v0 < ivar75) {
			goto LABEL1
		}
		if obj76[v0] != a1 {
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
	v1 = gvar78[a1]
	if v1 == gvar69 {
		goto LABEL1
	}
	if v1 == gvar71 {
		goto LABEL2
	}
	if v1 == gvar70 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v0 = ns.Distance(ns.GetObjectX(obj76[a1]), ns.GetObjectY(obj76[a1]), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	if !(v0 < 70) {
		goto LABEL5
	}
	gvar78[a1] = gvar70
	ns.Move(obj76[a1], wp77[a1])
LABEL5:
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj76[a1], ns.GetHost())
	ivar79[a1] += 1
	if !(ivar79[a1] > 4) {
		goto LABEL6
	}
	gvar78[a1] = gvar69
	ns.CreatureFollow(obj76[a1], ns.GetHost())
LABEL6:
	goto LABEL4
LABEL3:
	v0 = ns.Distance(ns.GetObjectX(obj76[a1]), ns.GetObjectY(obj76[a1]), ns.GetWaypointX(wp77[a1]), ns.GetWaypointY(wp77[a1]))
	if !(v0 < 30) {
		goto LABEL7
	}
	gvar78[a1] = gvar72
	ns.CreatureIdle(obj76[a1])
	return
LABEL7:
	goto LABEL4
LABEL4:
	ns.FrameTimerWithArg(ivar74, a1, WispAction)
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
	v1 = gvar78[v0]
	if v1 == gvar68 {
		goto LABEL3
	}
	if v1 == gvar71 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	gvar78[v0] = gvar69
	ns.CreatureFollow(ns.GetTrigger(), ns.GetCaller())
	gvar80[v0] = ns.FrameTimerWithArg(ivar74, v0, WispAction)
	goto LABEL5
LABEL4:
	gvar78[v0] = gvar70
	ns.Move(ns.GetTrigger(), wp77[v0])
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
	v1 = gvar78[v0]
	if v1 == gvar70 {
		goto LABEL3
	}
	goto LABEL4
LABEL3:
	gvar78[v0] = gvar71
	ivar79[v0] = 0
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
	gvar78[v0] = gvar73
	ns.CancelTimer(gvar80[v0])
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
