VERSION=latest
.PHONY: check_all
check_all: 
	@(./check_all)
build:
	@(GOOS=linux go build -o checker checker.go)
docker-build: build
	@(docker build -t eqemuchecker .)
docker-check_all: docker-build
	@(docker run -v $(PWD):/home -i eqemuchecker)
docker-push: docker-build
	@(docker tag eqemuchecker eqemu/checker:$(VERSION))
	@(docker push eqemu/checker:$(VERSION))