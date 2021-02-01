.PHONY: protobuf
## protobuf: generate protobuff code from definition
protobuf:
	protoc -I=loot_pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=loot_pb --go_out=loot_pb loot_pb/loot.proto

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
