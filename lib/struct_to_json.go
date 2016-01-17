package lib

import (
	"encoding/json"
	"fmt"
)

func StructToJson(v interface{}) string {
	b, err := json.Marshal(v)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	return string(b)
}
