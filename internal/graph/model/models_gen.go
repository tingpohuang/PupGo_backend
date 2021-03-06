// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Error interface {
	IsError()
}

type Payload interface {
	IsPayload()
}

// ProfileNode is a interface that stands for both pet and human
type ProfileNode interface {
	IsProfileNode()
}

type CoordinateInput struct {
	// when blur is setting that means the latitude and longitude is not the precise.
	IsBlur    bool    `json:"isBlur"`
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
}

type Event struct {
	ID          string        `json:"id"`
	Location    *Location     `json:"location"`
	TimeRange   *TimeRange    `json:"timeRange"`
	Limit       *EventsLimits `json:"limit"`
	Image       *string       `json:"image"`
	Description *string       `json:"description"`
	Type        *int          `json:"type"`
	// holder shuold be pet
	Holder *PetProfile `json:"holder"`
}

type EventRequest struct {
	ID          string      `json:"id"`
	Requester   ProfileNode `json:"requester"`
	Eventid     string      `json:"eventid"`
	Description *string     `json:"description"`
}

type EventsAcceptInput struct {
	Pid     string `json:"pid"`
	EventID string `json:"eventID"`
	Accept  bool   `json:"accept"`
}

type EventsAcceptPayload struct {
	Error     []Error       `json:"error"`
	Timestamp *string       `json:"timestamp"`
	Result    *EventRequest `json:"result"`
}

func (EventsAcceptPayload) IsPayload() {}

type EventsCreateInput struct {
	Pid         string             `json:"pid"`
	Location    *LocationInput     `json:"location"`
	TimeRange   *TimeRangeInput    `json:"timeRange"`
	Limit       *EventsLimitsInput `json:"limit"`
	Image       *string            `json:"image"`
	Description *string            `json:"description"`
}

type EventsCreatePayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    *Event  `json:"result"`
}

func (EventsCreatePayload) IsPayload() {}

type EventsJoinInput struct {
	Pid         string  `json:"pid"`
	EventID     string  `json:"eventID"`
	Description *string `json:"description"`
}

type EventsJoinPayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    *bool   `json:"result"`
}

func (EventsJoinPayload) IsPayload() {}

// Limitation of Events
type EventsLimits struct {
	LimitOfPet  *int `json:"limitOfPet"`
	LimitOfUser *int `json:"limitOfUser"`
}

type EventsLimitsInput struct {
	LimitOfDog   *int `json:"limitOfDog"`
	LimitOfHuman *int `json:"limitOfHuman"`
}

type EventsListGetInput struct {
	UID string `json:"uid"`
}

type EventsListGetPayload struct {
	Error     []Error  `json:"error"`
	Timestamp *string  `json:"timestamp"`
	Result    []*Event `json:"result"`
}

func (EventsListGetPayload) IsPayload() {}

type EventsUpdateInput struct {
	Eid         string             `json:"eid"`
	Pid         string             `json:"pid"`
	Location    *LocationInput     `json:"location"`
	TimeRange   *TimeRangeInput    `json:"timeRange"`
	Limit       *EventsLimitsInput `json:"limit"`
	Image       *string            `json:"image"`
	Description *string            `json:"description"`
}

type EventsUpdatePayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    bool    `json:"result"`
}

func (EventsUpdatePayload) IsPayload() {}

type FriendRemoveInput struct {
	PetID    string `json:"petID"`
	FriendID string `json:"friendID"`
}

type FriendRemovePayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    bool    `json:"result"`
}

func (FriendRemovePayload) IsPayload() {}

type FriendsListGetInput struct {
	Pid string `json:"pid"`
}

type FriendsListGetPayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    []*Pet  `json:"result"`
}

func (FriendsListGetPayload) IsPayload() {}

// location relation
type Location struct {
	Country   *string `json:"country"`
	State     *string `json:"state"`
	City      *string `json:"city"`
	Address   *string `json:"address"`
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
}

type LocationInput struct {
	Country    *string          `json:"country"`
	City       *string          `json:"city"`
	Address    *string          `json:"address"`
	State      *string          `json:"State"`
	Coordinate *CoordinateInput `json:"Coordinate"`
}

type NewEmail struct {
	Email string `json:"email"`
}

type Notification struct {
	NotificationID   string      `json:"notification_id"`
	NotificationType *int        `json:"notification_type"`
	UserID           *string     `json:"userId"`
	CreatedAt        *string     `json:"created_at"`
	EventInfo        *Event      `json:"eventInfo"`
	PetInfo          *PetProfile `json:"petInfo"`
	HasRead          *bool       `json:"has_read"`
}

type NotificationReadInput struct {
	Nid string `json:"nid"`
}

type NotificationReadPayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    bool    `json:"result"`
}

func (NotificationReadPayload) IsPayload() {}

type NotificationRemoveInput struct {
	Pid            string `json:"pid"`
	NotificationID string `json:"notificationID"`
}

type NotificationRemovePayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    bool    `json:"result"`
}

func (NotificationRemovePayload) IsPayload() {}

