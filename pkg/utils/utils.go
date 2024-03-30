package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// func ParseBody2(r *http.Request, x interface{}) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println(err)
// 		x = nil
// 		return
// 	}
// 	if err := json.Unmarshal(body, x); err != nil {
// 		log.Println(err)
// 		x = nil
// 	}
// }

func ParseBody(r *http.Request, x interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
