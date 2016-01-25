package commands
import (
	"net/http"
	"io/ioutil"
)



func GetStatuses() {

	//jsonR := getRawJSON()
	return
}

func getRawJSON() (rawjson string) {
	resp, err := http.Get("http://status.mojang.com/check")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return string(raw)
}