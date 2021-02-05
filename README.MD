# Chi Ip Banner
Chi ip banner is a chi middleware that bans some ips from your Chi http server.

It reads a `.txt` file in your project's root, called banip.txt, splits that file with linefeeds, and bans those ips. E.g.
```
0.0.0.0
127.0.0.1
```
### Getting started
To use this package, get it:
```sh
go get github.com/dragonDScript/chi-ip-banner
```
And import it and use it:
```go
package main

import "github.com/dragonDScript/chi-ip-banner"

func main(...) {
  r.Use(ipbanner.BanIPs)
}
```