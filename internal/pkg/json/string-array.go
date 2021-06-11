package json

import (
	"fmt"
	// "strings"
	// "encoding/json"
)

type String2Array struct {
	
}

func (sa String2Array) UnmarshalJSON(s []byte) (err error) {
	fmt.Println(s)
	fmt.Printf("string array unmarshal json v is %v\n", sa)
	// *(*[]string)(sa) = strings.Split(string(s), ",")

	return
}

func (sa String2Array) MarshalJSON() ([]byte, error) {
	fmt.Printf("string array marshal json v is %v\n", sa)
	return []byte{49}, nil
}
