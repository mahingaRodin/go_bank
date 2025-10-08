package main

import  ( "fmt"
	"runtime"
)

func main() {
	fmt.Println("🎉 Go is working perfectly on Windows!")
	fmt.Printf("Go version: %s\n", runtime.Version())
}