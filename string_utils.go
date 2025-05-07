package main

import (
	"strconv"
	"strings"
)

func IdentifyPrefixPostfix(userID, email string) bool {
	return strings.HasPrefix(userID, "id:") || strings.HasSuffix(email, "@educative.io")

}

func ContainsEducative(email string) bool {
	return strings.Contains(email, "educative")
}

func MaskUserName(email string) string {
	splitstr := strings.Split(email, "@")

	if len(splitstr) > 0 {
		str := splitstr[0]
		return str[0:1] + strings.Repeat("*", len(str)-2) + str[len(str)-1:] + "@" + splitstr[1]
	}

	return ""
}

func IndexOfAtSymbol(email string) int {
	return strings.Index(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	if len(userID) > 0 {
		return strings.TrimSpace(strings.Split(userID, "-")[1])

	}
	return ""
}

func ConvertStringToInt(str string) int {
	val, err := strconv.Atoi(str)

	if err == nil {
		return val
	}

	return 0
}
