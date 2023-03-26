package helpers

import "strconv"

func StrToInt(str string) (int, error) {
	Int, err := strconv.Atoi(str)
	return Int, err
}
