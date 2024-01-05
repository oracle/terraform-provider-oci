// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_apm_synthetics_dedicated_vantage_point", ApmSyntheticsDedicatedVantagePointDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_dedicated_vantage_points", ApmSyntheticsDedicatedVantagePointsDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_monitor", ApmSyntheticsMonitorDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_monitors", ApmSyntheticsMonitorsDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_on_premise_vantage_point", ApmSyntheticsOnPremiseVantagePointDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_on_premise_vantage_point_worker", ApmSyntheticsOnPremiseVantagePointWorkerDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_on_premise_vantage_point_workers", ApmSyntheticsOnPremiseVantagePointWorkersDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_on_premise_vantage_points", ApmSyntheticsOnPremiseVantagePointsDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_public_vantage_point", ApmSyntheticsPublicVantagePointDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_public_vantage_points", ApmSyntheticsPublicVantagePointsDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_result", ApmSyntheticsResultDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_script", ApmSyntheticsScriptDataSource())
	tfresource.RegisterDatasource("oci_apm_synthetics_scripts", ApmSyntheticsScriptsDataSource())
}
