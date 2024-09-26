#! /bin/bash


if [ $# -eq 0 ]; then
    echo "Error: Provide commit message"
    exit 1
fi

if [ $# -ne 1 ]; then
    echo "Error: Provide a single string. Use quotes"
    exit 1
fi
current_branch=$(git branch --show-current)

git add . && git commit -m "$1" && git push origin "$current_branch"

if [ $? -eq 0 ]; then
    echo && echo "Success: GitHub repository has been updated"
else
    echo && echo "Error: Github push failed!"
fi
