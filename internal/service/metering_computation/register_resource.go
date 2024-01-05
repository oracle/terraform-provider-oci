// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_metering_computation_custom_table", MeteringComputationCustomTableResource())
	tfresource.RegisterResource("oci_metering_computation_query", MeteringComputationQueryResource())
	tfresource.RegisterResource("oci_metering_computation_schedule", MeteringComputationScheduleResource())
	tfresource.RegisterResource("oci_metering_computation_usage", MeteringComputationUsageResource())
	tfresource.RegisterResource("oci_metering_computation_usage_carbon_emission", MeteringComputationUsageCarbonEmissionResource())
	tfresource.RegisterResource("oci_metering_computation_usage_carbon_emissions_query", MeteringComputationUsageCarbonEmissionsQueryResource())
}
