package commands

func ECC_IsOnline() (up bool) {
	return PingServer("mc.ecocitycraft.com", 25565) != ""
}

func ECC_PlayerCount() (online int) {
	resp, err := GetPingResponse("mc.ecocitycraft.com", 25565)
	if(err != nil) {
		return 0
	}
	return resp.Online
}

func ECC_Latency() (latency uint) {
	resp, err := GetPingResponse("mc.ecocitycraft.com", 25565)
	if(err != nil) {
		return 0
	}
	return resp.Latency
}