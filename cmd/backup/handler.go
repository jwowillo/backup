package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/jwowillo/cache/v2"
)

// MakeHomeHandler makes the home-page http.Handler.
func MakeHomeHandler(
	c cache.Cache,
	archivePath, templatePath string) http.Handler {
	return http.HandlerFunc(HomeHandler(
		MakeArchivesGetter(c), MakeTemplateGetter(c),
		archivePath, templatePath))
}

// HomeHandler returns the main http.HandlerFunc after injecting all
// dependencies.
func HomeHandler(
	archivesGetter ArchivesGetter,
	templateGetter TemplateGetter,
	archivePath, templatePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		as, err := archivesGetter(archivePath)
		if err != nil {
			log.Println(err)
			return
		}
		tmpl, err := templateGetter(templatePath)
		if err != nil {
			log.Println(err)
			return
		}
		sort.Slice(
			as,
			func(i, j int) bool {
				at := as[i].ModifiedTime
				bt := as[j].ModifiedTime
				return at.After(bt)
			})
		if err := tmpl.Execute(w, as); err != nil {
			log.Println(err)
		}
	}
}

// MakeBackupHandler makes a backup-script http.Handler.
func MakeBackupHandler(c cache.Cache, path, host string) http.Handler {
	return http.HandlerFunc(BackupHandler(
		MakeBackupGetter(c, host), path))
}

// BackupHandler returns the backup-script http.HandlerFunc after injecting all
// dependencies.
func BackupHandler(
	backupGetter BackupGetter, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := backupGetter(path)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(bs)
	}
}

// MakeStaticHandler makes a static-file http.Handler.
func MakeStaticHandler(staticPath string) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf("/%s/", staticPath),
		http.FileServer(http.Dir(staticPath)))
}
