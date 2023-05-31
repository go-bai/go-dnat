package main

import (
	"log"
	"os"

	"github.com/go-bai/go-dnat/cmd"
	"github.com/go-bai/go-dnat/db"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Ins.Close()

	app := cmd.InitCmd()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
