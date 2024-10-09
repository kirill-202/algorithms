#! /bin/bash

if [ $# -eq 0 ]; then
    echo "Error: No directory name provided"
    exit 1
fi

if [ $# -ne 1 ]; then
    echo "Error: Provide a single string. Use - for word separation"
    exit 1
fi

current_date=$(date +"%Y-%m-%d" |  awk "{print $2}")
directory_name="$1-${current_date}"
package_name="${directory_name//-/_}"

mkdir "${directory_name}" && cd "${directory_name}" && touch main.go && go mod init "${package_name}"