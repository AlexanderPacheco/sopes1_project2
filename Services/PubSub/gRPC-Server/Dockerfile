FROM golang:1.17.3-alpine3.13
WORKDIR /app/src
ENV GOPATH=/app
COPY . /app/src
RUN go build -o Server .
ENV GOOGLE_APPLICATION_CREDENTIALS='./key.json'
EXPOSE 50051
CMD ["./Server"]