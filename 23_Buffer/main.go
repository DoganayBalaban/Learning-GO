package main

import (
	"bytes"
	"fmt"
)

func main() {

	/* //Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Ranjan")
	strBuffer.WriteString("Kumar")
	fmt.Println("The string buffer output is", strBuffer.String()) */

	//Creating buffer variable to hold and manage the string data
	/* 	var strBuffer bytes.Buffer
	   	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")
	   	strBuffer.WriteString("[DONE]")
	   	fmt.Println("The string buffer output is", strBuffer.String()) */

	/* var byteString bytes.Buffer
	byteString.Write([]byte("Hello "))
	fmt.Fprintf(&byteString, "Hello friends how are you")
	byteString.WriteTo(os.Stdout) */

	/* var strByyte bytes.Buffer
	strByyte.Grow(64)
	strByytestrByyte := strByyte.Bytes()
	strByyte.Write([]byte("It is a 64 byte"))
	fmt.Printf("%b", strByytestrByyte[:strByyte.Len()]) */

	/* var strByyte bytes.Buffer
	strByyte.Grow(64)
	strByyte.Write([]byte("kumar"))
	fmt.Printf("%b", strByyte.Len()) */

	var strByyte = []byte("Ranjan, Kumar")
	strByyte = bytes.TrimPrefix(strByyte, []byte("Hello, "))
	strByyte = bytes.TrimPrefix(strByyte, []byte("Good we will meet again,"))
	fmt.Printf("Hello%s", strByyte)
}
