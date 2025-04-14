// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dblm

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_dblm_patch_management", DblmPatchManagementDataSource())
	tfresource.RegisterDatasource("oci_dblm_patch_management_databases", DblmPatchManagementDatabasesDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability", DblmVulnerabilityDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability_aggregated_vulnerability_data", DblmVulnerabilityAggregatedVulnerabilityDataDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability_notifications", DblmVulnerabilityNotificationsDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability_resources", DblmVulnerabilityResourcesDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability_scan", DblmVulnerabilityScanDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability_scans", DblmVulnerabilityScansDataSource())
	tfresource.RegisterDatasource("oci_dblm_vulnerability_vulnerabilities", DblmVulnerabilityVulnerabilitiesDataSource())
}
