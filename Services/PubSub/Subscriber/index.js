
const mysql = require('mysql');
const db_credentials = require('./db_creds');
var connection = mysql.createPool(db_credentials);

// Importar la libreria de Google PubSub
// Para instalar, utilizamos npm install --save @google-cloud/pubsub
// Generalmente esto lo hacemos en un fronted!
const { PubSub } = require('@google-cloud/pubsub');

// Importar axios para realizar una peticion http
// Para instalar utilizamos npm install --save axios
const axios = require('axios');

// Acá escribimos la suscripción que creamos en Google Pub/Sub
const SUB_NAME = 'projects/psychic-droplet-331301/subscriptions/suscription';

// Cantidad de segundos que estara prendido nuestro listener
// Solo para efectos practicos, realmente esto debería estar escuchando en todo momento!
const TIMEOUT = process.env.TIMEOUT || 6000;

// Crear un nuevo cliente de pubsub
const client = new PubSub();

// En este array guardaremos nuestros datos
const messages = [];

// Esta funcion se utilizara para leer un mensaje
// Se activara cuando se dispare el evento "message" del subscriber
const messageReader = async message => {

    console.log('¡Mensaje recibido!');
    //Json a cambiar
    console.log(`${message.id} - ${message.data}`);
    console.table(message.attributes);

    messages.push({ msg: String(message.data), id: message.id, ...message.attributes });

    // Con esto marcamos el acknowledgement de que recibimos el mensaje
    // Si no marcamos esto, los mensajes se nos seguirán enviando aunque ya los hayamos leído!
    message.ack();

    try {
        console.log(`Agregando mensaje al servidor...`);
        const jsonMessage = JSON.parse(message.data) || {};
        const request_body = { data: jsonMessage.Data || jsonMessage.data || "0" };

        // ++++++++++++++++++++ NOTA +++++++++++++++++++++++++++
        // Desde acá se enviará a la otra api que esté utilizando sockets o se haaría en la misma. Preguntarle al AUX
        console.log(`Realizando petición POST a ${process.env.API_URL}`);
        console.log(request_body);

        //Guardando informacion en BD

        var sql = "insert into juegos(post) value ('"+ String(request_body.data) +"')"; 
            connection.query(sql, function (err, result) {
                if (err) throw err;
                console.log("1 record inserted");
        });

        // ++++++++++++++++++++ NOTA +++++++++++++++++++++++++++
        //await axios.post(process.env.API_URL, request_body); //Esto no va porque Se enviaría a traves de sockets a angular creo :v
    }
    catch (e) {
        console.log(`Error al realizar POST ${e.message}`);
    }
};

// Empezamos nuestro manejador de notificaciones
const notificationListener = () => {

    // Creamos un subscriptor
    // Pasamos el nombre de nuestro subscriptor (que encontramos en Google Cloud)
    const sub = client.subscription(SUB_NAME);

    // Conectar el evento "message" al lector de mensajes
    sub.on('message', messageReader);

    console.log("Estoy a la espera de los mensajes...");

    setTimeout(() => {
        sub.removeListener('message', messageReader);

        if (messages.length > 0) {
            console.log(`${messages.length} mensajes recibidos: `);
            console.log("---------");
            console.table(messages);
        }
        else {
            console.log("No hubo ningún mensaje en este tiempo. :(")
        }

    }, TIMEOUT * 1000);
};

console.log(`Iniciando Subscriber, deteniendolo en ${TIMEOUT} segundos...`);

// Empezar a escuchar los mensajes
notificationListener();


//Para correrlo: node .