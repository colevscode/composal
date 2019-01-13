package main

import (
	"flag"

	"github.com/colevscode/composal/player"
	"github.com/colevscode/composal/server"
)

func main() {
	port := flag.Int("port", 8088, "port on which to run the player server")
	prefix := flag.String("prefix", "./samples", "folder containing samples")
	debug := flag.Bool("debug", false, "if set will print debug to stderr")
	flag.Parse()
	player.Setup(*prefix, *debug)
	server.RunServer(*port, *debug)
}
