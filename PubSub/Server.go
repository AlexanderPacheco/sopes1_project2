package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"
	//"strconv"
	"context"
	
	"github.com/gorilla/mux"

	//"text/template"
	_ "github.com/go-sql-driver/mysql"


	// Leer variables de entorno
	"github.com/joho/godotenv"
	// Libreria de Google PubSub
	"cloud.google.com/go/pubsub"
)

var MONGO_URL = "mongodb://mongoazzure:nSpxpeAiUugPuUxNhJWIwj9k5FNfXGoSXMR9I6T2whCQIWEoT0JrtDjIjenlIi7EGmvFkZB5OrAWB0WruUG0PA==@mongoazzure.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@mongoazzure@"


func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API GO PUBSUB")
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
func http_server(w http.ResponseWriter, r *http.Request) {
	// Comprobamos el tipo de peticion HTTP
    switch r.Method {
		// // Devolver una página sencilla con una forma html para enviar un mensaje
		// case "GET":     
		// 	// Leer y devolver el archivo form.html contenido en la carpeta del proyecto
		// 	http.ServeFile(w, r, "form.html")

		// Publicar un mensaje a Google PubSub
		case "POST":
			// Si existe un error con la forma enviada entonces no seguir
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			// // Obtener el nombre enviado desde la forma
			// saved := r.FormValue("Saved")
			// api := r.FormValue("Api")
			// loading_time := r.FormValue("Loading_time")
			// database := r.FormValue("Database")

			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Insert a Valid Task")
			}

			var publication Message
			json.Unmarshal(reqBody, &publication)

			// Obtener el nombre enviado desde la forma
			data := publication.Data

			// Crear un objeto JSON con los datos enviados desde la forma
			message, err := json.Marshal(Message{ Data: data})
			// Existio un error generando el objeto JSON
			if err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			// Publicar el mensaje, convertimos el objeto JSON a String
			publish(string(message))

			// Enviamos informacion de vuelta, indicando que fue generada la peticion
			fmt.Fprintf(w, "¡Mensaje Publicado desde Go!\n")
			fmt.Fprintf(w, "Data = %s\n", data)
			fmt.Fprintln(w, string(message))
		
		// Cualquier otro metodo no sera soportado
		default:
			fmt.Fprintf(w, "Metodo %s no soportado \n", r.Method)
			return
    }
}



func main()  {

	log.Println("Server started on: http://localhost:3001")


	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/SendNotification", http_server)
	

	log.Fatal(http.ListenAndServe(":3001", router))
}

//go run main.go
//go mod init main //Resuelve errores de requerimiento de modulos
/*
go get go.mongodb.org/mongo-driver/mongo
go get -u github.com/go-sql-driver/mysql
go get -v -u github.com/gorilla/mux
*/
