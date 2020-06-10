# aws_cloudfront_token

golang: aws cloudfront token signer

status: production ready

```golang
import (
	"io/ioutil"
	"log"
	"time"

	"github.com/getto-systems/aws_cloudfront_token-go"
)

pem, err := ioutil.ReadFile("path/to/aws_cloudfront_key_pair/pk.pem")
if err != nil {
	log.Fatal(err)
}

privateKey := aws_cloudfront_token.KeyPairPrivateKey(pem)

resource := "https://AWS_CLOUDFRONT_BASE_URL/*"
expires: := time.Now().Add(time.Duration(30 * 1_000_000_000)), // expires 30 seconds after

token, err := privateKey.Sign(resource, expires)
if err != nil {
	log.Fatal(err)
}

// token.Policy    // aws cloudfront fravor base64 encoded string
// token.Signature // aws cloudfront fravor base64 encoded string
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
