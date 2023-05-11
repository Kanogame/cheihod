FROM golang:1.20-alpine

RUN mkdir /build
ADD . /build
WORKDIR /build
EXPOSE 10234
EXPOSE 10235

RUN go build -o main .
CMD [ "/build/main" ]
