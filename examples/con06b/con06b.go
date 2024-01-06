package con06b

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4   int
	gvar5   int
	gvar6   int
	gvar7   int
	gvar8   int
	flag9   bool
	flag10  bool
	flag11  bool
	flag12  bool
	flag13  bool
	flag14  bool
	fvar15  float32
	fvar16  float32
	ivar17  int
	ivar18  int
	gvar19  int
	gvar20  int
	obj21   ns.ObjectID
	obj22   ns.ObjectID
	obj23   ns.ObjectID
	obj24   ns.ObjectID
	obj25   ns.ObjectID
	obj26   ns.ObjectID
	obj27   ns.ObjectID
	obj28   ns.ObjectID
	obj29   ns.ObjectID
	obj30   ns.ObjectID
	obj31   ns.ObjectID
	obj32   ns.ObjectID
	obj33   ns.ObjectID
	obj34   ns.ObjectID
	obj35   ns.ObjectID
	obj36   ns.ObjectID
	obj37   ns.ObjectID
	obj38   ns.ObjectID
	obj39   ns.ObjectID
	obj40   ns.ObjectID
	obj41   ns.ObjectID
	obj42   ns.ObjectID
	obj43   ns.ObjectID
	obj44   ns.ObjectID
	obj45   ns.ObjectID
	obj46   ns.ObjectID
	obj47   ns.ObjectID
	obj48   ns.ObjectID
	obj49   ns.ObjectID
	obj50   ns.ObjectID
	obj51   ns.ObjectID
	obj52   ns.ObjectID
	obj53   ns.ObjectID
	gvar54  ns.ObjectID
	obj55   ns.ObjectID
	obj56   ns.ObjectID
	obj57   ns.ObjectID
	obj58   ns.ObjectID
	obj59   ns.ObjectID
	obj60   ns.ObjectID
	obj61   ns.ObjectID
	obj62   ns.ObjectID
	obj63   ns.ObjectID
	obj64   ns.ObjectID
	gvar65  ns.ObjectGroupID
	gvar66  ns.ObjectGroupID
	gvar67  ns.ObjectGroupID
	gvar68  ns.ObjectGroupID
	gvar69  ns.ObjectGroupID
	gvar70  ns.ObjectGroupID
	gvar71  ns.ObjectGroupID
	gvar72  ns.WallGroupID
	gvar73  ns.WallGroupID
	gvar74  ns.WallGroupID
	gvar75  ns.WallGroupID
	gvar76  ns.WallGroupID
	gvar77  ns.WallGroupID
	wp78    ns.WaypointID
	wp79    ns.WaypointID
	wp80    ns.WaypointID
	wp81    ns.WaypointID
	wp82    ns.WaypointID
	wp83    ns.WaypointID
	wp84    ns.WaypointID
	wp85    ns.WaypointID
	wp86    ns.WaypointID
	wp87    ns.WaypointID
	wp88    ns.WaypointID
	wp89    ns.WaypointID
	wp90    ns.WaypointID
	wp91    ns.WaypointID
	flag92  bool
	flag93  bool
	flag94  bool
	flag95  bool
	ivar96  int
	ivar97  int
	obj98   ns.ObjectID
	gvar99  int
	gvar100 int
	gvar101 int
	gvar102 int
	wp103   ns.WaypointID
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	gvar7 = 3
	gvar8 = 4
	flag9 = true
	flag10 = true
	flag11 = true
	flag12 = true
	flag13 = false
	flag14 = true
	ivar17 = 0
	ivar18 = 0
	gvar19 = 0
	gvar20 = gvar4
	flag92 = true
	flag93 = false
	flag94 = true
	flag95 = false
	ivar96 = 76
	ivar97 = 70
}
func NullCallback() {
}
func fireKnightsRecognize() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ivar17 = ns.Random(1, 3)
	if ivar17 != 1 {
		goto LABEL2
	}
	ns.StoryPic(ns.GetTrigger(), "WarriorPic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, fireKnightsTalk1Start, fireKnightsTalk1End)
LABEL2:
	if ivar17 != 2 {
		goto LABEL3
	}
	ns.StoryPic(ns.GetTrigger(), "Warrior2Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, fireKnightsTalk2Start, fireKnightsTalk2End)
LABEL3:
	if ivar17 != 3 {
		goto LABEL1
	}
	ns.StoryPic(ns.GetTrigger(), "Warrior3Pic")
	ns.SetDialog(ns.GetTrigger(), ns.NORMAL, fireKnightsTalk3Start, fireKnightsTalk3End)
LABEL1:
	return
}
func fireKnightsTalk1Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar67)
	ns.TellStory(ns.SwordsmanRecognize, "Con06a:CityGuardSpeak1")
}
func fireKnightsTalk1End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar67)
	ns.SetCallback(ns.GetTrigger(), 3, NullCallback)
	ns.CancelDialog(ns.GetTrigger())
	ns.BecomePet(ns.GetTrigger())
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func fireKnightsTalk2Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar67)
	ns.TellStory(ns.SwordsmanRecognize, "Con06a:CityGuardSpeak2")
}
func fireKnightsTalk2End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar67)
	ns.SetCallback(ns.GetTrigger(), 3, NullCallback)
	ns.CancelDialog(ns.GetTrigger())
	ns.BecomePet(ns.GetTrigger())
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func fireKnightsTalk3Start() {
	ns.Frozen(ns.GetHost(), true)
	ns.ObjectGroupOff(gvar67)
	ns.TellStory(ns.SwordsmanRecognize, "War06a:CityGuardSpeak3")
}
func fireKnightsTalk3End() {
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar67)
	ns.SetCallback(ns.GetTrigger(), 3, NullCallback)
	ns.CancelDialog(ns.GetTrigger())
	ns.CreatureFollow(ns.GetTrigger(), ns.GetHost())
	ns.BecomePet(ns.GetTrigger())
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
}
func ownFriends() {
	ns.JournalEntry(ns.GetHost(), "War6Horrendous", 2)
	ns.JournalEdit(ns.GetHost(), "War6Fortress", 4)
	ns.SetOwner(ns.GetHost(), obj21)
	ns.SetOwner(ns.GetHost(), obj27)
	ns.SetOwner(ns.GetHost(), obj28)
	ns.SetOwner(ns.GetHost(), obj29)
	ns.SetOwner(ns.GetHost(), obj30)
	ns.StoryPic(obj31, "WarriorPic")
	ns.SetDialog(obj31, ns.NORMAL, fireKnightsTalk1Start, fireKnightsTalk1End)
	ns.StoryPic(obj32, "Warrior2Pic")
	ns.SetDialog(obj32, ns.NORMAL, fireKnightsTalk2Start, fireKnightsTalk2End)
	ns.StoryPic(obj33, "Warrior3Pic")
	ns.SetDialog(obj33, ns.NORMAL, fireKnightsTalk3Start, fireKnightsTalk3End)
	ns.StoryPic(obj34, "WarriorPic")
	ns.SetDialog(obj34, ns.NORMAL, fireKnightsTalk1Start, fireKnightsTalk1End)
}
func secretArea() {
	ns.MoveWaypoint(wp78, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp78)
	ns.GiveXp(ns.GetHost(), 500)
}
func toggleMainGates() {
	if gvar19 != 0 {
		goto LABEL1
	}
	gvar19 = 1
	openMainGates()
	goto LABEL2
LABEL1:
	gvar19 = 0
	closeMainGates()
LABEL2:
	ns.AudioEvent(ns.SecretWallStoneClose, wp84)
	ns.AudioEvent(ns.BoulderMove, wp84)
}
func openMainGates() {
	ns.Move(obj50, wp80)
	ns.Move(obj51, wp81)
	ns.Move(obj52, wp82)
	ns.Move(obj53, wp83)
}
func closeMainGates() {
	ns.Move(obj50, wp84)
	ns.Move(obj51, wp85)
	ns.Move(obj52, wp86)
	ns.Move(obj53, wp87)
}
func deathToDoorStops() {
	ns.Delete(ns.GetCaller())
}
func startHorrendousTalk() {
	ns.StartDialog(obj21, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj21)
}
func faceHorrendous() {
	ns.LookAtObject(ns.GetTrigger(), obj21)
}
func horrendousDie() {
	ns.JournalEntry(ns.GetHost(), "War6Necro", 2)
	ns.JournalEdit(ns.GetHost(), "War6Horrendous", 4)
	ns.JournalDelete(ns.GetHost(), "War6Fortress")
	flag14 = false
	ns.MoveObject(obj22, ns.GetObjectX(obj21), ns.GetObjectY(obj21))
	fvar15 = ns.GetObjectX(obj21) - 55
	fvar16 = ns.GetObjectY(obj21) + 20
	ns.MoveObject(obj21, 1161, 2794)
	ns.SetDialog(gvar54, ns.NORMAL, hecubahTalk4Start, hecubahTalk4End)
	ns.StartDialog(gvar54, ns.GetHost())
	ns.FrameTimer(110, horrendousDying)
}
func horrendousDying() {
	ns.MoveObject(obj23, ns.GetObjectX(obj22), ns.GetObjectY(obj22))
	ns.Delete(obj22)
	ns.FrameTimer(15, horrendousDropHalberd)
}
func horrendousDropHalberd() {
	ns.MoveObject(obj26, ns.GetObjectX(obj23)-10, ns.GetObjectY(obj23)+5)
}
func horrendousDead() {
	ns.Music(28, 100)
	ns.WideScreen(false)
	ns.MoveObject(obj24, ns.GetObjectX(obj23), ns.GetObjectY(obj23))
	ns.Delete(obj23)
	ns.AudioEvent(ns.DemonDie, wp90)
	if gvar54 == 0 {
		goto LABEL1
	}
	ns.Delete(gvar54)
LABEL1:
	return
}
func horrendousSpeak() {
	ns.Music(0, 100)
	ns.WideScreen(true)
	ns.SetDialog(obj21, ns.NORMAL, horrendousTalk1Start, horrendousTalk1End)
	ns.StartDialog(obj21, ns.GetHost())
}
func horrendousTalk1Start() {
	ns.ObjectGroupOn(gvar71)
	ns.Frozen(ns.GetHost(), true)
	ns.LockDoor(obj35)
	ns.LockDoor(obj36)
	ns.LockDoor(obj37)
	ns.GroupPauseObject(gvar67, 10)
	ns.ObjectGroupOff(gvar67)
	ns.TellStory(ns.DemonRecognize, "Con06a:BoastfulHorrendous")
}
func horrendousTalk1End() {
	ns.CancelDialog(obj21)
	ns.SetDialog(gvar54, ns.NEXT, hecubahTalk1Start, hecubahTalk1End)
	ns.FrameTimer(30, hecubahBustIn)
}
func horrendousTalk2Start() {
	ns.TellStory(ns.DemonRecognize, "War06a:HorrendousVsHec")
}
func horrendousTalk2End() {
	ns.SetDialog(obj21, ns.NEXT, horrendousTalk3Start, horrendousTalk3End)
	ns.FrameTimer(30, startHorrendousTalk)
}
func horrendousTalk3Start() {
	ns.TellStory(ns.DemonRecognize, "War06a:ChallengingHec")
}
func horrendousTalk3End() {
	ns.CancelDialog(obj21)
	ns.SetDialog(gvar54, ns.NEXT, hecubahTalk2Start, hecubahTalk2End)
	ns.FrameTimer(30, startHecubahTalk)
}
func horrendousTalk5Start() {
	ns.Music(0, 100)
	ns.LookAtObject(ns.GetHost(), obj23)
	ns.TellStory(ns.DemonRecognize, "War06a:HecDefeatsHorrendous")
}
func horrendousTalk5End() {
	ns.CancelDialog(obj21)
	ns.SetDialog(gvar54, ns.NEXT, hecubahTalk5Start, hecubahTalk5End)
	ns.FrameTimer(30, startHecubahTalk)
}
func horrendousTalk6Start() {
	ns.TellStory(ns.DemonRecognize, "War06a:HorrendousPride")
}
func horrendousTalk6End() {
	ns.ObjectGroupOn(gvar67)
	ns.EnchantOff(ns.GetHost(), ns.ENCHANT_INVULNERABLE)
	ns.CancelDialog(obj21)
	horrendousDead()
	ns.Frozen(ns.GetHost(), false)
	ns.ObjectGroupOn(gvar67)
	ns.AggressionLevel(obj56, 0.83)
	ns.AggressionLevel(obj57, 0.83)
	ns.AggressionLevel(obj58, 0.83)
	ns.AggressionLevel(obj59, 0.83)
	ns.AggressionLevel(obj60, 0.83)
	ns.AggressionLevel(obj61, 0.83)
	ns.AggressionLevel(obj62, 0.83)
	ns.AggressionLevel(obj63, 0.83)
	ns.AggressionLevel(obj64, 0.83)
}
func talkToHorrendous() {
	ns.SetDialog(obj21, ns.NORMAL, horrendousTalk6Start, horrendousTalk6End)
	ns.StartDialog(obj21, ns.GetHost())
}
func startHecubahTalk() {
	ns.StartDialog(gvar54, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), gvar54)
}
func hecubahBustIn() {
	ns.ObjectGroupOff(gvar71)
	ns.ObjectGroupOn(gvar66)
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar72)
	ns.NoWallSound(false)
	ns.Delete(obj46)
	ns.Delete(obj47)
	ns.Delete(obj48)
	ns.Delete(obj49)
	ns.FrameTimer(10, hecubahBreakIn)
}
func hecubahBreakIn() {
	ns.WallGroupBreak(gvar73)
	ns.Enchant(obj21, ns.ENCHANT_VILLAIN, 0)
	ns.ObjectGroupOn(gvar69)
	ns.AggressionLevel(gvar54, 0)
	ns.AggressionLevel(obj55, 0)
	ns.AggressionLevel(obj56, 0)
	ns.AggressionLevel(obj57, 0)
	ns.AggressionLevel(obj58, 0)
	ns.AggressionLevel(obj59, 0)
	ns.AggressionLevel(obj60, 0)
	ns.AggressionLevel(obj61, 0)
	ns.AggressionLevel(obj62, 0)
	ns.AggressionLevel(obj63, 0)
	ns.AggressionLevel(obj64, 0)
	ns.Wander(gvar54)
	ns.StartDialog(gvar54, ns.GetHost())
}
func hecKillsEveryone() {
	if !(ns.IsCaller(gvar54) || ns.IsCaller(obj55)) {
		goto LABEL1
	}
	ns.FrameTimer(3, hecIsKilling)
LABEL1:
	return
}
func isHecKilling() {
	if !flag9 {
		goto LABEL1
	}
	ns.Music(28, 75)
	ns.FrameTimer(3, hecIsKilling)
LABEL1:
	return
}
func hecIsKilling() {
	var v0 int
	if !flag10 {
		goto LABEL1
	}
	flag9 = false
	ns.LookAtObject(ns.GetHost(), gvar54)
	v0 = gvar20
	if v0 == gvar4 {
		goto LABEL2
	}
	if v0 == gvar5 {
		goto LABEL3
	}
	if v0 == gvar6 {
		goto LABEL4
	}
	if v0 == gvar7 {
		goto LABEL5
	}
	if v0 == gvar8 {
		goto LABEL6
	}
	goto LABEL7
LABEL2:
	ns.LookAtObject(gvar54, obj21)
	ns.LookAtObject(ns.GetHost(), obj21)
	ns.AudioEvent(ns.DeathRayCast, wp90)
	ns.Effect(ns.DEATH_RAY, ns.GetObjectX(gvar54), ns.GetObjectY(gvar54), ns.GetObjectX(obj21), ns.GetObjectY(obj21))
	ns.Damage(obj21, 0, 1000, 0)
	gvar20 = gvar5
	goto LABEL7
LABEL3:
	ns.LookAtObject(gvar54, obj27)
	ns.LookAtObject(ns.GetHost(), obj27)
	ns.AudioEvent(ns.DeathRayCast, wp90)
	ns.Effect(ns.DEATH_RAY, ns.GetObjectX(gvar54), ns.GetObjectY(gvar54), ns.GetObjectX(obj27), ns.GetObjectY(obj27))
	ns.Damage(obj27, 0, 1000, 0)
	gvar20 = gvar6
	goto LABEL7
LABEL4:
	ns.LookAtObject(gvar54, obj28)
	ns.LookAtObject(ns.GetHost(), obj28)
	ns.AudioEvent(ns.DeathRayCast, wp90)
	ns.Effect(ns.DEATH_RAY, ns.GetObjectX(gvar54), ns.GetObjectY(gvar54), ns.GetObjectX(obj28), ns.GetObjectY(obj28))
	ns.Damage(obj28, 0, 1000, 0)
	gvar20 = gvar7
	goto LABEL7
LABEL5:
	ns.LookAtObject(gvar54, obj29)
	ns.LookAtObject(ns.GetHost(), obj29)
	ns.AudioEvent(ns.DeathRayCast, wp90)
	ns.Effect(ns.DEATH_RAY, ns.GetObjectX(gvar54), ns.GetObjectY(gvar54), ns.GetObjectX(obj29), ns.GetObjectY(obj29))
	ns.Damage(obj29, 0, 1000, 0)
	gvar20 = gvar8
	goto LABEL7
LABEL6:
	ns.LookAtObject(gvar54, obj30)
	ns.LookAtObject(ns.GetHost(), obj30)
	ns.AudioEvent(ns.DeathRayCast, wp90)
	ns.Effect(ns.DEATH_RAY, ns.GetObjectX(gvar54), ns.GetObjectY(gvar54), ns.GetObjectX(obj30), ns.GetObjectY(obj30))
	ns.Damage(obj30, 0, 1000, 0)
	flag10 = false
	goto LABEL7
LABEL7:
	ns.FrameTimer(15, hecIsKilling)
LABEL1:
	ns.LookAtObject(ns.GetHost(), gvar54)
}
func hecubahTalk1Start() {
	ns.TellStory(ns.GhostRecognize, "War06a:HecubahThreat")
}
func hecubahTalk1End() {
	ns.CancelDialog(gvar54)
	ns.SetDialog(obj21, ns.NEXT, horrendousTalk2Start, horrendousTalk2End)
	ns.FrameTimer(30, startHorrendousTalk)
}
func hecubahTalk2Start() {
	ns.TellStory(ns.GhostRecognize, "War06a:HecubahAccepts")
}
func hecubahTalk2End() {
	ns.CancelDialog(gvar54)
	ns.Enchant(ns.GetHost(), ns.ENCHANT_INVULNERABLE, 0)
	ns.ObjectOn(obj38)
	ns.FrameTimer(240, isHecKilling)
}
func hecubahTalk4Start() {
	ns.ObjectOff(gvar54)
	ns.TellStory(ns.GhostRecognize, "War06a:HecubahWins")
}
func hecubahTalk4End() {
	ns.SetDialog(obj21, ns.NEXT, horrendousTalk5Start, horrendousTalk5End)
	ns.FrameTimer(30, startHorrendousTalk)
	ns.CancelDialog(gvar54)
}
func hecubahTalk5Start() {
	ns.TellStory(ns.GhostRecognize, "War06a:HecubahToNecro")
}
func hecubahTalk5End() {
	ns.CancelDialog(gvar54)
	ns.ObjectOn(obj55)
	ns.AggressionLevel(obj55, 0)
	ns.Walk(obj55, ns.GetWaypointX(wp79), ns.GetWaypointY(wp79))
	ns.SetDialog(obj55, ns.NEXT, necromancerTalk1Start, necromancerTalk1End)
	ns.FrameTimer(90, startNecromancerTalk)
}
func hecubahTalk6Start() {
	ns.TellStory(ns.GhostRecognize, "War06a:HecThreatJack")
}
func hecubahTalk6End() {
	ns.CancelDialog(gvar54)
	ns.ObjectOn(gvar54)
	ns.AggressionLevel(gvar54, 0)
	ns.Walk(gvar54, 931, 3036)
	ns.SetDialog(obj55, ns.NEXT, necromancerTalk2Start, necromancerTalk2End)
	ns.FrameTimer(30, startNecromancerTalk)
}
func startNecromancerTalk() {
	ns.StartDialog(obj55, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj55)
}
func necromancerTalk1Start() {
	ns.CreatureIdle(obj55)
	ns.LookAtObject(obj55, obj26)
	ns.TellStory(ns.EmberDemonRecognize, "War06a:NecroMockHorren")
}
func necromancerTalk1End() {
	ns.CancelDialog(obj55)
	ns.Pickup(obj55, obj26)
	ns.SetDialog(gvar54, ns.NEXT, hecubahTalk6Start, hecubahTalk6End)
	ns.FrameTimer(30, startHecubahTalk)
}
func necromancerTalk2Start() {
	ns.TellStory(ns.EmberDemonRecognize, "War06a:NecroThreatJack")
}
func necromancerTalk2End() {
	ns.CancelDialog(obj55)
	ns.ObjectOn(obj55)
	ns.AggressionLevel(obj55, 0)
	ns.Walk(obj55, 931, 3036)
	ns.FrameTimer(240, talkToHorrendous)
}
func lowerNecroWall() {
	ns.NoWallSound(true)
	ns.WallGroupOpen(gvar74)
}
func getHalberd() {
	ns.Pickup(ns.GetHost(), obj26)
	ns.FrameTimer(3, necroDies2)
}
func exitFade() {
	ns.Blind()
	ns.FrameTimer(60, exitTile)
}
func exitTile() {
	ns.Frozen(ns.GetHost(), false)
	ns.MoveObject(ns.GetHost(), 5394, 2955)
}
func FreezePlayer() {
	ns.CreatureIdle(obj55)
	ns.LookAtObject(ns.GetHost(), obj55)
	ns.LookAtObject(obj55, ns.GetHost())
	ns.Frozen(ns.GetHost(), true)
	ns.Delete(obj56)
	ns.Delete(obj57)
	ns.Delete(obj58)
	ns.Delete(obj59)
	ns.Delete(obj60)
	ns.Delete(obj61)
	ns.Delete(obj62)
	ns.Delete(obj63)
	ns.Delete(obj64)
	ns.FrameTimer(15, necroSetPiece)
}
func necroSetPiece() {
	ns.WideScreen(true)
	ns.ObjectOff(obj55)
	ns.SetDialog(obj55, ns.NORMAL, necromancerTalk3Start, necromancerTalk3End)
	ns.StartDialog(obj55, ns.GetHost())
}
func wallWalk() {
	if ivar17 != 0 {
		goto LABEL1
	}
	ns.AudioEvent(ns.SpellPhonemeUpLeft, wp103)
LABEL1:
	if ivar17 != 1 {
		goto LABEL2
	}
	ns.AudioEvent(ns.SpellPhonemeUp, wp103)
LABEL2:
	if ivar17 != 2 {
		goto LABEL3
	}
	ns.AudioEvent(ns.SpellPhonemeUpRight, wp103)
LABEL3:
	if !(ivar17 < 49) {
		goto LABEL4
	}
	ns.WallClose(ns.Wall(ivar96, ivar97))
	ivar17 += 1
	ivar96 -= 1
	ivar97 -= 1
	ns.FrameTimer(6, wallWalk)
LABEL4:
	return
}
func releasePlayer() {
	ns.NoWallSound(false)
	ns.CastSpellLocationLocation("SUMMON_SKELETON_LORD", ns.GetObjectX(obj55), ns.GetObjectY(obj55), 1575, 1506)
	ns.CastSpellLocationLocation("SUMMON_SKELETON_LORD", ns.GetObjectX(obj55), ns.GetObjectY(obj55), 1702, 1506)
	ns.Frozen(ns.GetHost(), false)
}
func endFightMusic() {
	ns.Music(27, 100)
}
func necromancerTalk3Start() {
	ns.GroupPauseObject(gvar67, 10)
	ns.ObjectGroupOff(gvar67)
	ns.TellStory(ns.ZombieRecognize, "Con06a:NecroAttackJack")
}
func necromancerTalk3End() {
	ns.ObjectGroupOn(gvar67)
	ns.CancelDialog(obj55)
	ns.WideScreen(false)
	ns.ObjectOn(obj55)
	ns.Frozen(obj55, false)
	ns.Wander(obj55)
	ns.FrameTimer(10, wallWalk)
	ns.FrameTimer(50, releasePlayer)
	ns.FrameTimer(60, lowerNecroWall)
}
func necromancerTalk4Start() {
	ns.GroupPauseObject(gvar67, 10)
	ns.ObjectGroupOff(gvar67)
	ns.TellStory(ns.ZombieRecognize, "Con06a:NecroDies")
}
func necromancerTalk4End() {
	ns.ObjectGroupOn(gvar67)
	ns.CancelDialog(obj55)
	ns.Frozen(obj55, false)
	ns.GroupDamage(gvar68, 0, 1000, 0)
	ns.FrameTimer(15, getHalberd)
}
func necroDies() {
	ivar18 += 1
	flag92 = false
	ns.LookAtObject(obj55, ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj55)
	ns.Frozen(obj55, true)
	ns.MoveObject(obj26, 1310, 2900)
	ns.SetDialog(obj55, ns.NORMAL, necromancerTalk4Start, necromancerTalk4End)
	ns.StartDialog(obj55, ns.GetHost())
}
func necroDies2() {
	if !(ivar18 > 4) {
		goto LABEL1
	}
	if !ns.HasItem(ns.GetHost(), obj26) {
		goto LABEL2
	}
	ns.JournalEdit(ns.GetHost(), "War6Necro", 4)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(45, exitFade)
	goto LABEL3
LABEL2:
	ns.FrameTimer(3, necroDies2)
LABEL3:
	goto LABEL4
LABEL1:
	ns.FrameTimer(3, necroDies2)
LABEL4:
	return
}
func MapInitialize() {
	obj21 = ns.Object("Horrendous")
	obj22 = ns.Object("Horrendous_")
	obj23 = ns.Object("Horrendous__")
	obj24 = ns.Object("Horrendous___")
	obj25 = ns.Object("Horrendous____")
	obj26 = ns.Object("OblivionHalberd")
	obj27 = ns.Object("EliteGuard1")
	obj28 = ns.Object("EliteGuard2")
	obj29 = ns.Object("EliteGuard3")
	obj30 = ns.Object("EliteGuard4")
	obj31 = ns.Object("F2Fire6")
	obj32 = ns.Object("F6FireGuard1")
	obj33 = ns.Object("F6FireGuard4")
	obj34 = ns.Object("F6Guard8")
	obj35 = ns.Object("ArenaDoor1")
	obj36 = ns.Object("ArenaDoor2")
	obj37 = ns.Object("ArenaDoor3")
	obj38 = ns.Object("KillTrigger")
	ns.ObjectOff(obj38)
	obj39 = ns.Object("CellDoor1")
	obj40 = ns.Object("CellDoor2")
	obj41 = ns.Object("CellDoor3")
	obj42 = ns.Object("CellDoor4")
	obj43 = ns.Object("CellDoor5")
	obj44 = ns.Object("CellDoor6")
	obj45 = ns.Object("CellDoor7")
	obj46 = ns.Object("BreakInObject1")
	obj47 = ns.Object("BreakInObject2")
	obj48 = ns.Object("BreakInObject3")
	obj49 = ns.Object("BreakInObject4")
	gvar54 = ns.Object("Hecubah")
	obj55 = ns.Object("Necromancer")
	obj56 = ns.Object("UndeadWar1")
	obj57 = ns.Object("UndeadWar2")
	obj58 = ns.Object("UndeadWar3")
	obj59 = ns.Object("UndeadWar4")
	obj60 = ns.Object("SkeLord1")
	obj61 = ns.Object("SkeLord2")
	obj62 = ns.Object("SkeLord3")
	obj63 = ns.Object("SkeLord4")
	obj64 = ns.Object("SkeLord5")
	obj50 = ns.Object("FG1Mover")
	obj51 = ns.Object("FG2Mover")
	obj52 = ns.Object("FG3Mover")
	obj53 = ns.Object("FG4Mover")
	gvar65 = ns.ObjectGroup("NecroGuards")
	ns.ObjectGroupOff(gvar65)
	gvar66 = ns.ObjectGroup("Wolves")
	ns.ObjectGroupOff(gvar66)
	gvar67 = ns.ObjectGroup("EveryMonsterOnMap")
	gvar68 = ns.ObjectGroup("EveryMonsterOnMap2")
	gvar69 = ns.ObjectGroup("PartyCrashers")
	gvar70 = ns.ObjectGroup("HecBreakInObjects")
	gvar71 = ns.ObjectGroup("DoorStopTriggers")
	ns.ObjectGroupOff(gvar71)
	gvar72 = ns.WallGroup("HecBreakIn")
	gvar73 = ns.WallGroup("HecBreakDown")
	gvar74 = ns.WallGroup("NecroWall")
	gvar75 = ns.WallGroup("NecroExitWall")
	gvar76 = ns.WallGroup("WizardCages")
	gvar77 = ns.WallGroup("DeathCage")
	wp78 = ns.Waypoint("PlayerSounds")
	wp79 = ns.Waypoint("NecroSpotWP")
	wp80 = ns.Waypoint("FG1Open")
	wp81 = ns.Waypoint("FG2Open")
	wp82 = ns.Waypoint("FG3Open")
	wp83 = ns.Waypoint("FG4Open")
	wp84 = ns.Waypoint("FG1Closed")
	wp85 = ns.Waypoint("FG2Closed")
	wp86 = ns.Waypoint("FG3Closed")
	wp87 = ns.Waypoint("FG4Closed")
	wp88 = ns.Waypoint("UpStairs")
	wp89 = ns.Waypoint("DownStairs")
	wp90 = ns.Waypoint("NecroNoise")
	wp91 = ns.Waypoint("WellWP")
	obj55 = ns.Object("Necromancer")
	obj98 = ns.Object("Halberd")
	wp103 = ns.Waypoint("SewerChase")
	ns.StoryPic(obj55, "BlackWizardPic")
	ns.StoryPic(obj21, "HorrendousPic")
	ns.StoryPic(gvar54, "HecubahPic")
	ns.StoryPic(obj55, "NecromancerPic")
	ns.AggressionLevel(obj21, 0)
	ns.AggressionLevel(obj27, 0)
	ns.AggressionLevel(obj28, 0)
	ns.AggressionLevel(obj29, 0)
	ns.AggressionLevel(obj30, 0)
	ns.ObjectGroupOff(gvar69)
	ns.LockDoor(obj39)
	ns.LockDoor(obj40)
	ns.LockDoor(obj41)
	ns.LockDoor(obj42)
	ns.LockDoor(obj43)
	ns.LockDoor(obj44)
	ns.LockDoor(obj45)
	ns.FrameTimer(3, ownFriends)
	ns.Music(16, 100)
}
func PlayerDeath() {
	ns.DeathScreen(6)
}
func unlockCells() {
	ns.UnlockDoor(obj39)
	ns.UnlockDoor(obj40)
	ns.UnlockDoor(obj41)
	ns.UnlockDoor(obj42)
	ns.UnlockDoor(obj43)
	ns.UnlockDoor(obj44)
	ns.UnlockDoor(obj45)
}
func lockCells() {
	ns.LockDoor(obj39)
	ns.LockDoor(obj40)
	ns.LockDoor(obj41)
	ns.LockDoor(obj42)
	ns.LockDoor(obj43)
	ns.LockDoor(obj44)
	ns.LockDoor(obj45)
}
func setNecroPosition() {
	if !ns.IsCaller(obj55) {
		goto LABEL1
	}
	ns.Frozen(obj55, true)
	ns.MoveObject(obj55, 1776, 1502)
LABEL1:
	if !ns.IsCaller(gvar54) {
		goto LABEL2
	}
	if gvar54 == 0 {
		goto LABEL2
	}
	ns.Delete(gvar54)
LABEL2:
	return
}
func setupNecroSetPiece() {
	ns.Music(18, 100)
	ns.PauseObject(ns.GetTrigger(), 10)
	ns.ObjectOff(ns.GetTrigger())
	ns.FrameTimer(1, FreezePlayer)
}
func setNecroPosition2() {
	if !ns.IsCaller(obj55) {
		goto LABEL1
	}
	ns.MoveObject(obj55, 5057, 3809)
LABEL1:
	return
}
func necroGuard() {
	ns.CreatureGuard(obj55, 5057, 3809, 5232, 3991, 500)
	ns.AggressionLevel(obj55, 0.5)
	ns.ObjectGroupOn(gvar65)
}
func NecroInjured() {
	if ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.RestoreHealth(obj55, 14)
LABEL1:
	return
}
func necroGuardDies() {
	ivar18 += 1
}
func MakeAWish() {
	var v0 int
	v0 = ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp91)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
