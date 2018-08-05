package main

import (
	"Curso/curso-go-gRPC/code_unary/greet/greetpb"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cliente não conectado: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)

}

func doUnary(client greetpb.GreetServiceClient) {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Name:  "Tiago Nunes",
			Email: "tiagonunes.developer@gmail.com",
		},
	}

	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro na chamada do RPC: %v", err)
	}

	log.Printf("Resposta: %v", res.Result)
}
