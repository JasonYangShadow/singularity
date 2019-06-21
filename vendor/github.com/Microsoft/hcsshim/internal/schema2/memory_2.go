/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type Memory2 struct {
	SizeInMB int32 `json:"SizeInMB,omitempty"`

	AllowOvercommit bool `json:"AllowOvercommit,omitempty"`

	EnableHotHint bool `json:"EnableHotHint,omitempty"`

	EnableColdHint bool `json:"EnableColdHint,omitempty"`

	EnableEpf bool `json:"EnableEpf,omitempty"`

	// EnableDeferredCommit is private in the schema. If regenerated need to add back.
	EnableDeferredCommit bool `json:"EnableDeferredCommit,omitempty"`
}
