package sdktest

import (
	"fmt"
	"log"
)

func printTest(msg string) {
	fmt.Println("Hello World from SDK:", msg)
}

func logInfo(msg string){
	log.Printf("Info: %v", msg)
}
