// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package time

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/chaos-mesh/chaos-mesh/pkg/mapreader"
)

// vdsoEntryName is the name of the vDSO entry
const vdsoEntryName = "[vdso]"

// FakeImage introduce the replacement of VDSO ELF entry and customizable variables.
// FakeImage could be constructed by LoadFakeImageFromEmbedFs(), and then used by FakeClockInjector.
type FakeImage struct {
	// symbolName is the name of the symbol to be replaced.
	symbolName string
	// content presents .text section which has been "manually relocation", the address of extern variables have been calculated manually
	content []byte
	// offset stores the table with variable name, and it's address in content.
	// the key presents extern variable name, ths value is the address/offset within the content.
	offset map[string]int
	// OriginFuncCode stores the raw func code like getTimeOfDay & ClockGetTime.
	OriginFuncCode []byte
	// OriginAddress stores the origin address of OriginFuncCode.
	OriginAddress uint64
	// fakeEntry stores the fake entry
	fakeEntry *mapreader.Entry

	logger logr.Logger
}

func NewFakeImage(symbolName string, content []byte, offset map[string]int, logger logr.Logger) *FakeImage {
	return &FakeImage{symbolName: symbolName, content: content, offset: offset, logger: logger}
}

func (it *FakeImage) AttachToProcess(pid int, variables map[string]uint64) (err error) {
	return errors.New("not implement fake image arm64")
}

func (it *FakeImage) Recover(pid int, vars map[string]uint64) error {
	return errors.New("not implement fake image arm64")
}
