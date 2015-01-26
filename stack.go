// Copyright 2015 Peter Mattis.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

package stack

import (
	"bytes"
	"runtime"
)

// trimLine returns a subslice of b by slicing off the bytes up to and
// including the first newline. Returns nil if b does not contain a
// newline. This method is safe to call with b==nil.
func trimLine(b []byte) []byte {
	index := bytes.IndexByte(b, '\n')
	if index == -1 {
		return nil
	}
	return b[index+1:]
}

// Stack formats the stack trace of the calling goroutine. The
// argument skip is the number of stack frames to skip before recoding
// the stack trace, with 0 identifying the caller of stack.
func Stack(skip int) []byte {
	// Grow buf until it's large enough to store the entire stack trace.
	buf := make([]byte, 1024)
	var n int
	for {
		n = runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, len(buf)*2)
	}

	// Skip over skip+1 stack frames. The output of runtime.Stack looks
	// like:
	//
	//   goroutine <num> [<state>]:
	//   <func1>(<addr1>)
	//     <file1>:<line1> +<offset1>
	//   <func2>(<addr2>)
	//     <file2>:<line2> +<offset2>
	//   ...
	//
	// We want to keep the first line identifying the goroutine, then
	// skip over skip+1 pairs of lines.
	start := trimLine(buf)
	end := start

	// There is a pair of lines per stack frame.
	for i := 0; i <= skip; i++ {
		end = trimLine(trimLine(end))
	}

	// Copy the bytes starting at "end" to the bytes starting at
	// "start", overwriting the first "skip+1" stack frames.
	copy(start, end)

	// We deleted the bytes between end and start and need to trim the
	// size of the buffer.
	n -= (len(start) - len(end))
	return buf[:n]
}
