package try

import (
	v1 "aigc-go/api/try/v1"
	"context"
)

type ItryV1 interface {
	Try(ctx context.Context, req *v1.TryReq) (resp *v1.TryRes, err error)
}
