package graph

import (
	"context"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/pkg/errors"
	"github.com/suzushin54/study-graphql-with-go/graph/model"
	"github.com/suzushin54/study-graphql-with-go/graph/services"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{
		Srv: Srv,
	}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("user not found."),
		}
	}

	indexes := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexes[ID] = i
	}

	users, err := u.Srv.ListUsersByIDs(ctx, IDs)

	for _, user := range users {
		var result *dataloader.Result[*model.User]
		if err != nil {
			result = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			result = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexes[user.ID]] = result
	}

	return results
}
