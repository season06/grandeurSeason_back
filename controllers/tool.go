package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Panic("responseWithJson error: " + err.Error())
	}
}

// if http method is GET, split parameter from URL
func getParameter(URL url.Values, key string, _type string) interface{} {
	param, ok := URL[key]
	if !ok {
		return ""
	} else {
		switch _type {
		case "_string":
			return string(param[0])
		case "_int":
			result, _ := strconv.Atoi(param[0])
			return result
		case "_uint64":
			result, _ := strconv.ParseUint(param[0], 10, 64)
			return result
		case "time":
			format := "2006-01-02T15:04:05.000Z"
			result, _ := time.Parse(format, param[0])
			return result
		}
	}
	return ""
}
