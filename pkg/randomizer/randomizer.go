package randomizer

import (
	"math/rand"
	"time"
)

var phrases = []string{
	"hello",
	"world",
	"go is fun",
}

func init() {

}

func GetRandomPhrase() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return phrases[r.Intn(len(phrases))]
}
