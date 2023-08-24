package infra

import (
	
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yamato0204/go-test.git/app/common"
	"github.com/yamato0204/go-test.git/app/entities"
)

 
func TestUserInfra_Create(tt *testing.T) {
	tt.Run(
		"ok:エラーなし",
		func(t *testing.T) {
			user := &entities.User{
				ID : "1",
				Name:  "Ken",
				
			}
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`INSERT INTO users`).
				ExpectExec().WithArgs(user.ID,  user.Name).
				WillReturnResult(sqlmock.NewResult(1, 1))

			
			r := NewUserInfra(db)
			assert.NoError(t, r.Create(user))
			assert.NoError(t, mock.ExpectationsWereMet())
		})
}

 func TestUserInfra_Update(tt *testing.T) {
	tt.Run(
		"ok: エラーなし",
		func(t *testing.T) {
			user := &entities.User{
				ID:        "1",
				Name:      "Ken",
				UpdatedAt: common.CurrentTime(),
			}
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`UPDATE users`).
		        ExpectExec().WithArgs(user.Name, user.UpdatedAt,user.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))

				
				i := NewUserInfra(db)
			assert.NoError(t, i.Update(user))
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

func  TestUserInfra_Delete(tt *testing.T) {
	tt.Run(
		"ok: エラーなし",
		func(t *testing.T) {
			id := "1"
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`DELETE FROM users`).
			    ExpectExec().WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))

				
			i := NewUserInfra(db)

			assert.NoError(t, i.Delete(id))
			assert.NoError(t, mock.ExpectationsWereMet())
        },)
}

