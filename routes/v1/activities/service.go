package activities

import (
	"Vitae/models"
	"Vitae/repositories"
	"Vitae/repositories/about"
	"Vitae/repositories/activities"
	"Vitae/tools/logging"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IReader interface {
	GetOne(dto *GetResponse, id string) error
	GetAll(userId string, amount int) ([]GetResponse, error)
}

type IWriter interface {
	AddOne(userId string, dto PostRequest) (string, error)
}

type IService interface {
	IReader
	IWriter
}

type Service struct {
	repo     activities.IRepository
	userRepo about.IRepository
}

func NewService(userRepo about.IRepository, activitiesRepo activities.IRepository) *Service {
	return &Service{
		repo:     activitiesRepo,
		userRepo: userRepo,
	}
}

func (s *Service) AddOne(userId string, dto PostRequest) (string, error) {
	ok, err := s.userRepo.Exists(userId)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", repositories.NewNotFoundError("User Id is incorrect")
	}
	parsedId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", repositories.NewInvalidIdError("User Id is not correctly format")
	}
	var entity models.Activity
	logging.Trace("DTO at service layer", map[string]interface{}{"dto": dto})
	copier.Copy(&entity, &dto)
	entity.UserId = parsedId
	logging.Trace("Entity at service layer", map[string]interface{}{"entity": entity})
	id, err := s.repo.AddOne(entity)
	logging.Trace("ID received at service layer", map[string]interface{}{"id": id})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) GetOne(dto *GetResponse, id string) error {
	var entity models.Activity
	logging.Trace("ID at service layer", map[string]interface{}{"id": id})
	err := s.repo.GetOne(&entity, id)
	logging.Trace("Entity at service layer", map[string]interface{}{"entity": entity})
	if err != nil {
		return err
	}
	err = copier.Copy(dto, &entity)
	if err != nil {
		return err
	}
	logging.Trace("DTO at service layer", map[string]interface{}{"dto": dto})
	return nil
}

func (s *Service) GetAll(userId string, amount int) ([]GetResponse, error) {
	ok, err := s.userRepo.Exists(userId)
	if err != nil {
		return []GetResponse{}, err
	}
	if !ok {
		return []GetResponse{}, repositories.NewNotFoundError("User Id is incorrect")
	}
	var dtos []GetResponse
	logging.Trace("Amount received in service layer", map[string]interface{}{"amount": amount})
	entities, err := s.repo.GetAll(userId, amount)
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
