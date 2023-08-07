package main

import (
	"flag"
	"fmt"
	"os/user"

	"github.com/Arch-4ng3l/GoServerHololens/assets"
	"github.com/Arch-4ng3l/GoServerHololens/server"
	"github.com/Arch-4ng3l/GoServerHololens/storage"
	"github.com/Arch-4ng3l/GoServerHololens/util"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	port := flag.String("p", "3000", "Port To Listen On")
	logs := flag.Bool("l", true, "Activation of Loging")
	flag.Parse()

	fmt.Println("[*] Building Assets ZIP")
	assets.Init(user.HomeDir)

	fmt.Println("[*] Connecting To Database")
	psql := storage.NewPostgres()
	psql.Init()
	fmt.Println()
	server := server.NewServer(psql, user.HomeDir)
	server.Run(":" + *port)
	if *logs {
		util.ActivateLogs(user.HomeDir)
	}

}
