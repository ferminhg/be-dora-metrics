# BE DORA Metrics

Endpoints for DORA metrics

The server is written in Go and uses the Gin framework.

Stack
- GoLang: https://golang.org/
- Gin Framework: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme

## Description

This is a simple API that returns the DORA metrics for a given repository.

DoraMetrics follow a clean architecture (port/adapter architecture), with the following layers:

**Domain**
- `internal/domain`: Contains the business logic of the application. It is the core of the application and does not depend on any other layer.

**Application**
- `cmd/api/main.go`: Contains the entrypoint of the application. `main.go` does not contain any business logic, it just wires the application, and doesn't depend on any framework, like Gin. It allows to change it if we find something better.
- `internal/platform/server/[health, metrics]`: Contains the uses cases of the endpoints. It depends on the domain layer, and on the framework used (Gin).

**Infrastructure**
- `internal/platform/server/server.go`: Contains the Gin server. Where the routes are defined, and the Gin server is started.
- `internal/platform/storage/`: Contains the storage layer. The implementation of the repository interfaces defined in the domain layer.
- `internal/shared`: Contains shared code between the different layers, like uuidGenerator or logger.

## Usage

### 👨‍💻 Running the dev server locally 

You need to have [Go installed in your machine](https://go.dev/doc/install).
If you don't have it, you could run it using docker

```bash
make run
# or
go run cmd/api/main.go
```

### 🧪 Running the tests

```bash
make test
# or
go test ./...
```


### 📈 Building the binary

```bash
make build
```

The binary will be available at `build/metrics`.
You could also run `go build -o build/metrics cmd/api/main.go`

If you want to run the binary, you can do it with `./build/metrics`


### 🐳 Running the dev server locally with Docker

_🚧 ?? In progress ?? 🚧_

```bash
make run-docker
```

### Endpoints
#### GET /metrics

Returns the DORA metrics on JSON format. Example:

```json
[
  {
    "id": "577942c9-4189-447c-aa31-cfa60dc3be1a",
    "name": "tc",
    "value": 13.006846250009444,
    "timestamp": "2023-05-21T22:01:07.717646+02:00"
  },
  {
    "id": "e08084b3-6b91-417d-be78-50e492b3406c",
    "name": "tc",
    "value": 17.47194416049814,
    "timestamp": "2023-05-21T22:01:07.717681+02:00"
  },
  {
    "id": "1873eb69-8f1b-448e-8d88-03816d9cd185",
    "name": "tc",
    "value": 12.52380143950969,
    "timestamp": "2023-05-21T22:01:07.717683+02:00"
  },
  {
    "id": "aa6e912e-1586-4456-8c79-7f7f9ad1f474",
    "name": "tc",
    "value": 10.843332067669111,
    "timestamp": "2023-05-21T22:01:07.717685+02:00"
  }
]
```

    