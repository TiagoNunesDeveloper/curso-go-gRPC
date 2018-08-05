package main

import (
	"Curso/curso-go-gRPC/code_calculator/sum/sumpb"
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

	c := sumpb.NewSumServiceClient(conn)

	doSum(c)
}

func doSum(client sumpb.SumServiceClient) {
	req := &sumpb.SumRequest{
		Suming: &sumpb.Suming{
			Number_1: 10,
			Number_2: 50,
		},
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro na chamada do RPC: %v", err)
	}

	log.Printf("Resposta: %v", res.Result)
}
