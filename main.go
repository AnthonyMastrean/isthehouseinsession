package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", inSession)
    http.ListenAndServe(":80", nil)
}

func inSession(writer http.ResponseWriter, request *http.Request) {
    upstream, err := http.Get("https://in-session.house.gov/")

    if err != nil {
        log.Fatal(err)
    }

    body, err := ioutil.ReadAll(upstream.Body)

    if err != nil {
        log.Fatal(err)
    }

    if upstream.StatusCode != http.StatusOK {
        writer.WriteHeader(upstream.StatusCode)
    }

    switch string(body) {
        case "0":
            writer.WriteHeader(http.StatusNoContent)
        case "1":
            writer.WriteHeader(http.StatusOK)
        default:
            writer.WriteHeader(http.StatusInternalServerError)
    }
}
