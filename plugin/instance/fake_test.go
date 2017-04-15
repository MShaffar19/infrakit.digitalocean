package instance

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/pkg/errors"
)

type fakeTagsService struct {
	expectedErr string
}

func (s *fakeTagsService) TagResources(context.Context, string, *godo.TagResourcesRequest) (*godo.Response, error) {
	if s.expectedErr != "" {
		return nil, errors.New(s.expectedErr)
	}
	return nil, nil
}

type fakeDropletsServices struct {
	createfunc  func(context.Context, *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error)
	expectedErr string
}

func (s *fakeDropletsServices) List(context.Context, *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
	if s.expectedErr != "" {
		return nil, nil, errors.New(s.expectedErr)
	}
	return []godo.Droplet{}, nil, nil
}

func (s *fakeDropletsServices) Get(context.Context, int) (*godo.Droplet, *godo.Response, error) {
	if s.expectedErr != "" {
		return nil, nil, errors.New(s.expectedErr)
	}
	return &godo.Droplet{}, nil, nil
}

func (s *fakeDropletsServices) Create(ctx context.Context, req *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error) {
	if s.expectedErr != "" {
		return nil, nil, errors.New(s.expectedErr)
	}
	if s.createfunc != nil {
		return s.createfunc(ctx, req)
	}
	return &godo.Droplet{}, nil, nil
}

func (s *fakeDropletsServices) Delete(context.Context, int) (*godo.Response, error) {
	if s.expectedErr != "" {
		return nil, errors.New(s.expectedErr)
	}
	return nil, nil
}