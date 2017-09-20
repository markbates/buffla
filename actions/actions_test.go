package actions

import (
	"testing"
	"time"

	"github.com/gobuffalo/suite"
	"github.com/markbates/buffla/models"
	"github.com/markbates/pop/nulls"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func (as *ActionSuite) CreateUser() *models.User {
	user := &models.User{
		Name:       "Mark",
		Email:      nulls.NewString("mark@example.com"),
		Provider:   "faux",
		ProviderID: time.Now().String(),
	}
	as.NoError(as.DB.Create(user))
	return user
}

func (as *ActionSuite) Login() *models.User {
	user := as.CreateUser()
	as.Session.Set("current_user_id", user.ID)
	return user
}

func (as *ActionSuite) CreateLink(user *models.User) *models.Link {
	link := &models.Link{
		Link:   "http://example.com",
		UserID: user.ID,
	}

	verrs, err := as.DB.ValidateAndCreate(link)
	as.NoError(err)
	as.False(verrs.HasAny())
	return link
}
