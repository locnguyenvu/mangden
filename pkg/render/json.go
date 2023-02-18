package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (r *Renderer) RenderJSON(w http.ResponseWriter, code int, data interface{}) {
	if !r.AllowedResponseCode(code) {
		r.logger.WithField("code", code).Errorln("unregistered response code")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := escapeJSON(fmt.Sprintf("%d is not a registered response code", code))
		fmt.Fprintf(w, jsonErrTmpl, msg)
		return
	}

	if data == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)

		// Return an OK response.
		if code >= http.StatusOK && code < http.StatusMultipleChoices {
			fmt.Fprint(w, jsonOKResp)
			return
		}

		// Return an error with the generic HTTP text as the error.
		msg := escapeJSON(http.StatusText(code))
		fmt.Fprintf(w, jsonErrTmpl, msg)
		return
	}

	// Special-case handle multi-errors.
	if typ, ok := data.([]error); ok {
		msgs := make([]string, 0, len(typ))
		for _, err := range typ {
			msgs = append(msgs, err.Error())
		}
		data = &multiError{Errors: msgs}
	}

	// If the provided value was an error, marshall accordingly.
	if typ, ok := data.(error); ok {
		data = &singleError{Error: typ.Error()}
	}

	// Acquire a renderer
	b := r.rendererPool.Get().(*bytes.Buffer)
	b.Reset()
	defer r.rendererPool.Put(b)

	// Render into the renderer
	if err := json.NewEncoder(b).Encode(data); err != nil {
		r.logger.WithField("error", err).Error("failed to marshal json")

		msg := "An internal error occurred."
		msg = escapeJSON(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, jsonErrTmpl, msg)
		return
	}

	// Rendering worked, flush to the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := b.WriteTo(w); err != nil {
		// We couldn't write the buffer. We can't change the response header or
		// content type if we got this far, so the best option we have is to log the
		// error.
		r.logger.WithField("error", err).Error("failed to write json to response")
	}
}

// RenderJSON500 renders the given error as JSON. In production mode, this always
// renders a generic "server error" message. In debug, it returns the actual
// error from the caller.
func (r *Renderer) RenderJSON500(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError

	r.RenderJSON(w, code, map[string]string{
		"error": http.StatusText(code),
	})
}

// escapeJSON does primitive JSON escaping.
func escapeJSON(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}

// jsonErrTmpl is the template to use when returning a JSON error. It is
// rendered using Printf, not json.Encode, so values must be escaped by the
// caller.
const jsonErrTmpl = `{"error":"%s"}`

// jsonOKResp is the return value for empty data responses.
const jsonOKResp = `{"ok":true}`

type singleError struct {
	Error string `json:"error,omitempty"`
}

type multiError struct {
	Errors []string `json:"errors,omitempty"`
}
