/*
* Copyright (C) 2017 Abrain, Ankur Yadav <ankurayadav@gmail.com>
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package knowledge

import "github.com/arbrain/abrain/base/alog"

// SyncNeo4j pushes data to neo4j by transforming
// data into desirable format of neo4j.
func SyncNeo4j() {
	alog.Printf(moduleName, "This will sync data to Neo4j!")
}

// InitSchemaNeo4j initializes schema based on meta onto neo4j.
func InitSchemaNeo4j() {
	alog.Printf(moduleName, "Initializing Schema on Neo4j!")

	/* TODO:
	1. Add code for Initializing schema on graphdb based on metadata.
	2. Add code for making a meta graph on graphdb.
	3. Make a generelized code for step 1 which will be applicable for all
		metas.
	*/

}
