// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/specvalue"

	// "fmt"
)

type svdep struct {
	SpecValue *specvalue.SpecValue
}

type SpecValueTest struct {
	Dependences svdep
}

func (s SpecValueTest) Test(ctx *cyako.Ctx) {
	// specvalue := s.Dependences.SpecValue
	// specvalue.SetInt("1")
	// r := specvalue.GetInt("1")
}

func init() {
	var m = SpecValueTest{
		Dependences: svdep{
			SpecValue: cyako.Svc["SpecValue"].(*specvalue.SpecValue),
		},
	}
	cyako.LoadModule(m)
}
