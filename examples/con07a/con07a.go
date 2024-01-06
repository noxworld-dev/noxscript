package con07a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4 ns.ObjectID
	obj5 ns.ObjectID
	obj6 ns.ObjectID
)

func MapSetup() {
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.FrameTimer(10, StartJandorTalk)
	ns.JournalEntry(ns.GetHost(), "GetHON", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
}
func JandorBriefStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con07A.scr:Jandor01")
}
func JandorBriefEnd() {
	ns.CancelDialog(obj4)
	ns.Frozen(ns.GetHost(), false)
}
func LanceDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "War01A.scr:Guard1Talk01")
}
func LanceDialogEnd() {
}
func StartJandorTalk() {
	ns.StartDialog(obj4, ns.GetHost())
}
func MapInitialize() {
	obj4 = ns.Object("Jandor")
	obj5 = ns.Object("Lance")
	obj6 = ns.Object("FacadeGate")
	ns.SetDialog(obj4, ns.NORMAL, JandorBriefStart, JandorBriefEnd)
	ns.StoryPic(obj4, "AirshipCaptainPic")
	ns.SetDialog(obj5, ns.NORMAL, LanceDialogStart, LanceDialogEnd)
	ns.StoryPic(obj5, "WizardGuard1Pic")
	ns.LockDoor(obj6)
	ns.StartupScreen(7)
	ns.FrameTimer(1, MapSetup)
}
func PlayerDeath() {
	ns.DeathScreen(7)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
