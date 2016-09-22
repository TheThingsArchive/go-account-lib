// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package claims

// contains checks for membership of a string in a stringmap
func contains(slice []string, el string) bool {
	for _, e := range slice {
		if e == el {
			return true
		}
	}
	return false
}
