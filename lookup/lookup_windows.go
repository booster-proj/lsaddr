// +build windows

// Copyright © 2019 booster authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package lookup

import (
	"regexp"
	"fmt"
)

func buildRgx(s string) (*regexp.Regexp, error) {
	rgx, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}

	return rgx, nil
}

func openNetFiles(rgx *regexp.Regexp) ([]NetFile, error) {
	return []NetFile{}, fmt.Errorf("not implemented yet")
}
