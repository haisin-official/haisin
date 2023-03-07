package utils

import "math/rand"

func GenSlug() (string, error) {
	return randomString(32), nil
}

func randomString(n int) string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	r := make([]rune, n)
	for i := range r {
		r[i] = runes[rand.Intn(len(runes))]
	}
	return string(r)
}
