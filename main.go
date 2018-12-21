package randurl

import (
	"fmt"
	"math/rand"
	"strings"
)

type PathComponent interface {
	String() string
}

type URLSpec struct {
	Scheme, Host string
	Components   []PathComponent
}

func (u URLSpec) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%s://%s", u.Scheme, u.Host))
	for _, s := range u.Components {
		b.WriteString(fmt.Sprintf("/%s", s.String()))
	}
	return b.String()
}

type StaticComponent string

func (s StaticComponent) String() string {
	return string(s)
}

const (
	LowercaseAlphabetChars = "abcdefghijklmnopqrstuvwxyz"
	UppercaseAlphabetChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetChars          = LowercaseAlphabetChars + UppercaseAlphabetChars
	DigitChars             = "0123456789"
	PunctuationChars       = ".,-_+!()[]{}*"
)

type RandomComponent struct {
	Chars                []rune
	Random               bool
	MinLength, MaxLength int
}

func (r RandomComponent) String() string {
	var targetLength int
	if r.MaxLength == r.MinLength {
		targetLength = r.MaxLength
	} else {
		targetLength = rand.Intn(r.MaxLength-r.MinLength) + r.MinLength
	}

	randomChars := make([]rune, targetLength)
	for i := 0; i < targetLength; i++ {
		randomChars[i] = r.Chars[rand.Intn(len(r.Chars))]
	}

	return string(randomChars)
}
