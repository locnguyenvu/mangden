package httpresponse

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"google.golang.org/protobuf/proto"
)

const (
	keyResponseAccept      = "Accept"
	keyResponseContentType = "Content-Type"
)

type jsonMarshal struct{}

type protobufMarshal struct{}

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

func (st *jsonMarshal) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (st *protobufMarshal) Marshal(v interface{}) ([]byte, error) {
	message, ok := v.(proto.Message)
	if !ok {
		return nil, errors.New("Wrong input for protobuf message")
	}
	return proto.Marshal(message)
}

var responseByContentType = map[string]Marshaler{
	"application/json":                &jsonMarshal{},
	"application/vnd.google.protobuf": &protobufMarshal{},
}

func Send(w http.ResponseWriter, r *http.Request, response interface{}, httpStatus int) {
	var responseBody []byte
	var err error
	accecptedType := r.Header.Get(keyResponseAccept)

	if marshaler, ok := responseByContentType[accecptedType]; ok {
		w.Header().Set(keyResponseContentType, accecptedType)
		responseBody, err = marshaler.Marshal(response)
	} else {
		marshaler = &jsonMarshal{}
		w.Header().Set(keyResponseContentType, "application/json")
		responseBody, err = marshaler.Marshal(response)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(responseBody)
}

func SendSuccess(w http.ResponseWriter, r *http.Request, response interface{}) {
	Send(w, r, response, 200)
}
