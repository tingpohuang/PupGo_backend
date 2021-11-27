package graph_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/tingpo/pupgobackend/internal/gorm"
	"github.com/tingpo/pupgobackend/internal/graph/model"
)

var (
	graphql_endpoint = "http://localhost:8080/query"
)

func TestMutationResolver(t *testing.T) {
	// var m struct {
	// }
	assert := assert.New(t)
	client := graphql.NewClient(graphql_endpoint)
	assert.NotNil(client)
	MutationResolver_EventsJoin(t, client, assert)
}

func TestMutationResolver_EventsJoin(t *testing.T) {
	assert := assert.New(t)
	client := graphql.NewClient(graphql_endpoint)
	assert.NotNil(client)
	MutationResolver_EventsJoin(t, client, assert)
}
func TestMutationResolver_EventsAccept(t *testing.T) {
	assert := assert.New(t)
	client := graphql.NewClient(graphql_endpoint)
	assert.NotNil(client)
	MutationResolver_EventsAccept(t, client, assert, true)
	MutationResolver_EventsAccept(t, client, assert, false)
}

func TestMutationResolver_EventsCreate(t *testing.T) {
	assert := assert.New(t)
	client := graphql.NewClient(graphql_endpoint)
	assert.NotNil(client)
	lat, long := fmt.Sprintf("%f", gorm.UserLocation1.Position.Lat), fmt.Sprintf("%f", gorm.UserLocation1.Position.Long)
	stime, etime := "123", "321"
	MutationResolver_EventsCreate(t, client, assert, &model.EventsCreateInput{
		Pid: gorm.Pet_2_id,
		Location: &model.LocationInput{
			Country: &gorm.UserLocation1.Country,
			City:    &gorm.UserLocation1.City,
			Address: &gorm.UserLocation1.Address,
			Coordinate: &model.CoordinateInput{
				Latitude:  &lat,
				Longitude: &long,
			},
		},
		TimeRange: &model.TimeRangeInput{
			StartTime: &stime,
			EndTime:   &etime,
		},
		Image: nil,
	})
}
func TestMutationResolver_RecommendationResponse(t *testing.T) {
	// "go test -timeout 30s -run ^TestMutationResolver_RecommendationResponse$ github.com/tingpo/pupgobackend/internal/graph/test"
	assert := assert.New(t)
	client := graphql.NewClient(graphql_endpoint)
	ctx := context.Background()
	mysqlConnector, err := gorm.GetConnectorFactory("mySQL")
	assert.Nil(err)
	db := mysqlConnector.NewDBConnection()
	sqlCnter := gorm.NewSQLCnter(db)
	assert.NotNil(sqlCnter)
	assert.NotNil(client)
	MutationResolver_RecommendationResponse(t, client, assert, &model.RecommendationResponseInput{
		Pid:         gorm.Pet_2_id,
		RecommendID: gorm.Pet_3_id,
		Result:      true,
	})
	res, err := sqlCnter.FindPetRecommendByID(ctx, gorm.Pet_2_id, gorm.Pet_3_id)
	assert.Nil(err)
	assert.Equal(1, res.Status)

	MutationResolver_RecommendationResponse(t, client, assert, &model.RecommendationResponseInput{
		Pid:         gorm.Pet_3_id,
		RecommendID: gorm.Pet_2_id,
		Result:      true,
	})
	res, err = sqlCnter.FindPetRecommendByID(ctx, gorm.Pet_3_id, gorm.Pet_2_id)
	assert.Nil(err)
	assert.Equal(3, res.Status)
	MutationResolver_RecommendationResponse(t, client, assert, &model.RecommendationResponseInput{
		Pid:         gorm.Pet_2_id,
		RecommendID: gorm.Pet_4_id,
		Result:      false,
	})
	res, err = sqlCnter.FindPetRecommendByID(ctx, gorm.Pet_3_id, gorm.Pet_2_id)
	assert.Nil(err)
	assert.Equal(-1, res.Status)
}

func TestMutationResolver_PetCreate(t *testing.T) {
	assert := assert.New(t)
	client := graphql.NewClient(graphql_endpoint)
	name := "testname"
	breed := "hahahusky"
	gender := model.PetGenderMale
	birthday := time.Now().UTC().String()
	// print(birthday)
	p := MutationResolver_PetCreate(t, client, assert, &model.PetCreateInput{
		Name:         &name,
		Image:        nil,
		Gender:       &gender,
		Breed:        &breed,
		IsCastration: true,
		Birthday:     &birthday,
		UID:          gorm.User_4_id,
	})
	assert.Equal(name, *p.Result.Name)
	assert.Equal(gender, *p.Result.Gender)
	assert.Equal(breed, *p.Result.Breed)
	assert.Equal(true, p.Result.IsCastration)
	assert.Equal(birthday, *p.Result.Birthday)
	assert.NotNil(p.Result.ID)
	q := MutationResolver_PetDelete(client, assert, &model.PetDeleteInput{
		Pid: *(p.Result.ID),
	})
	assert.True(q.Result)
}

