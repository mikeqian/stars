package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSubString(t *testing.T) {
	source := "你012妹子mm3456"
	begin := strings.Index(source, "妹子")
	l := len("妹子")

	fmt.Println(source[begin+l : begin+l+2])

	fmt.Println(begin)
	t.Log("pass")
}
