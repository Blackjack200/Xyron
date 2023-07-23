PHP_BUILD_PLUGIN=`which grpc_php_plugin`
GO_BUILD_PLUGIN=`which protoc-gen-go-grpc`

PHP_PROTO_OUT="php_src"
GO_PROTO_BASEDIR="./"
GO_PROTO_OUT="xyron"

proto_php:
	@mkdir -p $(PHP_PROTO_OUT)
	@rm -rdf $(PHP_PROTO_OUT)
	@mkdir -p $(PHP_PROTO_OUT)
	@cd xchange && protoc --plugin=protoc-gen-grpc=$(PHP_BUILD_PLUGIN) ./*.proto --php_out="../$(PHP_PROTO_OUT)" --grpc_out="../$(PHP_PROTO_OUT)"

proto_go:
	@rm -rdf $(GO_PROTO_OUT)
	@mkdir -p $(GO_PROTO_OUT)
	@cd xchange && protoc --plugin=protoc-gen-grpc=$(GO_BUILD_PLUGIN) ./*.proto --go_out="../$(GO_PROTO_BASEDIR)" --grpc_out="../$(GO_PROTO_BASEDIR)"

proto:
	@make proto_php
	@make proto_go