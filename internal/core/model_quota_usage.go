/*
 * Onyxia-api
 *
 * Swagger onyxia-api
 *
 * API version: snapshot
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package core

type QuotaUsage struct {
	Spec *Quota `json:"spec,omitempty"`

	Usage *Quota `json:"usage,omitempty"`
}