# `backup`

`backup` is a website which displays a table of stored backups.

## Install

`go get gopkg.in/jwowillo/backup`

# Run `backup`

`backup ?--port 80 ?--host 127.0.0.1`

Make sure host is the host the server will be accessed from on its network and
that backups will be deployed to.

## Run `script/backup.sh`

`backup 127.0.0.1 user '~'`

The script recursively zips the working directory into an archive named with the
working directory and current time and copies the archive into the passed remote
host, user, and directory.

## Run `run_tests`

`run_tests`

Runs the tests and shows the coverage.
