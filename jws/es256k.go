package jws

import (
	"github.com/lestrrat-go/jwx/v2/jwa"
)

func init() {
	addAlgorithmForKeyType(jwa.EC, jwa.ES256K)
}
