package main

import (
	"context"
	"log"
	"net"
	//"strconv"
	"google.golang.org/grpc"
	"os"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	pb "github.com/AlexanderPacheco/sopes1_project2/gRPC-Server/proto"
)

type server struct {
	pb.UnimplementedOperacionAritmeticaServer
}

func (s *server) OperarValores(ctx context.Context, in *pb.OperacionRequest) (*pb.OperacionReply, error) {
	log.Printf("Se va a %v : el valor %v con el valor %v", in.GetOperacion(),in.GetValor1(),in.GetValor2())

	//Paso el valor1 a int
	game_name := in.GetValor1()
	
	//Paso el valor2 a int
	/*
	intValor2, err := strconv.Atoi(in.GetValor2())
	if err != nil {
		log.Fatalf("Error al convertir a int: %v", err)
	}*/

	//**********************************************
	//********** RESOLUCION DEL JUEGO ACA **********
	//**********************************************

	strResultado := "{\"request_number\":30001, \"game\":"+in.GetOperacion()+", \"gamename\": \""+game_name+"\", \"winner\":\"001\", \"players\":"+in.GetValor2()+", \"worker\":\"RabbitMQ\" }" //+ game_name + "-" + in.GetValor2()

	//********** ENVIANDO INFORMACION A RABBITMQ **********
	sendResultado(strResultado)
	//******* FINALIZA ENVIO INFORMACION A RABBITMQ *******

	//log.Printf("Received: %v", in.GetOperacion())
	//log.Printf("Operacion exitosa")
	return &pb.OperacionReply{Resultado: "["+ in.GetOperacion() +":"+ game_name +"] = " + strResultado}, nil
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


// Funciones de RabbitMQ
func sendResultado(strResultado string) {

	host:= os.Getenv("HOST") //Obtengo el nombre del host donde esta corriendo actualmente mi imagen

	//conn, err := amqp.Dial("amqp://guest:guest@34.72.156.55:5672/")
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://guest:guest@"+host+":5672/")
	
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := bodyFrom(os.Args)
	body := strResultado
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "alexanderx3"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
