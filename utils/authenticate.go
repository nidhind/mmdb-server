package utils

import "fmt"

func Authenticate(accessToken string) bool {
	fmt.Println("Authentication successful", accessToken);
	return true
}