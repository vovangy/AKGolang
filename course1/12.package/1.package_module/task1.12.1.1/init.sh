#!/bin/bash
#chmod +rwx init.sh

if [ -z "$1" ]; then
 echo "Module name argument is missing"
 exit 1
fi

go mod init $1
