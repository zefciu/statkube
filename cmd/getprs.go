package main

import (
	"os"

	"github.com/Mirantis/statkube/importer"
)

func main() {
	limitStr := os.Args[1]
	limit, err := time.Parse("2006-01-02", limitStr)
	if err != nil {
		panic(err.Error())
	}
	db := db.GetDB()
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		panic("Set GITHUB_TOKEN")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	opt := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 1000},
		State:       "closed",
		Sort:        "updated",
		Direction:   "desc",
	}

	for {
		limitMet := false
		prs, resp, err := client.PullRequests.List(
			"kubernetes", "kubernetes", opt,
		)
		if err != nil {
			panic(err.Error())
		}
		for _, pr := range prs {
			//if pr is updated before limit, break as prs are sorted by updatedAt
			if pr.UpdatedAt.Before(limit) {
				limitMet = true
				break
			}
			handlePR(pr, client, db)
		}
		if resp.NextPage == 0 || limitMet {
			break
		}
		opt.ListOptions.Page = resp.NextPage

	}
}
