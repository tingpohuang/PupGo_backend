package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tingpo/pupgobackend/internal/graph/model"
)

func (r *notificationResolver) UserID(ctx context.Context, obj *model.Notification) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *notificationResolver) Userid(ctx context.Context, obj *model.Notification) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Notification returns generated.NotificationResolver implementation.
// func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

type notificationResolver struct{ *Resolver }
