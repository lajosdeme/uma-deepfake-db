package models

import (
	"gorm.io/gorm"
)

type ImageAssertion struct {
	gorm.Model
	ImageId      string
	ImageUrl     string
	Asserter     string
	AssertionId  string
	Resolved     bool
	Seen         bool
	LoggedAt     int64
	ResolvableAt int64
}
