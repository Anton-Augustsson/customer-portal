/*
 * Customer Portal API
 *
 * An example customer portal OpenAPI specification
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Camera struct {

	Id int64 `json:"id"`

	Name string `json:"name,omitempty"`

	Location Coordinates `json:"location,omitempty"`
}
