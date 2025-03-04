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

// Inserter is used to construct an insert query
type Inserter struct {

}

func (i *Inserter) Build() (*Query, error) {
	panic("implement me")
}

// Insert generate Inserter to builder insert query
func (db *DB) Insert() *Inserter {
	return &Inserter{}
}

// Columns specifies the columns that need to be inserted
// if cs is empty, all columns will be inserted except auto increment columns
func (db *DB) Columns(cs...string) *Inserter {
	panic("implements me")
}

// Values specify the rows
// all the elements must be the same structure
func (i *Inserter) Values(values...interface{}) *Inserter {
	panic("implement me")
}

// OnDuplicateKey generate MysqlUpserter
// if the dialect is not MySQL, it will panic
func (i *Inserter) OnDuplicateKey() *MysqlUpserter {
	panic("implement me")
}

// OnConflict generate PgSQLUpserter
// if the dialect is not PgSQL, it will panic
func (i *Inserter) OnConflict(cs ...string) *PgSQLUpserter {
	panic("implement me")
}

