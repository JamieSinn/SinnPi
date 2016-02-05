package irc

import (
	irc "github.com/fluffle/goirc/client"
	"fmt"
	"bufio"
	"os"
	"strings"
)

var (
	config = irc.NewConfig("SinnPi")
	c irc.Client
	channels = []string{"", ""}
)

func init() {
	config.SSL = true
	config.Pass = "nicetrylol"
	config.Server = ""
	c = irc.Client(config)

}

func readConfig() (config *irc.Config){
	// I probably should figure out how json works eventually...
	return
}

func JoinChannel(channel string, pass... string) {
	if pass != nil {
		c.Join(channel, pass)
	} else {
		c.Join(channel)
	}
}

func joinChannelsOnConnect() {
	c.HandleFunc(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			for i := 0; i <= len(channels); i++ {
				conn.Join(channels[i])
			}
		})
}
func Connect() {

	// And a signal on disconnect
	quit := make(chan bool)

	c.HandleFunc(irc.DISCONNECTED,
		func(conn *irc.Conn, line *irc.Line) { quit <- true })


	// Tell client to connect.
	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	// With a "simple" client, set Server before calling Connect...
	c.Config().Server = "irc.freenode.net"

	// ... or, use ConnectTo instead.
	if err := c.ConnectTo("irc.freenode.net"); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	// Wait for disconnect
	<-quit
}

func handleCommand() {
	in := make(chan string, 4)
	reallyquit := false
	go func() {
		con := bufio.NewReader(os.Stdin)
		for {
			s, err := con.ReadString('\n')
			if err != nil {
				// wha?, maybe ctrl-D...
				close(in)
				break
			}
			// no point in sending empty lines down the channel
			if len(s) > 2 {
				in <- s[0 : len(s)-1]
			}
		}
	}()

	// set up a goroutine to do parsey things with the stuff from stdin
	go func() {
		for cmd := range in {
			if cmd[0] == ':' {
				switch idx := strings.Index(cmd, " "); {
				case cmd[1] == 'd':
					fmt.Printf(c.String())
				case cmd[1] == 'f':
					if len(cmd) > 2 && cmd[2] == 'e' {
						// enable flooding
						c.Config().Flood = true
					} else if len(cmd) > 2 && cmd[2] == 'd' {
						// disable flooding
						c.Config().Flood = false
					}
					for i := 0; i < 20; i++ {
						c.Privmsg("#", "flood test!")
					}
				case idx == -1:
					continue
				case cmd[1] == 'q':
					reallyquit = true
					c.Quit(cmd[idx + 1 : len(cmd)])
				case cmd[1] == 'j':
					c.Join(cmd[idx + 1 : len(cmd)])
				case cmd[1] == 'p':
					c.Part(cmd[idx + 1 : len(cmd)])
				}
			} else {
				c.Raw(cmd)
			}
		}
	}()
}
