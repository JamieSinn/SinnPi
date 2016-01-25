package sinnpi

import (
	irc "github.com/fluffle/goirc/client"
)

func init() {
	config := irc.NewConfig("SinnPi")
	config.SSL = true
	config.Pass = "nicetrylol"
	config.Server = ""
}

func readConfig() (config *irc.Config){
	// I probably should figure out how json works eventually...
	return
}