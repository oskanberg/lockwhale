package lockwhale // import "github.com/oskanberg/lockwhale"

import (
	"bytes"
	"strings"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

type tnHandler struct{}

func (handler tnHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	p := make([]byte, 1)

	var line bytes.Buffer

	for {
		n, err := r.Read(p)
		if err != nil {
			break
		}

		if n == 0 {
			continue
		}

		line.WriteByte(p[0])

		if '\n' == p[0] && strings.HasSuffix(line.String(), "\r\n") {
			oi.LongWrite(w, []byte("eol"))
		}

	}
}

func Start() {
	var handler telnet.Handler = tnHandler{}

	err := telnet.ListenAndServe(":5555", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}
