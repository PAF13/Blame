package main

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func bestellnummerClean(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	return x
}

/*
func bestellnummerClean2(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "-", "")
	x = strings.ReplaceAll(x, "+", "")
	x = strings.ReplaceAll(x, "/", "")
	x = strings.ReplaceAll(x, ",", "")
	return x
}




func ValueRestrict(s string) error {
	switch s {
	case "Yes", "No", "I don't know":
		fmt.Println("Success!")
		return nil
	default:
		return fmt.Errorf("unsupported value: %q", s)
	}
}
*/
//dialog windows

// Generate UUID
func GenerateCryptoID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
