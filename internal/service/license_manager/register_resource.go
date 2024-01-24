// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_license_manager_configuration", LicenseManagerConfigurationResource())
	tfresource.RegisterResource("oci_license_manager_license_record", LicenseManagerLicenseRecordResource())
	tfresource.RegisterResource("oci_license_manager_product_license", LicenseManagerProductLicenseResource())
}
