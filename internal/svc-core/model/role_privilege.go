package model

import "github.com/konsultin/project-goes-here/dto"

type RolePrivilege struct {
	BaseField
	Id          int64 `db:"id"`
	RoleId      int32 `db:"roleId"`
	PrivilegeId int64 `db:"privilegeId"`
}

type RolePrivilegeJoinRow struct {
	RolePrivilege *RolePrivilege `db:"rolePrivilege"`
	Role          *Role          `db:"role"`
	Privilege     *Privilege     `db:"privilege"`
}

func NewRolePrivilege(privilegeId int64, s *dto.Subject) *RolePrivilege {
	return &RolePrivilege{
		BaseField:   NewBaseField(s),
		PrivilegeId: privilegeId,
	}
}
