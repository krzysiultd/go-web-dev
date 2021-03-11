package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Tyger Tyger, burning bright, In the forests of the night; What immortal hand or eye, Could frame thy fearful symmetry?"

	encodeStd :=
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(s64)
}
