# Distributed scheduler client

Welcome 👋

This repository contains client library/SDK for
the [distributed scheduler](https://github.com/GLCharge/distributed-scheduler).

Every version has OpenAPI specification that is used to generate the client using `opai-codegen` generator.
OpenApi specifications contain also documentation. For more detailed documentation about the modules and APIs please
follow
the official [OICP documentation](https://hubject.github.io/oicp-cpo-2.3-api-doc/)

## Installation

```go
import client "https://github.com/GLCharge/distributed-scheduler-client"
```

## Generation

This code base, was generated using [oapi-codegen](https://github.com/deepmap/oapi-codegen).

You can generate the code by running the Makefile script:

```bash
make gen-client
```