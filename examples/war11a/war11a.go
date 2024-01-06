package war11a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	gvar4  int
	gvar5  int
	gvar6  int
	gvar7  int
	gvar8  int
	wp9    [5]ns.WaypointID
	wp10   [4]ns.WaypointID
	wp11   [9]ns.WaypointID
	wp12   [9]ns.WaypointID
	wp13   [9]ns.WaypointID
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
	gvar25 int
	gvar26 int
	gvar27 int
	gvar28 int
	gvar29 int
	gvar30 int
	gvar31 int
	flag32 bool
	flag33 bool
	fvar34 float32
	fvar35 float32
	ivar36 int
	ivar37 int
	ivar38 int
	ivar39 int
	ivar40 int
	gvar41 int
	gvar42 int
	gvar43 int
	obj44  ns.ObjectID
	obj45  ns.ObjectID
	obj46  ns.ObjectID
	obj47  ns.ObjectID
	obj48  ns.ObjectID
	obj49  ns.ObjectID
	obj50  ns.ObjectID
	obj51  ns.ObjectID
	obj52  ns.ObjectID
	obj53  ns.ObjectID
	obj54  ns.ObjectID
	obj55  ns.ObjectID
	obj56  ns.ObjectID
	obj57  ns.ObjectID
	obj58  ns.ObjectID
	obj59  ns.ObjectID
	obj60  ns.ObjectID
	obj61  ns.ObjectID
	obj62  ns.ObjectID
	obj63  ns.ObjectID
	obj64  ns.ObjectID
	obj65  ns.ObjectID
	obj66  [4]ns.ObjectID
	obj67  [5]ns.ObjectID
	wp68   [4]ns.WaypointID
	wp69   ns.WaypointID
	wp70   ns.WaypointID
	wp71   ns.WaypointID
	wp72   ns.WaypointID
	wp73   ns.WaypointID
	wp74   ns.WaypointID
	wp75   ns.WaypointID
)

func init() {
	gvar4 = 0
	gvar5 = 1
	gvar6 = 2
	gvar7 = 3
	gvar8 = 4
	gvar14 = gvar4
	gvar15 = gvar5
	gvar16 = gvar4
	gvar17 = gvar5
	gvar18 = gvar4
	gvar19 = gvar5
	gvar20 = gvar4
	gvar21 = gvar5
	gvar22 = gvar4
	gvar23 = gvar5
	gvar24 = gvar4
	gvar25 = gvar5
	gvar26 = gvar4
	gvar27 = gvar5
	gvar28 = gvar4
	gvar29 = gvar5
	gvar30 = gvar4
	gvar31 = gvar5
	flag32 = false
	flag33 = true
	ivar36 = 0
	ivar37 = 0
	ivar39 = 0
	ivar40 = 0
	gvar41 = 0
	gvar42 = 0
	gvar43 = 1
}
func Flame0() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar15
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar15 = gvar5
	gvar14 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp9[0])
	v1[1] = ns.CreateObject("SmallFlame", wp11[0])
	v1[2] = ns.CreateObject("SmallFlame", wp12[0])
	v1[3] = ns.CreateObject("SmallFlame", wp13[0])
	if gvar14 != gvar4 {
		goto LABEL11
	}
	gvar15 = gvar6
	goto LABEL12
LABEL11:
	gvar15 = gvar4
LABEL12:
	gvar14 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp9[0])
	v1[1] = ns.CreateObject("MediumFlame", wp11[0])
	v1[2] = ns.CreateObject("MediumFlame", wp12[0])
	v1[3] = ns.CreateObject("MediumFlame", wp13[0])
	if gvar14 != gvar5 {
		goto LABEL15
	}
	gvar15 = gvar7
	goto LABEL16
LABEL15:
	gvar15 = gvar5
LABEL16:
	gvar14 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp9[0])
	v1[1] = ns.CreateObject("Flame", wp11[0])
	v1[2] = ns.CreateObject("Flame", wp12[0])
	v1[3] = ns.CreateObject("Flame", wp13[0])
	if gvar14 != gvar6 {
		goto LABEL19
	}
	gvar15 = gvar8
	goto LABEL20
LABEL19:
	gvar15 = gvar6
LABEL20:
	gvar14 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp9[0])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp9[0])
	v1[1] = ns.CreateObject("LargeFlame", wp11[0])
	v1[2] = ns.CreateObject("LargeFlame", wp12[0])
	v1[3] = ns.CreateObject("LargeFlame", wp13[0])
	gvar15 = gvar7
	gvar14 = gvar8
	goto LABEL6
LABEL6:
	if gvar14 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame0)
LABEL23:
	return
}
func Flame1() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar17
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar17 = gvar5
	gvar16 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp9[1])
	v1[1] = ns.CreateObject("SmallFlame", wp11[1])
	v1[2] = ns.CreateObject("SmallFlame", wp12[1])
	v1[3] = ns.CreateObject("SmallFlame", wp13[1])
	if gvar16 != gvar4 {
		goto LABEL11
	}
	gvar17 = gvar6
	goto LABEL12
LABEL11:
	gvar17 = gvar4
