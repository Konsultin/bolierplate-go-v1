package coreSql

import (
	"github.com/jmoiron/sqlx"
	"github.com/konsultin/project-goes-here/libs/sqlk"
	"github.com/konsultin/project-goes-here/libs/sqlk/pq/query"
)

type ClientAuth struct {
	FindByClientId *sqlx.Stmt
}

func NewClientAuth(db *sqlk.DatabaseContext) *ClientAuth {
	return &ClientAuth{
		FindByClientId: db.MustPrepareRebind(query.Select(query.Column("*")).
			From(ClientAuthSchema).
			Where(query.Equal(query.Column("clientId"))).
			Build()),
	}
}
