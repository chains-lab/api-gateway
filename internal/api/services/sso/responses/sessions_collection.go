package responses

import (
	"github.com/chains-lab/api-gateway/resources"
	pb "github.com/chains-lab/proto-storage/gen/go/svc/sso"
)

func SessionCollection(sessions *pb.SessionsList) *resources.SessionsCollection {
	var data []resources.SessionData
	for _, session := range sessions.Sessions {
		// Session — ваша функция-конвертер из pb.SessionResponse → ресурсы
		data = append(data, Session(session).Data)
	}
	return &resources.SessionsCollection{
		Data: resources.SessionsCollectionData{
			Type: "sessions",
			Attributes: resources.SessionsCollectionDataAttributes{
				Sessions: data,
			},
		},
	}
}
