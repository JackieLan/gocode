package main

import (
	_ "github.com/JackieLan/gocode/chapter2/sample/matchers"
	"github.com/JackieLan/gocode/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
