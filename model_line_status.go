/*
 * aibot open api
 *
 * aibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aibot
// LineStatus struct for LineStatus
type LineStatus struct {
	Status string `json:"status,omitempty"`
	IsLineAvailable bool `json:"isLineAvailable,omitempty"`
	Hint string `json:"hint,omitempty"`
	OccupiedJobList string `json:"occupiedJobList,omitempty"`
}
