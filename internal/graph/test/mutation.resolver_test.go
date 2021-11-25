package graph_test

import (
	"context"
	"fmt"
	"testing"

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
