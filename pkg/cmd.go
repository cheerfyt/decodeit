package pkg

import (
	"embed"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewApp(webFS embed.FS) *cli.App {
	return &cli.App{
		Name:  "decodeit",
		Usage: `start the server again`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "listen",
				Aliases: []string{"l"},
				Usage:   "server listen port",
				Value:   "0.0.0.0:9527",
			},
		},
		Action: func(ctx *cli.Context) error {
			var addr = ctx.String("listen")
			s := NewServer(addr, webFS)
			logrus.Infof("server start @%s", addr)
			return s.Start()
		},
	}
}
