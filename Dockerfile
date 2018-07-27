FROM iron/go

WORKDIR /app

ADD bin/helloworld-go /app/

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the helloworld command by default when the container starts.
ENTRYPOINT ["./helloworld-go"]