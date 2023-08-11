package dataintegration

import (
	"fmt"

	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDataintegrationWorkspaceProjectHints.GetIdFn = getDataintegrationWorkspaceProjectId
	exportDataintegrationWorkspaceFolderHints.GetIdFn = getDataintegrationWorkspaceFolderId
	exportDataintegrationWorkspaceApplicationHints.GetIdFn = getDataintegrationWorkspaceApplicationId
	tf_export.RegisterCompartmentGraphs("dataintegration", dataintegrationResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDataintegrationWorkspaceProjectId(resource *tf_export.OCIResource) (string, error) {

	projectKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find projectKey for Dataintegration WorkspaceProject")
	}
	workspaceId := resource.Parent.Id
	return GetWorkspaceProjectCompositeId(projectKey, workspaceId), nil
}

func getDataintegrationWorkspaceFolderId(resource *tf_export.OCIResource) (string, error) {

	folderKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find folderKey for Dataintegration WorkspaceFolder")
	}
	workspaceId := resource.Parent.Id
	return GetWorkspaceFolderCompositeId(folderKey, workspaceId), nil
}

func getDataintegrationWorkspaceApplicationId(resource *tf_export.OCIResource) (string, error) {

	applicationKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find applicationKey for Dataintegration WorkspaceApplication")
	}
	workspaceId := resource.Parent.Id
	return GetWorkspaceApplicationCompositeId(applicationKey, workspaceId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDataintegrationWorkspaceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace",
	DatasourceClass:        "oci_dataintegration_workspaces",
	DatasourceItemsAttr:    "workspaces",
	ResourceAbbreviation:   "workspace",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataintegration.WorkspaceLifecycleStateActive),
	},
}

var exportDataintegrationWorkspaceProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace_project",
	DatasourceClass:        "oci_dataintegration_workspace_projects",
	DatasourceItemsAttr:    "project_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "workspace_project",
	RequireResourceRefresh: true,
}

var exportDataintegrationWorkspaceFolderHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace_folder",
	DatasourceClass:        "oci_dataintegration_workspace_folders",
	DatasourceItemsAttr:    "folder_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "workspace_folder",
	RequireResourceRefresh: true,
}

var exportDataintegrationWorkspaceApplicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace_application",
	DatasourceClass:        "oci_dataintegration_workspace_applications",
	DatasourceItemsAttr:    "application_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "workspace_application",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataintegration.ApplicationLifecycleStateActive),
	},
}

var dataintegrationResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataintegrationWorkspaceHints},
	},
	"oci_dataintegration_workspace": {
		{
			TerraformResourceHints: exportDataintegrationWorkspaceApplicationHints,
			DatasourceQueryParams: map[string]string{
				"workspace_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDataintegrationWorkspaceFolderHints,
			DatasourceQueryParams: map[string]string{
				"workspace_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDataintegrationWorkspaceProjectHints,
			DatasourceQueryParams: map[string]string{
				"workspace_id": "id",
			},
		},
	},
}
