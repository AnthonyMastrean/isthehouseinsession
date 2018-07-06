FROM golang:alpine

LABEL maintainer="Anthony Mastrean <anthony.mastrean@gmail.com>"

WORKDIR /go/src/github.com/AnthonyMastrean/isthehouseinsession

COPY . .

RUN go build -o isthehouseinsession

EXPOSE 80

ENTRYPOINT ["./isthehouseinsession"]
