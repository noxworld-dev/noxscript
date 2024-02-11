package nstest

import (
	"testing"
	"time"

	"github.com/shoenig/test/must"
)

func TestTimers(t *testing.T) {
	r := NewRuntime()

	// Check basics
	ok1 := false
	ok2 := false
	r.Timers.New(0, func() {
		ok1 = true
	})
	r.Timers.New(1, func() {
		ok2 = true
	})
	must.False(t, ok1)
	must.False(t, ok2)
	r.Update()
	must.True(t, ok1) // Timers with 0 duration trigger instantly
	must.True(t, ok2) // Delayed by 1 frame
	// Timers are removed properly
	must.MapLen(t, 0, r.Timers.byID)

	// Check longer timers
	ok1 = false
	r.Timers.New(3, func() {
		ok1 = true
	})
	for i := 0; i < 2; i++ {
		r.Update()
		must.False(t, ok1)
	}
	r.Update()
	must.True(t, ok1)

	// Check that duration timers based on seconds trigger at exact frame
	start := r.Frame()
	ok1 = false
	ok2 = false
	r.Timers.NewDur(time.Second, func() {
		ok1 = true
		must.EqOp(t, start+FrameRate, r.Frame())
		r.Timers.NewDur(time.Second, func() {
			ok2 = true
			must.EqOp(t, start+2*FrameRate, r.Frame())
		})
	})

	for i := 0; i < FrameRate-1; i++ {
		r.Update()
		must.False(t, ok1)
		must.False(t, ok2)
	}
	r.Update()
	must.True(t, ok1)

	for i := 0; i < FrameRate-1; i++ {
		r.Update()
		must.False(t, ok2)
	}
	r.Update()
	must.True(t, ok2)
}
