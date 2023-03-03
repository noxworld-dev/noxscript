package noxscript

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/noxscript/ns/asm"
)

func TestRuntime(t *testing.T) {
	const path = "test.obj"

	f, err := os.Open(path)
	require.NoError(t, err)
	defer f.Close()

	r, err := LoadScript(f)
	require.NoError(t, err)

	err = r.CallByName("Add", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, 3, r.PopInt())

	err = r.CallByName("Sub", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, -1, r.PopInt())

	err = r.CallByName("If", nil, nil, 1)
	require.NoError(t, err)
	require.Equal(t, 1, r.PopInt())

	err = r.CallByName("If", nil, nil, 2)
	require.NoError(t, err)
	require.Equal(t, 1, r.PopInt())

	err = r.CallByName("If", nil, nil, 0)
	require.NoError(t, err)
	require.Equal(t, 0, r.PopInt())

	err = r.CallByName("AddLocal", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, 4, r.PopInt())
	err = r.CallByName("AddLocal", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, 4, r.PopInt())

	err = r.CallByName("AddGlobal", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, 3, r.PopInt())
	err = r.CallByName("AddGlobal", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, 6, r.PopInt())

	err = r.CallByName("AddLocalArr", nil, nil, 1, 2)
	require.NoError(t, err)
	require.Equal(t, 3, r.PopInt())
	err = r.CallByName("AddLocalArrFail", nil, nil, 1, 2)
	require.Error(t, err)
}

func TestDisasm(t *testing.T) {
	const path = "test.obj"

	f, err := os.Open(path)
	require.NoError(t, err)
	defer f.Close()

	s, err := asm.ReadScript(f)
	require.NoError(t, err)

	var buf bytes.Buffer
	for _, fnc := range s.Funcs {
		t.Run(fnc.Name, func(t *testing.T) {
			t.Logf("Args: %d, Returns: %d", fnc.Args, fnc.Return)
			for i, v := range fnc.Vars {
				sz := ""
				if v.Size > 1 {
					sz = fmt.Sprintf("[%d]", v.Size)
				}
				t.Logf("local_%d%s (+%d)", i, sz, v.Offs)
			}
			code, err := asm.Decode(fnc.Code)
			require.NoError(t, err)
			buf.Reset()
			asm.Print(&buf, code)
			t.Logf("\n%s", &buf)
			require.Equal(t, fnc.Code, asm.Encode(code))
		})
	}
}
