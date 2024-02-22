/*
 * Customer Portal API
 *
 * An example customer portal OpenAPI specification
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type SubscriberAPI struct {
}

// Post /subscribers
func (api *SubscriberAPI) AddSubscriber(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /subscriber/:subscriberId
func (api *SubscriberAPI) DeleteSubscriber(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /subscriber/:subscriberId
func (api *SubscriberAPI) GetSubscriberById(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /subscribers
func (api *SubscriberAPI) GetSubscribers(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /subscriber/:subscriberId/subscription
func (api *SubscriberAPI) GetSubscribersSubscriptions(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /subscriber/:subscriberId
func (api *SubscriberAPI) UpdateSubscriber(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

