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

	// ------------------
	// create shorter url
	// ------------------
	URL, err := cli.Submit(&kutt.SubmitInput{
		URL: "https://github.com/raahii/kutt-go",
		// CustomURL: "kutt-go",
		// Password:  "foobar",
		// Reuse: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">> created")
	printURL(URL)

	// -------------------
	// list registerd urls
	// -------------------
	URLs, err := cli.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n>> all urls")
	for _, u := range URLs {
		printURL(u)
	}

	// -------------------
	// delete url
	// -------------------
	err = cli.Delete(&kutt.DeleteInput{
		ID: URL.ID,
		// Domain: "xxx",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n>> url deleted")
	printURL(URL)
}
