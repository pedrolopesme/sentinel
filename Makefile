GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/sentinel
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	@echo "=============Building Sentinel============="
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	@echo "=============Running Sentinel============="
	go run main.go

test:
	@echo "=============Running Sentinel Tests============="
	$(GOTEST) -v ./...

clean: 
	@echo "=============Removing Sentinel============="
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

fmt:
	@echo "=============Running gofmt for all project files============="
	$(GOFMT) -w */*.go

docker-build:
	@echo "=============Building Local Sentinel Docker Image============="
	docker build -f ./Dockerfile -t sentinel .

docker-run: docker-build
	@echo "=============Starting Sentinel Container============="
	docker-compose up -d

docker-stop:
	@echo "=============Stopping Sentinel Container============="
	docker-compose down

docker-logs:
	@echo "=============Getting Sentinel Docker Logs============="
	docker-compose logs -f

docker-shell:
	@echo "=============Accessing Container Shell============="
	docker exec -t sentinel bash

docker-clean: docker-down
	@echo "=============Cleaning up============="
	rm -f sentinel
	docker system prune -f
	docker volume prune -f