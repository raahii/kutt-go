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



This is a Go wrapper for [the API](https://github.com/thedevs-network/kutt) and CLI tool. To get shorter url, you need to signup at [Kutt.it](https://kutt.it/login) and obtain API key from settings.

```
$ go get -u github.com/raahii/kutt-go
```



## Usage

This is a example to get shorter url. The full example is in `_example` directory. 

For API details, please refer [official repository](https://github.com/thedevs-network/kutt#api).

```go
package main

import (
	"fmt"
	"log"

	"github.com/raahii/kutt-go"
)

func main() {
	cli := kutt.NewClient("<api key>")

	// create shorter url for this repository
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

## CLI

```
$ go get -u github.com/raahii/kutt-go/cmd/kutt
```



You can register your API key to CLI with `kutt apikey <key>`.

```
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



## Licence

Code released under the [MIT License](LICENSE).



