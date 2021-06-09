// +build wasm,js

package main

import (
	"github.com/sapphi-red/node-midec/src"
)

func main() {
	var shutdownCh = make(chan struct{})

	d := 3

	src.Setup()

	<-shutdownCh
}
