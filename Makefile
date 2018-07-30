.PHONY: install test build serve clean pack deploy ship

PWD=$(shell pwd)
SRC=${PWD}

# For tagging a particular release
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

D_NAME=helloworld-go

pack: iron-build
	docker build -t docker.io/jdholdren/helloworld-go:$(TAG) .

ship: pack
	docker push docker.io/jdholdren/helloworld-go:$(TAG)

run:
	docker run -d -p 8888:8080 --rm -e TARGET=foobar --name ${D_NAME} docker.io/jdholdren/helloworld-go:$(TAG)

stop:
	docker stop ${D_NAME}

iron-build:
	docker run --rm -v "${PWD}":"${SRC}" -w "${SRC}" iron/go:dev go build -o bin/helloworld-go