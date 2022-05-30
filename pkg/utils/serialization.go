package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func Slice2String(argsList []string) string {
	/*Convert slice to serializable string*/
	data, err := json.Marshal(argsList)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	fmt.Printf("%s\n", string(data))
	return string(data)
}