// User Notification Setting.
type NotificationSetting struct {
	AllowedNotificationWhenNewEventsFromFriends    bool `json:"allowedNotificationWhenNewEventsFromFriends"`
	AllowedNotificationWhenEventsWillStartIn30Mins bool `json:"allowedNotificationWhenEventsWillStartIn30Mins"`
	AllowedNotificationWhenEventsStatusChanged     bool `json:"allowedNotificationWhenEventsStatusChanged"`
}

type NotificationsGetInput struct {
	UID string `json:"UID"`
}

type NotificationsGetPayload struct {
	Error     []Error         `json:"error"`
	Timestamp *string         `json:"timestamp"`
	Result    []*Notification `json:"result"`
}

func (NotificationsGetPayload) IsPayload() {}

type Pet struct {
	ID              string           `json:"id"`
	Owner           *PetOwner        `json:"owner"`
	PetProfile      *PetProfile      `json:"petProfile"`
	PetRelationShip *PetRelationship `json:"petRelationShip"`
	EventList       []*Event         `json:"eventList"`
}

// Pet Connection Friends List
type PetConnection struct {
	FriendList []*PetProfile `json:"friendList"`
}

type PetCreateInput struct {
	Name  *string `json:"name"`
	Image *string `json:"image"`
	// only two
	Gender *PetGender `json:"gender"`
	// breed of dog, cat etc
	Breed *string `json:"breed"`
	// is castration or not: true for castration
	IsCastration bool           `json:"isCastration"`
	Birthday     *string        `json:"birthday"`
	Location     *LocationInput `json:"location"`
	// tmp value, should also proof by JWT later
	UID string `json:"uid"`
}

type PetCreatePayload struct {
	Error     []Error     `json:"error"`
	Timestamp *string     `json:"timestamp"`
	Result    *PetProfile `json:"result"`
}

func (PetCreatePayload) IsPayload() {}

type PetDeleteInput struct {
	Pid string `json:"pid"`
}

type PetDeletePayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    bool    `json:"result"`
}

func (PetDeletePayload) IsPayload() {}

// PetOwner relationship. Hide behide user
type PetOwner struct {
	User     *User  `json:"user"`
	PetLists []*Pet `json:"petLists"`
}

type PetProfile struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Image *string `json:"image"`
	// only two
	Gender *PetGender `json:"gender"`
	// breed of dog, cat etc
	Breed *string `json:"breed"`
	// is castration or not: true for castration
	IsCastration bool      `json:"isCastration"`
	Birthday     *string   `json:"birthday"`
	Location     *Location `json:"location"`
	Description  *string   `json:"description"`
	Hobby        []*string `json:"hobby"`
}

func (PetProfile) IsProfileNode() {}

type PetProfileListGetInput struct {
	Pid []string `json:"pid"`
}

type PetProfileListGetPayload struct {
	Error     []Error       `json:"error"`
	Timestamp *string       `json:"timestamp"`
	Result    []*PetProfile `json:"result"`
}

func (PetProfileListGetPayload) IsPayload() {}

type PetProfileUpdatesInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Image *string `json:"image"`
	// only two
	Gender *PetGender `json:"gender"`
	// breed of dog, cat etc
	Breed *string `json:"breed"`
	// is castration or not: true for castration
	IsCastration bool           `json:"isCastration"`
	Birthday     *string        `json:"birthday"`
	Location     *LocationInput `json:"location"`
}

type PetProfileUpdatesPayload struct {
	Error     []Error     `json:"error"`
	Timestamp *string     `json:"timestamp"`
	Result    *PetProfile `json:"result"`
}

func (PetProfileUpdatesPayload) IsPayload() {}

// recommendation friends (always should be pet not user)
type PetRecommned struct {
	Recommend []*Recommendation `json:"recommend"`
}

// PetRelationship to other roles
type PetRelationship struct {
	Parentpet  *Pet           `json:"parentpet"`
	Connection *PetConnection `json:"connection"`
	Recommend  *PetRecommned  `json:"recommend"`
}

type PetsListGetInput struct {
	UID string `json:"uid"`
}

type PetsListGetPayload struct {
	Error     []Error       `json:"error"`
	Timestamp *string       `json:"timestamp"`
	Result    []*PetProfile `json:"result"`
}

func (PetsListGetPayload) IsPayload() {}

type ProfileListGetInput struct {
	ID []string `json:"id"`
}

type ProfileListGetPayload struct {
	Error     []Error       `json:"error"`
	Timestamp *string       `json:"timestamp"`
	Result    []ProfileNode `json:"result"`
}

func (ProfileListGetPayload) IsPayload() {}

type Recommendation struct {
	ID     string                `json:"id"`
	Pet    *PetProfile           `json:"pet"`
	Status *RecommendationStatus `json:"status"`
}

type RecommendationGetInput struct {
	Pid string `json:"pid"`
}

type RecommendationGetPayload struct {
	Error     []Error           `json:"error"`
	Timestamp *string           `json:"timestamp"`
	Result    []*Recommendation `json:"result"`
}

