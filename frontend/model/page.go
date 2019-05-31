package model

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevPage int
	NextPage int
	Items    []interface{}
}
