PROTO_DIR = api
OUT_DIR = gen

proto:
	protoc \
		-I $(PROTO_DIR) \
		--go_out=$(PROTO_DIR)/$(OUT_DIR) \
		--go-grpc_out=$(PROTO_DIR)/$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/goods/goods.proto \
		$(PROTO_DIR)/auth/auth.proto

swaga:
	swag init --dir cmd,internal/handler,api/gen/auth,api/gen/goods --output docs

clean:
	rm -rf $(OUT_DIR)/*