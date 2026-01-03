package model

import (
	"database/sql"

	"github.com/konsultin/project-goes-here/dto"
)

type ClientAuth struct {
	BaseField
	Id           int64                  `db:"id"`
	Name         string                 `db:"name"`
	ClientId     string                 `db:"clientId"`
	ClientTypeId dto.Role_Enum          `db:"clientTypeId"`
	Options      *ClientAuthOptions     `db:"options"`
	StatusId     dto.ControlStatus_Enum `db:"statusId"`
}

func NewClientAuth(name string, clientId string, clientTypeId dto.Role_Enum, clientSecret string, subject *dto.Subject, tokenLifeTime sql.NullInt64) *ClientAuth {
	if !tokenLifeTime.Valid {
		tokenLifeTime.Int64 = 2592000
		tokenLifeTime.Valid = true
	}

	return &ClientAuth{
		BaseField:    NewBaseField(subject),
		Name:         name,
		ClientId:     clientId,
		ClientTypeId: clientTypeId,
		Options: &ClientAuthOptions{
			ClientSecret:  clientSecret,
			TokenLifetime: tokenLifeTime.Int64,
		},
		StatusId: 1,
	}
}
