/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package verify

import (
	"fmt"
)

type logger struct{}

func (logger) errorf(format string, args ...interface{}) {
	fmt.Printf("::error::"+format+"\n", args...)
}

func (logger) debugf(format string, args ...interface{}) {
	fmt.Printf("::debug::"+format+"\n", args...)
}

func (logger) warningf(format string, args ...interface{}) {
	fmt.Printf("::warning::"+format+"\n", args...)
}