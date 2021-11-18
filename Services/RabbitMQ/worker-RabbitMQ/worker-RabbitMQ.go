package main

import (
	"bytes"
	"log"
	"time"
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	amqp "github.com/rabbitmq/amqp091-go"
)


//Los atributos del struct siempre inician en mayusculas
type mensaje struct{
	Status bool `json:Status`
	Name string `json:Name`
}

type allHashtags []string

//Los atributos del struct siempre inician en mayusculas
type Publicacion struct{
	Post string `json:Post`
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "admin"
    dbPass := "Ayd2_2021"
    dbName := "ayd2phase1"
	dbConnection := "database-ayd2.caxlo7fjv8ng.us-east-2.rds.amazonaws.com"
	dbPort := "3306"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbConnection+":"+dbPort+")/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
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

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//cadena = ""
			//cadena = String()
			saveGame(string(d.Body))//Guardando en bd
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


func saveGame(juego string) {

	log.Println("[POST]SaveR/")
	msg := mensaje{}
	msg.Name = "Go"

    //w.Header().Set("Content-Type", "application/json")
	
	//_ = json.NewDecoder(r.Body).Decode(&publicacion)
	//fmt.Printf("ImpresionXX : %+v \n", publicacion)
	
	var publicacion Publicacion

	//json.Unmarshal(reqBody, &publicacion)
	publicacion.Post = juego
	//fmt.Printf("New Post: %+v\n", publicacion)
	
	db := dbConn() //Iniciando conexion bd
	
	res, err := db.Exec("insert into juegos(post) value (?)",publicacion.Post)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Se registro juego: ", res)

	defer db.Close() //Finalizando conexion bd
	

}