LABEL12:
	gvar16 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp9[1])
	v1[1] = ns.CreateObject("MediumFlame", wp11[1])
	v1[2] = ns.CreateObject("MediumFlame", wp12[1])
	v1[3] = ns.CreateObject("MediumFlame", wp13[1])
	if gvar16 != gvar5 {
		goto LABEL15
	}
	gvar17 = gvar7
	goto LABEL16
LABEL15:
	gvar17 = gvar5
LABEL16:
	gvar16 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp9[1])
	v1[1] = ns.CreateObject("Flame", wp11[1])
	v1[2] = ns.CreateObject("Flame", wp12[1])
	v1[3] = ns.CreateObject("Flame", wp13[1])
	if gvar16 != gvar6 {
		goto LABEL19
	}
	gvar17 = gvar8
	goto LABEL20
LABEL19:
	gvar17 = gvar6
LABEL20:
	gvar16 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp9[1])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp9[1])
	v1[1] = ns.CreateObject("LargeFlame", wp11[1])
	v1[2] = ns.CreateObject("LargeFlame", wp12[1])
	v1[3] = ns.CreateObject("LargeFlame", wp13[1])
	gvar17 = gvar7
	gvar16 = gvar8
	goto LABEL6
LABEL6:
	if gvar16 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame1)
LABEL23:
	return
}
func Flame2() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar19
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar19 = gvar5
	gvar18 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp9[2])
	v1[1] = ns.CreateObject("SmallFlame", wp11[2])
	v1[2] = ns.CreateObject("SmallFlame", wp12[2])
	v1[3] = ns.CreateObject("SmallFlame", wp13[2])
	if gvar18 != gvar4 {
		goto LABEL11
	}
	gvar19 = gvar6
	goto LABEL12
LABEL11:
	gvar19 = gvar4
LABEL12:
	gvar18 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp9[2])
	v1[1] = ns.CreateObject("MediumFlame", wp11[2])
	v1[2] = ns.CreateObject("MediumFlame", wp12[2])
	v1[3] = ns.CreateObject("MediumFlame", wp13[2])
	if gvar18 != gvar5 {
		goto LABEL15
	}
	gvar19 = gvar7
	goto LABEL16
LABEL15:
	gvar19 = gvar5
LABEL16:
	gvar18 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp9[2])
	v1[1] = ns.CreateObject("Flame", wp11[2])
	v1[2] = ns.CreateObject("Flame", wp12[2])
	v1[3] = ns.CreateObject("Flame", wp13[2])
	if gvar18 != gvar6 {
		goto LABEL19
	}
	gvar19 = gvar8
	goto LABEL20
LABEL19:
	gvar19 = gvar6
LABEL20:
	gvar18 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp9[2])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp9[2])
	v1[1] = ns.CreateObject("LargeFlame", wp11[2])
	v1[2] = ns.CreateObject("LargeFlame", wp12[2])
	v1[3] = ns.CreateObject("LargeFlame", wp13[2])
	gvar19 = gvar7
	gvar18 = gvar8
	goto LABEL6
LABEL6:
	if gvar18 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame2)
LABEL23:
	return
}
func Flame3() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar21
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar21 = gvar5
	gvar20 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp9[3])
	v1[1] = ns.CreateObject("SmallFlame", wp11[3])
	v1[2] = ns.CreateObject("SmallFlame", wp12[3])
	v1[3] = ns.CreateObject("SmallFlame", wp13[3])
	if gvar20 != gvar4 {
		goto LABEL11
	}
	gvar21 = gvar6
	goto LABEL12
LABEL11:
	gvar21 = gvar4
LABEL12:
	gvar20 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp9[3])
	v1[1] = ns.CreateObject("MediumFlame", wp11[3])
	v1[2] = ns.CreateObject("MediumFlame", wp12[3])
	v1[3] = ns.CreateObject("MediumFlame", wp13[3])
	if gvar20 != gvar5 {
		goto LABEL15
	}
	gvar21 = gvar7
	goto LABEL16
LABEL15:
	gvar21 = gvar5
LABEL16:
	gvar20 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp9[3])
	v1[1] = ns.CreateObject("Flame", wp11[3])
	v1[2] = ns.CreateObject("Flame", wp12[3])
	v1[3] = ns.CreateObject("Flame", wp13[3])
	if gvar20 != gvar6 {
		goto LABEL19
	}
	gvar21 = gvar8
	goto LABEL20
LABEL19:
	gvar21 = gvar6
LABEL20:
	gvar20 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp9[3])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp9[3])
	v1[1] = ns.CreateObject("LargeFlame", wp11[3])
	v1[2] = ns.CreateObject("LargeFlame", wp12[3])
	v1[3] = ns.CreateObject("LargeFlame", wp13[3])
	gvar21 = gvar7
	gvar20 = gvar8
	goto LABEL6
LABEL6:
	if gvar20 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame3)
LABEL23:
	return
}
func Flame4() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar23
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar23 = gvar5
	gvar22 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp9[4])
	v1[1] = ns.CreateObject("SmallFlame", wp11[4])
	v1[2] = ns.CreateObject("SmallFlame", wp12[4])
	v1[3] = ns.CreateObject("SmallFlame", wp13[4])
	if gvar22 != gvar4 {
		goto LABEL11
	}
	gvar23 = gvar6
	goto LABEL12
