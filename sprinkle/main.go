package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-yaml/yaml"
)

const otherWord = "*"

var transforms []string = []string{}

func main() {
	configTransform()
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}

func configTransform() {
	var v struct {
		Transforms []string
	}
	buf, err := ioutil.ReadFile("./transforms.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := yaml.Unmarshal(buf, &v); err != nil {
		fmt.Println(err)
		return
	}
	transforms = v.Transforms
}
