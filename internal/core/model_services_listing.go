/*
 * Onyxia-api
 *
 * Swagger onyxia-api
 *
 * API version: snapshot
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package core

type ServicesListing struct {
	Apps []Service `json:"apps,omitempty"`

	Groups []Group `json:"groups,omitempty"`
}
