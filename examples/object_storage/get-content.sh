#!/usr/bin/env bash

# Example script to download an object using local-exec
echo "Downloading object $1 content and writing to $2..."
mkdir content
cd content
curl -o $2 $1
cd ..
echo "Finished Downloading object"