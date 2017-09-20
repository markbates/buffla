package models_test

import (
	"github.com/markbates/buffla/models"
	"github.com/satori/go.uuid"
)

func (ms *ModelSuite) Test_Link_BeforeValidations() {
	link := &models.Link{
		UserID: uuid.NewV4(),
		Link:   "http://gobuffalo.io",
	}
	err := link.BeforeValidations(ms.DB)
	ms.NoError(err)
	ms.NotZero(link.Code)
	ms.Len(link.Code, 7)
}
