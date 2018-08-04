gen_proto:
	protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

gen_server:
	go run greet/greet_server/server.go

gen_client:
	go run greet/greet_client/client.go