/*
 * aibot open api
 *
 * aibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aibot
// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	APIResponse ApiResponse `json:"APIResponse,omitempty"`
	// 导出文档的下载链接
	Data string `json:"data,omitempty"`
}
