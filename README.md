Release manager (development version)
=====================================

[![Build Status](https://travis-ci.org/itcreator/release-manager-micro.svg?branch=master)](https://travis-ci.org/itcreator/release-manager-micro)
[![License (3-Clause BSD)](https://img.shields.io/:license-BSD%203--Clause-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/itcreator/release-manager-micro)](https://goreportcard.com/report/github.com/itcreator/release-manager-micro)

This application generate version for your software (based on semver and gitflow).
 
Input: `project id`, `major version`, `minor version`, `branch name`

Output: semantic version http://semver.org/ (look like as `v1.2.0-rc.7`)

```
NOTICE: something does not work? Just remove all containers with consul agents

docker-compose rm -vf consul service.version.incremental service.semver service.project service.api
```



# Quick start (how to run and use)

- Install docker and docker-compose

- Clone sources
```bash
    git clone git@github.com:itcreator/release-manager.git .
```

- Run release manager
```bash
    devops/scripts/start.sh #or start_semver.sh or start_incremental.sh
    
    #waiting for the output: 2016/11/03 23:59:24 Serving release manager at http://[::]:80
```

- Create new project
```
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1/projects -d '{"name":"MyProject", "description":"demo project"}'
```
Response: 
```
    HTTP/1.1 201 Created
    Content-Type: application/release-manager.v1+json
    X-Error-Code: 
    Date: Fri, 04 Nov 2016 00:08:17 GMT
    Content-Length: 58
    
    {"description":"demo project","id":34,"name":"MyProject"}
```

- Generate semantic version
```
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1/projects/1/version/semantic -d '{"major":1, "minor": 2, "branch": "release"}'
```

Response:
```
    HTTP/1.1 201 Created
    Content-Type: application/release-manager.v1+json
    Date: Fri, 04 Nov 2016 00:09:53 GMT
    Content-Length: 24
    
    {"version":"v1.2.0-rc"}
```


- Generate incremental version
```
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1/increamental_version/test
```

Response:
```
    HTTP/1.1 201 Created
    Content-Type: application/release-manager.v1+json
    Date: Fri, 09 Dec 2016 12:15:24 GMT
    Content-Length: 14
    
    {"version":1}
```
- Use version number `v1.2.0-rc` (for example as `git tag` or ad `docker image tag`


# Development information
[API Documentation apiDoc/api_doc.yml](apiDoc/api_doc.yml)

```bash
    ##cd devops/docker/
    docker-compose up
    docker exec -it relm_service.api_1 bash
    go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects/1/version/semantic -d '{"major":1, "minor": 3, "branch": "release"}'
    curl -i http://127.0.0.1:9090/projects
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects -d '{"name":"MyProject", "description":"demo project"}'
    curl -i http://127.0.0.1:9090/projects/5
    curl -iX PUT -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects/5 -d '{"name":"Project 5!", "description":"demo project 5"}'
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
    docker exec -it relm_consul_1 consul join service.api
```
 
Check consul cluster
```bash
docker exec -it relm_consul_1 consul members
```


### project service
```bash
    docker exec -it relm_service.project_1 bash

    #server 
    go run main.go

```

### protobuf
```bash
docker exec -it relm_service.project_1 bash -c "protoc --go_out=plugins=micro:. proto/**/*.proto"
# or
protoc --go_out=plugins=micro:. proto/**/*.proto
docker-compose run --rm protoc protoc --go_out=plugins=micro:src/semver proto/semver/*.proto
```


### swagger
```
go-api ./swagger generate server -f /apiDoc/api_doc.yml
```

### db
```
docker exec -it relm_service.project.db_1 psql -U releasemanager -d release_manager
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
+implement versioning
+use gorp
Implement build scripts
Configure environment for CI and production
Implement UI (Angular 2)
add `OpenID Connect`
Implement ACL
folder/repo structure
dashboard for micro services
 

#### proto 
    - extend (inheritance)
