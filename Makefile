PROTOS=$(shell find proto -iname "*.proto")
.PHONY: protos
protos:
	protoc --proto_path=proto\
		--go_out=. --go_opt=module=github.com/gotbitoriginal/commontemp\
		--go-grpc_out=. --go-grpc_opt=module=github.com/gotbitoriginal/commontemp\
		$(PROTOS)