package repositories

import "RealEstate/models"

type GatePassRepository interface {
	Save(gatePass models.GatePass) (models.GatePass, error)
	FindAll() ([]models.GatePass, error)
	FindById(gatePassId string) (models.GatePass, error)
	FindByResidentId(residentId string) (models.GatePass, error)
}
