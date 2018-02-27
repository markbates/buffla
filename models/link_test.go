package models_test

import (
	"github.com/gobuffalo/uuid"
	"github.com/markbates/buffla/models"
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
