# NFT Sale Proxy

A proxy to hide NFT metadata during the sale and prevent people from sniping specific NFTs.

Examples: [alephao/nft-sale-proxy-examples](https://github.com/alephao/nft-sale-proxy-examples)

### Getting Started

#### AWS Lambda

To create the proxy using AWS Lambda is very simple, you just need a `go` file with the code below:

```go
package main

import (
	proxy "github.com/alephao/nft-sale-proxy/aws-lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(proxy.HandleRequest)
}
```

Then you need to make the environment variables listed below available during runtime.

### Configuration

The proxy is configured using environment variables:

|||
|-|-|
| `BASE_URL` | The baseURL that contains the actual token metadata |
| `INCOGNITO_IMAGE_URL` | The URL to the image that will show for non-revealed tokens |
| `INCOGNITO_NAME` | The `name` attribute that will show for non-revealed tokens. You can use the placeholder `{id}` and it will be replaced by the token id. |
| `INCOGNITO_DESCRIPTION` | The `description` attribute that will show for non-revealed tokens |
| `INCOGNITO_EXTERNAL_LINK` | The `external_link` attribute that will show for non-revealed tokens |
| `NUMBER_OF_TOKENS` | The maximum amount of tokens. The proxy will return 404 for incoming requests with a number highe than this value |
| `REVEAL_UP_TO` | The highest token id that will be revealed. Start with `-1`, to reveal none. To reveal the first `1000`, change to `999`. Etc. |
| `ERC1155` | Set this to `true` if the token is an ERC1155. This will use 32 bytes hex values padded with `0`s as the id. |