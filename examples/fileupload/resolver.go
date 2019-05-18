package fileupload

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/samlitowitz/graphqlc-gen-gqlgen/examples/fileupload/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	panic("not implemented")
}
