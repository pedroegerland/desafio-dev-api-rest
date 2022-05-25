FROM golang:1.18-alpine3.15 AS build
WORKDIR /app

COPY . /app

RUN go mod download
RUN go build -o build ./src/main.go

##
## Deploy
##
FROM golang:1.18-alpine3.15 as app

ARG COMMIT_HASH
ENV COMMIT_HASH=$COMMIT_HASH
ENV DD_VERSION=$COMMIT_HASH

WORKDIR /app

COPY --from=build /app/build ./build