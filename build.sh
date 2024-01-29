#!/usr/bin/env bash

docker build -t caller-service-test .
docker run -it -d \
    -e COUNTER_SERVICE_URL=http://host.docker.internal:8081 \
    -e LISTEN_PORT=8080 \
    -p 8080:8080 \
    caller-service-test