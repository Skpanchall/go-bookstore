package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Println(body)
		err := json.Unmarshal(body, x)
		if err != nil {
			fmt.Println(err)
		}
		return

	}

}
