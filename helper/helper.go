package helper

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var Config jsoniter.API

func init() {
	Config = jsoniter.Config{}.Froze()
	extra.RegisterFuzzyDecoders()
	fmt.Println("init called")
}
