package nstest

import "sync/atomic"

type Timers struct {
	last uint32
	byID map[uint32]*Timer
}

func (s *Timers) NewTimer(frames uint32, fnc func()) *Timer {
	id := atomic.AddUint32(&s.last, 1)
	t := &Timer{s: s, id: id, left: frames, fnc: fnc}
	if s.byID == nil {
		s.byID = make(map[uint32]*Timer)
	}
	s.byID[id] = t
	return t
}

func (s *Timers) Update() {
	for _, t := range s.byID {
		t.Update()
	}
}

type Timer struct {
	s    *Timers
	id   uint32
	left uint32
	fnc  func()
}

func (t *Timer) ScriptID() int {
	return int(t.id)
}

func (t *Timer) TimerScriptID() int {
	return int(t.id)
}

func (t *Timer) Cancel() bool {
	ok := t.fnc != nil
	t.fnc = nil
	delete(t.s.byID, t.id)
	return ok
}

func (t *Timer) Update() {
	if t.fnc == nil {
		return
	}
	if t.left > 0 {
		t.left--
	}
	if t.left == 0 {
		t.fnc()
		t.fnc = nil
		delete(t.s.byID, t.id)
	}
}
