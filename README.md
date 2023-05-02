# Muxxer

## Install

    go get github.com/flohansen/muxxer

## Usage

    package main

    import (
        "net/http"

        "github.com/flohansen/muxxer"
    )

    func main() {
        router := muxxer.New()
        router.GET("/hello", testHandler)
        http.ListenAndServe("0.0.0.0:3000", router)
    }

    func testHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "World!")
    }
