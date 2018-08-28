#!/usr/bin/env bash

# backup recursively zips the working directory into an archive named with the
# working directory and current time and copies the archive into the passed
# remote host, user, and directory.

set -e

FILE=$(basename $(pwd))-$(date +%s).zip

zip -r $FILE .

ssh $2@$1 mkdir -p $3
scp $FILE $2@$1:$3
