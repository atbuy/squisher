package lib

import (
	"fmt"
	"strings"
)

func RLECompression(bin string) string {
	var result strings.Builder
	prev := '2'
	counter := 0

	for _, bit := range bin {
		if bit == prev {
			prev = bit
			counter += 1
		} else {
			if prev == '0' {
				result.WriteString(fmt.Sprintf("%da", counter))
			} else {
				result.WriteString(fmt.Sprintf("%db", counter))
			}

			prev = bit
			counter = 1
		}
	}

	return result.String()
}
