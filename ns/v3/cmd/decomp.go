package cmd

import (
	"bytes"
	"errors"
	"go/format"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/noxworld-dev/opennox-lib/maps"

	"github.com/noxworld-dev/noxscript/ns/asm"
	"github.com/noxworld-dev/noxscript/ns/v3/noxast"
)

func Decomp(args []string, c *noxast.Config) error {
	if len(args) != 1 && len(args) != 2 {
		return errors.New("expected one or two argument")
	}
	fname := args[0]
	if c.Package == "" {
		base := filepath.Base(fname)
		base = strings.TrimSuffix(base, filepath.Ext(base))
		if base != "" && !strings.ContainsAny(base, " ") {
			c.Package = strings.ToLower(base)
		}
	}
	var out io.Writer = os.Stdout
	if len(args) == 2 {
		f, err := os.Create(args[1])
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	var raw []byte
	if filepath.Ext(fname) == ".map" {
		s, err := maps.ReadScript(f)
		if err != nil {
			return err
		}
		raw = s.Data
	} else {
		raw, err = io.ReadAll(f)
		if err != nil {
			return err
		}
	}
	_ = f.Close()

	scr, err := asm.ReadScript(bytes.NewReader(raw))
	if err != nil {
		return err
	}
	astf := noxast.Translate(scr, c)
	return format.Node(out, token.NewFileSet(), astf)
}
