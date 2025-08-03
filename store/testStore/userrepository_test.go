package testStore_test

import (
	"github.com/Rumpelstiltski1/restapi/internal/app/model"
	"github.com/Rumpelstiltski1/restapi/store"
	"github.com/Rumpelstiltski1/restapi/store/testStore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	s := testStore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))

	assert.NotNil(t, u)
}
func TestUserRepository_FindByEmail(t *testing.T) {

	s := testStore.New()

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	u := model.TestUser(t)
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
