package responses

import (
	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
)

func Cabinet(cabinet *electorcab.Cabinet) resources.Cabinet {
	return resources.Cabinet{
		Data: resources.CabinetData{
			Id:   cabinet.UserId,
			Type: "cabinet", //TODO
			Attributes: resources.CabinetAttributes{
				Profile:   profileAttributes(cabinet.Profile),
				Biography: biographyAttributes(cabinet.Biography),
				JobResume: jobResumeAttributes(cabinet.JobResume),
			},
		},
	}
}
