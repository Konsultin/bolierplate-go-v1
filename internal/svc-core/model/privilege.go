package model

type Privilege struct {
	BaseField
	Id      int64  `db:"id"`
	Xid     string `db:"xid"`
	Name    string `db:"name"`
	Exposed bool   `db:"exposed"`
	Sort    int32  `db:"sort"`
}
