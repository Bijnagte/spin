/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ConstraintState struct {
	Attributes *interface{} `json:"attributes,omitempty"`
	Type_ string `json:"type,omitempty"`
	Comment string `json:"comment,omitempty"`
	DeliveryConfigName string `json:"deliveryConfigName,omitempty"`
	Status string `json:"status,omitempty"`
	JudgedAt string `json:"judgedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	EnvironmentName string `json:"environmentName,omitempty"`
	ArtifactVersion string `json:"artifactVersion,omitempty"`
	JudgedBy string `json:"judgedBy,omitempty"`
}
