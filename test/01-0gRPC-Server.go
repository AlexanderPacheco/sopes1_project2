package main

import (
	"context"
	"log"
	"net"
	//"strconv"
	"google.golang.org/grpc"
	pb "github.com/AlexanderPacheco/sopes1_project2/gRPC-Server/proto"
)

type server struct {
	pb.UnimplementedOperacionAritmeticaServer
}

func (s *server) OperarValores(ctx context.Context, in *pb.OperacionRequest) (*pb.OperacionReply, error) {
	log.Printf("Se va a %v : el valor %v con el valor %v", in.GetOperacion(),in.GetValor1(),in.GetValor2())

	//Paso el valor1 a int
	intValor1 := in.GetValor1()
	
	//Paso el valor2 a int
	/*
	intValor2, err := strconv.Atoi(in.GetValor2())
	if err != nil {
		log.Fatalf("Error al convertir a int: %v", err)
	}*/

	//**********************************************
	//********** RESOLUCION DEL JUEGO ACA **********
	//**********************************************

	strResultado := "WINNER: " + intValor1 + "-" + in.GetValor2()
	//log.Printf("Received: %v", in.GetOperacion())
	//log.Printf("Operacion exitosa")
	return &pb.OperacionReply{Resultado: "["+ in.GetOperacion() +":"+ intValor1 +"] = " + strResultado}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOperacionAritmeticaServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

