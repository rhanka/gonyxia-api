/*
 * Onyxia-api
 *
 * Swagger onyxia-api
 *
 * API version: snapshot
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package core

type Event struct {
	Message string `json:"message,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`
}
