package mock_infra

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yamato0204/go-test.git/app/entities"
	"github.com/yamato0204/go-test.git/app/usecase"
)


func TestCreate(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// 
	
	expected := &entities.User{
		ID: "1",
		Name: "ken",
	}
	
	var err error
 	
	mockSample := NewMockIUserInfra(ctrl)
	mockSample.EXPECT().Create(expected).Return(err)

	UserUsecase := usecase.NewUserUsecase(mockSample)
	_, err = UserUsecase.Create(expected)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}