LABEL11:
	gvar23 = gvar4
LABEL12:
	gvar22 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp9[4])
	v1[1] = ns.CreateObject("MediumFlame", wp11[4])
	v1[2] = ns.CreateObject("MediumFlame", wp12[4])
	v1[3] = ns.CreateObject("MediumFlame", wp13[4])
	if gvar22 != gvar5 {
		goto LABEL15
	}
	gvar23 = gvar7
	goto LABEL16
LABEL15:
	gvar23 = gvar5
LABEL16:
	gvar22 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp9[4])
	v1[1] = ns.CreateObject("Flame", wp11[4])
	v1[2] = ns.CreateObject("Flame", wp12[4])
	v1[3] = ns.CreateObject("Flame", wp13[4])
	if gvar22 != gvar6 {
		goto LABEL19
	}
	gvar23 = gvar8
	goto LABEL20
LABEL19:
	gvar23 = gvar6
LABEL20:
	gvar22 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp9[4])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp9[4])
	v1[1] = ns.CreateObject("LargeFlame", wp11[4])
	v1[2] = ns.CreateObject("LargeFlame", wp12[4])
	v1[3] = ns.CreateObject("LargeFlame", wp13[4])
	gvar23 = gvar7
	gvar22 = gvar8
	goto LABEL6
LABEL6:
	if gvar22 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame4)
LABEL23:
	return
}
func Flame5() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar25
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar25 = gvar5
	gvar24 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp10[0])
	v1[1] = ns.CreateObject("SmallFlame", wp11[5])
	v1[2] = ns.CreateObject("SmallFlame", wp12[5])
	v1[3] = ns.CreateObject("SmallFlame", wp13[5])
	if gvar24 != gvar4 {
		goto LABEL11
	}
	gvar25 = gvar6
	goto LABEL12
LABEL11:
	gvar25 = gvar4
LABEL12:
	gvar24 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp10[0])
	v1[1] = ns.CreateObject("MediumFlame", wp11[5])
	v1[2] = ns.CreateObject("MediumFlame", wp12[5])
	v1[3] = ns.CreateObject("MediumFlame", wp13[5])
	if gvar24 != gvar5 {
		goto LABEL15
	}
	gvar25 = gvar7
	goto LABEL16
LABEL15:
	gvar25 = gvar5
LABEL16:
	gvar24 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp10[0])
	v1[1] = ns.CreateObject("Flame", wp11[5])
	v1[2] = ns.CreateObject("Flame", wp12[5])
	v1[3] = ns.CreateObject("Flame", wp13[5])
	if gvar24 != gvar6 {
		goto LABEL19
	}
	gvar25 = gvar8
	goto LABEL20
LABEL19:
	gvar25 = gvar6
LABEL20:
	gvar24 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp10[0])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp10[0])
	v1[1] = ns.CreateObject("LargeFlame", wp11[5])
	v1[2] = ns.CreateObject("LargeFlame", wp12[5])
	v1[3] = ns.CreateObject("LargeFlame", wp13[5])
	gvar25 = gvar7
	gvar24 = gvar8
	goto LABEL6
LABEL6:
	if gvar24 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame5)
LABEL23:
	return
}
func Flame6() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar27
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar27 = gvar5
	gvar26 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp10[1])
	v1[1] = ns.CreateObject("SmallFlame", wp11[6])
	v1[2] = ns.CreateObject("SmallFlame", wp12[6])
	v1[3] = ns.CreateObject("SmallFlame", wp13[6])
	if gvar26 != gvar4 {
		goto LABEL11
	}
	gvar27 = gvar6
	goto LABEL12
LABEL11:
	gvar27 = gvar4
LABEL12:
	gvar26 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp10[1])
	v1[1] = ns.CreateObject("MediumFlame", wp11[6])
	v1[2] = ns.CreateObject("MediumFlame", wp12[6])
	v1[3] = ns.CreateObject("MediumFlame", wp13[6])
	if gvar26 != gvar5 {
		goto LABEL15
	}
	gvar27 = gvar7
	goto LABEL16
LABEL15:
	gvar27 = gvar5
LABEL16:
	gvar26 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp10[1])
	v1[1] = ns.CreateObject("Flame", wp11[6])
	v1[2] = ns.CreateObject("Flame", wp12[6])
	v1[3] = ns.CreateObject("Flame", wp13[6])
	if gvar26 != gvar6 {
		goto LABEL19
	}
	gvar27 = gvar8
	goto LABEL20
LABEL19:
	gvar27 = gvar6
LABEL20:
	gvar26 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp10[1])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp10[1])
	v1[1] = ns.CreateObject("LargeFlame", wp11[6])
	v1[2] = ns.CreateObject("LargeFlame", wp12[6])
	v1[3] = ns.CreateObject("LargeFlame", wp13[6])
	gvar27 = gvar7
	gvar26 = gvar8
	goto LABEL6
LABEL6:
	if gvar26 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame6)
LABEL23:
	return
}
func Flame7() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar29
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar29 = gvar5
	gvar28 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp10[2])
	v1[1] = ns.CreateObject("SmallFlame", wp11[7])
	v1[2] = ns.CreateObject("SmallFlame", wp12[7])
	v1[3] = ns.CreateObject("SmallFlame", wp13[7])
	if gvar28 != gvar4 {
		goto LABEL11
	}
	gvar29 = gvar6
	goto LABEL12
