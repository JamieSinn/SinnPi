package commands
import (
	"github.com/ammario/mcping"
	"strconv"
	"regexp"
)

func PingServer(ip string, port uint16) (response string) {

	resp, err := mcping.Ping(ip, port)
	if err != nil {
		return "Error while pinging server."
	}
	re, _ := regexp.Compile("((\u00a7([0-9]|[a-f]))|\\n)")
	motd := re.ReplaceAllString(resp.Motd, "")
	return (motd + " has " + strconv.Itoa(resp.Online) + "/" + strconv.Itoa(resp.Max) + " players online." + " v" + resp.Version)
}