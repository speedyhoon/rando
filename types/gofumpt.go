package types

import (
	"mvdan.cc/gofumpt/format"
	"runtime"
)

// goFormat nicely formats the generated Go code.
func goFormat(src []byte) (code []byte, err error) {
	code, err = format.Source(src, format.Options{
		LangVersion: runtime.Version(),
		ExtraRules:  true,
	})
	if err != nil {
		return src, err
	}
	return
}
