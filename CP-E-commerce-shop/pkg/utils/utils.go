package utils

import (
	"encoding/json"
	"fmt"
)

func Debug(log any) {
	bytes, _ := json.MarshalIndent(log, "", "\t")
	fmt.Println(string(bytes))
}

func Outputtofile(log any) []byte {
	bytes, _ := json.Marshal(log)
	return bytes
}
