package helper

import (
	"strings"

	"github.com/ulule/deepcopier"
)

func Masking(value string) string {
	var firstDigits = value[0 : len(value)-4]
	data := []string{strings.Repeat("x", len(firstDigits)), "x", "x", "x", "x"}
	value = strings.Replace(value, firstDigits, strings.Join(data, ""), -1)
	return value
}

func DeepCopy(from, to interface{}) {

	deepcopier.Copy(from).To(to)
}
