PROJECT?=github.com/barrenao/tenerife
RELEASE?=0.0.1

COMMIT:=git-$(shell git rev-parse --short HEAD)
BUILD_TIME:=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	GO111MODULE=on CGO_ENABLED=0 go build \
  	-ldflags "-s -w -X ${PROJECT}/internal/diagostics.Version=${RELEASE} -X ${PROJECT}/internal/diagostics.Commit=${COMMIT} -X ${PROJECT}/internal/diagostics.BuildTime=${BUILD_TIME}"  -o bin/tenerife ${PROJECT}/cmd/tenerife

test:
	go test ./...


docker-build:
	docker build -t tenerife:$(RELEASE) .
