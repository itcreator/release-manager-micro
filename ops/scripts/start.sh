#!/usr/bin/env bash
set -e

export COMPOSE_FILE=ops/docker/compose/env/prod.yml:ops/docker/docker-compose.yml:ops/docker/compose/service/project.yml:ops/docker/compose/service/semver.yml
docker-compose stop
docker-compose rm -vf consul service.semver service.project service.api
docker-compose up -d

ops/scripts/wait.sh

mkdir -p ./var/db/project/data
mkdir -p ./var/db/semver/data

if [ "$1" == "--no-deps" ]
then
    (docker-compose exec -T service.project go run -mod=vendor cmd/service/main.go) &
    (docker-compose exec -T service.semver go run -mod=vendor cmd/service/main.go) &
    (docker-compose exec -T service.api go run -mod=vendor cmd/release-manager-server/main.go --host=0.0.0.0 --port=80) &

    wait
else
    (docker-compose exec -T service.project bash -c "go mod tidy && go mod vendor && go run -mod=vendor cmd/service/main.go") &
    (docker-compose exec -T service.semver bash -c "go mod tidy && go mod vendor && go run -mod=vendor cmd/service/main.go") &
    (docker-compose exec -T service.api bash -c "go mod tidy && go mod vendor && go run -mod=vendor cmd/release-manager-server/main.go --host=0.0.0.0 --port=80") &

    wait
fi
