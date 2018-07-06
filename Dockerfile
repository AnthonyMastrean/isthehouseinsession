FROM golang:alpine

LABEL maintainer="Anthony Mastrean <anthony.mastrean@gmail.com>"

WORKDIR /go/src/github.com/AnthonyMastrean/isthehouseinsession

COPY . .

RUN go build -v .

EXPOSE 80

ENTRYPOINT ["./isthehouseinsession"]
