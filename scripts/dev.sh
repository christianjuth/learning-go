#!/bin/sh
nodemon --signal SIGTERM -e go --quiet --exec "clear && go run $1 || true"