# Filesync

Filesync is a utility written in Golang which helps you to keep the files on the client up to date with the files on the server. Only the changed parts of files on the server are downloaded. Therefore it's great to synchronize your huge, and frequently changing files.

Forked from github.com/elgs/filesync/gsync

## Requirements

Needs access to `glibc` to compile properly, and so `busybox:ubuntu-14.04` is the selected base image.

## Local Usage

Install dependencies:

```sh
make deps
```

Run locally with config files (modify paths before running):

```sh
# Server
go run gsync.go gsync/server.json

# Client
go run gsync.go gsync/client.json
```

Alternatively, set environment variables (examples in [/.env](/.env)):

```sh
# Server
export FILESYNC_MODE=server
export FILESYNC_PORT=6776
export FILESYNC_IP=0.0.0.0
export FILESYNC_PATH=/tmp/share

go run gsync.go
```

Build package (requires `glibc`, won't work on OSX):

```sh
make build
```

Build Docker image:

```sh
make dev-image
make image
```

Run Docker-Compose cluster:

```sh
docker-compose up -d
```