// Copyright 2021 gotomicro
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

package eql

type Column struct {
	c string
	alias string
}

// C specify column
func C(c string) Column {
	return Column{
		c: c,
	}
}

// As means alias
func (Column) As(alias string) Selectable {
	panic("implement me")
}

func (Column) Inc(val interface{}) MathExpr {
	panic("implement me")
}

func (Column) Times(val interface{}) MathExpr {
	panic("implement me")
}

func (Column) expr() (string, error) {
	panic("implement me")
}

func (Column) selected() {
	panic("implement me")
}

type columns struct {
	cs []string
}

func (c columns) selected() {
	panic("implement me")
}

// Columns specify columns
func Columns(cs...string) Selectable {
	return columns{
		cs: cs,
	}
}

