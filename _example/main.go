package main

import (
	"fmt"
	"log"
	"os"

	"github.com/raahii/kutt-go"
)

func printURL(u *kutt.URL) {
	fmt.Printf("id: %s, target: %s, short_url: %s\n", u.ID, u.Target, u.ShortURL)
}

func main() {
	// new api client with api key
	key := os.Getenv("API_KEY")
	if key == "" {
		fmt.Println("api key is not set.")
		os.Exit(1)
	}
	cli := kutt.NewClient(key)

	// create shorter url
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

	fmt.Println(">> created")
	printURL(URL)

	// list registerd urls
	URLs, err := cli.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n>> all urls")
	for _, u := range URLs {
		printURL(u)
	}

	// delete url
	err = cli.Delete(
		URL.ID,
		// kutt.WithDomain("xxx"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n>> url deleted")
	printURL(URL)
}
