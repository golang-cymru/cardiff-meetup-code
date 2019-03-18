package main

import (
	"context"
	"log"
	"math"
	"os"
	"time"

	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
)

var (
	token string // Github API access token

	logE = log.New(os.Stderr, "", log.Ldate|log.Ltime)

	client *graphql.Client
)

const (
	// APIURL is the github API Graphql endpoint
	APIURL = "https://api.github.com/graphql"
)

type (
	response struct {
		Organization struct {
			MembersWithRole struct {
				TotalCount int `json:"totalCount"`
				PageInfo   struct {
					HasNextPage bool   `json:"hasNextPage"`
					EndCursor   string `json:"endCursor"`
				}
				Edges []user
			}
		}
	}

	user struct {
		HasTwoFactorEnabled bool   `json:"hasTwoFactorEnabled"`
		Role                string `json:"role"`
		Node                struct {
			Login                   string `json:"login"`
			Name                    string `json:"name"`
			ContributionsCollection struct {
				HasAnyContributions bool `json:"hasAnyContributions"`
			}
		}
	}

	userMap map[string]user
)

func main() {
	log.SetOutput(os.Stdout)

	if token = os.Getenv("TOKEN"); len(token) == 0 {
		logE.Fatal("Must set Github API TOKEN env var")
	}

	// Client is safe for reuse between calls
	client = graphql.NewClient(APIURL)

	// TODO
	org := "ONSdigital"
	orgID := "MDEyOk9yZ2FuaXphdGlvbjY4NjUxMzc="
	pageSize := 50

	req := graphql.NewRequest(`
		query($organization: String!, $pageSize: Int!, $from: DateTime, $cursor: String ) {
			organization(login: $organization) {
				membersWithRole(first: $pageSize, after: $cursor) {
					totalCount
					pageInfo{
						hasNextPage
						endCursor
					}
					edges {
						hasTwoFactorEnabled
						role
						node {
							login
							name
							contributionsCollection(from: $from) {
								hasAnyContributions
							}
						}
					}
				}
			}
		}
	`)
	req.Var("organization", org)
	req.Var("orgID", orgID)
	req.Var("pageSize", pageSize)
	req.Var("from", time.Now().AddDate(0, -1, 0).Format(time.RFC3339)) // One month ago

	req.Header.Set("Authorization", "bearer "+token)

	// All users returned from the graphql calls unfiltered
	users := make(map[string]user)

	// Counts
	numAdminUsers := 0

	// Specific lists
	adminsMissing2FA := make(userMap)
	usersWithoutFullname := make(userMap)
	usersWithNoContributionInLastMonth := make(userMap)

	hasNext := true
	var cursor *string // Allow the cursor to be nil to read the first page
	i := 1
	for hasNext {
		var res response
		req.Var("cursor", cursor)

		if err := client.Run(context.Background(), req, &res); err != nil {
			logE.Fatalf(`event="Exiting" error="%v"`, errors.Wrap(err, "client failed to call github api"))
		}

		numPages := math.Ceil(float64(res.Organization.MembersWithRole.TotalCount/pageSize)) + 1

		log.Printf("Read page %d/%v\n", i, numPages)
		i++

		for _, u := range res.Organization.MembersWithRole.Edges {
			users[u.Node.Login] = u
			if u.Role == "ADMIN" {
				numAdminUsers++
			}

			if u.Node.Name == "" {
				usersWithoutFullname[u.Node.Login] = u
			}

			if u.Role == "ADMIN" && !u.HasTwoFactorEnabled {
				adminsMissing2FA[u.Node.Login] = u
			}

			if !u.Node.ContributionsCollection.HasAnyContributions {
				usersWithNoContributionInLastMonth[u.Node.Login] = u
			}
		}

		if hasNext = res.Organization.MembersWithRole.PageInfo.HasNextPage; hasNext {
			cursor = &res.Organization.MembersWithRole.PageInfo.EndCursor
		}
	}

	// spew.Dump(users)

	log.Printf("Found %d users\n", len(users))
	log.Printf("Admin users: %d\n", numAdminUsers)
	log.Printf("No-name users: %d\n", len(usersWithoutFullname))
	log.Printf("Admins missing 2FA: %d\n", len(adminsMissing2FA))
	log.Printf("Users with no activity in last month: %d\n", len(usersWithNoContributionInLastMonth))

}
