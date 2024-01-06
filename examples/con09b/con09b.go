package con09b

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
	gvar37 ns.ObjectGroupID
	gvar38 ns.WallGroupID
	gvar39 ns.WallGroupID
	gvar40 ns.WallGroupID
	gvar41 ns.WallGroupID
	gvar42 ns.WallGroupID
	gvar43 ns.WallGroupID
	wp44   ns.WaypointID
	wp45   ns.WaypointID
	wp46   [9]ns.WaypointID
	wp47   [3]ns.WaypointID
	wp48   ns.WaypointID
	wp49   ns.WaypointID
	wp50   ns.WaypointID
	wp51   ns.WaypointID
	wp52   ns.WaypointID
	wp53   ns.WaypointID
	flag54 bool
	gvar55 int
	gvar56 int
	gvar57 int
	gvar58 int
	gvar59 int
	obj60  [9]ns.ObjectID
	obj61  ns.ObjectID
	obj62  ns.ObjectID
	gvar63 int
	flag64 bool
	flag65 bool
	flag66 bool
	flag67 bool
	flag68 bool
	gvar69 int
	gvar70 int
	gvar71 int
	gvar72 int
	gvar73 int
	gvar74 int
	ivar75 int
	ivar76 int
	obj77  [8]ns.ObjectID
	wp78   [8]ns.WaypointID
	gvar79 [8]int
	ivar80 [8]int
	gvar81 [8]ns.TimerID
)

