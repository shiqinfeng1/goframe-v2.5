# syntax = docker/dockerfile:experimental
# NOTE: Must be run in the context of the repo's root directory

####################################
## (1) Setup the build environment
# FROM golang:1.19-bullseye AS build-setup
# 使用代理地址， 下载速度更快
FROM dockerproxy.com/library/golang:1.19-bullseye AS build-setup

RUN apt-get update
# RUN apt-get -y install cmake zip

## (2) Setup crypto dependencies
FROM build-setup AS build-env

# Build the app binary in /app
RUN mkdir /app
WORKDIR /app

ARG TARGET
ARG COMMIT
ARG VERSION
ARG TAG 
ARG BRANCH
ARG TREE_STATE
ARG BUILD_TIME
ENV GOPRIVATE=

COPY . .

####################################
## (3) Build the production app binary
FROM build-env as build-production
WORKDIR /app

ARG GOARCH=amd64

# TAGS can be overriden to modify the go build tags (e.g. build without netgo)
ARG TAGS=""

RUN go env -w GOPROXY=https://goproxy.cn

# Keep Go's build cache between builds.
# https://github.com/golang/go/issues/27719#issuecomment-514747274
RUN --mount=type=cache,sharing=locked,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=1 GOOS=linux \
    go build --tags "${TAGS}" -ldflags "-extldflags -static" \
    -o ./app ${TARGET}

RUN chmod a+x /app/app

## (4) Add the statically linked production binary to a distroless image
# FROM gcr.io/distroless/base-debian11 as production
# 使用代理地址， 下载速度更快
# FROM gcr.dockerproxy.com/distroless/base-debian11:latest as production
FROM gcr.dockerproxy.com/distroless/base-debian11:debug as production

COPY --from=build-production /app/app /bin/app

ENTRYPOINT ["/bin/app"]

