package vmtest

import (
	"bytes"
	"testing"
	"time"

	"github.com/noxworld-dev/noxscript/vms"
)

var _ vms.Engine = &Engine{}

type Engine struct {
	T        testing.TB
	FrameVal int
	TimeVal  time.Duration
	stdout   *writer
	stderr   *writer
}

func (e *Engine) Close() {
	e.stdout.flush()
	e.stderr.flush()
}

func (e *Engine) Frame() int {
	return e.FrameVal
}

func (e *Engine) Time() time.Duration {
	return e.TimeVal
}

func (e *Engine) Console(error bool) vms.Writer {
	if e.stdout == nil {
		e.stdout = &writer{t: e.T, lvl: "info"}
	}
	if e.stderr == nil {
		e.stderr = &writer{t: e.T, lvl: "error"}
	}
	if error {
		return e.stderr
	}
	return e.stdout
}

type writer struct {
	t   testing.TB
	lvl string
	buf bytes.Buffer
}

func (w *writer) maybeFlush() {
	w.t.Helper()
	for {
		buf := w.buf.Bytes()
		i := bytes.IndexByte(buf, '\n')
		if i < 0 {
			break
		}
		w.t.Logf("%s: %s", w.lvl, string(buf[:i]))
		w.buf.Next(i + 1)
	}
}

func (w *writer) flush() {
	w.t.Helper()
	w.maybeFlush()
	if w.buf.Len() != 0 {
		w.t.Logf("%s: %s", w.lvl, w.buf.String())
		w.buf.Reset()
	}
}

func (w *writer) Write(p []byte) (int, error) {
	n, _ := w.buf.Write(p)
	w.maybeFlush()
	return n, nil
}

func (w *writer) WriteString(s string) (int, error) {
	n, _ := w.buf.WriteString(s)
	w.maybeFlush()
	return n, nil
}
