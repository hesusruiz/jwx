package jwa

const Secp256k1 EllipticCurveAlgorithm = "secp256k1"

func init() {
	allEllipticCurveAlgorithms[Secp256k1] = struct{}{}
}
