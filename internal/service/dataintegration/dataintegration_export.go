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
	exportDataintegrationWorkspaceImportRequestHints.GetIdFn = getDataintegrationWorkspaceImportRequestId
	exportDataintegrationWorkspaceExportRequestHints.GetIdFn = getDataintegrationWorkspaceExportRequestId
	exportDataintegrationWorkspaceApplicationPatchHints.GetIdFn = getDataintegrationWorkspaceApplicationPatchId
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

func getDataintegrationWorkspaceImportRequestId(resource *tf_export.OCIResource) (string, error) {

	importRequestKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find importRequestKey for Dataintegration WorkspaceImportRequest")
	}
	workspaceId := resource.Parent.Id
	return GetWorkspaceImportRequestCompositeId(importRequestKey, workspaceId), nil
}

func getDataintegrationWorkspaceExportRequestId(resource *tf_export.OCIResource) (string, error) {

	exportRequestKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find exportRequestKey for Dataintegration WorkspaceExportRequest")
	}
	workspaceId := resource.Parent.Id
	return GetWorkspaceExportRequestCompositeId(exportRequestKey, workspaceId), nil
}

func getDataintegrationWorkspaceApplicationPatchId(resource *tf_export.OCIResource) (string, error) {

	applicationKey, ok := resource.Parent.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find applicationKey for Dataintegration WorkspaceApplicationPatch")
	}
	patchKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find patchKey for Dataintegration WorkspaceApplicationPatch")
	}
	workspaceId := resource.Parent.SourceAttributes["workspace_id"].(string)
	return GetWorkspaceApplicationPatchCompositeId(applicationKey, patchKey, workspaceId), nil
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

var exportDataintegrationWorkspaceImportRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace_import_request",
	DatasourceClass:        "oci_dataintegration_workspace_import_requests",
	DatasourceItemsAttr:    "import_request_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "workspace_import_request",
	RequireResourceRefresh: true,
}

var exportDataintegrationWorkspaceExportRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace_export_request",
	DatasourceClass:        "oci_dataintegration_workspace_export_requests",
	DatasourceItemsAttr:    "export_request_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "workspace_export_request",
	RequireResourceRefresh: true,
}

var exportDataintegrationWorkspaceApplicationPatchHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace_application_patch",
	DatasourceClass:        "oci_dataintegration_workspace_application_patches",
	DatasourceItemsAttr:    "patch_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "workspace_application_patch",
	RequireResourceRefresh: true,
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
			TerraformResourceHints: exportDataintegrationWorkspaceExportRequestHints,
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
			TerraformResourceHints: exportDataintegrationWorkspaceImportRequestHints,
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
	"oci_dataintegration_workspace_application": {
		{
			TerraformResourceHints: exportDataintegrationWorkspaceApplicationPatchHints,
			DatasourceQueryParams: map[string]string{
				"application_key": "key",
				"workspace_id":    "workspace_id",
			},
		},
	},
}
