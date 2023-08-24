package infra

import (
	"database/sql"

	"github.com/yamato0204/go-test.git/app/entities"
)


type UserInfra struct {

	db *sql.DB
}

type IUserInfra interface {

	FindByID(id string) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id string) error

}



func NewUserInfra(db *sql.DB) IUserInfra {
	return &UserInfra{db}
}


func (i *UserInfra) FindByID(id string) (*entities.User, error){
	stmt, err := i.db.Prepare(`SELECT id, name, created_at, updated_at FROM users WHERE id = ?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	user := &entities.User{}
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (i *UserInfra) Create(user *entities.User) error {
	stmt, err := i.db.Prepare(`INSERT INTO users(id, name) VALUES (?, ?) `)

	if err != nil {
		return err
	}

	defer stmt.Close()
	if _, err :=  stmt.Exec(user.ID, user.Name) ; err != nil{
		return err
	}
	return nil
}


func (i *UserInfra) Update(user *entities.User) error {
	stmt, err := i.db.Prepare(`UPDATE users SET name = ?, updated_at = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.UpdatedAt, user.ID); err != nil {
		return err
	}

	return nil
}


func (i *UserInfra) Delete(id string) error {
	stmt, err := i.db.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}