package printjson

import (
	"encoding/json"
	"os"
)

func PrintJson(patch string, data interface{}) error {
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = f.Write(json_data)
	if err != nil {
		return err
	}

	return nil
}
