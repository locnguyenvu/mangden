package render

import (
    "bytes"
    "net/http"
    "sync"

	"github.com/sirupsen/logrus"
)

var allowedResponseCodes = map[int]struct{}{
	http.StatusOK:                    {},
	http.StatusBadRequest:            {},
	http.StatusUnauthorized:          {},
	http.StatusNotFound:              {},
	http.StatusMethodNotAllowed:      {},
	http.StatusConflict:              {},
	http.StatusPreconditionFailed:    {},
	http.StatusRequestEntityTooLarge: {},
	http.StatusTooManyRequests:       {},
	http.StatusInternalServerError:   {},
}

type Renderer struct {
    logger logrus.FieldLogger
	rendererPool *sync.Pool
}

func NewRenderer(logger logrus.FieldLogger) *Renderer {
    return &Renderer{
        logger: logger,
        rendererPool: &sync.Pool{
            New: func() interface{} {
                return bytes.NewBuffer(make([]byte, 0, 1024))
            },
        },

    }
}


func (r *Renderer) AllowedResponseCode(code int) bool {
	_, ok := allowedResponseCodes[code]
	return ok
}
