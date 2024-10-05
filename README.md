# QuaverGo

This is an api wrapper for the rhythm game Quaver. 

## Usage
`go get github.com/AndrewwwGoodwin/quaverGo`

Initialize a client, and get some info

```go
package main

import (
	"fmt"
	"github.com/AndrewwwGoodwin/quaverGo"
)

func main() {
	quaverClient := quaverGo.Init()
	user := quaverClient.Users.ID(371737)
	fmt.Println(user.Username)
}
```

### Rate Limiting
as per the [docs](https://wiki.quavergame.com/docs/api-v2), the rate limit is set at 100 req/min

This is enforced via errors, and controlled bia the RateLimitManager

