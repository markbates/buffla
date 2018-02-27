package actions

import (
	"net/http"

	"github.com/gobuffalo/pop/nulls"
	"github.com/markbates/buffla/models"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (as *ActionSuite) Test_AuthCallback_NewUser() {
	cau := gothic.CompleteUserAuth
	defer func() {
		gothic.CompleteUserAuth = cau
	}()
	gothic.CompleteUserAuth = func(http.ResponseWriter, *http.Request) (goth.User, error) {
		return goth.User{
			Name:     "Mark",
			Provider: "github",
			UserID:   "123",
			Email:    "mark@example.com",
		}, nil
	}
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	res := as.HTML("/auth/github/callback").Get()
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	u := &models.User{}
	err = as.DB.First(u)
	as.NoError(err)
	as.Equal("Mark", u.Name)
	as.Equal("github", u.Provider)
	as.Equal("123", u.ProviderID)
	as.Equal("mark@example.com", u.Email.String)

	as.Equal(u.ID, as.Session.Get("current_user_id"))
}

func (as *ActionSuite) Test_AuthCallback_ExistingUser() {
	u := &models.User{
		Name:       "Mark",
		Provider:   "github",
		ProviderID: "123",
		Email:      nulls.NewString("mark@example.com"),
	}
	err := as.DB.Create(u)
	as.NoError(err)

	cau := gothic.CompleteUserAuth
	defer func() {
		gothic.CompleteUserAuth = cau
	}()
	gothic.CompleteUserAuth = func(http.ResponseWriter, *http.Request) (goth.User, error) {
		return goth.User{
			Name:     "Ringo",
			Provider: "github",
			UserID:   "123",
			Email:    "ringo@example.com",
		}, nil
	}

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	res := as.HTML("/auth/github/callback").Get()
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	u = &models.User{}
	err = as.DB.First(u)
	as.NoError(err)
	as.Equal("Ringo", u.Name)
	as.Equal("github", u.Provider)
	as.Equal("123", u.ProviderID)
	as.Equal("ringo@example.com", u.Email.String)

	as.Equal(u.ID, as.Session.Get("current_user_id"))
}
