#!/bin/bash
#chmod +rwx build.sh

echo "Compiling started..."
go build main.go
echo "Compiling complete."
echo "Trying to launch program"
./main
echo "Program exited"
