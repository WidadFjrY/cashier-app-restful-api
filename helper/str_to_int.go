package helper

import "strconv"

func StringToInt(userId string) int {
	userIdInt, err := strconv.Atoi(userId)
	PanicIfError(err)

	return userIdInt
}
