package utils

import "encoding/json"

func CopyStruct(src interface{}, dest interface{}) error {
	srcByte, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(srcByte, &dest)
	if err != nil {
		return err
	}
	return nil
}
