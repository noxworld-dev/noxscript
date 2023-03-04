package eud

import ns3 "github.com/noxworld-dev/noxscript/ns/v3"

func WallUtilWallIsOpened(wallID ns3.WallID) bool {
	panic("not implemented")
}

func WallUtilWallIsDestroyed(wallID ns3.WallID) bool {
	panic("not implemented")
}

func WallUtilCreateMagicWall(sPtr int, sWallX, sWallY int, sWallDir int, sUnk2 int) ns3.WallID {
	panic("not implemented")
}

func WallUtilAddBreakableWall(wallID ns3.WallID) {
	panic("not implemented")
}

func WallUtilGetWallAtObjectPosition(anyunit int) ns3.WallID {
	var xPos, yPos int
	unit := ns3.ObjectID(anyunit)
	if unit >= ns3.GetHost() || unit == ns3.SELF || unit == ns3.OTHER {
		xPos = int(ns3.GetObjectX(unit))
		yPos = int(ns3.GetObjectY(unit))
	} else {
		wp := ns3.WaypointID(anyunit)
		xPos = int(ns3.GetWaypointX(wp))
		yPos = int(ns3.GetWaypointY(wp))
	}
	if xPos > yPos {
		xPos += 23
	} else {
		yPos += 23
	}
	spX := (xPos + yPos - 22) / 46
	spY := (xPos - yPos) / 46
	tx := spX * 46
	ty := spY * 46
	rx := (tx + ty) / 2
	return ns3.Wall(rx/23, (rx-ty)/23)
}

func WallUtilCloseWallAtObjectPosition(anyunit int) {
	w := WallUtilGetWallAtObjectPosition(anyunit)
	if w != 0 {
		ns3.WallClose(w)
	}
}

func WallUtilOpenWallAtObjectPosition(anyunit int) {
	w := WallUtilGetWallAtObjectPosition(anyunit)
	if w != 0 {
		ns3.WallOpen(w)
	}
}

func WallUtilDestroyWallAtObjectPosition(anyunit int) {
	w := WallUtilGetWallAtObjectPosition(anyunit)
	if w != 0 {
		ns3.WallBreak(w)
	}
}

func WallUtilToggleWallAtObjectPosition(anyunit int) {
	w := WallUtilGetWallAtObjectPosition(anyunit)
	if w != 0 {
		ns3.WallToggle(w)
	}
}
