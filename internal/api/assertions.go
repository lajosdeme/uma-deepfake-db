package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepfake-db/internal/config"
	"github.com/deepfake-db/internal/models"
	"github.com/deepfake-db/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func GetAssertions(c *gin.Context) {
	pagination, query := utils.ProcessQuery(c)
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	var assertions []models.ImageAssertion
	assertion := utils.ImageAssertionFromQuery(query)
	fmt.Println(assertion.AssertionId)

	var result *gorm.DB

	if query.ResolvedQueried {
		result = queryBuilder.Model(&models.ImageAssertion{}).Where(assertion).Where("Resolved = ?", query.Resolved).Find(&assertions)
	}

	if query.SeenQueried {
		result = queryBuilder.Model(&models.ImageAssertion{}).Where(assertion).Where("Seen = ?", query.Seen).Find(&assertions)
	}

	if !query.ResolvedQueried && !query.SeenQueried {
		result = queryBuilder.Model(&models.ImageAssertion{}).Where(assertion).Find(&assertions)
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": assertions,
	})
}

func GetNotifications(c *gin.Context) {
	//var db *gorm.DB
	pagination, _ := utils.ProcessQuery(c)
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	var assertions []models.ImageAssertion

	asserter := c.Params.ByName("asserter")
	fmt.Println("asserter: ", asserter)

	result := queryBuilder.Model(&models.ImageAssertion{}).Where("Asserter = ?", asserter).Where("Seen = ?", false).Where("Disputed = ?", false).Find(&assertions)

	fmt.Println("assertions length: ", len(assertions))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	t := time.Now().Unix()

	var resolvableAssertions []models.ImageAssertion

	for _, assertion := range assertions {
		if assertion.ResolvableAt < t {
			resolvableAssertions = append(resolvableAssertions, assertion)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resolvableAssertions,
	})
}

func SetNotificationSeen(c *gin.Context) {
	var assertion models.ImageAssertion
	if err := c.BindJSON(&assertion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("assertion id: ", assertion.AssertionId)

	var localAssertion models.ImageAssertion
	res := config.DB.First(&localAssertion, "assertion_id = ?", assertion.AssertionId)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	assertion.Seen = true

	res = config.DB.Model(&localAssertion).Updates(assertion)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
