package types

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/speedyhoon/rando"
	"github.com/speedyhoon/utl"
	"math/rand"
	"strings"
)

// Package generates a Go main package with 1 to 6 randomly generated structs.
// It returns the package and test code required to test MarshalJ() & UnmarshalJ().
func Package(name string) (pkg, test []byte, err error) {
	b := bytes.NewBufferString(fmt.Sprintf("package %s\n\n", name))
	structNames := make(uniqueNames)

	tests := bytes.NewBuffer(b.Bytes())
	tests.WriteString(`

import (
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
	"testing"
)

`)
	structs := bytes.NewBuffer(nil)
	mx := rand.Intn(5) + 1
	for i := 0; i < mx; i++ {
		typeName := structNames.add(StringExported)
		structLines, testLines := Struct(typeName, 150)
		structs.Write(structLines)

		tests.WriteString(fmt.Sprintf(`
func TestFuzz_%d(t *testing.T) {
	var expected, actual %s
	require.NoError(t, actual.UnmarshalJ(expected.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, %[2]s{}, expected)
	require.Equal(t, %[2]s{}, actual)

	expected = %[2]s{
		%[3]s
	}
	src := expected.MarshalJ()
	require.NoError(t, actual.UnmarshalJ(src))
	require.NotEqual(t, %[2]s{}, expected)
	require.NotEqual(t, %[2]s{}, actual)
	require.Equal(t, expected, actual)
}
`, i, typeName, testLines))
	}
	if bytes.Contains(structs.Bytes(), []byte("time.Time")) {
		b.WriteString("import \"time\"\n\n")
	}
	b.Write(structs.Bytes())

	pkg, err1 := goFormat(b.Bytes())
	test, err2 := goFormat(tests.Bytes())

	return pkg, test, errors.Join(err1, err2)
}

type uniqueNames map[string]struct{}

// add executes function f until a unique string is generated, not already in uniqueNames.
func (u uniqueNames) add(f func() string) string {
	for i := 0; i <= 999; i++ {
		tmp := f()
		_, ok := u[tmp]
		if !ok {
			u[tmp] = struct{}{}
			return tmp
		}
	}
	panic("function `f` does not generate enough randomness")
}

// Struct generates a random struct with a random number of fields.
func Struct(name string, fieldsQty uint) (fields, testLines []byte) {
	b := bytes.NewBufferString(fmt.Sprintf("type %s struct{\n", name))
	tl := bytes.NewBuffer(nil)

	fieldNames := make(uniqueNames)

	mx := rand.Intn(int(fieldsQty)) + 1
	for i := 0; i < mx; i++ {
		//b.WriteString(ExportedNames())
		n := fieldNames.add(FieldName)
		typ := RandomType()
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
	b.WriteString("}\n")
	return b.Bytes(), tl.Bytes()
}

func TypeRandomFunc(t string) string {
	var n interface{}
	switch t {
	case "time.Duration":
		n = rando.Duration
	case "[]time.Duration":
		n = rando.Durations
	case "time.Time":
		n = rando.Time
	case "[]time.Time":
		n = rando.Times
	case "struct{}":
		return "struct{}{}"
	case "[]byte":
		n = rando.BytesNil
	case "[]complex64":
		n = rando.Complex64s
	case "[]complex128":
		n = rando.Complex128s
	case "[]uint":
		n = rando.Uints
	case "[]uint8":
		n = rando.Uint8s
	case "[]uint16":
		n = rando.Uint16s
	case "[]uint32":
		n = rando.Uint32s
	case "[]uint64":
		n = rando.Uint64s
	case "[]int":
		n = rando.Ints
	case "[]int8":
		n = rando.Int8s
	case "[]int16":
		n = rando.Int16s
	case "[]int32", "[]rune":
		n = rando.Int32s
	case "[]int64":
		n = rando.Int64s
	case "[]bool":
		n = rando.Bools
	case "[]float32":
		n = rando.Float32s
	case "[]float64":
		n = rando.Float64s
	case "[]string":
		n = rando.Strings
	default:
		return "rando." + strings.Title(t) + "()"
	}
	return utl.NameOfFunc(n) + "()"
}

func SupportedList() []string {
	return []string{
		"bool",
		"byte",
		"complex64",
		"complex128",
		"float32",
		"float64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"rune",
		"string",
		"time.Duration",
		"time.Time",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"[]bool",
		"[]byte",
		"[]complex64",
		"[]complex128",
		"[]float32",
		"[]float64",
		"[]int",
		"[]int8",
		"[]int16",
		"[]int32",
		"[]int64",
		"[]rune",
		"[]string",
		"[]time.Duration",
		"[]time.Time",
		"[]uint",
		"[]uint8",
		"[]uint16",
		"[]uint32",
		"[]uint64",
	}
}

func RandomType() string {
	list := SupportedList()
	return list[rand.Intn(len(list))]
}

func ExportedNames() string {
	fieldNames := make([]string, rand.Intn(5)+1)
	for j := range fieldNames {
		fieldNames[j] = StringExported()
	}
	return strings.Join(fieldNames, ", ")
}

func StringExported() string {
	s := UppercaseChar()

	mx := rand.Intn(100)
	for i := 0; i < mx; i++ {
		s += AlphaNumeric()
	}
	return s
}

func PackageName() string {
	s := LowercaseChar()

	mx := rand.Intn(100)
	for i := 0; i < mx; i++ {
		s += AlphaNumeric()
	}
	return s
}
func FieldName() string {
	s := AlphaChar()

	mx := rand.Intn(100)
	for i := 0; i < mx; i++ {
		s += AlphaNumeric()
	}
	return s
}

// AlphaNumeric returns an uppercase, lowercase letter or number digit.
func AlphaNumeric() string {
	switch rand.Intn(3) {
	case 1:
		return UppercaseChar()
	case 2:
		return LowercaseChar()
	default:
		return NumberChar()
	}
}

// AlphaChar returns an uppercase or lowercase letter.
func AlphaChar() string {
	if rand.Intn(1) == 0 {
		return UppercaseChar()
	}
	return LowercaseChar()
}

func NumberChar() string {
	return fmt.Sprintf("%c", rand.Intn('9'-'0')+'0')
}

func LowercaseChar() string {
	return fmt.Sprintf("%c", rand.Intn('z'-'a')+'a')
}

func UppercaseChar() string {
	return fmt.Sprintf("%c", rand.Intn('Z'-'A')+'A')
}
