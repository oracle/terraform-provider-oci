// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_stack_monitoring_baselineable_metric", StackMonitoringBaselineableMetricDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_baselineable_metrics", StackMonitoringBaselineableMetricsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_baselineable_metrics_evaluate", StackMonitoringBaselineableMetricsEvaluateDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_config", StackMonitoringConfigDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_configs", StackMonitoringConfigsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_discovery_job", StackMonitoringDiscoveryJobDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_discovery_job_logs", StackMonitoringDiscoveryJobLogsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_discovery_jobs", StackMonitoringDiscoveryJobsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_metric_extension", StackMonitoringMetricExtensionDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_metric_extensions", StackMonitoringMetricExtensionsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resource", StackMonitoringMonitoredResourceDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resources", StackMonitoringMonitoredResourcesDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resource_task", StackMonitoringMonitoredResourceTaskDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resource_tasks", StackMonitoringMonitoredResourceTasksDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resource_type", StackMonitoringMonitoredResourceTypeDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resource_types", StackMonitoringMonitoredResourceTypesDataSource())
}
