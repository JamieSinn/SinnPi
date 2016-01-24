package commands
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type MojangServices struct {
	Services []struct {
		Status map[string]string
	}
}


func GetStatuses()(result MojangServices) {

	jsonR := getRawJSON()
	fmt.Println(jsonR)

	result = MojangServices{}
	json.Unmarshal([]byte(jsonR), &result)

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