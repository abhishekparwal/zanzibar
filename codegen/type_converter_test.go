// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package codegen_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/codegen"
	"go.uber.org/thriftrw/compile"
)

var thriftExtensionLength int = len(".thrift")

type dummyFS map[string][]byte

// Read returns the contents for the specified file.
func (fs dummyFS) Read(filename string) ([]byte, error) {
	if contents, ok := fs[filename]; ok {
		return contents, nil
	}
	return nil, os.ErrNotExist
}

// Abs returns the absolute path for the specified file.
// The dummy implementation always returns the original path.
func (dummyFS) Abs(filename string) (string, error) {
	return filename, nil
}

type naivePackageNameResolver struct {
}

func (r *naivePackageNameResolver) TypePackageName(
	thriftFile string,
) (string, error) {
	if thriftFile[0] == '.' {
		return "", errors.Errorf("Naive does not support relative imports")
	}

	_, fileName := filepath.Split(thriftFile)

	return fileName[0 : len(fileName)-thriftExtensionLength], nil
}

func newTypeConverter() *codegen.TypeConverter {
	return &codegen.TypeConverter{
		Lines:  []string{},
		Helper: &naivePackageNameResolver{},
	}
}

func compileProgram(
	content string,
	otherFiles map[string][]byte,
) (*compile.Module, error) {
	if otherFiles == nil {
		otherFiles = map[string][]byte{}
	}
	otherFiles["structs.thrift"] = []byte(content)

	program, err := compile.Compile(
		"structs.thrift",
		compile.Filesystem(dummyFS(otherFiles)),
	)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func convertTypes(
	fromStruct string,
	toStruct string,
	content string,
	otherFiles map[string][]byte,
) (string, error) {
	converter := newTypeConverter()
	program, err := compileProgram(content, otherFiles)
	if err != nil {
		return "", err
	}

	err = converter.GenStructConverter(
		program.Types[fromStruct].(*compile.StructSpec).Fields,
		program.Types[toStruct].(*compile.StructSpec).Fields,
	)
	if err != nil {
		return "", err
	}

	return trim(strings.Join(converter.Lines, "\n")), nil
}

func countTabs(line string) int {
	count := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '\t' {
			count++
		} else {
			break
		}
	}
	return count
}

func trim(text string) string {
	lines := strings.Split(text, "\n")
	tabs := -1
	newLines := []string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		if tabs == -1 {
			tabs = countTabs(line)
		}

		if tabs >= len(line) {
			newLines = append(newLines, "")
		} else {
			newLines = append(newLines, line[tabs:])
		}
	}

	newText := strings.Join(newLines, "\n")
	return strings.TrimSpace(newText)
}

func TestConverStrings(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional string one
			2: required string two
		}

		struct Bar {
			1: optional string one
			2: required string two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*string)(in.One)
		out.Two = string(in.Two)
	`), lines)
}

func TestConvertBools(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional bool one
			2: required bool two
		}

		struct Bar {
			1: optional bool one
			2: required bool two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*bool)(in.One)
		out.Two = bool(in.Two)
	`), lines)
}

func TestConvertInt8(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional i8 one
			2: required i8 two
		}

		struct Bar {
			1: optional i8 one
			2: required i8 two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*int8)(in.One)
		out.Two = int8(in.Two)
	`), lines)
}

func TestConvertInt16(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional i16 one
			2: required i16 two
		}

		struct Bar {
			1: optional i16 one
			2: required i16 two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*int16)(in.One)
		out.Two = int16(in.Two)
	`), lines)
}

func TestConvertInt32(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional i32 one
			2: required i32 two
		}

		struct Bar {
			1: optional i32 one
			2: required i32 two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*int32)(in.One)
		out.Two = int32(in.Two)
	`), lines)
}

func TestConvertInt64(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional i64 one
			2: required i64 two
		}

		struct Bar {
			1: optional i64 one
			2: required i64 two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*int64)(in.One)
		out.Two = int64(in.Two)
	`), lines)
}

func TestConvertDouble(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional double one
			2: required double two
		}

		struct Bar {
			1: optional double one
			2: required double two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*float64)(in.One)
		out.Two = float64(in.Two)
	`), lines)
}

func TestConvertBinary(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct Foo {
			1: optional binary one
			2: required binary two
		}

		struct Bar {
			1: optional binary one
			2: required binary two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = []byte(in.One)
		out.Two = []byte(in.Two)
	`), lines)
}

func TestConvertStruct(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct NestedFoo {
			1: required string one
			2: optional string two
		}

		struct NestedBar {
			1: required string one
			2: optional string two
		}

		struct Foo {
			3: optional NestedFoo three
			4: required NestedFoo four
		}

		struct Bar {
			3: optional NestedBar three
			4: required NestedBar four
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		if in.Three != nil {
			out.Three = &structs.NestedBar{}
			out.Three.One = string(in.Three.One)
			out.Three.Two = (*string)(in.Three.Two)
		} else {
			out.Three = nil
		}
		if in.Four != nil {
			out.Four = &structs.NestedBar{}
			out.Four.One = string(in.Four.One)
			out.Four.Two = (*string)(in.Four.Two)
		} else {
			out.Four = nil
		}
	`), lines)
}

func TestHandlesMissingFields(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct NestedFoo {
			1: required string one
		}

		struct NestedBar {
			1: required string one
			2: optional string two
		}

		struct Foo {
			3: optional NestedFoo three
			4: required NestedFoo four
		}

		struct Bar {
			3: optional NestedBar three
			4: required NestedBar four
		}`,
		nil,
	)

	assert.Equal(t, "cannot map by name for the field two", err.Error())
	assert.Equal(t, "", lines)
}

func TestStructTypeMisMatch(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`struct NestedFoo {
			1: required string one
			2: optional string two
		}

		struct NestedBar {
			1: required string one
			2: optional string two
		}

		struct Foo {
			3: optional NestedFoo three
			4: required string four
		}

		struct Bar {
			3: optional NestedBar three
			4: required NestedBar four
		}`,
		nil,
	)

	assert.Equal(t, "", lines)
	assert.Equal(t, "could not convert struct fields, "+
		"incompatible type for four :", err.Error())
}

func TestConvertTypeDef(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`typedef string UUID

		struct Foo {
			1: optional UUID one
			2: required UUID two
		}

		struct Bar {
			1: optional UUID one
			2: required UUID two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*structs.UUID)(in.One)
		out.Two = structs.UUID(in.Two)
	`), lines)
}

func TestConvertEnum(t *testing.T) {
	lines, err := convertTypes(
		"Foo", "Bar",
		`enum ItemState {
			REQUIRED,
			OPTIONAL
		}

		struct Foo {
			1: optional ItemState one
			2: required ItemState two
		}

		struct Bar {
			1: optional ItemState one
			2: required ItemState two
		}`,
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, trim(`
		out.One = (*structs.ItemState)(in.One)
		out.Two = structs.ItemState(in.Two)
	`), lines)
}
