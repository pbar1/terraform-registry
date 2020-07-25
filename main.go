package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const helpMsg = `Terraform Registry API

Usage:
  terraform-registry [FLAGS]

Flags:
      --debug     set log level to debug
  -h, --help      help for terraform-registry
  -v, --version   print program version`

var (
	ver                   string
	port                  = 8080
	backend               Backend
	moduleArchiveFilename = "module.tar.gz"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" || arg == "-help" {
			fmt.Println(helpMsg)
			os.Exit(0)
		}
		if arg == "-v" || arg == "--version" || arg == "-version" {
			fmt.Println(ver)
			os.Exit(0)
		}
		if arg == "--debug" || arg == "-debug" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
				With().Caller().Logger()
			log.Debug().Msg("log level set to debug")
		}
	}

	backend = NewFilesystemBackend(".", moduleArchiveFilename)

	server := NewServer()
	if err := server.Listen(port); err != nil {
		log.Fatal().Err(err)
	}
}
