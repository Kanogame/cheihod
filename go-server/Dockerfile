FROM golang:1.20-alpine

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN go build -o main .
CMD [ "/build/main" ]
