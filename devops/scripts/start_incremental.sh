#!/usr/bin/env bash

export COMPOSE_FILE=devops/docker/compose/env/prod.yml:devops/docker/docker-compose.yml:devops/docker/compose/service/incremental.yml
export CONSUL_AGENTS=service.api:service.version.incremental
docker-compose stop
docker-compose rm -vf consul service.version.incremental service.api
docker-compose up -d

i="0"
while [ $i -lt 25 ] #waiting for cluster 25 seconds
do
    i=$[$i+1]
    sleep 1
    echo -e ".\c"
done

echo ""

(docker exec -i relm_service.version.incremental_1 bash -c "glide install && go run cmd/service/main.go") &
(docker exec -i relm_service.api_1 bash -c "glide install && go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80") &


wait

