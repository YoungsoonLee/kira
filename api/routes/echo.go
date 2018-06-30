package routes

import (
	"net/http"

	"github.com/YoungsoonLee/kira/api/utils"
)

// Echo ...
// for healty check
func Echo(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, nil, "okay")
}
