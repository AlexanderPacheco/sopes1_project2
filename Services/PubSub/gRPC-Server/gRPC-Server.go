package main

import (
	"fmt"
	"context"
	"log"
	"net"
	//"strconv"
	"google.golang.org/grpc"
	"os"
	"encoding/json"

	// Leer variables de entorno
	"github.com/joho/godotenv"
	// Libreria de Google PubSub
	"cloud.google.com/go/pubsub"

	pb "github.com/AlexanderPacheco/sopes1_project2/gRPC-Server/proto"
)

var MONGO_URL = "mongodb://mongoazzure:nSpxpeAiUugPuUxNhJWIwj9k5FNfXGoSXMR9I6T2whCQIWEoT0JrtDjIjenlIi7EGmvFkZB5OrAWB0WruUG0PA==@mongoazzure.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@mongoazzure@"

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

	strResultado := "{\"request_number\":30001, \"game\":"+in.GetOperacion()+", \"gamename\": \""+game_name+"\", \"winner\":\"001\", \"players\":"+in.GetValor2()+", \"worker\":\"PubSub\" }" //+ game_name + "-" + in.GetValor2()

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

// ------------------------------------------------------------- PUB SUB ------------------------------------------------------------------------
// Con esta funcion obtendremos variables de entorno
// Desde el archivo de configuracion
func goDotEnvVariable(key string) string {

	// Leer el archivo .env ubicado en la carpeta actual
	err := godotenv.Load(".env")
	
	// Si existio error leyendo el archivo
	if err != nil {
	  log.Fatalf("Error cargando las variables de entorno")
	}
	
	// Enviar la variable de entorno que se necesita leer
	return os.Getenv(key)
}


// Esta funcion es utilizada para publicar un mensaje
// Como parametro se manda el mensaje que publicaremos a PubSub
func publish(msg string) error {
	// Definimos el ProjectID del proyecto
	// Este dato lo sacamos de Google Cloud
	projectID := goDotEnvVariable("PROJECT_ID")

	// Definimos el TopicId del proyecto
	// Este dato lo sacamos de Google Cloud
	topicID := goDotEnvVariable("TOPIC_ID")

	// Definimos el contexto en el que ejecutaremos PubSub
	ctx := context.Background()
	// Creamos un nuevo cliente
	client, err := pubsub.NewClient(ctx, projectID)
	// Si un error ocurrio creando el nuevo cliente, entonces imprimimos un error y salimos
	if err != nil {
		fmt.Println("Error al crear el cliente")
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	
	// Obtenemos el topico al que queremos enviar el mensaje
	t := client.Topic(topicID)

	// Publicamos los datos del mensaje
	result := t.Publish(ctx, &pubsub.Message { Data: []byte(msg), })
	
	// Bloquear el contexto hasta que se tenga una respuesta de parte de GooglePubSub
	id, err := result.Get(ctx)
	
	// Si hubo un error creando el mensaje, entonces mostrar que existio un error
	if err != nil {
		fmt.Println("Error al crear un mensaje")
		fmt.Println(err)
		return fmt.Errorf("Error: %v", err)
	}

	// El mensaje fue publicado correctamente
	fmt.Println("Published a message; msg ID: %v\n", id)
	return nil
}


// Esta estructura almacenara la forma en la que se enviaran los datos al servidor
type Message struct {
	Data string
}


// Creamos un server sencillo que unicamente acepte peticiones GET y POST a '/'
func sendResultado(strResultado string) {
	
	// Obtener el nombre enviado desde la forma
	data := strResultado

	// Crear un objeto JSON con los datos enviados desde la forma
	message, err := json.Marshal(Message{ Data: data})
	// Existio un error generando el objeto JSON
	if err != nil {
		fmt.Println("Error en informacion")
		return
	}

	// Publicar el mensaje, convertimos el objeto JSON a String
	publish(string(message))

	fmt.Println("¡Mensaje Publicado desde Go!\n")

}