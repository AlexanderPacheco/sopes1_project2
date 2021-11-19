package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"bytes"



	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	kafka "github.com/segmentio/kafka-go"
)



func main() {

	r := getKafkaReader("kafka:9092", "byKafka", "byKafkaSender")

	defer r.Close()

	fmt.Println("Inicio del consumo ... !!")
	Regresar := make(chan bool)
	go func(){
		for{
			msg, err := r.ReadMessage(context.Background())
			errorOcurrido(err,"Error leyendo kafka")
			log.Printf("Mensaje Recivido : %s " ,string(msg.Value));
			//Redis
			url := "http://34.66.140.125:3000/saveData"
			req, err := http.Post(url, "application/json", bytes.NewBuffer(msg.Value))
			req.Header.Set("Content-Type", "application/json")
			errorOcurrido(err, "Ingreso del nuevo documento")
			defer req.Body.Close()
			newBody, err := ioutil.ReadAll(req.Body)
			errorOcurrido(err, "Leer la respuesta HTTP Post")
			sb := string(newBody)
			log.Printf(sb)
			//Mongo
			saveOnServer(string(msg.Value))
			errorOcurrido(err, "Insertando la respuesta HTTP Post")
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-Regresar
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, 
		MaxBytes: 10e6, 
	})
}
func errorOcurrido(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func saveOnServer(req string) {
	fmt.Println("methood save infected")

	clientOptions := options.Client().ApplyURI("mongodb+srv://chris:amor4219@cluster0.3plrc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	errorOcurrido(err, "Cliente Mongo")

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("sopes").Collection("personas")

	var bdoc interface{}
	errb := bson.UnmarshalExtJSON([]byte(req), true, &bdoc)
	fmt.Println(errb)

	insertResult, err := collection.InsertOne(context.Background(), bdoc)
	errorOcurrido(err, "Eror insertar Mongo")


	fmt.Println("Inserted a single document: ", insertResult)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	if err != nil {
		log.Fatal(err)
	}
}