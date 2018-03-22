package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    r, err := http.Get("https://in-session.house.gov/")

    if err != nil {
        log.Fatal(err)
    }

    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatal(err)
    }

    var status int
    switch string(body) {
        case "0":
            status = 404
        case "1":
            status = 200
        default:
            log.Fatal(err)
    }

    fmt.Printf("%d\n", status)
}
