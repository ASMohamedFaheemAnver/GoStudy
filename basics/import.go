package main

import (
	fmtCustom "fmt"
	axios "net/http"
)

func main() {
	fmtCustom.Println("Hello, Go!")
	resp, error := axios.Get("https://www.google.com")
	if error != nil {
		fmtCustom.Println("Error:", error)
		return
	}
	defer resp.Body.Close()
	fmtCustom.Println("Response Status:", resp.Status)
	fmtCustom.Println("Response Body:", resp.Body)
}
