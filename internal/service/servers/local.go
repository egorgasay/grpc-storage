package servers

import (
	"context"

	"github.com/egorgasay/gost"
	"itisadb/config"
	"itisadb/internal/constants"
	"itisadb/internal/domains"
	"itisadb/internal/models"
	"itisadb/internal/service/logic"
	"itisadb/pkg"
)

type LocalServer struct {
	*logic.Logic
	storage domains.Storage
	config  config.Config
	ram     gost.RwLock[models.RAM]
}

func NewLocalServer(uc *logic.Logic) *LocalServer {
	return &LocalServer{
		Logic: uc,
		ram:   gost.NewRwLock(models.RAM{}),
	}
}

func (s *LocalServer) IsOffline() bool { return false }

func (s *LocalServer) Reconnect(_ context.Context) (res gost.ResultN) {
	return
}

func (s *LocalServer) ResetTries() {}

func (s *LocalServer) Number() int32 {
	return constants.MainStorageNumber
}

func (s *LocalServer) RAM() models.RAM {
	r := s.ram.RBorrow()
	defer s.ram.Release()

	return r.Read()
}

func (s *LocalServer) RefreshRAM(_ context.Context) (res gost.Result[gost.Nothing]) {
	r := pkg.CalcRAM()
	if r.IsErr() {
		return res.Err(r.Error())
	}

	s.ram.SetWithLock(r.Unwrap())

	return res.Ok(gost.Nothing{})
}

func (s *LocalServer) NewUser(ctx context.Context, claims gost.Option[models.UserClaims], user models.User) (r gost.ResultN) {
	if s.config.Balancer.On {
		return r.Ok()
	}

	if rUser := s.storage.NewUser(user); r.IsErr() {
		return r.Err(rUser.Error())
	}

	return r.Ok()
}

func (s *LocalServer) DeleteUser(ctx context.Context, claims gost.Option[models.UserClaims], login string) (r gost.Result[bool]) {
	if s.config.Balancer.On {
		return r.Ok(false)
	}

	rUser := s.storage.GetUserIDByName(login)
	if rUser.IsErr() {
		return r.Err(rUser.Error())
	}

	return s.storage.DeleteUser(rUser.Unwrap())
}

func (s *LocalServer) ChangePassword(ctx context.Context, claims gost.Option[models.UserClaims], login string, password string) (r gost.ResultN) {
	if s.config.Balancer.On {
		return r.Ok()
	}

	rUser := s.storage.GetUserByName(login)
	if rUser.IsErr() {
		return r.Err(rUser.Error())
	}

	user := rUser.Unwrap()
	user.Password = password

	return s.storage.SaveUser(user.ID, user)
}

func (s *LocalServer) ChangeLevel(ctx context.Context, claims gost.Option[models.UserClaims], login string, level models.Level) (r gost.ResultN) {
	return r.Ok()
}
