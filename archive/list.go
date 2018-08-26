// Package archive exports a function that finds archives in a directory.
package archive

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

// Archive is a zip file represented by its name, size, and last
// modification-time.
type Archive struct {
	Name         string
	Size         int64
	ModifiedTime time.Time
}

// List of of Archives in the directory.
//
// Returns an error if the directory couldn't be read.
func List(d string) ([]Archive, error) {
	fs, err := ioutil.ReadDir(d)
	if err != nil {
		return nil, err
	}
	var as []Archive
	for _, f := range fs {
		if filepath.Ext(f.Name()) != ".zip" {
			continue
		}
		as = append(as, Archive{
			Name:         f.Name(),
			Size:         f.Size(),
			ModifiedTime: f.ModTime()})
	}
	return as, nil
}
