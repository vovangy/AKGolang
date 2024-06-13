#!/bin/bash
#chmod +rwx debug.sh

echo "Debug started..."
echo $1
go build -o myprogram $1
dlv exec myprogram
echo "Debug ended."

