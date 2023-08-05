package main

import (
	"flag"
	"fmt"
	"os/user"

	"github.com/Arch-4ng3l/GoServerHololens/assets"
	"github.com/Arch-4ng3l/GoServerHololens/server"
	"github.com/Arch-4ng3l/GoServerHololens/storage"
)

func main() {
	user, err := user.Current()
	if err != nil {
		return
	}

	port := flag.String("p", "3000", "Port To Listen On")
	flag.Parse()

	fmt.Println("[*] Building Assets ZIP")
	assets.Init(user.HomeDir)

	fmt.Println("[*] Connecting To Database")
	psql := storage.NewPostgres()
	psql.Init()

	//file, err := os.Create(user.HomeDir + "/Hololens/logs/" + time.Now().Format(time.DateOnly) + ".log")
	if err != nil {
		return
	}
	//os.Stdin = file
	//os.Stdout = file
	//os.Stderr = file

	server := server.NewServer(psql, user.HomeDir)
	server.Run(":" + *port)

}
