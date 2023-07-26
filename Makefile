PHP_BUILD_PLUGIN=`which grpc_php_plugin`
GO_BUILD_PLUGIN=`which protoc-gen-go-grpc`
JAVA_BUILD_PLUGIN=`which protoc-gen-grpc-java`

PHP_PROTO_OUT="src/main/php"

GO_PROTO_BASEDIR="./"
GO_PROTO_OUT="xyron"

JAVA_PROTO_OUT="src/main/java"

php:
	@mkdir -p $(PHP_PROTO_OUT)
	@rm -rdf $(PHP_PROTO_OUT)
	@mkdir -p $(PHP_PROTO_OUT)
	@cd src/main/proto && protoc --plugin=protoc-gen-grpc=$(PHP_BUILD_PLUGIN) ./*.proto --php_out="../../../$(PHP_PROTO_OUT)" --grpc_out="../../../$(PHP_PROTO_OUT)"

go:
	@rm -rdf $(GO_PROTO_OUT)
	@mkdir -p $(GO_PROTO_OUT)
	@cd src/main/proto && protoc --plugin=protoc-gen-grpc=$(GO_BUILD_PLUGIN) ./*.proto --go_out="../../../$(GO_PROTO_BASEDIR)" --grpc_out="../../../$(GO_PROTO_BASEDIR)"

java:
	@mkdir -p $(JAVA_PROTO_OUT)
	@rm -rdf $(JAVA_PROTO_OUT)
	@mkdir -p $(JAVA_PROTO_OUT)
	@cd src/main/proto && protoc --plugin=protoc-gen-grpc=$(JAVA_BUILD_PLUGIN) ./*.proto --java_out="../../../$(JAVA_PROTO_OUT)" --grpc_out="../../../$(JAVA_PROTO_OUT)"

all:
	@make php
	@make go
	@make java