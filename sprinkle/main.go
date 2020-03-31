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

var defaultTransform []string = []string{
	"*",
	"go ",
	"app",
	"*site",
	"*time",
}

var transforms []string = []string{}

func main() {
	if err := configTransform(); err != nil {
		transforms = defaultTransform
	}
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}

func configTransform() error {
	var v struct {
		Transforms []string
	}
	buf, err := ioutil.ReadFile("./sprinkle/transforms.yml")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := yaml.Unmarshal(buf, &v); err != nil {
		fmt.Println(err)
		return err
	}
	transforms = v.Transforms
	return nil
}
