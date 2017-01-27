package pako

import (
	"fmt"
	"io"

	"github.com/gopherjs/gopherjs/js"
)

type Writer struct {
	deflate *js.Object
}

func NewWriter(w io.Writer) *Writer {
	i := js.Global.Get("pako").Get("Deflate").New()
	i.Call("onData", func(chunk *js.Object) {
		buf, ok := js.Global.Get("Uint8Array").New(chunk).Interface().([]byte)
		if !ok {
			panic("chunk isn't a Uint8Array or something")
		}
		_, err := w.Write(buf)
		if err != nil {
			panic("failed to write: " + err.Error())
		}
	})
	return &Writer{
		deflate: i,
	}
}

func (w *Writer) Close() error {
	w.deflate.Call("push", "", true)
	return nil
}

func (w *Writer) Flush() error {
	return nil
}

func (w *Writer) Write(p []byte) (int, error) {
	buf := js.Global.Get("Uint8Array").New(p)
	w.deflate.Call("push", buf, false)
	if errNo := w.deflate.Get("err").Int(); errNo != 0 {
		msg := w.deflate.Get("msg").String()
		return 0, fmt.Errorf("Zlib Error #%d: %s", errNo, msg)
	}
	return len(p), nil
}
