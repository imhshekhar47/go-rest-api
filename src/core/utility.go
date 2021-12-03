package core

import "strings"

// NVL: Returns primary if not empty else otherwise
func OneOfTheVal(primary string, otherwise string) string {
	if len(strings.TrimSpace(primary)) == 0 {
		return otherwise
	}

	return primary
}
