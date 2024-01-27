package types

import (
	"bytes"
	"errors"
	"fmt"
)

// PackageSequence generates a Go main package a random struct with the given field types.
// It returns the package and test code required to test MarshalJ() & UnmarshalJ().
func PackageSequence(name string, typ string) (pkg, test []byte, err error) {
	b := bytes.NewBufferString(fmt.Sprintf(`// Code generated by rando; DO NOT EDIT.

package %s

`, name))
	tests := bytes.NewBuffer(b.Bytes())
	tests.WriteString(`

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

`)
	structs := bytes.NewBuffer(nil)
	typeNames := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
		"Twenty",
		"TwentyOne",
		"TwentyTwo",
		"TwentyThree",
	}
	for i, typeName := range typeNames {
		structLines, testLines := StructSequence(typeName, uint(i+1), typ)
		structs.Write(structLines)

		tests.WriteString(fmt.Sprintf(`
func TestFuzz_%[3]d(t *testing.T) {
	var expected, actual %[1]s
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, %[1]s{}, expected)
	require.Equal(t, %[1]s{}, actual)

	actual = %[1]s{
		%[2]s
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	//require.NotEqual(t, %[1]s{}, expected)
	//require.NotEqual(t, %[1]s{}, actual)
	require.Equal(t, expected, actual)
}
`, typeName, testLines, i+1))
	}

	if bytes.Contains(structs.Bytes(), []byte("time.Time")) {
		b.WriteString("import \"time\"\n\n")
	}
	b.Write(structs.Bytes())

	pkg, err1 := goFormat(b.Bytes())
	test, err2 := goFormat(tests.Bytes())

	return pkg, test, errors.Join(err1, err2)
}

// StructSequence generates a struct with random field names with the given types.
func StructSequence(name string, qty uint, typ string) (fields, testLines []byte) {
	b := bytes.NewBufferString(fmt.Sprintf("type %s struct{\n", name))
	tl := bytes.NewBuffer(nil)

	fieldNames := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
		"Twenty",
		"TwentyOne",
		"TwentyTwo",
		"TwentyThree",
	}

	for _, n := range fieldNames[:qty] {
		b.WriteString(n)
		b.WriteString("\t")
		b.WriteString(typ)
		b.WriteString("\n")

		// Test lines.
		tl.WriteString(n)
		tl.WriteString(":\t")
		tl.WriteString(TypeRandomFunc(typ))
		tl.WriteString(",\n")
	}
	b.WriteString("}\n\n")
	return b.Bytes(), tl.Bytes()
}