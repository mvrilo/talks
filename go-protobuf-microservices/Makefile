present: generated
	@present

deps:
	@go get -x golang.org/x/tools/cmd/present

generated: clean
	@mkdir generated
	@protoc -I ./proto -I ./third_party/googleapis --go_out=./generated proto/message.proto --go-grpc_out=./generated proto/message.proto
	@protoc -I ./proto -I ./third_party/googleapis --grpc-gateway_out=./generated  proto/message.proto

clean:
	@rm -rf generated
