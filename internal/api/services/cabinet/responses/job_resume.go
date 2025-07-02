package responses

import (
	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
)

func JobResume(jobResume *electorcab.JobResume) resources.JobResume {
	return resources.JobResume{
		Data: resources.JobResumeData{
			Id:         jobResume.UserId,
			Type:       "job_resume",
			Attributes: jobResumeAttributes(jobResume),
		},
	}
}

func jobResumeAttributes(jobResume *electorcab.JobResume) resources.JobResumeAttributes {
	return resources.JobResumeAttributes{
		Degree:   jobResume.Degree,
		Industry: jobResume.Industry,
		Income:   jobResume.Income,

		DegreeUpdatedAt:   parseAndFormat(jobResume.DegreeUpdatedAt),
		IndustryUpdatedAt: parseAndFormat(jobResume.IndustryUpdatedAt),
		IncomeUpdatedAt:   parseAndFormat(jobResume.IncomeUpdatedAt),
	}
}
