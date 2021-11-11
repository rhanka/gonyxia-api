/*
 * Onyxia-api
 *
 * Swagger onyxia-api
 *
 * API version: snapshot
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package core

type CatalogWrapper struct {
	Catalog *CatalogWrapper `json:"catalog,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Maintainer string `json:"maintainer,omitempty"`

	Location string `json:"location,omitempty"`

	Status string `json:"status,omitempty"`

	LastUpdateTime int64 `json:"lastUpdateTime,omitempty"`

	Scm string `json:"scm,omitempty"`

	Type_ string `json:"type,omitempty"`
}
