package main

import (
	"flag"
	"fmt"

	"github.com/Arch-4ng3l/GoServerHololens/assets"
	"github.com/Arch-4ng3l/GoServerHololens/server"
	"github.com/Arch-4ng3l/GoServerHololens/storage"
)

func main() {
	port := flag.String("p", "3000", "Port To Listen On")
	flag.Parse()

	fmt.Println("[*] Building Assets ZIP")
	assets.Init()

	fmt.Println("[*] Connecting To Database")
	psql := storage.NewPostgres()
	psql.Init()

	server := server.NewServer(psql)
	server.Run(":" + *port)

}
