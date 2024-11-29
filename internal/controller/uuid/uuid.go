package uuid

import "aigc-go/api/uuid"

type UuidController struct{}

func Uuidv1() uuid.Iuuidv1 {
	return &UuidController{}
}
