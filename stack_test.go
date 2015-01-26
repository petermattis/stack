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
	"fmt"
	"regexp"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack(0)
	fmt.Printf("%s\n", s)
	p := `(?m)^goroutine .+\ngithub.com/petermattis/stack\.TestStack\(.+$`
	if ok, _ := regexp.Match(p, s); !ok {
		t.Errorf("Expected %s, but got\n%s", p, s)
	}
}
