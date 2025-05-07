package main

import "strings"

func IdentifyPrefixPostfix(userID, email string) bool {
	return strings.HasPrefix(userID, "id:") || strings.HasSuffix(email, "@educative.io")

}

func ContainsEducative(email string) bool {
	return strings.Contains(email, "educative")
}

func MaskUserName(email string) string {
	// Implement this function
	splitstr := strings.Split(email, "@")

	if len(splitstr) > 0 {
		return splitstr[0]
	}

	return ""
}
