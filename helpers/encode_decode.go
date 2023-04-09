package helpers

import (
	"encoding/json"
	"net/http"
)

func DecodeReq(r *http.Request, result any) any {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	} else {
		return decoder.Decode(result)
	}
}

func EncodeRes(w http.ResponseWriter, response any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}
