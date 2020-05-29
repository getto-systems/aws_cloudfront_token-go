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

type Token struct {
	Policy    string
	Signature string
}

type Param struct {
	PrivateKey []byte
	BaseURL    string
	Expires    time.Time
}

func Sign(param Param) (Token, error) {
	var nullToken Token

	policy := cloudfrontPolicy(param.BaseURL, param.Expires)

	block, _ := pem.Decode(param.PrivateKey)

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nullToken, err
	}

	rng := rand.Reader

	hashed := sha1.Sum(policy)
	signed, err := rsa.SignPKCS1v15(rng, key, crypto.SHA1, hashed[:])
	if err != nil {
		return nullToken, err
	}

	return Token{
		Policy:    cloudfrontBase64(policy),
		Signature: cloudfrontBase64(signed),
	}, nil
}

func cloudfrontPolicy(baseURL string, expires time.Time) []byte {
	// AWS CloudFront flavored json : no extra spaces, strict formatted
	return []byte(fmt.Sprintf(
		"{\"Statement\":[{\"Resource\":\"%s\",\"Condition\":{\"DateLessThan\":{\"AWS:EpochTime\":%d}}}]}",
		baseURL,
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