package main 

import (
    "log"
    "os"
    "fmt"
)

func main() {
    key := "DB_CONN"
    
    connString, ex := os.LookupEnv(key)
    
    if !ex {
        log.Printf("The env variable %s is not set \n", key)
    }
    fmt.Println(connString)
}
