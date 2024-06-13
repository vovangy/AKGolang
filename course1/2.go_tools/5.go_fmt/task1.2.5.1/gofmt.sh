#!/bin/bash
#chmod +rwx gofmt.sh

if [ -z "$1" ]; then
 echo "Usage: gofmt.sh <filename>"
 exit 1
fi

go fmt $1
