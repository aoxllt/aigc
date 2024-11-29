package uuid

import (
	v1 "aigc-go/api/uuid/v1"
	"aigc-go/internal/dao"
	"aigc-go/internal/service/uuid"
	"aigc-go/utility"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

func init() {
	uuid.RegisterUuid(Reg{})
}

type user struct {
	Uid  string    `json:"uid"`
	Time time.Time `json:"time"`
}
type Reg struct{}

func (r Reg) Getuuid(ctx context.Context) (string, error) {
	//TODO implement me
	get, err := ghttp.RequestFromCtx(ctx).Session.Get("uuid")
	if err != nil {
		return "", err
	}
	return get.String(), nil
}

func (r Reg) Setuuid(ctx context.Context, req *v1.UuidReq) (string, error) {
	var u user
	rCtx := ghttp.RequestFromCtx(ctx)
	if rCtx == nil {
		return "false", gerror.New("Invalid HTTP context")
	}
	existingUUID, _ := rCtx.Session.Get("uuid")
	if existingUUID != nil {

		return "uuid already exists ", nil
	}
	uid := utility.GenerateUUID()
	err := rCtx.Session.Set("uuid", uid)

	if err != nil {
		return "false", err
	}
	u.Uid = uid
	u.Time = time.Now()
	_, err = dao.Users.Ctx(ctx).Insert(u)
	if err != nil {
		return "", err
	}
	return "true", nil
}
