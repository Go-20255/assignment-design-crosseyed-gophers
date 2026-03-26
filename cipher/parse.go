package cipher

import (
	"fmt"
	"strconv"
	"strings"
)

func getRequired(params map[string]string, name string) (string, error) {
	v, ok := params[name]
	if !ok || strings.TrimSpace(v) == "" {
		return "", fmt.Errorf("missing required parameter --%s", name)
	}
	return v, nil
}

func getOptional(params map[string]string, name string) (string, bool) {
	v, ok := params[name]
	if !ok {
		return "", false
	}
	return v, true
}

func parseIntParam(raw string) (int, error) {
	n, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil {
		return 0, fmt.Errorf("expected int, got %q", raw)
	}
	return n, nil
}

func parseByteParam(raw string) (byte, error) {
	n, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil {
		return 0, fmt.Errorf("expected byte (0-255), got %q", raw)
	}
	if n < 0 || n > 255 {
		return 0, fmt.Errorf("byte out of range 0-255: %d", n)
	}
	return byte(n), nil
}

func encodeInt(n int) string { return strconv.Itoa(n) }

func encodeByte(b byte) string { return strconv.Itoa(int(b)) }