func (RecommendationGetPayload) IsPayload() {}

type RecommendationResponseInput struct {
	Pid         string `json:"pid"`
	RecommendID string `json:"recommendID"`
	Result      bool   `json:"result"`
}

type RecommendationResponsePayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	// If the other already accepts, return petprofile otherwise return null
	Result *PetProfile `json:"result"`
}

func (RecommendationResponsePayload) IsPayload() {}

type TimeRange struct {
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}

type TimeRangeInput struct {
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}

type UpdatesNotificationSettings struct {
	Error     []Error              `json:"error"`
	Timestamp *string              `json:"timestamp"`
	Result    *NotificationSetting `json:"result"`
}

func (UpdatesNotificationSettings) IsPayload() {}

type UpdatesNotificationSettingsInput struct {
	UID                                            string `json:"uid"`
	AllowedNotificationWhenNewEventsFromFriends    bool   `json:"allowedNotificationWhenNewEventsFromFriends"`
	AllowedNotificationWhenEventsWillStartIn30Mins bool   `json:"allowedNotificationWhenEventsWillStartIn30Mins"`
	AllowedNotificationWhenEventsStatusChanged     bool   `json:"allowedNotificationWhenEventsStatusChanged"`
}

// User total profile, if want to get other user data. Always use userprofile!
type User struct {
	ID                  string               `json:"id"`
	Cooldown            *string              `json:"cooldown"`
	CreatedTime         *string              `json:"createdTime"`
	Email               string               `json:"email"`
	PetOwner            *PetOwner            `json:"petOwner"`
	UserProfile         *UserProfile         `json:"userProfile"`
	NotificationSetting *NotificationSetting `json:"notificationSetting"`
	EventList           []*Event             `json:"eventList"`
}

type UserCreateByIDInput struct {
	ID string `json:"id"`
}

type UserCreateByIDPayload struct {
	Error     []Error `json:"error"`
	Timestamp *string `json:"timestamp"`
	Result    *User   `json:"result"`
}

func (UserCreateByIDPayload) IsPayload() {}

type UserNotification struct {
	NotificationList []*Notification `json:"notificationList"`
}

// User profile let open to public
type UserProfile struct {
	ID     *string     `json:"id"`
	Name   *string     `json:"name"`
	Gender *UserGender `json:"gender"`
	// only use year as unit not month
	Birthday *string   `json:"birthday"`
	Email    *string   `json:"email"`
	Location *Location `json:"location"`
}

func (UserProfile) IsProfileNode() {}

type UserProfileListGetInput struct {
	UID []string `json:"uid"`
}

type UserProfileListGetPayload struct {
	Error     []Error        `json:"error"`
	Timestamp *string        `json:"timestamp"`
	Result    []*UserProfile `json:"result"`
}

func (UserProfileListGetPayload) IsPayload() {}

// Pet have no not to declared
type PetGender string

const (
	PetGenderMale   PetGender = "MALE"
	PetGenderFemale PetGender = "FEMALE"
)

var AllPetGender = []PetGender{
	PetGenderMale,
	PetGenderFemale,
}

func (e PetGender) IsValid() bool {
	switch e {
	case PetGenderMale, PetGenderFemale:
		return true
	}
	return false
}

func (e PetGender) String() string {
	return string(e)
}

func (e *PetGender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PetGender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PetGender", str)
	}
	return nil
}

func (e PetGender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RecommendationStatus string

const (
	RecommendationStatusNoOneAnswered RecommendationStatus = "NO_ONE_ANSWERED"
	RecommendationStatusAgree         RecommendationStatus = "AGREE"
	RecommendationStatusDisagree      RecommendationStatus = "DISAGREE"
)

var AllRecommendationStatus = []RecommendationStatus{
	RecommendationStatusNoOneAnswered,
	RecommendationStatusAgree,
	RecommendationStatusDisagree,
}

func (e RecommendationStatus) IsValid() bool {
	switch e {
	case RecommendationStatusNoOneAnswered, RecommendationStatusAgree, RecommendationStatusDisagree:
		return true
	}
	return false
}

func (e RecommendationStatus) String() string {
	return string(e)
}

func (e *RecommendationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RecommendationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RecommendationStatus", str)
	}
	return nil
}

func (e RecommendationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// only user 3 cases
type UserGender string

const (
	UserGenderMale   UserGender = "MALE"
	UserGenderFemale UserGender = "FEMALE"
	// put all other to this
	UserGenderNotToDeclared UserGender = "NOT_TO_DECLARED"
)

var AllUserGender = []UserGender{
	UserGenderMale,
	UserGenderFemale,
	UserGenderNotToDeclared,
}

func (e UserGender) IsValid() bool {
	switch e {
	case UserGenderMale, UserGenderFemale, UserGenderNotToDeclared:
		return true
	}
	return false
}

func (e UserGender) String() string {
	return string(e)
}

func (e *UserGender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserGender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserGender", str)
	}
	return nil
}

func (e UserGender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
