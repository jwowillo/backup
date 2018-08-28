package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jwowillo/cache/v2/standard"
)

// main starts the server.
func main() {
	if *host == "" {
		log.Fatal("must set host")
	}

	const archivePath = "archive"
	const staticPath = "static"
	for _, p := range []string{archivePath, staticPath} {
		http.Handle(fmt.Sprintf("/%s/", p), MakeStaticHandler(p))
	}

	c := standard.ChangedCache("cache")

	const backupPath = "script/backup.sh"
	http.Handle(
		fmt.Sprintf("/%s", backupPath),
		MakeBackupHandler(c, backupPath, *host))

	const templatePath = "tmpl/index.html"
	http.Handle("/", MakeHomeHandler(c, archivePath, templatePath))

	log.Printf("listening on :%d", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

// init parses flags.
func init() {
	flag.Parse()
}

// port to listen on.
var port = flag.Int("port", 8080, "port to listen on")

// host the server is running on.
var host = flag.String("host", "", "host the server is running on")
