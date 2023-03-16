package vms

import "time"

// TimeSource is an interface for getting relative time.
type TimeSource interface {
	// Frame returns current server tick/frame number.
	Frame() int
	// Time returns relative time from server start.
	Time() time.Duration
}

// Frames returns duration in game frames.
func Frames(num int) Duration {
	return Duration{frames: num}
}

// Time returns duration given time.Duration.
func Time(dt time.Duration) Duration {
	return Duration{time: dt}
}

// Seconds returns duration in seconds.
func Seconds(sec float64) Duration {
	return Time(time.Duration(sec * float64(time.Second)))
}

// Duration for script events.
//
// It's separate from time.Duration because it can be expressed in game frames.
type Duration struct {
	frames int
	time   time.Duration
}

func (d Duration) LessOrEq(d2 Duration) bool {
	if d.IsFrames() && d2.IsFrames() {
		return d.frames <= d2.frames
	}
	if d.IsTime() && d2.IsTime() {
		return d.time <= d2.time
	}
	return false
}

func (d Duration) Add(d2 Duration) Duration {
	if d.IsFrames() && d2.IsFrames() {
		d.frames += d2.frames
	}
	if d.IsTime() && d2.IsTime() {
		d.time += d2.time
	}
	return d
}

// IsTime checks if duration is set in seconds.
func (d Duration) IsTime() bool {
	return d.time != 0
}

// IsFrames checks if duration is set in frames.
func (d Duration) IsFrames() bool {
	return d.frames != 0
}

// Time checks if duration is set in seconds and returns it.
func (d Duration) Time() (time.Duration, bool) {
	return d.time, d.IsTime()
}

// Frames checks if duration is set in frames and returns it.
func (d Duration) Frames() (int, bool) {
	return d.frames, d.IsFrames()
}
