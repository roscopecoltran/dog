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

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/arbrain/abrain/base/alog"
)

// BrainMeta should be a generic structure for ArBrain.
// Every resource should be usable and accesible via brainQL only.
// should be in such a way that its generic enough with graph schema
// architecture.
// It should also capture out concepts action associated with the data
type BrainMeta struct {
	Nodes []NodeMeta `json:"nodes"`
}

// NodeMeta structure for generic node type.
type NodeMeta struct {
	Node       string      `json:"node"`
	Properties []PropsMeta `json:"properties"`
	EdgesOut   []EdgesMeta `json:"edgesOut"`
}

// PropsMeta structure for generic properties type for each node.
type PropsMeta struct {
	Property   string `json:"property"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Constraint string `json:"constraint"`
	Value      string `json:"value"`
}

// EdgesMeta structure for generic edgesOut type for each node.
type EdgesMeta struct {
	Edge              string      `json:"edge"`
	To                string      `json:"to"`
	Properties        []PropsMeta `json:"properties"`
	NodeProperty      string      `json:"nodeProperty"`
	NodePropertyValue string      `json:"nodePropertyValue"`
}

// GetBrainMeta returns BrainMeta datastructure object
// which contains schema of nodes and edges in the system.
func GetBrainMeta() BrainMeta {
	raw, err := ioutil.ReadFile("knowledge/schema_meta.json")
	if err != nil {
		alog.Printf(moduleName, "File error: %v\n", err)
		os.Exit(1)
	}

	var schema BrainMeta
	err2 := json.Unmarshal(raw, &schema)
	if err2 != nil {
		alog.Printf(moduleName, "error: %v", err2)
		os.Exit(1)
	}

	return schema
}

// GraphConfig structure for storing graphDB connection
// configuration.
type GraphConfig struct {
	GraphDB  string `json:"graphDB"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetGraphConfig returns GraphConfig datastructure object
// which contains connection info for connecting to graphdb.
func GetGraphConfig() GraphConfig {
	raw, err := ioutil.ReadFile("knowledge/graphdb_config.json")
	if err != nil {
		alog.Printf(moduleName, "File error: %v\n", err)
		os.Exit(1)
	}

	var config GraphConfig
	err2 := json.Unmarshal(raw, &config)
	if err2 != nil {
		alog.Printf(moduleName, "error: %v", err2)
		os.Exit(1)
	}

	return config
}
