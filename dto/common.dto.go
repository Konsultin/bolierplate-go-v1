package dto

type Response[T any] struct {
	Message   string `json:"message"`
	Code      Code   `json:"code"`
	Data      T      `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

type ErrorTrace struct {
	Source string   `json:"source"`
	Trace  []string `json:"trace"`
	Err    error    `json:"err"`
}

type Subject struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

type SimpleData struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type Empty struct{}

type ControlStatus_Result struct {
	Id   ControlStatus_Enum `json:"id,omitempty"`
	Name string             `json:"name,omitempty"`
}

type ControlStatus_Enum int32

const (
	ControlStatus_UNKNOWN_CONTROL_STATUS ControlStatus_Enum = 0
	ControlStatus_ACTIVE                 ControlStatus_Enum = 1
	ControlStatus_INACTIVE               ControlStatus_Enum = 2
	ControlStatus_PENDING                ControlStatus_Enum = 3
	ControlStatus_LOCKED                 ControlStatus_Enum = 4
	ControlStatus_DRAFT                  ControlStatus_Enum = 5
)

var (
	ControlStatus_Enum_name = map[int32]string{
		0: "UNKNOWN_CONTROL_STATUS",
		1: "ACTIVE",
		2: "INACTIVE",
		3: "PENDING",
		4: "LOCKED",
		5: "DRAFT",
	}
	ControlStatus_Enum_value = map[string]int32{
		"UNKNOWN_CONTROL_STATUS": 0,
		"ACTIVE":                 1,
		"INACTIVE":               2,
		"PENDING":                3,
		"LOCKED":                 4,
		"DRAFT":                  5,
	}
)
