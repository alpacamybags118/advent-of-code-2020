#!/usr/bin/env bash

if [ -z $1 ]; then
  echo "You must provide a day to run."
  exit 1
fi

go run $1/main.go