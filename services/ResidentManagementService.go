package services

import (
	"RealEstate/models"
	"RealEstate/repositories"
)

type ResidentManagementService interface {
	OnboardResident()
}

type ResidentManagementServiceImpl struct {
}
