package hook

import (
	"fmt"
	"lig"
)

func WithLevelPrefix(level lig.Level, msg string) string {
	return fmt.Sprintf("[%-5v] ", level.ToString()) + msg
}
