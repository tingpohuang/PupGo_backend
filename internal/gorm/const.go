package gorm

import "time"

const (
	User_1_id                         = "bf448152-bf2c-421e-af4c-ae458737da0e"
	User_2_id                         = "982a33ee-2792-4ede-b307-f38d187a2015"
	User_3_id                         = "74c479bd-4076-40e1-84dd-228ab8c183fb"
	User_4_id                         = "fb3e1393-1aed-4b92-b637-24b491e7e4a2"
	Pet_1_id                          = "dad9e58f-42dd-4118-9765-03990ca89ee9"
	Pet_2_id                          = "d26d7d6b-4b29-4c50-ba13-0eeaff957923"
	Pet_3_id                          = "26906919-7cbe-4e41-bd6b-fafd560b4ca6"
	Pet_4_id                          = "0ede6c06-18b7-44c1-b5fe-4c01de57c770"
	Event_1_id                        = "eba52c68-15a7-4883-9ad4-a3ecc09b0fe7"
	Event_2_id                        = "1ba52c68-15a7-4883-9ad4-a1ec09b0fefq"
	Event_3_id                        = "2ba52c68-15a7-4883-9ad4-a1ec09b0few3"
	Petsize                           = 11
	Notification_NewFriend            = 0
	Notification_NewParticipants      = 1
	Notification_EventJoined          = 2
	Notification_EventContentUpdate   = 3
	Notification_EventsToNotification = 4
	Notification_EventsToFriends      = 5
)

var (
	Pet_ids = [Petsize]string{
		"caefd1f0-a4fc-4ba3-81d8-1d0b0fbec730",
		"149464c2-f8ee-4e6f-a551-260b6467fa95",
		"ca035fdd-bc77-4d76-879a-1588c394cd74",
		"a743cc24-c619-4d5a-b65d-2b41be2aeb31",
		"77433673-35b1-4d12-b04a-9cfcf0d5ed9c",
		"5d76c3ad-d286-4c82-9ff0-6e043389f00d",
		"d206d148-f800-4f90-bd86-ea40f3b694b9",
		"df34398a-e6b2-4ba6-8e82-a87c669ea7c6",
		"aa73746c-78f1-48bf-ae5b-8a7b0e8d741a",
		"a29fb9c3-4bc4-438f-9589-5b9143098da7",
		"eba52c68-15a7-4883-9ad4-a1ec20930f67"}
	Pet_imgs = [Petsize]string{
		"https://i.ibb.co/DVqdYWP/IMG-3228-min.jpg",
		"https://i.ibb.co/TcLZTXm/IMG-3227-min.jpg",
		"https://i.ibb.co/JB31CsL/IMG-3225-min.jpg",
		"https://i.ibb.co/YXYHtB8/IMG-3224-min.jpg",
		"https://i.ibb.co/ccK4g5R/IMG-3235-min.jpg",
		"https://i.ibb.co/cFJN2sd/IMG-3234-min.jpg",
		"https://i.ibb.co/7XKCfbW/IMG-3233-min.jpg",
		"https://pupgo-testing-pictures-1.s3.us-west-1.amazonaws.com/IMG-3232-min.jpeg",
		"https://i.ibb.co/bJ2s4tc/IMG-3231-min.jpg",
		"https://i.ibb.co/x3NDrwq/IMG-3230-min.jpg",
		"https://i.ibb.co/C2tFDw7/IMG-3229-min.jpg"}
	Pet_names         = [Petsize]string{"Tuple", "Array", "Sequel", "Class", "struct", "Abstract", "Hoare", "Queue", "Stack", "List", "String"}
	Pet_genders       = [Petsize]int{0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0}
	Pet_breeds        = [Petsize]string{"Husky", "Husky", "Corgi", "Russell Terrier", "French BuildDog", "Autsralian Shepherd", "Bichon Frise", "Golden Retriever", "Sennenhunde", "American Eskimo", "Affenpinscher"}
	Pet_isCastrations = [Petsize]bool{true, true, false, true, true, true, false, true, true, true, true}
	loc               = time.FixedZone("UTC-8", -8*60*60)

	Pet_Birthdays = [11]time.Time{
		time.Date(2020, 1, 1, 0, 0, 0, 0, loc),
		time.Date(2017, 4, 1, 0, 0, 0, 0, loc),
		time.Date(2018, 1, 1, 0, 0, 0, 0, loc),
		time.Date(2015, 5, 1, 0, 0, 0, 0, loc),
		time.Date(2021, 3, 1, 0, 0, 0, 0, loc),
		time.Date(2013, 9, 1, 0, 0, 0, 0, loc),
		time.Date(2016, 12, 1, 0, 0, 0, 0, loc),
		time.Date(2011, 11, 1, 0, 0, 0, 0, loc),
		time.Date(2020, 6, 1, 0, 0, 0, 0, loc),
		time.Date(2018, 7, 1, 0, 0, 0, 0, loc),
		time.Date(2019, 8, 1, 0, 0, 0, 0, loc),
	}
	UserLocation1 = UserLocation{
		User_id:   User_1_id,
		Latitude:  100,
		Longitude: 100,
		Country:   "USA",
		State:     "CA",
		City:      "Los Angeles",
		Address:   "1878 Greenfield Avenue",
	}
	UserLocation2 = UserLocation{
		User_id:   User_2_id,
		Latitude:  300,
		Longitude: 304.56,
		Country:   "USA",
		State:     "CA",
		City:      "Los Angeles",
		Address:   "1878 Greenfield Avenue",
	}
	UserLocation3 = UserLocation{
		User_id:   User_3_id,
		Latitude:  400,
		Longitude: 500,
		Country:   "USA",
		State:     "CA",
		City:      "Los Angeles",
		Address:   "1878 Greenfield Avenue",
	}
)
