package try

import (
	v1 "aigc-go/api/try/v1"
	"aigc-go/internal/service/try"
	"context"
)

func (t TryController) Try(ctx context.Context, req *v1.TryReq) (resp *v1.TryRes, err error) {
	//TODO implement me
	resp = &v1.TryRes{}
	process, err := try.ItryforController().Process(ctx, req.Text)
	if err != nil {
		return resp, err
	}
	resp.Mes = process
	return resp, nil
}
