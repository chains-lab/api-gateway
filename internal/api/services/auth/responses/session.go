package responses

import (
	"time"

	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/sso"
)

func Session(session *sso.SessionResponse) resources.Session {
	createdAt, err := time.Parse(time.RFC3339, session.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}

	lastUsed, err := time.Parse(time.RFC3339, session.LastUsed)
	if err != nil {
		lastUsed = time.Time{}
	}

	return resources.Session{
		Data: resources.SessionData{
			Type: "session",
			Id:   session.Id,
			Attributes: resources.SessionAttributes{
				UserId:    session.UserId,
				Client:    session.Client,
				Ip:        session.Ip,
				CreatedAt: lastUsed,
				LastUsed:  createdAt,
			},
		},
	}
}
