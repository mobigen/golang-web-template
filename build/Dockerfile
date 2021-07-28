# Copyright (c) 2021 jblim Mobigen, Inc. All Rights Reserved.

# Build the docker
ARG DOCKER_BUILD_IMAGE=golang:1.16
ARG DOCKER_BASE_IMAGE=alpine:3.12

FROM ${DOCKER_BUILD_IMAGE} AS build
WORKDIR /app/
COPY . /app/
RUN apt-get update -yq && apt-get install -yq unzip
RUN make build

# Final Image
FROM ${DOCKER_BASE_IMAGE}

LABEL name="web-server" \
  maintainer="irisdev@mobigen.com" \
  vendor="Mobigen" \
  architecture="x86_64" \
  url="https://mobigen.com" \
  description="Web Server Sample. Created Platform Team Mobigen" 

#ENV SERVER=/app/bin/server \
ENV USER_UID=1000 \
    USER_NAME=server

RUN  apk update && apk add libc6-compat && apk add ca-certificates
COPY --from=build /app/build/bin /usr/local/bin
COPY --from=build /app/build/_output/ /app/bin
RUN  /usr/local/bin/user_setup
WORKDIR /app

USER ${USER_UID}

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/entrypoint"]
