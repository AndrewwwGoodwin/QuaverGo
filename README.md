# QuaverGo

This is an API wrapper for the rhythm game, [Quaver](https://quavergame.com/). 

The project is VERY incomplete and is not suitable for use in anything
## Usage
`go get github.com/AndrewwwGoodwin/QuaverGo`

Initialize a client, and get some info

```go
package main

import (
	"fmt"
	"github.com/AndrewwwGoodwin/QuaverGo"
)

func main() {
	client := QuaverGo.Init()
	data, err := client.Users.ID(371737)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
```

### Rate Limiting
As per the [docs](https://wiki.quavergame.com/docs/api-v2), the rate limit is currently set at 100 req/min

This is enforced via errors, and controlled via the RateLimitManager

### Aside
Better docs coming soon! The project is far from done!