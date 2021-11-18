FROM golang:1.16.2-alpine3.13
WORKDIR /app/src
ENV GOPATH=/app
COPY . /app/src
RUN go get -u github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o Server .
ENV GOOGLE_APPLICATION_CREDENTIALS='./key.json'
EXPOSE 3001
#CMD ["go","run","Server.go"]
CMD ["./Server"]