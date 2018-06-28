package routes

import (
	"net/http"

	"github.com/YoungsoonLee/mrecun/api/utils"
)

// Echo ...
// for healty check
func Echo(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, "okay")
}
