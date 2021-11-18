package main

import (
	"context"
	"log"
	"time"
	"net/http"
	"encoding/json" 
	"github.com/gorilla/mux" 
	"os"
	"github.com/Pallinder/go-randomdata"
	"google.golang.org/grpc"
	pb "github.com/AlexanderPacheco/sopes1_project2/gRPC-Client-api/proto"
)

var nombre_api = "default"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API GO - gRPC Client!\n"))
}

func endgame(w http.ResponseWriter, r *http.Request) {
	operacion := mux.Vars(r)["game"] //Obtengo la operacion a realizar
	num1 := mux.Vars(r)["gamename"] //Obtengo el valor 1
	num2 := mux.Vars(r)["players"] //Obtengo el valor 2
	host:= os.Getenv("HOST") //Obtengo el nombre del host donde esta corriendo actualmente mi imagen

	log.Printf("Iniciando nuevo juego")

	/********************************** gRPC llamada al servidor ********************************/
	conn, err := grpc.Dial(host+":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOperacionAritmeticaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Funcion de gRPC-Server
	reply, err := c.OperarValores(ctx, &pb.OperacionRequest{
		Operacion: operacion,
		Valor1:num1,
		Valor2:num2,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//log.Printf("Greeting: %s", reply.GetResultado())
	/********************************** gRPC ********************************/

	/********************************** Respuesta ********************************/
	w.Header().Set("Content-Type", "application/json")
   	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(struct {
		Mensaje string `json:"mensaje"`
		Server string `json:"server"`
	}{Mensaje: reply.GetResultado(),Server:nombre_api})
}

func main() {
	nombre_api = randomdata.SillyName()
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", IndexHandler)
	//router.HandleFunc("/operacion/{op}/valor1/{val1}/valor2/{val2}",operacionesAritmeticas).Methods("POST")
	router.HandleFunc("/game/{game}/gamename/{gamename}/players/{players}", endgame).Methods("POST")
    log.Println("Listening at port 2000") 
	log.Fatal(http.ListenAndServe(":2000", router))
}