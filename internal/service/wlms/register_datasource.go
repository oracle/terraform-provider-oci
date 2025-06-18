// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_wlms_managed_instance", WlmsManagedInstanceDataSource())
	tfresource.RegisterDatasource("oci_wlms_managed_instance_scan_results", WlmsManagedInstanceScanResultsDataSource())
	tfresource.RegisterDatasource("oci_wlms_managed_instance_server", WlmsManagedInstanceServerDataSource())
	tfresource.RegisterDatasource("oci_wlms_managed_instance_server_installed_patches", WlmsManagedInstanceServerInstalledPatchesDataSource())
	tfresource.RegisterDatasource("oci_wlms_managed_instance_servers", WlmsManagedInstanceServersDataSource())
	tfresource.RegisterDatasource("oci_wlms_managed_instances", WlmsManagedInstancesDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain", WlmsWlsDomainDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_agreement_records", WlmsWlsDomainAgreementRecordsDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_applicable_patches", WlmsWlsDomainApplicablePatchesDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_scan_results", WlmsWlsDomainScanResultsDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_server", WlmsWlsDomainServerDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_server_backup", WlmsWlsDomainServerBackupDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_server_backup_content", WlmsWlsDomainServerBackupContentDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_server_backups", WlmsWlsDomainServerBackupsDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_server_installed_patches", WlmsWlsDomainServerInstalledPatchesDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domain_servers", WlmsWlsDomainServersDataSource())
	tfresource.RegisterDatasource("oci_wlms_wls_domains", WlmsWlsDomainsDataSource())
}
