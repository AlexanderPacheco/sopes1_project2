package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	//"io/ioutil"

	kafka "github.com/segmentio/kafka-go"
)





func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type Person struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Gender       string `json:"gender"`
	Age          string    `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	Origin       string `json:"origin"`
}
func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	return err
}



func nuevoPaciente(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		wrt.Header().Set("Content-Type", "application/json")
		if req.Method == "GET" {
			wrt.WriteHeader(http.StatusOK)
			wrt.Write([]byte("{\"message\": \"Sender:Ok\"}"))
			return;
		}

		var p Person
		//var pp Persona
		err := decodeJSONBody(wrt, req, &p)
		errorOcurrido(err, "Parseo JSON")
		//bodymayus := Persona{name: p.Name,location: p.Location,age: p.Age,gender: p.Gender,vaccine_type:p.Vaccine_type,origin:"Kafka"}
		bodymayus := Person{Name: p.Name,Location: p.Location,Age: p.Age,Gender: p.Gender,Vaccine_type:p.Vaccine_type,Origin:"Kafka"}
		data, err := json.Marshal(bodymayus)
		errorOcurrido(err, "Parseo JSON DATA")

		//body, err := ioutil.ReadAll(data)
		//errorOcurrido(err, "Parseo JSON iotil")
		fmt.Printf("el body es: ")
		//fmt.Printf((bodymayus))
		fmt.Printf("el data es: ")
		fmt.Printf(string(data))
		if err != nil {
			log.Fatalln(err)
		}
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
			Value: data,
		}
		err = kafkaWriter.WriteMessages(req.Context(), msg)
		wrt.WriteHeader(http.StatusCreated)
		wrt.Write([]byte(string(data)))
		if err != nil {
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}


func errorOcurrido(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func main (){
	kafkaWriter := getKafkaWriter("kafka:9092", "byKafka")
	defer kafkaWriter.Close()
	http.HandleFunc("/", nuevoPaciente(kafkaWriter))
	fmt.Println("EL PRODUCER ESTA CORRIENDO TALEGA.. !!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}