func MutationResolver_EventsJoin(t *testing.T, c *graphql.Client, assert *assert.Assertions) {
	ctx := context.Background()
	req := graphql.NewRequest(`
    mutation($Pid:ID!, $Eid:ID!){
		eventsJoin(eventsJoinInput:{
		  pid:$Pid,
		  eventID:$Eid,
		  description:""
		}){
		  timestamp,
		  result
		},
	},
`)
	req.Var("Pid", gorm.Pet_4_id)
	req.Var("Eid", gorm.Event_1_id)
	var respData model.EventsJoinPayload
	if err := c.Run(ctx, req, &respData); err != nil {
		assert.Nil(err)
	}
}

func MutationResolver_EventsAccept(t *testing.T, c *graphql.Client, assert *assert.Assertions, accept bool) {
	ctx := context.Background()
	req := graphql.NewRequest(`
    mutation($Pid:ID!, $Eid:ID!,$Accept: Boolean!){
		eventsAccept(eventsAcceptInput:{
		  pid:$Pid,
		  eventID:$Eid,
		  accept: $Accept
		}){
		  timestamp,
		  result{
			eventid
		  }
		},
	  },
`)
	req.Var("Pid", gorm.Pet_4_id)
	req.Var("Eid", gorm.Event_1_id)
	req.Var("Accept", accept)
	var respData model.EventsAcceptPayload
	if err := c.Run(ctx, req, &respData); err != nil {
		assert.Nil(err)
	}
}

func MutationResolver_EventsCreate(t *testing.T, c *graphql.Client, assert *assert.Assertions, m *model.EventsCreateInput) {
	ctx := context.Background()
	req := graphql.NewRequest(`
    mutation($input:EventsCreateInput!){
		eventsCreate(eventsCreateInput:$input){
		  timestamp,
		  result{
			id
			image
			holder{
				id
			  name
			}
		  }
		},
	  },
`)
	req.Var("input", m)
	var respData model.EventsAcceptPayload
	if err := c.Run(ctx, req, &respData); err != nil {
		assert.Nil(err)
	}
}

func MutationResolver_RecommendationResponse(t *testing.T, c *graphql.Client, assert *assert.Assertions, m *model.RecommendationResponseInput) {
	ctx := context.Background()
	req := graphql.NewRequest(`
    mutation($input: RecommendationResponseInput!){
		recommendationResponse(recommendationResponseInput:$input){
		  result{
			  id
			  image
			  gender
			}
		}
	  }
`)
	req.Var("input", m)
	respData := model.RecommendationResponsePayload{
		Result: &model.PetProfile{},
	}
	if err := c.Run(ctx, req, &respData); err != nil {
		assert.Nil(err)
	}
	if m.Result {
		assert.NotNil(&respData)
		assert.NotNil(respData.Result)
		assert.NotNil(respData.Result.ID)
		assert.Equal(respData.Result.ID, m.RecommendID)
	} else {
		assert.Nil(respData)
	}
}

func MutationResolver_PetCreate(t *testing.T, c *graphql.Client, assert *assert.Assertions, m *model.PetCreateInput) *model.PetCreatePayload {
	ctx := context.Background()
	req := graphql.NewRequest(`
	mutation($input: PetCreateInput!){
		petCreate(petCreateInput:$input){
		  result{
			  id
			  image
			  gender
			  breed
			  birthday
			  isCastration
			}
		}
	}
	`)
	req.Var("input", m)
	id := ""
	name := ""
	image := ""
	pg := model.PetGender("")
	breed := ""
	bd := ""
	respData := model.PetCreatePayload{
		Result: &model.PetProfile{
			ID:       &id,
			Name:     &name,
			Image:    &image,
			Gender:   &pg,
			Breed:    &breed,
			Birthday: &bd,
		},
	}

	if err := c.Run(ctx, req, &respData); err != nil {
		assert.Nil(err)
	}
	fmt.Println(respData)
	assert.NotNil(respData.Result)
	assert.NotNil(respData.Result.ID)
	return &respData
}

func MutationResolver_PetDelete(c *graphql.Client, assert *assert.Assertions, m *model.PetDeleteInput) *model.PetDeletePayload {
	ctx := context.Background()
	req := graphql.NewRequest(`
	mutation($input: PetDeleteInput!){
		petDelete(petDeleteInput:$input){
		  result
    	timestamp
		}
	}
	`)
	req.Var("input", m)
	var respData model.PetDeletePayload
	if err := c.Run(ctx, req, &respData); err != nil {
		assert.Nil(err)
	}
	return &respData
}
