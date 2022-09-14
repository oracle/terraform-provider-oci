package artifacts

import (
	"fmt"

	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportArtifactsContainerRepositoryHints.GetIdFn = getArtifactsContainerRepositoryId
	exportArtifactsContainerImageSignatureHints.GetIdFn = getArtifactsContainerImageSignatureId
	exportArtifactsRepositoryHints.GetIdFn = getArtifactsRepositoryId
	tf_export.RegisterCompartmentGraphs("artifacts", artifactsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getArtifactsRepositoryId(resource *tf_export.OCIResource) (string, error) {
	repositoryId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find repositoryId for Artifacts Respository")
	}
	return repositoryId, nil
}

func getArtifactsContainerRepositoryId(resource *tf_export.OCIResource) (string, error) {

	repositoryId, ok := resource.SourceAttributes["repository_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find repositoryId for Artifacts ContainerRepository")
	}
	return repositoryId, nil
}

func getArtifactsContainerImageSignatureId(resource *tf_export.OCIResource) (string, error) {

	imageSignatureId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find imageSignatureId for Artifacts ContainerImageSignature")
	}
	return imageSignatureId, nil
}

var exportArtifactsContainerConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_artifacts_container_configuration",
	DatasourceClass:      "oci_artifacts_container_configuration",
	ResourceAbbreviation: "container_configuration",
}

// Hints for discovering and exporting this resource to configuration and state files
var exportArtifactsContainerRepositoryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_artifacts_container_repository",
	DatasourceClass:        "oci_artifacts_container_repositories",
	DatasourceItemsAttr:    "container_repository_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "container_repository",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_artifacts.ContainerRepositoryLifecycleStateAvailable),
	},
}

var exportArtifactsContainerImageSignatureHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_artifacts_container_image_signature",
	DatasourceClass:        "oci_artifacts_container_image_signatures",
	DatasourceItemsAttr:    "container_image_signature_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "container_image_signature",
	RequireResourceRefresh: true,
}

var exportArtifactsRepositoryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_artifacts_repository",
	DatasourceClass:        "oci_artifacts_repositories",
	DatasourceItemsAttr:    "repository_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "repository",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_artifacts.RepositoryLifecycleStateAvailable),
	},
}

var artifactsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportArtifactsContainerRepositoryHints},
		{TerraformResourceHints: exportArtifactsContainerImageSignatureHints},
		{TerraformResourceHints: exportArtifactsRepositoryHints},
	},
}
