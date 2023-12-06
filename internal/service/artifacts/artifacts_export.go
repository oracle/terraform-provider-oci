package artifacts

import (
	"fmt"

	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportArtifactsContainerRepositoryHints.GetIdFn = getArtifactsContainerRepositoryId
	exportArtifactsContainerImageSignatureHints.GetIdFn = getArtifactsContainerImageSignatureId
	exportArtifactsGenericArtifactHints.GetIdFn = getArtifactsGenericArtifactId
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
	repositoryId, ok := resource.SourceAttributes["id"].(string)
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

func getArtifactsGenericArtifactId(resource *tf_export.OCIResource) (string, error) {

	artifactId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find artifactId for Artifacts GenericArtifact")
	}
	return artifactId, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportArtifactsContainerConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_artifacts_container_configuration",
	DatasourceClass:      "oci_artifacts_container_configuration",
	ResourceAbbreviation: "container_configuration",
}

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
	DiscoverableLifecycleStates: []string{
		string(oci_artifacts.ContainerImageSignatureLifecycleStateAvailable),
	},
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

var exportArtifactsGenericArtifactHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_artifacts_generic_artifact",
	DatasourceClass:        "oci_artifacts_generic_artifacts",
	DatasourceItemsAttr:    "generic_artifact_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "generic_artifact",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_artifacts.GenericArtifactLifecycleStateAvailable),
	},
}

var artifactsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportArtifactsContainerConfigurationHints},
		{TerraformResourceHints: exportArtifactsContainerRepositoryHints},
		{TerraformResourceHints: exportArtifactsContainerImageSignatureHints},
		{TerraformResourceHints: exportArtifactsRepositoryHints},
	},
	"oci_artifacts_repository": {
		{
			TerraformResourceHints: exportArtifactsGenericArtifactHints,
			DatasourceQueryParams: map[string]string{
				"repository_id": "id",
			},
		},
	},
}
