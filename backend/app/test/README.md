# How to use test API server

## Contents

1. [Install Docker Image](#install-docker-image)
1. [Run API Server in Container](#run-api-server-in-container)
1. [Stop API Server](#stop-api-server)

## Install Docker Image

1. Move the directory where this file located.
```bash
cd /<path-your-environment>/splitter/backend/app/test
```

2. Make Docker image from `./Dockerfile`
```bash
docker build . -t splitter-mock-server:1.0
```

### Uninstall This Docker Image (Appendix)

When you delete this image, run under the command.

```bash
docker image rm splitter-mock-server:1.0
```

## Run API server in Container

Run the API server in Container using Prism.

Prism is a software running server on Node environment.

This command run the server for your api test of Splitter.

```bash
docker container run --rm -it --name splitter-mock-server -p 4010:4010 splitter-mock-server:1.0 prism mock -h 0.0.0.0 splitter-api.yaml
```

You can access API server in localhost.

```bash
curl localhost:4010/<path>
```

## Stop API server

After open another terminal, run under command.

```bash
docker container stop splitter-mock-server
```