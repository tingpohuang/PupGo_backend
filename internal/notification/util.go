package notification

import (
	"context"

	"github.com/tingpo/pupgobackend/internal/gorm"
)

func UserDeviceToTokens(ctx context.Context, devices []gorm.User_device) ([]string, error) {
	if len(devices) == 0 {
		return nil, nil
	}
	ret := make([]string, len(devices))
	for i := 0; i < len(devices); i++ {
		ret[i] = devices[i].Device_id
	}
	return ret, nil
}

func PetIDToTokens(ctx context.Context, pid string, s *gorm.SQLCnter) ([]string, error) {
	panic("not yet ready")
	// devices, err := s.FindUserDeviceID(ctx, pid)
	// if err != nil {
	// 	return nil, err
	// }
	// tokens, err := UserDeviceToTokens(ctx, devices)
	// if err != nil {
	// 	return nil, err
	// }
	// return tokens, nil
}

func EventIDToTokens(ctx context.Context, eventId string, s *gorm.SQLCnter) ([]string, error) {
	// devices, err := s.FindUserDeviceID(ctx, eventId)
	// if err != nil {
	// 	return nil, err
	// }
	// tokens, err := UserDeviceToTokens(ctx, devices)
	// if err != nil {
	// 	return nil, err
	// }
	// return tokens, nil}
	panic("not implemented")
}
