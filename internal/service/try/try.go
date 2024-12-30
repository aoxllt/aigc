package try

import "context"

type Itry interface {
	Process(ctx context.Context, text string) (string, error)
}

var localtry Itry

func ItryforController() Itry {
	if localtry == nil {
		panic("接口未实现")
	}
	return localtry
}

func RegisterItry(try Itry) {
	localtry = try
}
