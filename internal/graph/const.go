package graph

type Season int64
type EventStatus int
type RecommendationStatus int
type PetGender int

const (
	EventStatusNoAnswer           EventStatus          = 0
	EventStatusAccept             EventStatus          = 1
	EventStatusDecline            EventStatus          = -1
	RecommendationStatusNoAnswer  RecommendationStatus = 0
	RecommendationStatusLowAgree  RecommendationStatus = 1
	RecommendationStatusHighAgree RecommendationStatus = 2
	RecommendationStatusBothAgree RecommendationStatus = 3
	RecommendationStatusDecline   RecommendationStatus = -1
	PetGenderMale                 PetGender            = 1
	PetGenderFemale               PetGender            = 0
)
