// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package health_checks

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_health_checks_http_monitor", HealthChecksHttpMonitorDataSource())
	tfresource.RegisterDatasource("oci_health_checks_http_monitors", HealthChecksHttpMonitorsDataSource())
	tfresource.RegisterDatasource("oci_health_checks_http_probe_results", HealthChecksHttpProbeResultsDataSource())
	tfresource.RegisterDatasource("oci_health_checks_ping_monitor", HealthChecksPingMonitorDataSource())
	tfresource.RegisterDatasource("oci_health_checks_ping_monitors", HealthChecksPingMonitorsDataSource())
	tfresource.RegisterDatasource("oci_health_checks_ping_probe_results", HealthChecksPingProbeResultsDataSource())
	tfresource.RegisterDatasource("oci_health_checks_vantage_points", HealthChecksVantagePointsDataSource())
}
