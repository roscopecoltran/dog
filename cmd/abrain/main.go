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

package main

import (
	"net/http"

	"github.com/arbrain/abrain/base/alog"
	"github.com/arbrain/abrain/knowledge"
)

var (
	moduleName = "abrain"
)

func main() {

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		alog.Printf(moduleName, "Query Received")
	})

	// Display some basic instructions
	alog.Printf(moduleName, "Server is running on port 8080")

	http.ListenAndServe(":8080", nil)

	/* Adding sample graphQL style query that
	is expected from client side
	which should be parsed and broken properly
	and then should be sent to respective modules
	for processing.
	*/

	// To create data in a batch.
	// This should contain sample data for graph
	// and also its associated functions.
	queryCreate := `{
		// Will add some sample graphQL query.
	}`

	alog.Printf(moduleName, queryCreate)

	knowledge.Sync()
}
