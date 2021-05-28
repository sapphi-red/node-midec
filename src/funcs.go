package src


import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/sapphi-red/midec"
	_ "github.com/sapphi-red/midec/gif"
	_ "github.com/sapphi-red/midec/png"
	_ "github.com/sapphi-red/midec/webp"
	_ "github.com/sapphi-red/midec/isobmff"
)

func Setup() {
	js.Global().Set("_p", js.FuncOf(p))
	js.Global().Set("_transfer", js.FuncOf(transfer))
	js.Global().Set("_isAnimated", js.FuncOf(isAnimated))
}

func p(this js.Value, args []js.Value) interface{} {
	go func() {
		fmt.Println(args[0].String())
		args[1].Invoke()
	}()
	return nil
}

func transfer(this js.Value, args []js.Value) interface{} {
	go func() {
		bts := make([]byte, args[0].Length())
		js.CopyBytesToGo(bts, args[0])

		fmt.Println(bts)

		args[1].Invoke(nil, len(bts))
	}()
	return nil
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
