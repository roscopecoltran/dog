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

var (
	moduleName  = "knowledge"
	graphConfig GraphConfig
)

// Sync pushes data to graph store by transforming
// data into desirable format of graph store.
func Sync() {
	graphConfig = GetGraphConfig() // Read graphdb configuration data from json file.

	if graphConfig.GraphDB == "neo4j" {
		// If configured to use neo4j graph backend.
		alog.Printf(moduleName, "Neo4j backend selected.")

		SyncNeo4j()
	}

	alog.Printf(moduleName, "This will sync data to graphdb!")
}

// InitSchema initializes schema based on meta onto graphdb.
func InitSchema() {
	alog.Printf(moduleName, "Initializing Schema on graphdb!")

	/* TODO:
	1. Add code for Initializing schema on graphdb based on metadata.
	2. Add code for making a meta graph on graphdb.
	3. Make a generelized code for step 1 which will be applicable for all
		metas.
	*/

}
