package utils

import (
	"testing"
)

func TestDBNew(t *testing.T) {
	db := DBNew()
	if db == nil {
		t.Error("fail connect to DB")
	}
}
