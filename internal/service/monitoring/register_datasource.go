// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_monitoring_alarm", MonitoringAlarmDataSource())
	tfresource.RegisterDatasource("oci_monitoring_alarm_history_collection", MonitoringAlarmHistoryCollectionDataSource())
	tfresource.RegisterDatasource("oci_monitoring_alarm_statuses", MonitoringAlarmStatusesDataSource())
	tfresource.RegisterDatasource("oci_monitoring_alarm_suppression", MonitoringAlarmSuppressionDataSource())
	tfresource.RegisterDatasource("oci_monitoring_alarm_suppressions", MonitoringAlarmSuppressionsDataSource())
	tfresource.RegisterDatasource("oci_monitoring_alarms", MonitoringAlarmsDataSource())
	tfresource.RegisterDatasource("oci_monitoring_metric_data", MonitoringMetricDataDataSource())
	tfresource.RegisterDatasource("oci_monitoring_metrics", MonitoringMetricsDataSource())
}
