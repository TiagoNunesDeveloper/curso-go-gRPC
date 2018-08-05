gen_proto:
	protoc sum/sumpb/sum.proto --go_out=plugins=grpc:.

gen_server:
	go run sum/sum_server/server.go

gen_client:
	go run sum/sum_client/client.go