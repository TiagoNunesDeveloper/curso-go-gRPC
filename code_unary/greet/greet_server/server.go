package main

import (
	"Curso/curso-go-gRPC/code_unary/greet/greetpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	nome := req.GetGreeting().GetName()
	email := req.GetGreeting().GetEmail()
	result := "Olá " + nome + ", seu Email é " + email
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Printf("Servidor falhou: %v", err)
	}

	fmt.Println("Servidor Rodando na porta :50051")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Servidor server falhou: %v", err)
	}

}
