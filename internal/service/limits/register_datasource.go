// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package limits

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_limits_limit_definitions", LimitsLimitDefinitionsDataSource())
	tfresource.RegisterDatasource("oci_limits_limit_values", LimitsLimitValuesDataSource())
	tfresource.RegisterDatasource("oci_limits_quota", LimitsQuotaDataSource())
	tfresource.RegisterDatasource("oci_limits_quotas", LimitsQuotasDataSource())
	tfresource.RegisterDatasource("oci_limits_resource_availability", LimitsResourceAvailabilityDataSource())
	tfresource.RegisterDatasource("oci_limits_services", LimitsServicesDataSource())
}
