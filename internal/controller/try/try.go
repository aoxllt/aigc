package try

import "aigc-go/api/try"

type TryController struct{}

func NewTryController() try.ItryV1 {
	return &TryController{}
}
