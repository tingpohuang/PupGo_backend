package gorm

const (
	User_1_id  = "bf448152-bf2c-421e-af4c-ae458737da0e"
	User_2_id  = "982a33ee-2792-4ede-b307-f38d187a2015"
	User_3_id  = "74c479bd-4076-40e1-84dd-228ab8c183fb"
	Pet_1_id   = "dad9e58f-42dd-4118-9765-03990ca89ee9"
	Pet_2_id   = "d26d7d6b-4b29-4c50-ba13-0eeaff957923"
	Pet_3_id   = "26906919-7cbe-4e41-bd6b-fafd560b4ca6"
	Pet_4_id   = "0ede6c06-18b7-44c1-b5fe-4c01de57c770"
	Event_1_id = "eba52c68-15a7-4883-9ad4-a3ecc09b0fe7"
)

var (
	UserLocation1 = UserLocation{
		User_id: User_1_id,
		Position: Location{
			Lat:  1.23,
			Long: 4.56,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "1878 Greenfield Avenue",
	}
	UserLocation2 = UserLocation{
		User_id: User_2_id,
		Position: Location{
			Lat:  1.23,
			Long: 4.56,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "1878 Greenfield Avenue",
	}
	UserLocation3 = UserLocation{
		User_id: User_3_id,
		Position: Location{
			Lat:  1.232,
			Long: 4.56,
		},
		Country: "USA",
		State:   "CA",
		City:    "Los Angeles",
		Address: "1878 Greenfield Avenue",
	}
)
