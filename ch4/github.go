package ch4

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error occured %s", err)
	}
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func SortIssueResultByCreatedDate(issues *[]*Issue) map[string][]*Issue {
	sorted := make(map[string][]*Issue)
	now := time.Now()
	for _, v := range *issues {
		if v.CreatedAt.Month() == now.Month() && v.CreatedAt.Year() == now.Year() {
			sorted["month"] = append(sorted["month"], v)
		}
		if v.CreatedAt.Year() == now.Year() {
			sorted["lastYear"] = append(sorted["lastYear"], v)
		}
		if v.CreatedAt.Year() < now.Year() {
			sorted["past"] = append(sorted["past"], v)
		}
	}

	return sorted
}
