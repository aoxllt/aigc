package getret

import "context"

type IGetret interface {
	Getcontent(ctx context.Context, sk string) (string, error)
}

var localgetret IGetret

func GetretforController() IGetret {
	if localgetret == nil {
		panic("接口未注册")
	}
	return localgetret
}
func Registergetret(i IGetret) {
	localgetret = i
}
