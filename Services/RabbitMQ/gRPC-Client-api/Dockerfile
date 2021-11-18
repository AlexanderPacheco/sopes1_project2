FROM golang:1.17.3-alpine3.13
WORKDIR /app/src
ENV GOPATH=/app
COPY . /app/src
RUN go build -o Server .
EXPOSE 2000
CMD ["./Server"]