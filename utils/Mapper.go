package utils

import (
	"RealEstate/dtos/requests"
	"RealEstate/dtos/responses"
	"RealEstate/models"
)

func MapOnboardResidentRequestToResident(request requests.OnboardResidentRequest) models.Resident {
	return models.Resident{
		Name:         request.Name,
		Email:        request.Email,
		PhoneNumber:  request.PhoneNumber,
		HouseAddress: request.Address,
	}
}

func MapResidentToOnboardResidentResponse(resident models.Resident) responses.OnboardResidentResponse {
	return responses.OnboardResidentResponse{
		ResidentID:     resident.ID.Hex(),
		ResidentName:   resident.Name,
		DateRegistered: resident.DateRegistered.String(),
	}
}
