package toolkit

import "crypto/rand"

const randomeStringSource = "abcdefghijklmnopqrstuxvwzABCDEFGHIJKLMNOPQRSTUWXYZ1234567890_+"

// Tools is tye type used to instantiate this module.
//Any variable of this type will have access to all the methods with the reciever *Tools.
type Tools struct{}

// RandomString returns a string of random characters of length n, using randomeStringSource as the source for the string.
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomeStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x&y]
	}

	return string(s)
}