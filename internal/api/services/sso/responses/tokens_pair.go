package responses

import (
	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/sso"
)

func TokensPair(pair *sso.TokensPairResponse) *resources.TokensPair {
	return &resources.TokensPair{
		Data: resources.TokensPairData{
			Type: "tokens_pair",
			Attributes: resources.TokensPairDataAttributes{
				AccessToken:  pair.AccessToken,
				RefreshToken: pair.RefreshToken,
			},
		},
	}
}
