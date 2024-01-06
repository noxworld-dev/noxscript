package noxast_test

import (
	"bytes"
	"go/format"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/maps"
	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/shoenig/test/must"

	"github.com/noxworld-dev/noxscript/ns/asm"
	"github.com/noxworld-dev/noxscript/ns/v3/noxast"
)

func TestTranslate(t *testing.T) {
	const path = "../../nsvm/test.obj"

	f, err := os.Open(path)
	must.NoError(t, err)
	defer f.Close()

	s, err := asm.ReadScript(f)
	must.NoError(t, err)

	af := noxast.Translate(s, nil)
	var buf bytes.Buffer
	format.Node(&buf, token.NewFileSet(), af)
	t.Logf("\n%s", &buf)
}

func TestTranslateMaps(t *testing.T) {
	path := noxtest.DataPath(t, "maps")
	names, err := os.ReadDir(path)
	must.NoError(t, err)
	const outDir = ".out"
	_ = os.MkdirAll(outDir, 0755)
	for _, fi := range names {
		name := fi.Name()
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(path, name, name+".map")
			f, err := ifs.Open(path)
			must.NoError(t, err)
			defer f.Close()

			ss, err := maps.ReadScript(f)
			must.NoError(t, err)
			if len(ss.Data) == 0 {
				t.Skip("no scripts")
			}

			s, err := asm.ReadScript(bytes.NewReader(ss.Data))
			must.NoError(t, err)

			af := noxast.Translate(s, nil)
			var buf bytes.Buffer
			format.Node(&buf, token.NewFileSet(), af)
			t.Logf("\n%s", &buf)
			err = os.MkdirAll(filepath.Join(outDir, name), 0755)
			must.NoError(t, err)
			err = os.WriteFile(filepath.Join(outDir, name, name+".go"), buf.Bytes(), 0644)
			must.NoError(t, err)
			err = os.WriteFile(filepath.Join(outDir, name, "script_test.go"), []byte(`package script

import "testing"

func TestBuild(t *testing.T) {
}
`), 0644)
			must.NoError(t, err)
		})
	}
}
