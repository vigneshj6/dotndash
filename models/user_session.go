package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type UserSession struct {
	UserId       string    `json:"UserId" db:"UserId"`
	PermissionId string    `json:"PermissionId" db:"PermissionId"`
	SessionToken string    `json:"SessionToken" db:"SessionToken"`
	Timestamp    time.Time `json:"Timestamp" db:"Timestamp"`
}

// String is not required by pop and may be deleted
func (u UserSession) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserSessions is not required by pop and may be deleted
type UserSessions []UserSession

// String is not required by pop and may be deleted
func (u UserSessions) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserSession) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserSession) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserSession) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
