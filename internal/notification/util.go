package notification

import "github.com/tingpo/pupgobackend/internal/gorm"

func UserDeviceToTokens(devices []gorm.User_device) ([]string, error) {
	if devices == nil || len(devices) == 0 {
		return nil, nil
	}
	ret := make([]string, len(devices))
	for i := 0; i < len(devices); i++ {
		ret[i] = devices[i].Device_id
	}
	return ret, nil
}
