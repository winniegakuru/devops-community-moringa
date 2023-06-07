package main

import "fmt"

var version = "dev1"

func main() {

 fmt.Printf("Version: %s\n", version)

 fmt.Println(hello())
}

func hello() string {
 return "Hello Golang"
}