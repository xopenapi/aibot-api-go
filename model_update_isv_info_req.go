/*
 * aibot open api
 *
 * aibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aibot
// UpdateIsvInfoReq struct for UpdateIsvInfoReq
type UpdateIsvInfoReq struct {
	TenantSign string `json:"tenantSign"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
	SmsCallbackUrl string `json:"smsCallbackUrl,omitempty"`
}
