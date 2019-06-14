package main

import (
    "encoding/json"
    "bufio"
    "fmt"
    "net/http"
    "time"
    "strings"
    "bytes"
)

func identity() {
    timeout := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeout,
    }

    resp, err := client.Get("https://localhost:3000/identity")

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    // Print the first 5 lines of the response body.
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func open(dbname, dbtype string) {
    fmt.Println("open")
    var url strings.Builder
    url.WriteString("https://localhost:3000/db/")
    url.WriteString(dbname)

    values := map[string]string{"create": "true", "type": dbtype}
    jsonValue, _ := json.Marshal(values)

    timeout := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeout,
    }

    resp, err := client.Post(url.String(), "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    // Print the first 5 lines of the response body.
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
