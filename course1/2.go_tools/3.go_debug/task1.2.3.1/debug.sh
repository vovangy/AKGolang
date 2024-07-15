#!/bin/bash
#chmod +rwx debug.sh

echo "Debug started..."
dlv debug main.go
echo "Debug ended."
