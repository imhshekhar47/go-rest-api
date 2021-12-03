#!/usr/bin/env bash

ACTION=${1:-help}

case "${ACTION}" in
    run)
        cd src \
          && go run main.go \
          && cd -
        ;;

    clean)
        [[ -d "build" ]] && rm -rf "build"
        ;;
    build)
        mkdir -p build \
          && go build -o ./build/go-rest-api \
        ;;
    *)
        echo "Bad usage"
        exit 1
        ;;
esac

exit $?

