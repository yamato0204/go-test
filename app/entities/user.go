package entities

import "time"

type User struct {
	ID     string 
	Name   string 
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

type UserRes struct {
	ID        string      `json:"id" form:"id" validate:"required"`
	Name      string    `json:"name" form:"name" validate:"required,max=32"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}

func NewUserRes(id string, name string, createdAt time.Time, updatedAt time.Time) *UserRes {
	return &UserRes{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewUser(id string, name string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