func init() {
	flag54 = false
	gvar55 = 0
	gvar56 = 1
	gvar57 = 2
	gvar58 = 3
	gvar59 = 4
	gvar63 = gvar55
	flag64 = false
	flag65 = false
	flag66 = false
	flag67 = false
	flag68 = false
	gvar69 = 0
	gvar70 = 1
	gvar71 = 2
	gvar72 = 3
	gvar73 = 4
	gvar74 = 5
	ivar75 = 15
	ivar76 = 2
}
func PlayerDeath() {
	ns.DeathScreen(9)
}
func AidanDialogStart() {
	var v0 int
	v0 = gvar63
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
	ns.DestroyChat(obj4)
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanGreet1")
	flag54 = true
	goto LABEL6
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanGreet2")
	goto LABEL6
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanWaiting")
	goto LABEL6
LABEL4:
	ns.DestroyChat(obj4)
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanThankful")
	goto LABEL6
LABEL5:
	ns.DestroyChat(obj4)
	ns.TellStory(ns.SwordsmanHurt, "War09a:AidanLeaving")
	goto LABEL6
LABEL6:
	return
}
func AidanDialogEnd() {
	var v0 int
	v0 = gvar63
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
	gvar63 = gvar56
	goto LABEL6
LABEL2:
	gvar63 = gvar57
	goto LABEL6
LABEL3:
	goto LABEL6
LABEL4:
	flag54 = true
	flag66 = true
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.CancelDialog(obj4)
	ns.GiveXp(ns.GetHost(), 1000)
	ns.Pickup(ns.GetHost(), obj5)
	ns.PrintToAll("GeneralPrint:FindSpell")
	ns.CreatureFollow(obj4, ns.GetHost())
	ns.AggressionLevel(obj4, 0.5)
	ns.BecomePet(obj4)
	ns.SetOwner(ns.GetHost(), obj4)
	goto LABEL6
LABEL5:
	goto LABEL6
LABEL6:
	return
}
func WispInitialize() {
	obj77[0] = ns.Object("Wisp2")
	gvar79[0] = gvar69
	wp78[0] = ns.Waypoint("Wisp2Dest")
	obj77[1] = ns.Object("Wisp3")
	gvar79[1] = gvar69
	wp78[1] = ns.Waypoint("Wisp3Dest")
}
func MapInitialize() {
	obj4 = ns.Object("Aidan")
	obj5 = ns.Object("SlowBook")
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
	obj17 = ns.Object("CryptWarrior")
	obj18 = ns.Object("UrchinCageDoor")
	obj19 = ns.Object("UrchinShaman")
	obj20 = ns.Object("Urchin1")
	obj21 = ns.Object("Urchin2")
	obj22 = ns.Object("Urchin3")
	obj23 = ns.Object("Urchin4")
	obj24 = ns.Object("Necromancer")
	obj25 = ns.Object("OgreDude")
	gvar38 = ns.WallGroup("Surprise02Walls")
	gvar39 = ns.WallGroup("Surprise02DestWalls")
	gvar40 = ns.WallGroup("GraveyardTrapWalls")
	gvar41 = ns.WallGroup("GraveyardTrapWalls2")
	gvar42 = ns.WallGroup("CryptWall01")
	gvar43 = ns.WallGroup("SubGraveyardWalls")
	gvar26 = ns.ObjectGroup("GraveyardTrapEnemies")
	gvar27 = ns.ObjectGroup("TriggerGroup01")
	gvar28 = ns.ObjectGroup("TriggerGroup02")
	gvar29 = ns.ObjectGroup("SubGraveyardTriggers")
	gvar30 = ns.ObjectGroup("SubGraveyardCreatures")
	gvar31 = ns.ObjectGroup("CryptGuardianTriggers")
	gvar32 = ns.ObjectGroup("CryptWarriorTriggers")
	gvar33 = ns.ObjectGroup("UrchinPrisoners")
	gvar34 = ns.ObjectGroup("Secret03Enemies")
	gvar35 = ns.ObjectGroup("Secret03Triggers")
	gvar36 = ns.ObjectGroup("Secret04Triggers")
	gvar37 = ns.ObjectGroup("Secret06Triggers")
	wp46[0] = ns.Waypoint("CreateWP1")
	wp46[1] = ns.Waypoint("CreateWP2")
	wp46[2] = ns.Waypoint("CreateWP3")
	wp46[3] = ns.Waypoint("CreateWP4")
	wp46[4] = ns.Waypoint("CreateWP5")
	wp46[5] = ns.Waypoint("CreateWP6")
	wp46[6] = ns.Waypoint("CreateWP7")
	wp46[7] = ns.Waypoint("CreateWP8")
	wp46[8] = ns.Waypoint("CreateWP9")
	wp44 = ns.Waypoint("VileZombieCreateWP01")
	wp45 = ns.Waypoint("GhostCreateWP01")
	wp50 = ns.Waypoint("FaceHereWP")
	wp47[0] = ns.Waypoint("TeleWP1")
	wp47[1] = ns.Waypoint("TeleWP2")
	wp47[2] = ns.Waypoint("TeleWP3")
	wp48 = ns.Waypoint("CryptGuardianEntrance")
	wp49 = ns.Waypoint("CryptWarriorEntrance")
	wp51 = ns.Waypoint("AidanWait")
	wp52 = ns.Waypoint("AidanLook")
	wp53 = ns.Waypoint("PlayerSounds")
	ns.LockDoor(obj18)
	ns.UnBlind()
	WispInitialize()
	ns.StoryPic(obj4, "AidanPic")
	ns.SetDialog(obj4, ns.NORMAL, AidanDialogStart, AidanDialogEnd)
	ns.SetOwner(obj25, obj4)
	ns.Damage(obj6, 0, 50, 8)
}
func AidanRescued() {
	ns.ObjectOff(obj9)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	gvar63 = gvar58
	ns.StartDialog(obj4, ns.GetHost())
}
func PlayOgre() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(7, 100)
LABEL1:
	return
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
	if !(flag54 == false && ns.IsCaller(ns.GetHost())) {
		goto LABEL1
	}
	ns.Chat(obj4, "War09a:AidanPlea")
	if flag67 {
		goto LABEL1
	}
	ns.ClearOwner(obj4)
	ns.SetOwner(ns.GetHost(), obj4)
	flag67 = true
