package models

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Link struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	Link       string    `json:"link" db:"link"`
	Code       string    `json:"code" db:"code"`
	ClickCount int       `json:"-" db:"click_count" form:"-" select:"(select count(*) from clicks where link_id = links.id) as click_count" rw:"r"`
}

// String is not required by pop and may be deleted
func (l Link) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Links is not required by pop and may be deleted
type Links []Link

// String is not required by pop and may be deleted
func (l Links) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

func (l *Link) BeforeValidations(tx *pop.Connection) error {
	if l.Code != "" {
		return nil
	}
	h := sha1.New()
	h.Write(l.UserID.Bytes())
	h.Write([]byte(l.Link))
	h.Write([]byte(time.Now().String()))
	code := fmt.Sprintf("%x", h.Sum(nil))[:7]
	exists, err := tx.Where("code = ?", code).Exists("links")
	if err != nil {
		return errors.WithStack(err)
	}
	if exists {
		return l.BeforeValidations(tx)
	}
	l.Code = code
	return nil
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Link) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.Link, Name: "Link"},
		&validators.StringIsPresent{Field: l.Code, Name: "Code"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Link) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Link) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
