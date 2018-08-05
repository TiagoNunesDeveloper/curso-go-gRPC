package main

import (
	"Curso/curso-go-gRPC/code_calculator/sum/sumpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	number1 := req.GetSuming().GetNumber_1()
	number2 := req.GetSuming().GetNumber_2()

	result := fmt.Sprintf("A soma de %d + %d é %d", number1, number2, (number1 + number2))
	res := &sumpb.SumResponse{
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
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Servidor server falhou: %v", err)
	}

}
