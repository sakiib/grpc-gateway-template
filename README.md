### gRPC Gateway Template
#### gRPC to JSON proxy generator following the gRPC HTTP spec.

---
#### Project structure:
```bash
├── buf.gen.yaml
├── buf.yaml
├── gateway
│   └── gateway.go
├── gen
│   └── go
│       └── helloworld
│           ├── hello_world_grpc.pb.go
│           ├── hello_world.pb.go
│           └── hello_world.pb.gw.go
├── go.mod
├── go.sum
├── Makefile
├── proto
│   ├── google
│   │   └── api
│   │       ├── annotations.proto
│   │       └── http.proto
│   └── helloworld
│       └── hello_world.proto
├── README.md
├── server
│   └── server.go
└── service
    └── greeter_service.go

```

---
#### Usage:
```bash
# install the dependencies
$ make install
# generate pb.go, pb.gw.go, grpc.pb.go
$ make gen
# run the gRPC server & the gateway
# gRPC server running on port: 8080 & the gateway server running on port: 8000
$ make run
# format go code
$ make fmt
# remove the generated pb.go, pb.gw.go, grpc.pb.go
$ make clean
```

---
#### Use cURL to send HTTP requests:
```bash
# request
$ curl -X POST -H "Content-Type:application/json" -k http://localhost:8000/v1/hello -d '{"name": "sakib"}'
```
```bash
# response
{"message":"Hello sakib"}
```
