#!/bin/bash

if [[ $# -lt 2 ]]; then
    echo "usage: $0 提交git的comment信息  v3.0.1(tag的名字)"
    exit 1
fi

git add -A
git commit -a -m "$1" || exit 1
git pull || exit 1
git push || exit
git tag "$2" || exit 1
git push origin "$2"
