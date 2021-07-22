### gRPC Gateway Template
#### gRPC to JSON proxy generator following the gRPC HTTP spec.

---
#### Project structure:
```bash
├── buf.gen.yaml
├── buf.yaml
├── Dockerfile
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
├── grpc
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
│   └── main.go
├── service
│   └── greeter_service.go
└── vendor

```

---
#### Clone the repo & run locally:
```bash
# install the dependencies
$ make install
# generate pb.go, pb.gw.go, grpc.pb.go
$ make gen
# run the gRPC server & the gateway
# format go code
$ make fmt
# gRPC server running on port: 8080 & the gateway server running on port: 8000
$ make run
# remove the generated pb.go, pb.gw.go, grpc.pb.go
$ make clean
```

---
#### or Use the docker image:
```bash
$ docker pull sakibalamin/grpc-gateway:v0.0.1
$ docker run -p 8000:8000 -it sakibalamin/grpc-gateway:v0.0.1
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
