package services

import (
	"context"
	"github.com/suzushin54/study-graphql-with-go/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	// Add more services here
}

type services struct {
	*userService
	*repositoryService
	*issueService
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	ListUsersByIDs(ctx context.Context, ids []string) ([]*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepositoryID(ctx context.Context, repositoryID string, number int) (*model.Issue, error)
}

func NewServices(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{
			exec: exec,
		},
		repositoryService: &repositoryService{
			exec: exec,
		},
	}
}
