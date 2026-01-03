package model

import "github.com/konsultin/project-goes-here/dto"

type Role struct {
	BaseField
	Id          int32                  `db:"id"`
	Xid         string                 `db:"xid"`
	Name        string                 `db:"name"`
	Description string                 `db:"description"`
	RoleTypeId  dto.RoleType_Enum      `db:"roleTypeId"`
	StatusId    dto.ControlStatus_Enum `db:"statusId"`
}

type FindRoleResult struct {
	Rows  []Role
	Count int64
}

func NewRole(xid string, role *dto.Role, s *dto.Subject) *Role {
	return &Role{
		BaseField:   NewBaseField(s),
		Xid:         xid,
		Name:        role.Name,
		Description: role.Description,
		RoleTypeId:  dto.RoleType_ADMIN,
		StatusId:    dto.ControlStatus_ACTIVE,
	}
}
