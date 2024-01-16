# Distributed scheduler client

Welcome 👋

This repository contains client library/SDK for
the [distributed scheduler](https://github.com/GLCharge/distributed-scheduler).

Every version has OpenAPI specification that is used to generate the client using `oapi-codegen` generator.

## Installation

```bash
 go get github.com/GLCharge/distributed-scheduler-client@latest"
```

## Usage

```go
import client "https://github.com/GLCharge/distributed-scheduler-client
```

## Generation

This code base, was generated using [oapi-codegen](https://github.com/deepmap/oapi-codegen).

You can generate the SDK manually by running the Makefile script:

```bash
make gen-client
```
