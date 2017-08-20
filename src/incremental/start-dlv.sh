#!/usr/bin/env bash

set -x

docker-compose exec service.api pkill -f debug
docker-compose exec service.api dlv debug cmd/service/main.go --headless --listen=:2345 --log --api-version=2 &
