#! /bin/bash

if [ $# -eq 0 ]; then
    echo "Error: No directory name provided"
    exit 1
fi

if [ $# -ne 1 ]; then
    echo "Error: Provide a single string. Use - for word separation"
    exit 1
fi

directory_name="$1"
package_name="${directory_name//-/_}"

mkdir "$directory_name" && cd "$directory_name" && touch main.go && go mod init "$package_name"