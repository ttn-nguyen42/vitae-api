package about

import (
	"Vitae/models"
	"Vitae/repositories/about"
	"Vitae/tools/logging"

	"github.com/jinzhu/copier"
)

type IReader interface {
	GetOne(dto *GetResponse, id string) error
	GetAll(amount int) ([]GetResponse, error)
}

type IWriter interface {
	AddOne(dto PostRequest) (string, error)
}

type IService interface {
	IReader
	IWriter
}

type Service struct {
	repo about.IRepository
}

func NewService(repo about.IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddOne(dto PostRequest) (string, error) {
	var entity models.About
	logging.Trace("DTO at service layer", map[string]interface{}{"dto": dto})
	copier.Copy(&entity, &dto)
	logging.Trace("Entity at service layer", map[string]interface{}{"entity": entity})
	id, err := s.repo.AddOne(entity)
	logging.Trace("ID received at service layer", map[string]interface{}{"id": id})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) GetOne(dto *GetResponse, id string) error {
	var entity models.About
	logging.Trace("ID at service layer", map[string]interface{}{"id": id})
	err := s.repo.GetOne(&entity, id)
	logging.Trace("Entity at service layer", map[string]interface{}{"entity": entity})
	if err != nil {
		return err
	}
	err = copier.Copy(dto, &entity)
    dto.Id = id
	if err != nil {
		return err
	}
	logging.Trace("DTO at service layer", map[string]interface{}{"dto": dto})
	return nil
}

func (s *Service) GetAll(amount int) ([]GetResponse, error) {
	var dtos []GetResponse
	logging.Trace("Amount received in service layer", map[string]interface{}{"amount": amount})
	entities, err := s.repo.GetAll(amount)
	if err != nil {
		return nil, err
	}
	if entities == nil {
		return []GetResponse{}, nil
	}
	logging.Trace("Entities received from data layer", map[string]interface{}{"entities_length": len(entities)})
	err = copier.Copy(&dtos, &entities)
	for at, data := range entities {
		dtos[at].Id = data.Id.Hex()
	}
	logging.Trace("DTOs received from data layer", map[string]interface{}{"dtos_length": len(dtos)})
	if err != nil {
		return nil, err
	}
	return dtos, nil
}
