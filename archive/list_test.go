package archive_test

import (
	"testing"

	"github.com/jwowillo/backup/archive"
)

// TestList tests that List returns a list of the correct archives.
func TestList(t *testing.T) {
	as, err := archive.List("data")
	if err != nil {
		t.Error(err)
	}
	expected := []string{"a.zip", "b.zip", "c.zip"}
	same := true
	if len(as) != len(expected) {
		same = false
	} else {
		for i, a := range as {
			// Only check if names are equal. This is good enough to
			// make sure that files are at least read and the
			// correct files are ignored.
			if a.Name != expected[i] {
				same = false
			}
		}
	}
	if !same {
		t.Errorf("archive.List(%s) = %v, want %v", "data", as, expected)
	}
}

// TestListNoDir tests that List returns an error when given directories that
// don't exist.
func TestListNoDir(t *testing.T) {
	if _, err := archive.List(""); err == nil {
		t.Errorf("err = nil, want not nil")
	}
}
