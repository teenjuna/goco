package internal

import (
	"strings"
)

func String[R Renderer](nodes ...R) (string, error) {
	buf := strings.Builder{}
	for _, node := range nodes {
		if err := node.Render(&buf); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func MustString[R Renderer](nodes ...R) string {
	s, err := String(nodes...)
	if err != nil {
		panic(err)
	}
	return s
}
