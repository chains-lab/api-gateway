package responses

import (
	"time"

	"github.com/chains-lab/api-gateway/resources"
	"github.com/chains-lab/proto-storage/gen/go/svc/electorcab"
)

func Biography(biography *electorcab.Biography) resources.Biography {
	return resources.Biography{
		Data: resources.BiographyData{
			Id:         biography.UserId,
			Type:       "biography",
			Attributes: biographyAttributes(biography),
		},
	}
}

func parseAndFormat(dateStr *string) *time.Time {
	if dateStr == nil {
		return nil
	}

	parsedTime, err := time.Parse(time.RFC3339, *dateStr)
	if err != nil {
		return nil
	}

	res := parsedTime.UTC()

	return &res
}

func biographyAttributes(biography *electorcab.Biography) resources.BiographyAttributes {
	return resources.BiographyAttributes{
		Birthday:        parseAndFormat(biography.Birthday),
		Sex:             biography.Sex,
		City:            biography.City,
		Region:          biography.Region,
		Country:         biography.Country,
		Nationality:     biography.Nationality,
		PrimaryLanguage: biography.PrimaryLanguage,

		SexUpdatedAt:             parseAndFormat(biography.SexUpdatedAt),
		NationalityUpdatedAt:     parseAndFormat(biography.NationalityUpdatedAt),
		PrimaryLanguageUpdatedAt: parseAndFormat(biography.PrimaryLanguageUpdatedAt),
		ResidenceUpdatedAt:       parseAndFormat(biography.ResidenceUpdatedAt),
	}
}
