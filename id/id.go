package id

import "regexp"

type Email string

func (e *Email) Validate() bool {
	return regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`).MatchString(string(*e))
}

// type Keypair struct {
// 	Pubkey  []byte
// 	Privkey []byte
// }
