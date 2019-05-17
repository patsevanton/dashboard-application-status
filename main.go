package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    url := "http://10.2.64.110:8080/monitoring/selftest"
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    html, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n",html)

}

