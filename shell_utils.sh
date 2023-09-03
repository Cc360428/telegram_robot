#!/bin/bash

function docker_build() {
  go mod tidy
  go mod vendor
  docker-compose down
  docker rmi telegramcc:tg_v1
  docker build -t telegramcc:tg_v1 .
}

if [ $# == 0 ]; then
  docker_build
else
  case $1 in
  1)
    generateProtobuf
    ;;

  esac
fi
exit 0
