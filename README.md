Kutt.it API Client for Go
--

[Kutt.it](https://kutt.it/) is a **Modern Open Source URL shortener.** 

- Custom domain
- Password for the URL
- Managing links
- Free & Open Source
- **50** URLs shortening per day.



This is a a Go wrapper for [the API](https://github.com/thedevs-network/kutt). 

To get shorter url,  you need to signup at [Kutt.it](https://kutt.it/login) and obtain API key from settings.

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
	URL, err := cli.Submit(&kutt.SubmitInput{
		URL: "https://github.com/raahii/kutt-go",
		// CustomURL: "kutt-go",
		// Password:  "foobar",
		// Reuse: true,
	})
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



## Licence

Code released under the [MIT License](LICENSE).



