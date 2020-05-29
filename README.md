# aws_cloudfront_token

golang: aws cloudfront token signer

status: production ready

```golang
import (
	"io/ioutil"
	"log"

	"github.com/getto-systems/aws_cloudfront_token-go"
)

privateKey, err := ioutil.ReadFile("path/to/aws_cloudfront_key_pair/pk.pem")
if err != nil {
	log.Fatal(err)
}

token, err := aws_cloudfront_token.Sign(aws_cloudfront_token.Param{
	PrivateKey: privateKey,
	BaseURL: "https://AWS_CLOUDFRONT_BASE_URL",
	Expires: time.Now().Add(time.Duration(30 * 1_000_000_000)), // expires 30 seconds after
})
if err != nil {
	log.Fatal(err)
}

// token.Policy:    aws cloudfront fravor base64 encoded string
// token.Signature: aws cloudfront fravor base64 encoded string
```


###### Table of Contents

- [Requirements](#Requirements)
- [Usage](#Usage)
- [License](#License)

## Requirements

- golang: 1.14


## License

[MIT](LICENSE) license.

Copyright &copy; shun-fix9
