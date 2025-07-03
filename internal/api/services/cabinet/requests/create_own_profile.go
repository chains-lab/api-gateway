package requests

import (
	"encoding/json"
	"net/http"

	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/gatekit/jsonkit"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func NewCreateOwnProfile(r *http.Request) (req resources.CreateOwnProfile, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = jsonkit.NewDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In("create_own_profile")), //TODO
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	return req, errs.Filter()
}
