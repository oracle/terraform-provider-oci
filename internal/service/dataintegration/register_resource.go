// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_dataintegration_workspace", DataintegrationWorkspaceResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_application", DataintegrationWorkspaceApplicationResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_application_patch", DataintegrationWorkspaceApplicationPatchResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_application_schedule", DataintegrationWorkspaceApplicationScheduleResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_application_task_schedule", DataintegrationWorkspaceApplicationTaskScheduleResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_export_request", DataintegrationWorkspaceExportRequestResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_folder", DataintegrationWorkspaceFolderResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_import_request", DataintegrationWorkspaceImportRequestResource())
	tfresource.RegisterResource("oci_dataintegration_workspace_project", DataintegrationWorkspaceProjectResource())
}
