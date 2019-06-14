package main

import (
    "os"
)

func main() {
    args := os.Args

    if len(args) > 1 {
        switch args[1] {
        case "identity":
            identity()
        case "open":
            open(args[2], args[3])
        }
    }
}
