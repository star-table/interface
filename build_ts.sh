#!/bin/bash

CURRENT_PATH=$(cd `dirname $0` && pwd)
PROTOC_GEN_TS_PATH=$(which protoc-gen-ts) # https://github.com/improbable-eng/ts-protoc-gen

if [[ $# -lt 1 ]]; then
    echo "usage: $0 helloworld/v1"
    exit 1
fi

# ts
if [ -d "${CURRENT_PATH}/proto/$1" ];then
  rm -rf "${CURRENT_PATH}/ts/$1"
  mkdir -p "${CURRENT_PATH}/ts/$1"
  cp -rf ${CURRENT_PATH}/proto/$1/*.proto "${CURRENT_PATH}/ts/$1/"
  cd "${CURRENT_PATH}/ts/$1" || exit 1

  protoc --proto_path=. \
    --proto_path=${CURRENT_PATH}/proto \
    --proto_path=${CURRENT_PATH}/proto/third_party \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:." --ts_out=. *.proto

  rm -rf *.proto
else
  echo "文件不存在"
  exit 1
fi

