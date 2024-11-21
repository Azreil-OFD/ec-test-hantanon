package main

import (
	"backend/internal/routs"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
	routs.RegisterRoutes()
}