LABEL1:
	return
}
func CheckRescue() {
	if ns.IsLocked(obj8) {
		goto LABEL1
	}
	if !(ns.CurrentHealth(obj4) > 0) {
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
	if !(ns.IsCaller(ns.GetHost()) && flag66 == true && ns.CurrentHealth(obj4) > 0 && flag68 == false) {
		goto LABEL3
	}
	flag68 = true
	ns.RestoreHealth(obj4, 120)
	ns.SetDialog(obj4, ns.NORMAL, AidanDialogStart, AidanDialogEnd)
	gvar63 = gvar59
	ns.CreatureGuard(obj4, ns.GetWaypointX(wp51), ns.GetWaypointY(wp51), ns.GetWaypointX(wp52), ns.GetWaypointY(wp52), 0)
	ns.BecomeEnemy(obj4)
	ns.SetOwner(ns.GetHost(), obj4)
LABEL3:
	return
}
func DestroyAidan() {
	if !flag68 {
		goto LABEL1
	}
	ns.Delete(obj4)
LABEL1:
	return
}
func Secret03Declare() {
	ns.ObjectGroupOff(gvar35)
	ns.ObjectGroupOn(gvar34)
	ns.MoveWaypoint(wp53, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp53)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 5)
}
func Secret04Declare() {
	ns.ObjectGroupOff(gvar36)
	ns.MoveWaypoint(wp53, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp53)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 150)
}
func Secret05Declare() {
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.MoveWaypoint(wp53, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp53)
	ns.GiveXp(ns.GetHost(), 100)
}
func Secret06Declare() {
	ns.ObjectGroupOff(gvar37)
	ns.MoveWaypoint(wp53, ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()))
	ns.AudioEvent(ns.SecretFound, wp53)
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
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
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
	v0 = ns.GetWaypointX(wp44)
	v1 = ns.GetWaypointY(wp44)
	v2 = ns.GetWaypointX(wp45)
	v3 = ns.GetWaypointY(wp45)
	v4 = ns.GetWaypointX(wp50)
	v5 = ns.GetWaypointY(wp50)
	ns.LookWithAngle(obj61, 224)
	ns.LookWithAngle(obj62, 192)
	ns.CreatureGuard(obj61, v0, v1, v4, v5, 500)
	ns.CreatureGuard(obj62, v2, v3, v4, v5, 500)
}
func GraveyardTrapTrigger() {
	ns.ObjectGroupOff(gvar27)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar40)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar41)
	ns.ObjectGroupOn(gvar26)
	obj61 = ns.CreateObject("VileZombie", wp44)
	obj62 = ns.CreateObject("Ghost", wp45)
	ns.FrameTimer(1, GraveyardCreateOrders)
}
func SubGraveSurprise() {
	ns.ObjectGroupOff(gvar29)
	ns.ObjectGroupOn(gvar30)
	ns.WallGroupBreak(gvar43)
}
func CryptGuardian() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp48)
	v1 = ns.GetWaypointY(wp48)
	ns.ObjectGroupOff(gvar31)
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
		ns.SetRoamFlag(obj60[v0], 1)
		ns.Wander(obj60[v0])
		ns.AggressionLevel(obj60[v0], 0.83)
		v0 += 1
	}
LABEL1:
	return
}
func CreateZombies01() {
	var v0 int
	ns.ObjectGroupOff(gvar28)
	ns.ObjectOn(obj24)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar38)
	ns.NoWallSound(false)
	ns.WallGroupBreak(gvar39)
	v0 = 0
	for {
		if !(v0 < 5) {
			goto LABEL1
		}
		obj60[v0] = ns.CreateObject("Zombie", wp46[v0])
		v0 += 1
	}
LABEL1:
	v0 = 5
	for {
		if !(v0 < 7) {
			goto LABEL3
		}
		obj60[v0] = ns.CreateObject("Ghost", wp46[v0])
		v0 += 1
	}
