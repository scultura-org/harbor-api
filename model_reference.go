/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Reference struct {
	// The parent ID of the reference
	ParentId int64 `json:"parent_id,omitempty"`
	// The child ID of the reference
	ChildId int64 `json:"child_id,omitempty"`
	// The digest of the child artifact
	ChildDigest string       `json:"child_digest,omitempty"`
	Platform    *Platform    `json:"platform,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	// The download URLs
	Urls []string `json:"urls,omitempty"`
}
