package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tenlisboa/pulley-challenge/pkg/challenge"
	"github.com/tenlisboa/pulley-challenge/pkg/parser"
	"github.com/tenlisboa/pulley-challenge/pkg/pulley"
)

var ctx context.Context

const (
	url string = "https://ciphersprint.pulley.com/%s"
	challenger string = "gabriellisboa.rx@gmail.com"
)

func init() {
	ctx = context.Background()
}

func f(str, change string) string {
	return fmt.Sprintf(str, change)
}

func main() {
	c := pulley.NewClient(&http.Client{})	

	body := c.Request(f(url, challenger))
	ch := parser.DecodeTo[challenge.Challenge](body)

	for len(body) != 0 {
		nhash := challenge.Decode(ch.EncryptionMethod, ch.EncryptedPath[5:])
		body = c.Request(f(url, f("task_%s", nhash)))
		ch = parser.DecodeTo[challenge.Challenge](body)
		fmt.Printf("Level: %d\n", ch.Level)
		if ch.Level == 6 {
			break
		}
	}

	fmt.Println("Done!")
}
