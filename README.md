<div align="center">
  <img width="150" src="https://imgur.com/dk0StSP.png" />
</div>
<h1 align="center">kutt-go</h1>
<p align="center">
  <a href="https://kutt.it">Kutt.it</a> API Client for Go and CLI tool
</p>
<div align="center"></div>

[Kutt.it](https://kutt.it/) is a **Modern Open Source URL shortener.** 

- Custom domain
- Password for the URL
- Managing links
- Free & Open Source
- **50** URLs shortening per day.



This repo contains a CLI and golang package for the service.



## CLI

### Installation

```shell
$ go install github.com/raahii/kutt-go/cmd/kutt@latest
```



### Usage

```sh
$ kutt register <your api key>
$ kutt submit <url to shorten>
https://kutt.it/ztPDmK
```

```sh
❯ kutt --help
CLI tool for Kutt.it (URL Shortener)

Usage:
  kutt [command]

Available Commands:
  apikey      Register your api key to cli
  delete      Delete a shorted link (Give me url id or url shorted)
  list        List of last 5 URL objects.
  submit      Submit a new short URL
  help        Help about any command

Flags:
  -k, --apikey string   api key for Kutt.it
  -h, --help            help for kutt

Use "kutt [command] --help" for more information about a command.
```



## Go Package

### Installation

```sh
$ go get -u github.com/raahii/kutt-go
```



### Example

This is a example to get shorter url. See also codes in `_example` directory. 

```go
package main

import (
	"fmt"
	"log"

	"github.com/raahii/kutt-go"
)

func main() {
	cli := kutt.NewClient("<api key>")

	// create a short url
	target := "https://github.com/raahii/kutt-go"
	URL, err := cli.Submit(
		target,
		// kutt.WithCustomURL("kutt-go"),
		// kutt.WithPassword("foobar"),
		// kutt.WithReuse(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(URL.ShortURL) // https://kutt.it/kutt-go
}
```

```go
type URL struct {
	ID         string    `json:"id"`
	Target     string    `json:"target"`
	ShortURL   string    `json:"shortUrl"`
	Password   bool      `json:"password"`
	Reuse      bool      `json:"reuse"`
	DomainID   string    `json:"domain_id"`
	VisitCount int       `json:"visit_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
```



## Reference

- [thedevs-network/kutt: Free Modern URL Shortener.](https://github.com/thedevs-network/kutt#api)
- [Kutt API v2 documentation](https://docs.kutt.it/)



## Licence

Code released under the [MIT License](LICENSE).



