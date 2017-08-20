#!/usr/bin/env bash

set -x

docker-compose exec service.api pkill -f debug
docker-compose exec service.api dlv debug cmd/release-manager-server/main.go --headless --listen=:2345 --log --api-version=2 -- --host=0.0.0.0 --port=80 &