LABEL11:
	gvar29 = gvar4
LABEL12:
	gvar28 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp10[2])
	v1[1] = ns.CreateObject("MediumFlame", wp11[7])
	v1[2] = ns.CreateObject("MediumFlame", wp12[7])
	v1[3] = ns.CreateObject("MediumFlame", wp13[7])
	if gvar28 != gvar5 {
		goto LABEL15
	}
	gvar29 = gvar7
	goto LABEL16
LABEL15:
	gvar29 = gvar5
LABEL16:
	gvar28 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp10[2])
	v1[1] = ns.CreateObject("Flame", wp11[7])
	v1[2] = ns.CreateObject("Flame", wp12[7])
	v1[3] = ns.CreateObject("Flame", wp13[7])
	if gvar28 != gvar6 {
		goto LABEL19
	}
	gvar29 = gvar8
	goto LABEL20
LABEL19:
	gvar29 = gvar6
LABEL20:
	gvar28 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp10[2])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp10[2])
	v1[1] = ns.CreateObject("LargeFlame", wp11[7])
	v1[2] = ns.CreateObject("LargeFlame", wp12[7])
	v1[3] = ns.CreateObject("LargeFlame", wp13[7])
	gvar29 = gvar7
	gvar28 = gvar8
	goto LABEL6
LABEL6:
	if gvar28 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame7)
