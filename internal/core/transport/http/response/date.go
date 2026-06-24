package core_http_response

import "strings"

func FormatDate(s string) string {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return s
	}
	return parts[1] + "-" + parts[0]
}
