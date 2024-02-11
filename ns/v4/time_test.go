package ns

import (
	"testing"
	"time"

	"github.com/shoenig/test/must"
)

func TestTime(t *testing.T) {
	must.EqOp(t, time.Second/2, Seconds(0.5).time)
	must.EqOp(t, time.Second, Seconds(1).time)
	must.EqOp(t, time.Second*2, Seconds(2).time)
}
