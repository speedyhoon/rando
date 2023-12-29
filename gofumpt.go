package rando

import (
	"mvdan.cc/gofumpt/format"
	"runtime"
	"strings"
)

// goFormat nicely formats the generated Go code.
func goFormat(src []byte) (code []byte, err error) {
	code, err = format.Source(src, format.Options{
		LangVersion: strings.TrimPrefix(runtime.Version(), "go"),
		ExtraRules:  true,
	})
	if err != nil {
		return src, err
	}
	return
}
