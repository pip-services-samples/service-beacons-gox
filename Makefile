.PHONY: all build clean install uninstall fmt simplify check run test protogen docgen benchmark

install:
	@go install ./bin/run.go

run: install
	@go run ./bin/run.go

test:
	@go clean -testcache && go test -v ./test/...

protogen: env go.sum
	protoc --go_out=plugins=grpc:. protos/beacons_v1.proto

docgen: env go.sum
	gold -gen -nouses -dir=docs -emphasize-wdpkgs ./...

benchmark:
	@go run ./benchmark/main.go