package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/konsultin/project-goes-here/dto"
	"github.com/konsultin/project-goes-here/internal/svc-core/pkg/httpk"
	unaryHttpk "github.com/konsultin/project-goes-here/internal/svc-core/pkg/httpk/unary"
	"github.com/konsultin/project-goes-here/internal/svc-core/pkg/svck"
	"github.com/konsultin/project-goes-here/libs/errk"
	logkOption "github.com/konsultin/project-goes-here/libs/logk/option"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateAnonymousUserSession(payload *unaryHttpk.BasicAuth, clientTypeId dto.Role_Enum) (*dto.CreateAnonymousSession_Result, error) {
	clientAuth, err := s.Repo.FindClientAuthByClientId(payload.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.Log.Errorf("clientAuth is not found. Username = %s", payload.Username)
			return nil, httpk.UnauthorizedError
		}
		s.Log.Error("Failed to FindClientAuthByClientId. Username = %s", logkOption.Error(err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(clientAuth.Options.ClientSecret), []byte(payload.Password))
	if err != nil {
		s.Log.Error("Failed to compare the password", logkOption.Error(err))
		return nil, httpk.UnauthorizedError.Wrap(err).Trace()
	}

	if clientAuth.ClientTypeId != clientTypeId {
		s.Log.Error("Invalid clientTypeId")
		return nil, httpk.UnauthorizedError.Trace()
	}

	var subjectType int32
	switch clientAuth.ClientTypeId {
	case dto.Role_ANONYMOUS_ADMIN, dto.Role_ANONYMOUS_USER:
		subjectType = int32(clientAuth.ClientTypeId)
	default:
		return nil, fmt.Errorf("invalid Client Type. clientTypeId = %v", clientAuth.ClientTypeId)
	}

	rolePrivileges, fErr := s.Repo.FindRolePrivilegeByRoleId(subjectType)
	if fErr != nil {
		s.Log.Error("Failed to FindRolePrivilegeByRoleId", logkOption.Error(fErr))
		return nil, errk.Trace(fErr)
	}

	var audience []string
	for _, val := range rolePrivileges {
		audience = append(audience, val.Privilege.Xid)
	}

	jwtAdapter := s.NewJwtAdapter()
	session, err := jwtAdapter.Issue(IssueJwtPayload{
		Subject:     clientAuth.ClientId,
		Audience:    audience,
		Lifetime:    clientAuth.Options.TokenLifetime,
		SessionId:   gonanoid.MustGenerate(svck.AlphaNumUpperCharSet, 6),
		SubjectType: subjectType,
	})

	if err != nil {
		s.Log.Error("Failed to issue jwt payload", logkOption.Error(err))
		return nil, errk.Trace(err)
	}

	return &dto.CreateAnonymousSession_Result{
		Session: session,
		Scopes:  audience,
	}, nil
}
