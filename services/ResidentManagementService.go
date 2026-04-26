package services

import (
	"RealEstate/dtos/requests"
	"RealEstate/dtos/responses"
	"RealEstate/repositories"
	"RealEstate/utils"
	"errors"
)

type ResidentService struct {
	residentRepository repositories.ResidentRepository
}

func NewResidentManagementService(r repositories.ResidentRepository) *ResidentService {
	return &ResidentService{residentRepository: r}
}

func (rs *ResidentService) OnboardNewResident(request requests.OnboardResidentRequest) (responses.OnboardResidentResponse, error) {
	_, err := rs.residentRepository.FindByEmail(request.Email)

	if err == nil {
		return responses.OnboardResidentResponse{}, errors.New("resident already exists")
	}

	resident := utils.MapOnboardResidentRequestToResident(request)

	savedUser, err := rs.residentRepository.Save(resident)
	if err != nil {
		return responses.OnboardResidentResponse{}, err
	}

	return utils.MapResidentToOnboardResidentResponse(savedUser), nil

}
