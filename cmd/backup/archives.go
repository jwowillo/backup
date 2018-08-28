package main

import (
	"errors"

	"github.com/jwowillo/backup/archive"
	"github.com/jwowillo/cache/v2"
)

// ArchivesGetter gets all zip files at a path.
type ArchivesGetter func(string) ([]archive.Archive, error)

// MakeArchivesGetter creates the ArchivesGetter.
func MakeArchivesGetter(c cache.Cache) ArchivesGetter {
	return MakeArchivesGetterFromGetter(cache.NewFallbackGetter(
		c, MakeGetterFromArchivesGetter(archive.List)))
}

// MakeArchivesGetterFromGetter adapts a cache.Getter to a ArchivesGetter.
func MakeArchivesGetterFromGetter(g cache.Getter) ArchivesGetter {
	return func(p string) ([]archive.Archive, error) {
		x := g.Get(cache.Key(p))
		if x == nil {
			return nil, errors.New("couldn't get archives")
		}
		return x.([]archive.Archive), nil
	}
}

// MakeGetterFromArchivesGetter adapts a ArchivesGetter to a cache.Getter.
func MakeGetterFromArchivesGetter(tg ArchivesGetter) cache.Getter {
	return cache.GetterFunc(func(k cache.Key) cache.Value {
		bs, err := tg(string(k))
		if err != nil {
			return nil
		}
		return bs
	})
}
