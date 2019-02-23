<div align="center">
    <img src="https://raw.githubusercontent.com/pedrolopesme/sentinel/master/sentinel.png?raw=true" />
    <h1> sentinel </h1>
</div>

A Golang utility for crawling historical stocks data.

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
docker-clean:    terminate Sentinel container and remove all data related to them
```
 