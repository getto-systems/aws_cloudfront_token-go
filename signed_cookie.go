package aws_cloudfront_token

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"

	"fmt"
	"time"
)

// AWS CloudFront Signed Cookie Token
type SignedCookieToken struct {
	// CloudFront flavored base64 string
	Policy string
	// CloudFront flavored base64 string
	Signature string
}

type KeyPairPrivateKey []byte

func (privateKey KeyPairPrivateKey) Sign(resource string, expires time.Time) (SignedCookieToken, error) {
	policy := cloudfrontPolicy(resource, expires)

	block, _ := pem.Decode(privateKey)

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return SignedCookieToken{}, err
	}

	rng := rand.Reader

	hashed := sha1.Sum(policy)
	signed, err := rsa.SignPKCS1v15(rng, key, crypto.SHA1, hashed[:])
	if err != nil {
		return SignedCookieToken{}, err
	}

	return SignedCookieToken{
		Policy:    cloudfrontBase64(policy),
		Signature: cloudfrontBase64(signed),
	}, nil
}

func cloudfrontPolicy(resource string, expires time.Time) []byte {
	// AWS CloudFront flavored json : no extra spaces, strict formatted
	return []byte(fmt.Sprintf(
		"{\"Statement\":[{\"Resource\":\"%s\",\"Condition\":{\"DateLessThan\":{\"AWS:EpochTime\":%d}}}]}",
		resource,
		expires.Unix(),
	))
}

func cloudfrontBase64(raw []byte) string {
	// AWS CloudFront flavored base64
	encoded := base64.StdEncoding.EncodeToString(raw)
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "=", "_")
	encoded = strings.ReplaceAll(encoded, "/", "~")
	return encoded
}
