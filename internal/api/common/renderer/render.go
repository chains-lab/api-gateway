package renderer

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
)

func Render(w http.ResponseWriter, res interface{}) {
	httpkit.Render(w, res)
}
