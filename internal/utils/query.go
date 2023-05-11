package utils

import (
	"strconv"

	"github.com/deepfake-db/internal/models"
	"github.com/gin-gonic/gin"
)

func ProcessQuery(c *gin.Context) (*models.Pagination, *models.Query) {
	// setting up default values
	limit := 25
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()

	localQuery := models.Query{}

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		case "imageId":
			localQuery.ImageId = queryValue
		case "imageUrl":
			localQuery.ImageUrl = queryValue
		case "asserter":
			localQuery.Asserter = queryValue
		case "assertionId":
			localQuery.AssertionId = queryValue
		case "resolved":
			resolved, _ := strconv.ParseBool(queryValue)
			localQuery.Resolved = resolved
			localQuery.ResolvedQueried = true
		case "seen":
			seen, _ := strconv.ParseBool(queryValue)
			localQuery.Seen = seen
			localQuery.SeenQueried = true
		}
	}

	pagination := models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
	return &pagination, &localQuery
}

func ImageAssertionFromQuery(q *models.Query) *models.ImageAssertion {
	var imgAssertion models.ImageAssertion

	if q.ImageId != "" {
		imgAssertion.ImageId = q.ImageId
	}
	if q.ImageUrl != "" {
		imgAssertion.ImageId = q.ImageUrl
	}
	if q.Asserter != "" {
		imgAssertion.Asserter = q.Asserter
	}
	if q.AssertionId != "" {
		imgAssertion.AssertionId = q.AssertionId
	}
	if q.ResolvedQueried {
		imgAssertion.Resolved = q.Resolved
	}
	if q.SeenQueried {
		imgAssertion.Seen = q.Seen
	}

	return &imgAssertion
}
