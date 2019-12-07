Release manager (development version)
=====================================

[![Build Status](https://travis-ci.org/itcreator/release-manager-micro.svg?branch=master)](https://travis-ci.org/itcreator/release-manager-micro)
[![License (3-Clause BSD)](https://img.shields.io/:license-BSD%203--Clause-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/itcreator/release-manager-micro)](https://goreportcard.com/report/github.com/itcreator/release-manager-micro)
[![BCH compliance](https://bettercodehub.com/edge/badge/itcreator/release-manager-micro?branch=master)](https://bettercodehub.com/)

This application generate version for your software (based on semver and gitflow).
 
Input: `project id`, `major version`, `minor version`, `branch name`

Output: semantic version http://semver.org/ (look like as `v1.2.0-rc.7`)

```
NOTICE: something does not work? Just remove all containers with consul agents

docker-compose rm -vf consul service.semver service.project service.api
```


[Changelog](changelog.md)

# Quick start (how to run and use)

- Install docker and docker-compose

- Clone sources
```bash
    git clone https://github.com/itcreator/release-manager-micro.git .
```

- Run release manager
```bash
    ops/scripts/start.sh
    
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
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1/projects/{uuid}/version/semantic -d '{"major":1, "minor": 2, "branch": "release"}'
```

Response:
```
    HTTP/1.1 201 Created
    Content-Type: application/release-manager.v1+json
    Date: Fri, 04 Nov 2016 00:09:53 GMT
    Content-Length: 24
    
    {"version":"v1.2.0-rc"}
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
    docker-compose up
    docker-compose exec service.api bash
    go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects/{uuid}/version/semantic -d '{"major":1, "minor": 3, "branch": "release"}'
    curl -i http://127.0.0.1:9090/projects
    curl -iX POST -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects -d '{"name":"MyProject", "description":"demo project"}'
    curl -i http://127.0.0.1:9090/projects/{uuid}
    curl -iX PUT -H "Content-Type: application/release-manager.v1+json" http://127.0.0.1:9090/projects/{uuid} -d '{"name":"Project 5!", "description":"demo project 5"}'
```

Shortcuts
```bash
    export PATH=./ops/shortcuts:$PATH
    go-project go run cmd/service/main.go
    go-api go run cmd/release-manager-server/main.go --host=0.0.0.0 --port=80
```


Add new agent: go to `ops/consul/scripts/boot.sh` and add line  consul join service.***
Add new agent manually:
```bash
    docker-compose exec consul consul join service.api
```
 
Check consul cluster
```bash
    docker-compose exec consul consul members
```


### Project service
```bash
    docker-compose exec service.project bash

    #server 
    go run main.go

```

### Protobuf
```bash
    protoc --go_out=plugins=micro:. proto/**/*.proto
    docker-compose run --rm protoc protoc --go_out=plugins=micro:./src/project proto/project/*.proto
```


### Swagger

#### Generate server
```bash
    docker-compose run --rm go_swagger generate server -f /apiDoc/api_doc.yml
```

#### Generate client
```bash
    docker-compose run --rm go_swagger generate client -f /apiDoc/api_doc.yml
```

### Serve API documentation
```bash
    docker-compose run --rm go_swagger serve --no-open --port=8070 /apiDoc/api_doc.yml
```

And open [http://127.0.0.1:8070/docs]

Or

```bash
docker-compose run --rm go_swagger serve --no-open --port=8070 --flavor=swagger /apiDoc/api_doc.yml
```

And open [http://petstore.swagger.io/?url=http%3A%2F%2Flocalhost%3A8070%2Fswagger.json]

### DB
```
    docker-compose exec service.project.db psql -U releasemanager -d release_manager
```

### Tests
```
    go test  -mod=vendor handler/* --cover
```

### Docker containers ip
```
    docker inspect -f '{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)
```

### TODO

Implement build scripts

Configure environment for CI and production

Implement UI (Angular 2)

Dashboard for micro services

Add `OpenID Connect`

Implement ACL

folder/repo structure

 

#### Proto
    - extend (inheritance)
