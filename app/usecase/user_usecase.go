package usecase

import (
	

	"github.com/yamato0204/go-test.git/app/entities"
	"github.com/yamato0204/go-test.git/app/infra"
)


type UserUsecase struct{

	i  infra.IUserInfra
}

type IUserUsecase interface {

	FindByID(id string) (*entities.UserRes, error)
	Create(user *entities.User) (entities.UserRes, error)
	// Update(user *entities.User) error
	// Delete(id string) error


}

//to Res
func NewUserUsecase(i infra.IUserInfra) IUserUsecase{
	return &UserUsecase{i}
}



func (s *UserUsecase)convertTo(user *entities.User) *entities.UserRes {
	return entities.NewUserRes(user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

}


func (s *UserUsecase) FindByID(id string) (*entities.UserRes, error) {
	user, err := s.i.FindByID(id)

	if err != nil {
		return nil, err
	}

	return s.convertTo(user), nil
}

func (s *UserUsecase) Create(user *entities.User) (entities.UserRes, error) {


	if err := s.i.Create(user); err != nil{
		return entities.UserRes{}, err
	}
	
	resUser := entities.UserRes{
	    ID:    user.ID,            
	    Name:  user.Name,    
	    CreatedAt: user.CreatedAt,
	    UpdatedAt: user.UpdatedAt, 
	}

	return resUser, nil

}