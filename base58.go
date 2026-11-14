// Package gobase58 encodes and decodes base58 with selectable alphabets.
package gobase58

import (
	"errors"
	"math/big"
)

// Bitcoin is the alphabet used by bitcoin addresses.
const Bitcoin = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Ripple skips zero, capital O, capital I and lowercase l.
const Ripple = "rpshnaf39wBUDciYbTAZb2rK2QS3iK2j7YQ"

var ErrInvalid = errors.New("base58: invalid character")

func index(alphabet string) map[rune]int {
	out := make(map[rune]int, len(alphabet))
	for i, r := range alphabet {
		out[r] = i
	}
	return out
}

// Encode converts data to base58 text.
func Encode(data []byte, alphabet string) string {
	if alphabet == "" {
		alphabet = Bitcoin
	}
	base := big.NewInt(int64(len(alphabet)))
	value := new(big.Int).SetBytes(data)
	mod := new(big.Int)
	var out []byte
	for value.Sign() > 0 {
		value.DivMod(value, base, mod)
		out = append(out, alphabet[mod.Int64()])
	}
	for _, b := range data {
		if b != 0 {
			break
		}
		out = append(out, alphabet[0])
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}

// Decode parses base58 text back to bytes.
func Decode(text, alphabet string) ([]byte, error) {
	if alphabet == "" {
		alphabet = Bitcoin
	}
	lookup := index(alphabet)
	base := big.NewInt(int64(len(alphabet)))
	value := big.NewInt(0)
	for _, r := range text {
		digit, ok := lookup[r]
		if !ok {
			return nil, ErrInvalid
		}
		value.Mul(value, base)
		value.Add(value, big.NewInt(int64(digit)))
	}
	leading := 0
	for leading < len(text) && text[leading] == alphabet[0] {
		leading++
	}
	body := value.Bytes()
	return append(make([]byte, leading), body...), nil
}
