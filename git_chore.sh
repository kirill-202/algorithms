#! /bin/bash



if [ $# -eq 0 ]; then
    echo "Error: Provide commit message"
    exit 1
fi

current_branch=$(git branch --show-current)

git add . && git commit -m "$1" && git push origin "$current_branch" && echo success