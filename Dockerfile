FROM golang:1.10.1
WORKDIR /go/src/github.com/jdholdren/helloworld-go/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM scratch

# Document that the service listens on port 8080.
EXPOSE 8080

COPY --from=0 /go/src/github.com/jdholdren/helloworld-go/app .

ENTRYPOINT ["/app"]