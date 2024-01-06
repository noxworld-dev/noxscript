package wiz07a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	obj4 ns.ObjectID
	obj5 ns.ObjectID
	obj6 ns.ObjectID
)

func CreatureSetup() {
	ns.SetOwner(ns.GetHost(), obj4)
	ns.SetOwner(ns.GetHost(), obj5)
	ns.LockDoor(obj6)
	ns.SetDialog(obj4, ns.NORMAL, JandorStart, JandorEnd)
	ns.SetDialog(obj5, ns.NORMAL, LanceDialogStart, LanceDialogEnd)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.JournalEntry(ns.GetHost(), "GetHON", 2)
	ns.FrameTimer(1, StartJandorTalk)
}
func StartJandorTalk() {
	ns.StartDialog(obj4, ns.GetHost())
}
func JandorStart() {
	ns.Music(21, 100)
	ns.TellStory(ns.HumanMaleEatFood, "Wiz07.scr:JandorTalk01")
}
func JandorEnd() {
	ns.CancelDialog(obj4)
}
func LanceDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "War01A.scr:Guard1Talk01")
}
func LanceDialogEnd() {
}
func MapInitialize() {
	obj4 = ns.Object("Jandor")
	obj5 = ns.Object("Lance")
	obj6 = ns.Object("FacadeGate")
	ns.StoryPic(obj4, "AirshipCaptainPic")
	ns.StoryPic(obj5, "WizardGuard2Pic")
	ns.StartupScreen(7)
	ns.FrameTimer(1, CreatureSetup)
}
func PlayerDeath() {
	ns.DeathScreen(10)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
