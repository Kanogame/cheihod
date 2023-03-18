FROM golang:1.20

ENV GO111MODULE=auto

WORKDIR /build

COPY go.mod ./

COPY ./HttpServer ./HttpServer
COPY ./static ./static
COPY main.go ./

RUN go build -o go-server ./

EXPOSE 10234

CMD [ "/go-server" ]