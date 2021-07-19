FROM golang 

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o main.go

EXPOSE 8080

CMD [ "/app/main" ]
