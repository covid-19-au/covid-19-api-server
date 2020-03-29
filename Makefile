.PHONY: dep dep-build

all: lib-go lib-js

lib-go: clean-proto dep dep-build mkdir-build
	protoc -I proto \
		-I third_party/grpc-gateway/third_party/googleapis \
		--proto_path=$(GOPATH)/src \
		--proto_path=$(GOPATH)/src/github.com/golang/protobuf \
		--proto_path=proto/ \
		--go_out=plugins=grpc:protogen/go \
		--govalidators_out=protogen/go \
		proto/*.proto

lib-js: clean-proto dep dep-build mkdir-build
	protoc -I proto \
		--proto_path=proto/ \
		--js_out=import_style=commonjs+dts,binary:protogen/cjs \
		--grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:protogen/cjs \
		proto/*.proto

	protoc -I proto \
		--proto_path=proto/ \
		--js_out=import_style=closure:protogen/closure \
		--grpc-web_out=import_style=closure,mode=grpcwebtext:protogen/closure \
		proto/*.proto

	protoc -I proto \
		--proto_path=proto/ \
		--js_out=import_style=typescript:protogen/ts \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:protogen/ts \
		proto/*.proto

dep: dep-vendor
	go mod download

dep-build:
	GO111MODULE=off go get -u github.com/golang/protobuf/proto
	GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go
	GO111MODULE=off go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

dep-vendor:
	git submodule update --init --recursive

mkdir-build:
	mkdir -p protogen/{cjs,closure,go,ts}

clean: clean-build clean-proto

clean-build:
	rm -rf build

clean-proto:
	rm -rf protogen
