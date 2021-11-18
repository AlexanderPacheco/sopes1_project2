#Obtener la imagen que necesitamos
#en este caso, node v.14.x.x
FROM node:14.18-alpine3.12

#Asignar el workdir de la api
WORKDIR /usr/src/app

#Copiar el package-json
COPY package*.json ./

#Correr 'npm install' para instalar todas las dependencias que necesitamos
RUN npm install

#Aca pueden existir muchos mas comandos "run"


#Copiar todos los archivos al workdir
COPY . .

ENV GOOGLE_APPLICATION_CREDENTIALS='./key.json'

#Exponer el puerto 3500 al mundo
#EXPOSE 3500

#Correr NodeJS
#CMD [ "node", "server.js" ]
CMD [ "node", "." ]