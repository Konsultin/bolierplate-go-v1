package repository

import (
	"context"

	"github.com/konsultin/project-goes-here/internal/svc-core/model"
	"github.com/konsultin/project-goes-here/libs/errk"
)

func (r *Repository) FindClientAuthByClientId(ctx context.Context, id string) (*model.ClientAuth, error) {
	var m model.ClientAuth
	err := r.sql.ClientAuth.FindByClientId.GetContext(ctx, &m, id)
	if err != nil {
		return nil, errk.Trace(err)
	}
	return &m, nil
}
