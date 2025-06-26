package responses

import (
	"time"

	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/auth"
)

func User(user *auth.UserResponse) resources.User {
	updatedAt, err := time.Parse(time.RFC3339, user.UpdatedAt)
	if err != nil {
		updatedAt = time.Time{}
	}

	createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	return resources.User{
		Data: resources.UserData{
			Id:   user.Id,
			Type: "user",
			Attributes: resources.UserDataAttributes{
				Email:     user.Email,
				Role:      string(user.Role),
				UpdatedAt: updatedAt,
				CreatedAt: createdAt,
			},
		},
	}
}
