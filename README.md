Release manager (development version)
=====================================
_**Will be use instead https://github.com/itcreator/release-manager**_

[![Build Status](https://travis-ci.org/itcreator/release-manager-micro.svg?branch=master)](https://travis-ci.org/itcreator/release-manager-micro)
[![License (3-Clause BSD)](https://img.shields.io/:license-BSD%203--Clause-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/itcreator/release-manager-micro)](https://goreportcard.com/report/github.com/itcreator/release-manager-micro)


```bash
    ##cd devops/docker/
    docker-compose up
    docker exec -it relmmicro_service.api_1 bash
    go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80
    curl -i http://127.0.0.1:9090/projects
    curl -X POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects -d '{"name":"MyProject", "description":"demo project"}'
    curl -i http://127.0.0.1:9090/projects/5
    curl -X PUT -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects/5 -d '{"name":"Project 5!", "description":"demo project 5"}'
```

shortcuts
```bash
    export PATH=./devops/shortcuts:$PATH
    go-project go run cmd/service/main.go
    go-api go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80
```


Add new agent: go to `devops/consul/scripts/boot.sh` and add line  consul join service.***
Add new agent manually:
```bash
    docker exec -it relmmicro_consul_1 consul join service.api
```
 
Check consul cluster
```bash
docker exec -it relmmicro_consul_1 consul members
```


### project service
```bash
    docker exec -it relmmicro_service.project_1 bash

    #server 
    go run main.go

```

### protobuf
```bash
docker exec -it relmmicro_service.project_1 bash -c "protoc --go_out=plugins=micro:. proto/**/*.proto"
# or
protoc --go_out=plugins=micro:. proto/**/*.proto
```


### swagger
```
go-api ./swagger generate server -f /apiDoc/api_doc.yml
```

### db
```
docker exec -it relmmicro_service.project.db_1 psql -U releasemanager -d release_manager
```

### tests
```
go test $(glide novendor) --cover 
```

### docker containers ip
```
docker inspect -f '{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)
```

###TODO:
+project - implement CRUD
+Consul health check
+go-swagger - return correct response and error messages
implement versioning
add `OpenID Connect`
map project proto to swagger model automatically
folder/repo structure
dashboard for micro services
add library https://github.com/grpc-ecosystem/grpc-gateway to API layer
 

#### proto 
    - extend (inheritance)
