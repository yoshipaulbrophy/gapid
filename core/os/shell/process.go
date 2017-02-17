// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shell

import "github.com/google/gapid/core/log"

// Process is the interface to a running process, as started by a Target.
type Process interface {
	// Kill is called to ask the process to stop immediately.
	// It must be safe to call Kill more than once, and on a non running process.
	Kill() error
	// Wait blocks until the process has completed or the context is cancelled,
	// returning an error if the process failed for some reason.
	Wait(ctx log.Context) error
}