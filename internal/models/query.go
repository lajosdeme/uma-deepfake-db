package models

type Query struct {
	ImageId         string `json:"imageId"`
	ImageUrl        string `json:"imageUrl"`
	Asserter        string `json:"asserter"`
	AssertionId     string `json:"assertionId"`
	Resolved        bool   `json:"resolved"`
	Seen            bool   `json:"seen"`
	ResolvedQueried bool
	SeenQueried     bool
}
