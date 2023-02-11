package utils

import "math/rand"

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}
