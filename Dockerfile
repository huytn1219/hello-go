FROM registry.access.redhat.com/ubi8/go-toolset:1.15.13-4 

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o main.go

EXPOSE 8080

CMD [ "/main" ]
