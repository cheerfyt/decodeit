package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/cheerfyt/decodeit/pkg"
)

//go:embed web/dist
var webFs embed.FS

func main() {
	err := pkg.NewApp(webFs).Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
