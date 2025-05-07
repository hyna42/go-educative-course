package main

import "strings"

func IdentifyPrefixPostfix(userID, email string) bool {
	return strings.HasPrefix(userID, "id:") || strings.HasSuffix(email, "@educative.io")

}

func ContainsEducative(email string) bool {
    return strings.Contains(email,"educative")
}

// func MaskUserName(email string) string {
//     // Implement this function
// }



