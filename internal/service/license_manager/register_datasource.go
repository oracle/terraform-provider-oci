// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_license_manager_configuration", LicenseManagerConfigurationDataSource())
	tfresource.RegisterDatasource("oci_license_manager_license_metric", LicenseManagerLicenseMetricDataSource())
	tfresource.RegisterDatasource("oci_license_manager_license_record", LicenseManagerLicenseRecordDataSource())
	tfresource.RegisterDatasource("oci_license_manager_license_records", LicenseManagerLicenseRecordsDataSource())
	tfresource.RegisterDatasource("oci_license_manager_product_license", LicenseManagerProductLicenseDataSource())
	tfresource.RegisterDatasource("oci_license_manager_product_license_consumers", LicenseManagerProductLicenseConsumersDataSource())
	tfresource.RegisterDatasource("oci_license_manager_product_licenses", LicenseManagerProductLicensesDataSource())
	tfresource.RegisterDatasource("oci_license_manager_top_utilized_product_licenses", LicenseManagerTopUtilizedProductLicensesDataSource())
	tfresource.RegisterDatasource("oci_license_manager_top_utilized_resources", LicenseManagerTopUtilizedResourcesDataSource())
}
