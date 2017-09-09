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

package alog

import (
	"fmt"
	"log"
)

// Printf function to print logging data in proper format
// Pass module name as parameter so that logging could show
// Second param is msg that needs to be displayed
// Or format string need to parse subsequent parameters
// Third params that will be passed to format string
// From which module logs were produced.
func Printf(module string, format string, v ...interface{}) {
	s := fmt.Sprintf("<%s> %s", module, format)
	log.Printf(s, v...)
}