LABEL23:
	return
}
func Flame8() {
	var (
		v0 int
		v1 [4]ns.ObjectID
		v2 int
	)
	v2 = gvar31
	if v2 == gvar4 {
		goto LABEL1
	}
	if v2 == gvar5 {
		goto LABEL2
	}
	if v2 == gvar6 {
		goto LABEL3
	}
	if v2 == gvar7 {
		goto LABEL4
	}
	if v2 == gvar8 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL7
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL7:
	gvar31 = gvar5
	gvar30 = gvar4
	goto LABEL6
LABEL2:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL9
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL9:
	v1[0] = ns.CreateObject("SmallFlame", wp10[3])
	v1[1] = ns.CreateObject("SmallFlame", wp11[8])
	v1[2] = ns.CreateObject("SmallFlame", wp12[8])
	v1[3] = ns.CreateObject("SmallFlame", wp13[8])
	if gvar30 != gvar4 {
		goto LABEL11
	}
	gvar31 = gvar6
	goto LABEL12
LABEL11:
	gvar31 = gvar4
LABEL12:
	gvar30 = gvar5
	goto LABEL6
LABEL3:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL13
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL13:
	v1[0] = ns.CreateObject("MediumFlame", wp10[3])
	v1[1] = ns.CreateObject("MediumFlame", wp11[8])
	v1[2] = ns.CreateObject("MediumFlame", wp12[8])
	v1[3] = ns.CreateObject("MediumFlame", wp13[8])
	if gvar30 != gvar5 {
		goto LABEL15
	}
	gvar31 = gvar7
	goto LABEL16
LABEL15:
	gvar31 = gvar5
LABEL16:
	gvar30 = gvar6
	goto LABEL6
LABEL4:
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL17
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL17:
	v1[0] = ns.CreateObject("Flame", wp10[3])
	v1[1] = ns.CreateObject("Flame", wp11[8])
	v1[2] = ns.CreateObject("Flame", wp12[8])
	v1[3] = ns.CreateObject("Flame", wp13[8])
	if gvar30 != gvar6 {
		goto LABEL19
	}
	gvar31 = gvar8
	goto LABEL20
LABEL19:
	gvar31 = gvar6
LABEL20:
	gvar30 = gvar7
	goto LABEL6
LABEL5:
	ns.AudioEvent(ns.DemonBreath, wp10[3])
	v0 = 0
	for {
		if !(v0 < 4) {
			goto LABEL21
		}
		ns.Delete(v1[v0])
		v0 += 1
	}
LABEL21:
	v1[0] = ns.CreateObject("LargeFlame", wp10[3])
	v1[1] = ns.CreateObject("LargeFlame", wp11[8])
	v1[2] = ns.CreateObject("LargeFlame", wp12[8])
	v1[3] = ns.CreateObject("LargeFlame", wp13[8])
	gvar31 = gvar7
	gvar30 = gvar8
	goto LABEL6
LABEL6:
	if gvar30 == gvar4 {
		goto LABEL23
	}
	ns.FrameTimer(2, Flame8)
LABEL23:
	return
}
func LightTrigger0() {
	ns.ObjectOff(obj60)
	ns.ObjectOn(obj56)
	ns.ObjectOn(obj52)
	ns.ObjectOff(obj44)
	ns.ObjectOn(obj48)
}
func LightTrigger1() {
	ns.ObjectOff(obj61)
	ns.ObjectOn(obj57)
	ns.ObjectOn(obj53)
	ns.ObjectOff(obj45)
	ns.ObjectOn(obj49)
}
func LightTrigger2() {
	ns.ObjectOff(obj62)
	ns.ObjectOn(obj58)
	ns.ObjectOn(obj54)
	ns.ObjectOff(obj46)
	ns.ObjectOn(obj50)
}
func LightTrigger3() {
	ns.ObjectOff(obj63)
	ns.ObjectOn(obj59)
	ns.ObjectOn(obj55)
	ns.ObjectOff(obj47)
	ns.ObjectOn(obj51)
}
func LightTrigger0b() {
	ns.CreateObject("RedPotion", wp72)
	ns.ObjectOn(obj60)
	ns.ObjectOff(obj56)
	ns.ObjectOff(obj52)
	ns.ObjectOn(obj44)
	ns.ObjectOff(obj48)
}
func LightTrigger1b() {
	ns.CreateObject("RedPotion", wp73)
	ns.ObjectOn(obj61)
	ns.ObjectOff(obj57)
	ns.ObjectOff(obj53)
	ns.ObjectOn(obj45)
	ns.ObjectOff(obj49)
}
func LightTrigger2b() {
	ns.CreateObject("RedPotion", wp74)
	ns.ObjectOn(obj62)
	ns.ObjectOff(obj58)
	ns.ObjectOff(obj54)
	ns.ObjectOn(obj46)
	ns.ObjectOff(obj50)
}
func LightTrigger3b() {
	ns.CreateObject("RedPotion", wp75)
	ns.ObjectOn(obj63)
	ns.ObjectOff(obj59)
	ns.ObjectOff(obj55)
	ns.ObjectOn(obj47)
	ns.ObjectOff(obj51)
}
func DemonCreate(a1 int) {
	ivar37 += 1
	checkTotalDead()
	if !(ivar37 < 5) {
		goto LABEL1
	}
	if gvar41 != 0 {
		goto LABEL1
	}
	ivar38 = ns.Random(1, 100)
	if !(ivar38 <= 25 && gvar42 == 0) {
		goto LABEL2
	}
	gvar42 = 1
	obj66[a1] = ns.CreateObject("Demon", wp10[a1])
	ns.SetCallback(obj66[a1], 5, DemonDie)
	goto LABEL3
LABEL2:
	obj66[a1] = ns.CreateObject("EmberDemon", wp10[a1])
	ns.SetCallback(obj66[a1], 5, EmberDie)
LABEL3:
	ns.RetreatLevel(obj66[a1], 0)
	ns.SetCallback(obj66[a1], 4, DemonIdle)
	ns.SetCallback(obj66[a1], 4, DemonIdle)
	ns.LookAtObject(obj66[a1], ns.GetHost())
	ns.FrameTimerWithArg(5, a1, DemonInit)
LABEL1:
	return
}
func DemonInit(a1 int) {
	ns.SetCallback(obj66[a1], 4, DemonIdle)
	ns.LookAtObject(obj66[a1], ns.GetHost())
	ns.AggressionLevel(obj66[a1], 0.83)
	ns.Wander(obj66[a1])
}
func DemonIdle() {
	ns.AggressionLevel(ns.GetTrigger(), 0.83)
	ns.Wander(ns.GetTrigger())
}
func DemonDie() {
	gvar42 = 0
	if ivar39 != 0 {
		goto LABEL1
	}
	ivar39 += 1
	ns.FrameTimer(590, Flame6)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
	goto LABEL2
LABEL1:
	if ivar39 != 1 {
		goto LABEL3
	}
	ivar39 += 1
	ns.FrameTimer(590, Flame7)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
	goto LABEL2
LABEL3:
	if ivar39 != 2 {
		goto LABEL4
	}
	ivar39 += 1
	ns.FrameTimer(590, Flame8)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
	goto LABEL2
LABEL4:
	if ivar39 != 3 {
		goto LABEL2
	}
	ivar39 = 0
	ns.FrameTimer(590, Flame5)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
LABEL2:
	return
}
func EmberDie() {
	if ivar39 != 0 {
		goto LABEL1
	}
	ivar39 += 1
	ns.FrameTimer(590, Flame6)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
	goto LABEL2
LABEL1:
	if ivar39 != 1 {
		goto LABEL3
	}
	ivar39 += 1
	ns.FrameTimer(590, Flame7)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
	goto LABEL2
LABEL3:
	if ivar39 != 2 {
		goto LABEL4
	}
	ivar39 += 1
	ns.FrameTimer(590, Flame8)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
	goto LABEL2
LABEL4:
	if ivar39 != 3 {
		goto LABEL2
	}
	ivar39 = 0
	ns.FrameTimer(590, Flame5)
	ns.FrameTimerWithArg(600, ivar39, DemonCreate)
LABEL2:
	return
}
func SetupGuards() {
	ivar40 = 0
	for {
		if !(ivar40 < 5) {
			goto LABEL1
		}
		if ivar40 != 4 {
			goto LABEL2
		}
		Flame4()
	LABEL2:
		obj67[ivar40] = ns.CreateObject("EmberDemon", wp9[ivar40])
		ns.RetreatLevel(obj67[ivar40], 0)
		ns.CreatureGuard(obj67[ivar40], ns.GetWaypointX(wp9[ivar40]), ns.GetWaypointY(wp9[ivar40]), 2900, 2900, 300)
		ns.LookAtObject(obj67[ivar40], ns.GetHost())
		ns.FrameTimerWithArg(2, ivar40, GuardInit)
		ivar40 += 1
	}
LABEL1:
	ns.SetCallback(obj67[0], 5, Guard0Die)
	ns.SetCallback(obj67[1], 5, Guard1Die)
	ns.SetCallback(obj67[2], 5, Guard2Die)
	ns.SetCallback(obj67[3], 5, Guard3Die)
	ns.SetCallback(obj67[4], 5, Guard4Die)
}
func Guard0Die() {
	ns.FrameTimer(590, Flame0)
	ns.FrameTimerWithArg(600, nil, GuardCreate)
}
func Guard1Die() {
	ns.FrameTimer(590, Flame1)
	ns.FrameTimerWithArg(600, 1, GuardCreate)
}
func Guard2Die() {
	ns.FrameTimer(590, Flame2)
	ns.FrameTimerWithArg(600, 2, GuardCreate)
}
func Guard3Die() {
	ns.FrameTimer(590, Flame3)
	ns.FrameTimerWithArg(600, 3, GuardCreate)
}
func Guard4Die() {
	ns.FrameTimer(590, Flame4)
	ns.FrameTimerWithArg(600, 4, GuardCreate)
}
func GuardCreate(a1 int) {
	ivar36 += 1
	checkTotalDead()
	if !(ivar36 < 5) {
		goto LABEL1
	}
	if gvar41 != 0 {
		goto LABEL1
	}
	obj67[a1] = ns.CreateObject("EmberDemon", wp9[a1])
	ns.RetreatLevel(obj67[a1], 0)
	ns.CreatureGuard(obj67[a1], ns.GetWaypointX(wp9[a1]), ns.GetWaypointY(wp9[a1]), 2900, 2900, 300)
	ns.LookAtObject(obj67[a1], ns.GetHost())
	if a1 != 0 {
		goto LABEL2
	}
	ns.SetCallback(obj67[0], 5, Guard0Die)
LABEL2:
	if a1 != 1 {
		goto LABEL3
	}
	ns.SetCallback(obj67[1], 5, Guard1Die)
LABEL3:
	if a1 != 2 {
		goto LABEL4
	}
	ns.SetCallback(obj67[2], 5, Guard2Die)
LABEL4:
	if a1 != 3 {
		goto LABEL5
	}
	ns.SetCallback(obj67[3], 5, Guard3Die)
LABEL5:
	if a1 != 4 {
		goto LABEL6
	}
	ns.SetCallback(obj67[4], 5, Guard4Die)
LABEL6:
	ns.FrameTimerWithArg(2, a1, GuardInit)
LABEL1:
	return
}
func GuardInit(a1 int) {
	ns.CreatureGuard(obj67[a1], ns.GetWaypointX(wp9[a1]), ns.GetWaypointY(wp9[a1]), 2900, 2900, 300)
	ns.LookAtObject(obj67[a1], ns.GetHost())
}
func SetupDemons() {
	if !(ivar39 < 4) {
		goto LABEL1
	}
	ns.MoveObject(obj66[ivar39], ns.GetWaypointX(wp68[ivar39]), ns.GetWaypointY(wp68[ivar39]))
	ivar39 += 1
	ns.FrameTimer(30, SetupDemons)
LABEL1:
	return
}
func SetupFight() {
	ns.WideScreen(true)
	ns.LookAtObject(ns.GetHost(), obj64)
	ns.LookAtObject(obj64, ns.GetHost())
	ns.FrameTimer(50, HecTalk1)
}
func StartFight() {
	ns.WideScreen(false)
	ns.AggressionLevel(obj64, 0.83)
	ns.CreatureHunt(obj64)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(40, EnableDemons)
}
func HecTalk1() {
	ns.Enchant(obj64, ns.ENCHANT_PROTECT_FROM_FIRE, 0)
	ns.StoryPic(obj64, "HecubahPic")
	ns.SetDialog(obj64, ns.NEXT, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj64, ns.GetHost())
}
func HecTalk2() {
	gvar43 = 2
	ns.SetDialog(obj64, ns.NORMAL, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj64, ns.GetHost())
}
func hecTalkDemonStart() {
	if !flag33 {
		goto LABEL1
	}
	ns.Chat(obj64, "War11a:Bah")
	ns.FrameTimer(60, hecTalkDemonEnd)
LABEL1:
	return
}
func hecTalkDemonEnd() {
	flag32 = true
}
func hecTaunts() {
	if !flag32 {
		goto LABEL1
	}
	flag32 = false
	ns.Chat(obj64, "War11a:HecubahTaunts")
	ns.FrameTimer(600, resetTaunts)
LABEL1:
	return
}
func resetTaunts() {
	flag32 = true
}
func hecRetreatTalk() {
	if !ns.IsVisibleTo(ns.GetHost(), obj64) {
		goto LABEL1
	}
	ns.Frozen(ns.GetHost(), true)
	ns.SetDialog(obj64, ns.NORMAL, hecTalkRetreatStart, hecTalkRetreatEnd)
	ns.StartDialog(obj64, ns.GetHost())
LABEL1:
	return
}
func hecTalkRetreatStart() {
	ns.WideScreen(true)
	ns.Frozen(obj64, true)
	ns.Frozen(obj66[0], true)
	ns.Frozen(obj66[1], true)
	ns.Frozen(obj66[2], true)
	ns.Frozen(obj66[3], true)
	ns.Frozen(obj67[0], true)
	ns.Frozen(obj67[1], true)
	ns.Frozen(obj67[2], true)
	ns.Frozen(obj67[3], true)
	ns.Frozen(obj67[4], true)
	ns.TellStory(ns.HecubahRecognize, "War11a:HecubahIntro")
}
func hecTalkRetreatEnd() {
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(obj64, false)
	ns.Frozen(obj66[0], false)
	ns.Frozen(obj66[1], false)
	ns.Frozen(obj66[2], false)
	ns.Frozen(obj66[3], false)
	ns.Frozen(obj67[0], false)
	ns.Frozen(obj67[1], false)
	ns.Frozen(obj67[2], false)
	ns.Frozen(obj67[3], false)
	ns.Frozen(obj67[4], false)
	ns.CancelDialog(obj64)
}
func checkTotalDead() {
	if !(ivar37 >= 8 && ivar36 >= 9) {
		goto LABEL1
	}
	hecTalkDemonStart()
LABEL1:
	return
}
func EnableDemons() {
	SetupGuards()
	ivar39 = 0
	for {
		if !(ivar39 < 4) {
			goto LABEL1
		}
		ns.ObjectOn(obj66[ivar39])
		ns.RetreatLevel(obj66[ivar39], 0)
		if ivar39 == 2 {
			goto LABEL2
		}
		ns.SetCallback(obj66[ivar39], 5, EmberDie)
	LABEL2:
		ns.SetCallback(obj66[ivar39], 4, DemonIdle)
		ivar39 += 1
	}
LABEL1:
	return
}
func EndScene() {
	ns.Music(25, 100)
	ns.UnBlind()
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(wp71), ns.GetWaypointY(wp71))
	ns.LookAtObject(ns.GetHost(), obj65)
	ns.Frozen(ns.GetHost(), false)
	ns.FrameTimer(2, EndScene2)
}
func EndScene2() {
	ns.LookAtObject(ns.GetHost(), obj65)
	ns.Frozen(ns.GetHost(), true)
	ns.Frozen(obj65, true)
	gvar43 = 3
	HecubahDies()
	ns.SetDialog(obj64, ns.NORMAL, HecubahDialogStart, HecubahDialogEnd)
	ns.StartDialog(obj64, ns.GetHost())
}
func EndScene3() {
	ns.EndGame(0)
}
func HecubahDies() {
	ns.Frozen(obj65, false)
	ns.ObjectOn(obj65)
	ns.Damage(obj65, 0, 5000, 0)
	fvar34 = ns.GetObjectX(obj65)
	fvar35 = ns.GetObjectY(obj65)
	ns.MoveWaypoint(wp70, fvar34, fvar35)
	ns.Music(0, 100)
	ns.AudioEvent(ns.HecubahDieFrame0A, wp70)
	ns.FrameTimer(66, HecubahDie1)
	ns.FrameTimer(130, HecubahDie2)
	ns.FrameTimer(170, HecubahDie3)
	ns.FrameTimer(280, HecubahDie4)
}
func HecubahDie1() {
	ns.AudioEvent(ns.HecubahDieFrame98, wp70)
}
func HecubahDie2() {
	ns.AudioEvent(ns.HecubahDieFrame194, wp70)
}
func HecubahDie3() {
	ns.AudioEvent(ns.HecubahDieFrame283, wp70)
}
func HecubahDie4() {
	ns.AudioEvent(ns.HecubahDieFrame439, wp70)
}
func HecubahDie5() {
	ns.Blind()
	ns.FrameTimer(100, EndScene3)
}
func HecubahDialogStart() {
	if gvar43 != 1 {
		goto LABEL1
	}
	ns.TellStory(ns.DemonRecognize, "War11a:PunyFlea")
LABEL1:
	if gvar43 != 2 {
		goto LABEL2
	}
	ns.TellStory(ns.DemonRecognize, "War11a:CatchMe")
LABEL2:
	if gvar43 != 3 {
		goto LABEL3
	}
	ns.TellStory(ns.DemonRecognize, "War11a:Nooooooooo")
LABEL3:
	return
}
func HecubahDialogEnd() {
	if gvar43 != 1 {
		goto LABEL1
	}
	ns.CancelDialog(obj64)
	SetupDemons()
	ns.FrameTimer(150, HecTalk2)
LABEL1:
	if gvar43 != 2 {
		goto LABEL2
	}
	ns.CancelDialog(obj64)
	ns.Move(obj64, wp69)
LABEL2:
	if gvar43 != 3 {
		goto LABEL3
	}
	ns.CancelDialog(obj64)
	HecubahDie5()
LABEL3:
	return
}
func MapInitialize() {
	ns.Music(24, 100)
	obj64 = ns.Object("Hecubah")
	obj65 = ns.Object("Hecubah2")
	obj44 = ns.Object("LightTrigger0")
	obj45 = ns.Object("LightTrigger1")
	obj46 = ns.Object("LightTrigger2")
	obj47 = ns.Object("LightTrigger3")
	obj48 = ns.Object("LightTrigger0b")
	obj49 = ns.Object("LightTrigger1b")
	obj50 = ns.Object("LightTrigger2b")
	obj51 = ns.Object("LightTrigger3b")
	obj52 = ns.Object("InnerLight0")
	obj53 = ns.Object("InnerLight1")
	obj54 = ns.Object("InnerLight2")
	obj55 = ns.Object("InnerLight3")
	obj56 = ns.Object("MiddleLight0")
	obj57 = ns.Object("MiddleLight1")
	obj58 = ns.Object("MiddleLight2")
	obj59 = ns.Object("MiddleLight3")
	obj60 = ns.Object("OuterLight0")
	obj61 = ns.Object("OuterLight1")
	obj62 = ns.Object("OuterLight2")
	obj63 = ns.Object("OuterLight3")
	obj66[0] = ns.Object("Demon0")
	obj66[1] = ns.Object("Demon1")
	obj66[2] = ns.Object("Demon2")
	obj66[3] = ns.Object("Demon3")
	obj67[0] = ns.Object("Guard0")
	obj67[1] = ns.Object("Guard1")
	obj67[2] = ns.Object("Guard2")
	obj67[3] = ns.Object("Guard3")
	obj67[4] = ns.Object("Guard4")
	wp69 = ns.Waypoint("HecRun1")
	wp70 = ns.Waypoint("HecubahWP")
	wp71 = ns.Waypoint("EndPlayer")
	wp68[0] = ns.Waypoint("DemonStart0")
	wp68[1] = ns.Waypoint("DemonStart1")
	wp68[2] = ns.Waypoint("DemonStart2")
	wp68[3] = ns.Waypoint("DemonStart3")
	wp10[0] = ns.Waypoint("Loc0")
	wp10[1] = ns.Waypoint("Loc1")
	wp10[2] = ns.Waypoint("Loc2")
	wp10[3] = ns.Waypoint("Loc3")
	wp9[0] = ns.Waypoint("GuardLoc0")
	wp9[1] = ns.Waypoint("GuardLoc1")
	wp9[2] = ns.Waypoint("GuardLoc2")
	wp9[3] = ns.Waypoint("GuardLoc3")
	wp9[4] = ns.Waypoint("GuardLoc4")
	wp72 = ns.Waypoint("PotionLoc0")
	wp73 = ns.Waypoint("PotionLoc1")
	wp74 = ns.Waypoint("PotionLoc2")
	wp75 = ns.Waypoint("PotionLoc3")
	wp11[0] = ns.Waypoint("FireLocA0")
	wp11[1] = ns.Waypoint("FireLocA1")
	wp11[2] = ns.Waypoint("FireLocA2")
	wp11[3] = ns.Waypoint("FireLocA3")
	wp11[4] = ns.Waypoint("FireLocA4")
	wp11[5] = ns.Waypoint("FireLocA5")
	wp11[6] = ns.Waypoint("FireLocA6")
	wp11[7] = ns.Waypoint("FireLocA7")
	wp11[8] = ns.Waypoint("FireLocA8")
	wp12[0] = ns.Waypoint("FireLocB0")
	wp12[1] = ns.Waypoint("FireLocB1")
	wp12[2] = ns.Waypoint("FireLocB2")
	wp12[3] = ns.Waypoint("FireLocB3")
	wp12[4] = ns.Waypoint("FireLocB4")
	wp12[5] = ns.Waypoint("FireLocB5")
	wp12[6] = ns.Waypoint("FireLocB6")
	wp12[7] = ns.Waypoint("FireLocB7")
	wp12[8] = ns.Waypoint("FireLocB8")
	wp13[0] = ns.Waypoint("FireLocC0")
	wp13[1] = ns.Waypoint("FireLocC1")
	wp13[2] = ns.Waypoint("FireLocC2")
	wp13[3] = ns.Waypoint("FireLocC3")
	wp13[4] = ns.Waypoint("FireLocC4")
	wp13[5] = ns.Waypoint("FireLocC5")
	wp13[6] = ns.Waypoint("FireLocC6")
	wp13[7] = ns.Waypoint("FireLocC7")
	wp13[8] = ns.Waypoint("FireLocC8")
	ns.StartupScreen(11)
	ns.Frozen(ns.GetHost(), true)
	ns.FrameTimer(2, SetupFight)
}
func HecubahDie() {
	gvar41 = 1
	flag33 = false
	ns.Frozen(ns.GetHost(), true)
	ns.WideScreen(true)
	ns.ObjectOff(obj64)
	ns.DestroyEveryChat()
	ns.CancelDialog(obj64)
	ivar40 = 0
	for {
		if !(ivar40 < 5) {
			goto LABEL1
		}
		ns.Damage(obj67[ivar40], 0, 1000, 0)
		ivar40 += 1
	}
LABEL1:
	ivar39 = 0
	for {
		if !(ivar39 < 4) {
			goto LABEL3
		}
		ns.Damage(obj66[ivar39], 0, 1000, 0)
		ivar39 += 1
	}
LABEL3:
	ns.Blind()
	ns.FrameTimer(100, EndScene)
}
func killHecubah() {
	ns.Damage(obj64, 0, 5000, 0)
}
func PlayerDeath() {
	ns.DeathScreen(11)
}
func OnEvent(typ string) {
	switch typ {
	case "MapInitialize":
		MapInitialize()
	case "PlayerDeath":
		PlayerDeath()
	}
}
