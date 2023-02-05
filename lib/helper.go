package lib

import (
	"strconv"
)

type StringHelper struct{}

func (e *StringHelper) ToInt(value string) int {
	out, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int(out)
}
