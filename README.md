# covid-19-api-server

This repository holds the backend implementation for [covid-19-au.github.io](https://github.com/covid-19-au/covid-19-au.github.io). The project is written in [Go](https://golang.org) and primarily uses [gRPC](https://grpc.io/about) for service-to-service communication.

Docker image will be available in future releases.

## Contributing

#### Build Instructions

* Follow [gRPC Quickstart](https://grpc.io/docs/quickstart/go) to install dependencies for gRPC and Protobuf.

* Update Protobuf specifications will require re-compilation. Run `protoc -I proto/ proto/*.proto --go_out=plugins=grpc:protogen` and generated artifacts will be available under `protogen` directory.

* Use `go build` to build the binary. Follow [this document](https://golang.org/cmd/go/#hdr-Build_modes) to learn more about additional build options.

## Teams

#### Engineering
* Robin Liu: opensource@greenvine.dev
* Han Wang: freddie.wanah@gmail.com

#### Infrastructure
* Robin Liu: opensource@greenvine.dev

#### PM
* Chunyang Chen: Chunyang.Chen@monash.edu

## Contact

You can contact us by:

* Email opensource@greenvine.dev or create [GitHub Issues](https://github.com/covid-19-au/covid-19-api-server/issues/new) if there are any the bugs or feature requests.
* Email Chunyang.Chen@monash.edu if you want to join the team for the development.
