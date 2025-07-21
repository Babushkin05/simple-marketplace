PROTO_DIR = api
OUT_DIR = $(PROTO_DIR)/gen

proto:
	@mkdir -p $(OUT_DIR)
	protoc \
		-I $(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/goods.proto

clean:
	rm -rf $(OUT_DIR)/*