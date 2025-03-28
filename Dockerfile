FROM golang:1.24.1-alpine3.21

WORKDIR /kv_app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o kv_app ./cmd

EXPOSE 8080

CMD [ "./kv_app" ]
