<div align="center">
    <img src="https://raw.githubusercontent.com/pedrolopesme/sentinel/master/sentinel.png?raw=true" />
    <h1> sentinel </h1>
</div>

A Golang utility for crawling historical stocks data.

<p align="center">
  <a href="https://travis-ci.org/pedrolopesme/sentinel"> <img src="https://api.travis-ci.org/pedrolopesme/sentinel.svg?branch=master" /></a>
  <a href="https://goreportcard.com/report/github.com/pedrolopesme/sentinel"> <img src="https://goreportcard.com/badge/github.com/pedrolopesme/sentinel" /></a>
  <a href="https://codeclimate.com/github/pedrolopesme/sentinel/maintainability"> <img src="https://api.codeclimate.com/v1/badges/b7cee6500978112b2910/maintainability" /></a>
  <a href="https://sonarcloud.io/dashboard?id=pedrolopesme_sentinel"> <img src="https://sonarcloud.io/api/project_badges/measure?project=pedrolopesme_sentinel&metric=alert_status" /></a>
  <a href="https://sonarcloud.io/dashboard?id=pedrolopesme_sentinel"> <img src="https://sonarcloud.io/api/project_badges/measure?project=pedrolopesme_sentinel&metric=coverage" /></a>
</p>
<br>

#### Executing
 
 This project comes with a useful [Makefile](Makefile) containing all needed targets to build, run docker containers, 
 run tests, etc. To list all available targets, just type:
 
 ```
 $ make help
 ```   
  
 So far, the available commands are:
 
 ```
build:           compiles Sentinel binary
run:             run main func
test:            run unit tests
clean:           clean all Sentinel binaries
fmt:             run go fmt on all go files
docker-build:    build Sentinel docker image
docker-run:      build Sentinel docker image and execute docker compose up
docker-stop:     execute a docker compose down
docker-logs:     make a tail -f on Sentinel running containers
docker-shell:    login on Sentinel running container
docker-prune:    terminate Sentinel container and prune volume and system
```

#### TO DO

Roadmap to V0.1

- [ ] use a message broker (NATS ?) as the default source of scheduling
- [ ] implement all TODOs  
- [ ] create kubernetes deployment yml
- [x] use a message broker (NATS ?) to send all stock quotations collected on each run 
