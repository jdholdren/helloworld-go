.PHONY: install test build serve clean pack deploy ship

PWD=$(shell pwd)
SRC=${PWD}

D_NAME=helloworld-go

pack: iron-build
	docker build -t docker.io/jdholdren/helloworld-go .

ship: pack
	docker push docker.io/jdholdren/helloworld-go

run:
	docker run -d -p 8888:8080 --rm -e TARGET=foobar --name ${D_NAME} docker.io/jdholdren/helloworld-go

stop:
	docker stop ${D_NAME}

iron-build:
	docker run --rm -v "${PWD}":"${SRC}" -w "${SRC}" iron/go:dev go build -o bin/helloworld-go