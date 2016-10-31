package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe(":3030", nil))

}
