package main

import "fmt"

var version = "devs"

func main() {

 fmt.Printf("Version: %s\n", version)

 fmt.Println(hello())
}

func hello() string {
 return "Hello Golang"
}