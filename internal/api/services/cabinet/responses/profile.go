package responses

import (
	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
)

func Profile(profile *electorcab.Profile) resources.ProfileData {
	return resources.ProfileData{
		Id:         profile.UserId,
		Type:       "profile",
		Attributes: profileAttributes(profile),
	}
}

func profileAttributes(profile *electorcab.Profile) resources.ProfileAttributes {
	return resources.ProfileAttributes{
		Username:    profile.Username,
		Pseudonym:   profile.Pseudonym,
		Description: profile.Description,
		Avatar:      profile.Avatar,
		Official:    profile.Official,

		UpdatedAt: *parseAndFormat(&profile.UpdatedAt),
		CreatedAt: *parseAndFormat(&profile.CreatedAt),
	}
}
