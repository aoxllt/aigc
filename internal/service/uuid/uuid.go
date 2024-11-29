package uuid

import (
	v1 "aigc-go/api/uuid/v1"
	"context"
)

type Iuuidservice interface {
	Setuuid(ctx context.Context, req *v1.UuidReq) (string, error)
	Getuuid(ctx context.Context) (string, error)
}

var localUuid Iuuidservice

func UuidforController() Iuuidservice {
	if localUuid == nil {
		panic("localUuid is nil")
	}
	return localUuid
}
func RegisterUuid(c Iuuidservice) {
	localUuid = c
}
