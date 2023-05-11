package services

import (
	"context"
	"github.com/suzushin54/study-graphql-with-go/graph/db"
	"github.com/suzushin54/study-graphql-with-go/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

type issueService struct {
	exec boil.ContextExecutor
}

func (i *issueService) GetIssueByNumber(ctx context.Context, repositoryID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repositoryID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}
	return convertIssue(issue), nil
}

func convertIssue(issue *db.Issue) *model.Issue {
	issueURL, err := model.UnmarshalURI(issue.URL)
	if err != nil {
		log.Println("invalid URI", issue.URL)
	}

	return &model.Issue{
		ID:     issue.ID,
		URL:    issueURL,
		Title:  issue.Title,
		Closed: issue.Closed != 0,
		Number: int(issue.Number),
		Author: &model.User{
			ID: issue.Author,
		},
		Repository: &model.Repository{
			ID: issue.Repository,
		},
	}
}
