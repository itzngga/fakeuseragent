## Fake User Agent
`fakeuseragent` is a Go package that provides fake user agent strings for popular web browsers. It can be used for testing, web scraping, or any other scenario where you need to simulate different user agents.

## Usage
The package provides two main functions for retrieving user agents:

- RandomUserAgent: Retrieves a random user agent.
- MobileUserAgent: Retrieves a random user agent from mobile devices.
- DesktopUserAgent: Retrieves a random user agent from popular Desktops.
- RandomUserAgent: Retrieves a random user agent from the supported browsers.

Here's an example of how to use the package:

```go
package main

import (
	"fmt"

	"github.com/itzngga/fakeuseragent"
)

func main() {
	// Get a random user agent for Desktop
	desktopUserAgent := fakeuseragent.DesktopUserAgent()
	fmt.Println("Desktop User Agent:", desktopUserAgent)
}
```

## Acknowledgements

This package was ported from [rand_user_agent](https://github.com/xyanyue/rand_user_agent).