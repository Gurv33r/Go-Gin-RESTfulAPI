package server

import "strings"

type Date struct {
	Str string `json:"currentDate"`
}

func IsHalloween(request *Date) bool {
	monthAndDay := strings.Split(request.Str, "-")[1:]
	if strings.Compare(monthAndDay[0], "10") == 0 && strings.Compare(monthAndDay[1], "31") == 0 {
		return true
	}
	return false
}
