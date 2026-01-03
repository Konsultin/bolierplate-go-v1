package dto

type Role struct {
	Xid         string                `json:"xid,omitempty"`
	Version     int64                 `json:"version,omitempty"`
	Name        string                `json:"name,omitempty"`
	RoleType    *RoleType_Result      `json:"roleType,omitempty"`
	CreatedAt   int64                 `json:"createdAt,omitempty"`
	UpdatedAt   int64                 `json:"updatedAt,omitempty"`
	ModifiedBy  *Subject              `json:"modifiedBy,omitempty"`
	Description string                `json:"description,omitempty"`
	Privileges  []string              `json:"privileges,omitempty"`
	Status      *ControlStatus_Result `json:"status,omitempty"`
}

type Role_Enum int32

const (
	Role_UNKNOWN_ROLE Role_Enum = 0

	// Admin
	Role_ANONYMOUS_ADMIN Role_Enum = 1
	Role_ADMIN           Role_Enum = 3

	// User
	Role_ANONYMOUS_USER Role_Enum = 2
	Role_USER           Role_Enum = 4
)

var (
	Role_Enum_name = map[int32]string{
		0: "UNKNOWN_ROLE",
		1: "ANONYMOUS_ADMIN",
		2: "ANONYMOUS_USER",
		3: "ADMIN",
		4: "USER",
	}

	Role_Enum_value = map[string]int32{
		"UNKNOWN_ROLE":    0,
		"ANONYMOUS_ADMIN": 1,
		"ANONYMOUS_USER":  2,
		"ADMIN":           3,
		"USER":            4,
	}
)

type RoleType_Result struct {
	Id   RoleType_Enum `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
}

type RoleType_Enum int32

const (
	RoleType_UNKNOWN RoleType_Enum = 0
	RoleType_ADMIN   RoleType_Enum = 1
	RoleType_USER    RoleType_Enum = 2
	RoleType_SYSTEM  RoleType_Enum = 9
)

var (
	RoleType_Enum_name = map[int32]string{
		0: "UNKNOWN",
		1: "ADMIN",
		2: "USER",
		9: "SYSTEM",
	}
	RoleType_Enum_value = map[string]int32{
		"UNKNOWN": 0,
		"ADMIN":   1,
		"USER":    2,
		"SYSTEM":  9,
	}
)
