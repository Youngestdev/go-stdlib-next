package main

import (
    "log"
    "os"
)

func main() {
    connString := os.Getenv("DB_CONN")
    log.Printf("Connection string: %s \n", connString)
}
