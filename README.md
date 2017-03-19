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

## Demo

The [docker-compose.yml](/docker-compose.yml) details a setup with a server and 2 clients. It will create a `./demo/` directory with 3 subfolders: `server`, `client1`, and `client2`. Each subfolder is volumed as `/share/` inside each container.

As files in `./demo/server` are modified, they will be altered in `./demo/client1` and `./demo/client2`.

Run Docker-Compose cluster:

```sh
docker-compose up -d
```

Check directories:

```sh
ls -al ./demo/server
ls -al ./demo/client1
ls -al ./demo/client2
```

Add a file to server volume:

```sh
echo "testing 123" > ./demo/server/test1
```

Check directories again:

```sh
ls -al ./demo/server
ls -al ./demo/client1
ls -al ./demo/client2
```

At this point, the `./demo/client*` directories should contain a `test1` file.

## Notes

### Docker Environment

While the server configuration can be set to an IP of `0.0.0.0` (accepts traffic from anywhere), the clients need a specific address to connect to. If running locally, the clients can be set to connect to `127.0.0.1` - but this will not work in a Dockerized environment.

The Docker-Compose.yml file contains `links: [fs-server:fs-server]` which enables the clients to contact the server container at `http://fs-server`. This is supported in a Rancher environment as well.

### Rancher Environment

- todo