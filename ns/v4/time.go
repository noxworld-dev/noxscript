package ns

import "time"

// FrameFunc is a callback type for OnFrame.
type FrameFunc func()

// OnFrame registers a callback function that will be called each server frame.
// See FrameFunc for details.
func OnFrame(fnc FrameFunc) {
	if impl == nil {
		return
	}
	impl.OnFrame(fnc)
}

// OnEachFrame registers a callback function that will be called each N-th server frame.
// See FrameFunc for details.
func OnEachFrame(each int, fnc FrameFunc) {
	if impl == nil {
		return
	}
	if each <= 1 {
		impl.OnFrame(fnc)
		return
	}
	start := impl.Frame()
	impl.OnFrame(func() {
		if (impl.Frame()-start)%each == 0 {
			fnc()
		}
	})
}

// Frame returns current server frame number.
func Frame() int {
	if impl == nil {
		return 0
	}
	return impl.Frame()
}

// FrameRate returns current server update rate (aka server frame rate).
func FrameRate() int {
	if impl == nil {
		return 30
	}
	return impl.FrameRate()
}

// Now returns relative time from the server/map start.
// For absolute time, use time.Now().
func Now() time.Duration {
	if impl == nil {
		return 0
	}
	return impl.Time()
}

type Timer interface {
	TimerHandle
	// Cancel a timer. Returns true if successful.
	Cancel() bool
}

// NewTimer creates a timer that calls the given script function after a given delay.
//
// Example:
//
//	// Trigger by function reference:
//	NewTimer(Frames(10), myCallback)
//	// Trigger by function name (can call original NoxScript as well):
//	NewTimer(Seconds(10), "myCallback")
//	// Passing arguments to the callback:
//	NewTimer(Time(3*time.Second), "myCallback", obj)
func NewTimer(dt Duration, fnc Func, args ...any) Timer {
	if impl == nil {
		return nil
	}
	return impl.NewTimer(dt, fnc, args...)
}

// TimeSource is an interface for getting relative time.
type TimeSource interface {
	// Frame returns current server tick/frame number.
	Frame() int
	// Time returns relative time from server start.
	Time() time.Duration
}

// Frames returns duration in game frames.
func Frames(num int) Duration {
	return Duration{frames: num, kind: durFrames}
}

// Time returns duration given time.Duration.
func Time(dt time.Duration) Duration {
	return Duration{time: dt, kind: durTime}
}

// Seconds returns duration in seconds.
func Seconds(sec float64) Duration {
	return Time(time.Duration(sec * float64(time.Second)))
}

// Infinite returns an infinite duration.
func Infinite() Duration {
	return Duration{kind: durInfinite}
}

// NowWithSource returns current timestamp of the time source.
func NowWithSource(s TimeSource) Duration {
	return Duration{
		frames: s.Frame(),
		time:   s.Time(),
		kind:   durFrames | durTime,
	}
}

type durKind byte

func (k durKind) Has(k2 durKind) bool {
	return k&k2 != 0
}

const (
	durFrames = durKind(1 << iota)
	durTime
	durInfinite
)

// Duration for script events.
//
// It's separate from time.Duration because it can be expressed in game frames.
type Duration struct {
	frames int
	time   time.Duration
	kind   durKind
}

func (d Duration) LessOrEq(d2 Duration) bool {
	if d.IsInfinite() || d2.IsInfinite() {
		return false
	}
	if d.IsFrames() && d2.IsFrames() {
		return d.frames <= d2.frames
	}
	if d.IsTime() && d2.IsTime() {
		return d.time <= d2.time
	}
	return false
}

func (d Duration) Add(d2 Duration) Duration {
	if d.IsInfinite() || d2.IsInfinite() {
		return Infinite()
	}
	if d.IsFrames() && d2.IsFrames() {
		d.frames += d2.frames
	} else {
		d.frames = 0
		d.kind &^= durFrames
	}
	if d.IsTime() && d2.IsTime() {
		d.time += d2.time
	} else {
		d.time = 0
		d.kind &^= durTime
	}
	return d
}

func (d Duration) Sub(d2 Duration) Duration {
	if d.IsInfinite() || d2.IsInfinite() {
		return Infinite()
	}
	if d.IsFrames() && d2.IsFrames() {
		d.frames -= d2.frames
	} else {
		d.frames = 0
		d.kind &^= durFrames
	}
	if d.IsTime() && d2.IsTime() {
		d.time -= d2.time
	} else {
		d.time = 0
		d.kind &^= durTime
	}
	return d
}

// IsInfinite checks if duration is infinite.
func (d Duration) IsInfinite() bool {
	return d.kind.Has(durInfinite)
}

// IsTime checks if duration is set in seconds.
func (d Duration) IsTime() bool {
	return d.kind.Has(durTime)
}

// IsFrames checks if duration is set in frames.
func (d Duration) IsFrames() bool {
	return d.kind.Has(durFrames)
}

// Time checks if duration is set in seconds and returns it.
func (d Duration) Time() (time.Duration, bool) {
	if !d.IsTime() {
		return 0, false
	}
	return d.time, true
}

// Frames checks if duration is set in frames and returns it.
func (d Duration) Frames() (int, bool) {
	if !d.IsFrames() {
		return 0, false
	}
	return d.frames, true
}
