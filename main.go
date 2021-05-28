// +build wasm,js

package main

import (
	"github.com/sapphi-red/node-midec/src"
)

func main() {
	var shutdownCh = make(chan struct{}, 0)

	src.Setup()

	<-shutdownCh
}
