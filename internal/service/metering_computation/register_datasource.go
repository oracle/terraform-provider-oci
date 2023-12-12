// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_metering_computation_average_carbon_emission", MeteringComputationAverageCarbonEmissionDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_clean_energy_usage", MeteringComputationCleanEnergyUsageDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_configuration", MeteringComputationConfigurationDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_custom_table", MeteringComputationCustomTableDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_custom_tables", MeteringComputationCustomTablesDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_queries", MeteringComputationQueriesDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_query", MeteringComputationQueryDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_schedule", MeteringComputationScheduleDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_scheduled_run", MeteringComputationScheduledRunDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_scheduled_runs", MeteringComputationScheduledRunsDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_schedules", MeteringComputationSchedulesDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_usage_carbon_emissions_config", MeteringComputationUsageCarbonEmissionsConfigDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_usage_carbon_emissions_queries", MeteringComputationUsageCarbonEmissionsQueriesDataSource())
	tfresource.RegisterDatasource("oci_metering_computation_usage_carbon_emissions_query", MeteringComputationUsageCarbonEmissionsQueryDataSource())
}
