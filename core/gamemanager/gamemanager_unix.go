// Copyright 2024 bbaa
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unix

package main

import (
	"io"
	"syscall"
)

var MinecraftProcess_SysProcAttr = &syscall.SysProcAttr{
	Setpgid: true,
}

func (pty *MinecraftPty) readerWrapper(r io.Reader) io.Reader {
	return r
}

func (pty *MinecraftPty) writerWrapper(w io.WriteCloser) io.WriteCloser {
	return w
}
