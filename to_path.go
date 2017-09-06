package godash

import (
	"regexp"
	"strconv"
	"fmt"
)

func ToPath(path string) []interface{} {
	r := regexp.MustCompile(`\[['"](.+)["']\]|^([A-Za-z0-1-_]+)|\.([A-Za-z0-1-_]+)|\[(\d+)\]`)
	d := regexp.MustCompile(`^\d+$`)
	parts := make([]interface{}, 0)

	matches := r.FindAllStringSubmatch(path, -1)
	for _, v := range matches {
		for j, m := range v {
			if m != "" && j != 0 {
				if d.MatchString(m) {
					if n, err := strconv.Atoi(m); err == nil {
						parts = append(parts, n)
					}
				} else {
					parts = append(parts, fmt.Sprintf("%v", m))
				}
			}
		}
	}

	return parts
}