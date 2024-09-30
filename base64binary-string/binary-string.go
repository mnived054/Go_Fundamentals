package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "91_*()#$@&anmoqmmmsE!W"
	fmt.Println("Data:", data)

	encoded_data := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded Data:", encoded_data)

	decoded_data, _ := b64.StdEncoding.DecodeString(encoded_data)
	fmt.Println("Decoded Data:", string(decoded_data))
}
