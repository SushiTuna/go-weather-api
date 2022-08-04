/*
 * Weather Aggregator API
 *
 * Aggregator API for collecting weather info of a location.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// Pagination - Pagination of long array
type Pagination struct {
	Total float32 `json:"total,omitempty"`

	Page float32 `json:"page,omitempty"`

	PerPage float32 `json:"perPage,omitempty"`

	HasNext bool `json:"hasNext,omitempty"`

	HasPrevious bool `json:"hasPrevious,omitempty"`

	Results []interface{} `json:"results,omitempty"`
}
