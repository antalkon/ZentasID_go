package randstring

import "crypto/rand"

func GenerateRandomFixedString(l int) string {
	chars := "qwertyuiopasdfghjklzxcvbnm"
	ll := len(chars)
	b := make([]byte, l)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < l; i++ {
		b[i] = chars[int(b[i])%ll]
	}

	return string(b)
}
