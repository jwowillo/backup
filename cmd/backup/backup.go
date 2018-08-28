package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/jwowillo/cache/v2"
)

// BackupGetter gets the backup-script after injecting arguments.
type BackupGetter func(string) ([]byte, error)

// MakeBackupGetter creates the BackupGetter.
func MakeBackupGetter(c cache.Cache, host string) BackupGetter {
	b := func(p string) ([]byte, error) {
		return Backup(p, host)
	}
	return MakeBackupGetterFromGetter(cache.NewFallbackGetter(
		c, MakeGetterFromBackupGetter(b)))
}

// MakeBackupGetterFromGetter adapts a cache.Getter to a BackupGetter.
func MakeBackupGetterFromGetter(g cache.Getter) BackupGetter {
	return func(p string) ([]byte, error) {
		x := g.Get(cache.Key(p))
		if x == nil {
			return nil, errors.New("couldn't read script")
		}
		return x.([]byte), nil
	}
}

// MakeGetterFromBackupGetter adapts a BackupGetter to a cache.Getter.
func MakeGetterFromBackupGetter(tg BackupGetter) cache.Getter {
	return cache.GetterFunc(func(k cache.Key) cache.Value {
		bs, err := tg(string(k))
		if err != nil {
			return nil
		}
		return bs
	})
}

// Backup ...
func Backup(path, host string) ([]byte, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	u, err := currentUser()
	if err != nil {
		return nil, err
	}
	dir, err := directory()
	if err != nil {
		return nil, err
	}
	bs = bytes.Replace(bs, []byte("$2"), []byte(u), -1)
	bs = bytes.Replace(bs, []byte("$1"), []byte(host), -1)
	bs = bytes.Replace(bs, []byte("$3"), []byte(dir), -1)
	return bs, nil
}

func currentUser() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.Username, nil
}

func directory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "archive"), nil
}
