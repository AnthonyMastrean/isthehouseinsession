FROM golang:alpine

WORKDIR /go/src/github.com/AnthonyMastrean/isthehouseinsession

COPY . .

RUN go build -o isthehouseinsession

EXPOSE 80

ENTRYPOINT ["./isthehouseinsession"]
