// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waa

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_waa_web_app_acceleration", WaaWebAppAccelerationDataSource())
	tfresource.RegisterDatasource("oci_waa_web_app_acceleration_policies", WaaWebAppAccelerationPoliciesDataSource())
	tfresource.RegisterDatasource("oci_waa_web_app_acceleration_policy", WaaWebAppAccelerationPolicyDataSource())
	tfresource.RegisterDatasource("oci_waa_web_app_accelerations", WaaWebAppAccelerationsDataSource())
}
