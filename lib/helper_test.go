package lib

import (
	"testing"
)

func HelperTest(t *testing.T) {
	helper := StringHelper{}
	if helper.ToInt("123") != 123 {
		t.Error("Numeric conversion failed")
	}
}
