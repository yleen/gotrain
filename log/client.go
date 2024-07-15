package log

import (
	"bytes"
	"fmt"
	stlog "log"
	"net/http"

	"github.com/allen/Go/registry"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

// // Write implements io.Writer.
// func (cl *clientLogger) Write(p []byte) (n int, err error) {
// 	panic("unimplemented")
// }

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to send log message. Service responded write error: %s", res.StatusCode)
	}
	return len(data), nil
}
