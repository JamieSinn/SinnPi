package commands
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type MojangServices struct {
	ServiceName,
	ServiceStatus string
}

func GetStatuses() (statuses []MojangServices){
	var out []map[string]string
	json.Unmarshal([]byte(getRawJSON()), &out)

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