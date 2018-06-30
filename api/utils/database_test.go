package utils

import (
	"testing"
)

// TestDBNew
// Database Connection Test
func TestDBNew(t *testing.T) {
	db := DBNew()
	if db == nil {
		t.Error("fail connect to DB")
	}
}
