package main
import (
	"github.com/jamiesinn/commands"
	"fmt"
)

func main()  {
	fmt.Println(commands.PingServer("mc.ecocitycraft.com", 25565))
	fmt.Println(commands.GetStatuses())
}
