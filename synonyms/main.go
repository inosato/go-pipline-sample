package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/inosato/go-pipline-sample/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{
		APIKey: apiKey,
	}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonims(word)
		if err != nil {
			log.Fatalf("cannot search %q : %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("no synonims : %v\n", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
