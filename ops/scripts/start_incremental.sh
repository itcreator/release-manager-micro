#!/usr/bin/env bash

export COMPOSE_FILE=ops/docker/compose/env/prod.yml:ops/docker/docker-compose.yml:ops/docker/compose/service/incremental.yml
export CONSUL_AGENTS=service.api:service.version.incremental
docker-compose stop
docker-compose rm -vf consul service.version.incremental service.api
docker-compose up -d

ops/scripts/wait.sh

if [ "$1" == "--no-glide" ]
then
    (docker-compose exec -T service.version.incremental go run cmd/service/main.go) &
    (docker-compose exec -T service.api go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80) &

    wait
else
    (docker-compose exec -T service.version.incremental  bash -c "glide install && go run cmd/service/main.go") &
    (docker-compose exec -T service.api bash -c "glide install && go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80") &

    wait
fi
