package graphql

import (
	"context"

	"github.com/elritz/go-gqlgen/models"
)

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}
