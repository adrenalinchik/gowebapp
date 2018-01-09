package controller

import (
	"net/http"
	"fmt"
	"strconv"
)

// parseID takes id variable from the given request url,
// parses the obtained text and returns the result like int64.
// Used in most controllers.
func parseID(r *http.Request) (id int64, err error) {
	str := r.FormValue("id")
	if str == "" {
		err = fmt.Errorf("id field is empty")
		return
	}
	id, err = strconv.ParseInt(str, 10, 0)
	if err != nil {
		err = fmt.Errorf("not a number: " + str)
		return
	}
	return
}
