// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package models

import (
	fmt "fmt"
	io "io"
	strconv "strconv"

	bug "github.com/MichaelMure/git-bug/bug"
)

type Authored interface{}
type BugConnection struct {
	Edges      []BugEdge      `json:"edges"`
	Nodes      []bug.Snapshot `json:"nodes"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}
type BugEdge struct {
	Cursor string       `json:"cursor"`
	Node   bug.Snapshot `json:"node"`
}
type CommentConnection struct {
	Edges      []CommentEdge `json:"edges"`
	Nodes      []bug.Comment `json:"nodes"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}
type CommentEdge struct {
	Cursor string      `json:"cursor"`
	Node   bug.Comment `json:"node"`
}
type OperationConnection struct {
	Edges      []OperationEdge `json:"edges"`
	Nodes      []bug.Operation `json:"nodes"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}
type OperationEdge struct {
	Cursor string        `json:"cursor"`
	Node   bug.Operation `json:"node"`
}
type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusClosed Status = "CLOSED"
)

func (e Status) IsValid() bool {
	switch e {
	case StatusOpen, StatusClosed:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
