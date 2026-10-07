package vm

import "strings"

type VM struct {
	Name   string
	CPU    int
	RAM    int
	Status string
}

func buildVMName(prefix string, id int) string {
	return prefix + "-" + strings.Repeat("x", id)
}