LABEL3:
	v0 = 7
	for {
		if !(v0 < 9) {
			goto LABEL5
		}
		obj60[v0] = ns.CreateObject("VileZombie", wp46[v0])
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
	v0 = ns.GetWaypointX(wp47[v2])
	v1 = ns.GetWaypointY(wp47[v2])
	if !ns.IsCaller(obj24) {
		goto LABEL1
	}
	ns.Enchant(ns.GetCaller(), ns.ENCHANT_INVISIBLE, 2)
	ns.Effect(ns.BLUE_SPARKS, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.Effect(ns.SMOKE_BLAST, ns.GetObjectX(ns.GetCaller()), ns.GetObjectY(ns.GetCaller()), 0, 0)
	ns.MoveObject(obj24, v0, v1)
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
	v2 = ns.Random(0, 2)
	v0 = ns.GetWaypointX(wp47[v2])
	v1 = ns.GetWaypointY(wp47[v2])
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
	ns.UnlockDoor(obj18)
	if flag65 {
		goto LABEL1
	}
	flag65 = true
	ns.SetOwner(ns.GetHost(), obj20)
	ns.SetOwner(ns.GetHost(), obj21)
	ns.SetOwner(ns.GetHost(), obj22)
	ns.SetOwner(ns.GetHost(), obj23)
	ns.SetOwner(ns.GetHost(), obj19)
	ns.SetRoamFlag(obj19, 1)
	ns.SetRoamFlag(obj20, 1)
	ns.SetRoamFlag(obj21, 1)
	ns.SetRoamFlag(obj22, 1)
	ns.SetRoamFlag(obj23, 1)
	ns.AggressionLevel(obj19, 0.83)
	ns.AggressionLevel(obj20, 0.83)
	ns.AggressionLevel(obj21, 0.83)
	ns.AggressionLevel(obj22, 0.83)
	ns.AggressionLevel(obj23, 0.83)
LABEL1:
	return
}
func LockUrchinCage() {
	ns.LockDoor(obj18)
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
		if !(v0 < ivar76) {
			goto LABEL1
		}
		if obj77[v0] != a1 {
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
	v1 = gvar79[a1]
	if v1 == gvar70 {
		goto LABEL1
	}
	if v1 == gvar72 {
		goto LABEL2
	}
	if v1 == gvar71 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	v0 = ns.Distance(ns.GetObjectX(obj77[a1]), ns.GetObjectY(obj77[a1]), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	if !(v0 < 70) {
		goto LABEL5
	}
	gvar79[a1] = gvar71
	ns.Move(obj77[a1], wp78[a1])
LABEL5:
	goto LABEL4
LABEL2:
	ns.LookAtObject(obj77[a1], ns.GetHost())
	ivar80[a1] += 1
	if !(ivar80[a1] > 4) {
		goto LABEL6
	}
	gvar79[a1] = gvar70
	ns.CreatureFollow(obj77[a1], ns.GetHost())
LABEL6:
	goto LABEL4
LABEL3:
	v0 = ns.Distance(ns.GetObjectX(obj77[a1]), ns.GetObjectY(obj77[a1]), ns.GetWaypointX(wp78[a1]), ns.GetWaypointY(wp78[a1]))
	if !(v0 < 30) {
		goto LABEL7
	}
	gvar79[a1] = gvar73
	ns.CreatureIdle(obj77[a1])
	return
LABEL7:
	goto LABEL4
LABEL4:
	ns.FrameTimerWithArg(ivar75, a1, WispAction)
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
	v1 = gvar79[v0]
	if v1 == gvar69 {
		goto LABEL3
	}
	if v1 == gvar72 {
		goto LABEL4
	}
	goto LABEL5
LABEL3:
	gvar79[v0] = gvar70
	ns.CreatureFollow(ns.GetTrigger(), ns.GetCaller())
	gvar81[v0] = ns.FrameTimerWithArg(ivar75, v0, WispAction)
	goto LABEL5
LABEL4:
	gvar79[v0] = gvar71
	ns.Move(ns.GetTrigger(), wp78[v0])
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
	v1 = gvar79[v0]
	if v1 == gvar71 {
		goto LABEL3
	}
	goto LABEL4
LABEL3:
	gvar79[v0] = gvar72
	ivar80[v0] = 0
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
	gvar79[v0] = gvar74
	ns.CancelTimer(gvar81[v0])
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
