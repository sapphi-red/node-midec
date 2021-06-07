package src

import (
	"bytes"
	"syscall/js"

	"github.com/sapphi-red/midec"
	_ "github.com/sapphi-red/midec/gif"
	_ "github.com/sapphi-red/midec/isobmff"
	_ "github.com/sapphi-red/midec/png"
	_ "github.com/sapphi-red/midec/webp"
)

func Setup() {
	js.Global().Set("_isAnimated", js.FuncOf(isAnimated))
}

func isAnimated(this js.Value, args []js.Value) interface{} {
	go func() {
		bts := make([]byte, args[0].Length())
		js.CopyBytesToGo(bts, args[0])

		r := bytes.NewReader(bts)
		res, err := midec.IsAnimated(r)
		if err != nil {
			args[1].Invoke(err.Error(), res)
		} else {
			args[1].Invoke(nil, res)
		}
	}()
	return nil
}
