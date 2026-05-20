// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package costad

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_costad_cost_alert_subscription", CostadCostAlertSubscriptionDataSource())
	tfresource.RegisterDatasource("oci_costad_cost_alert_subscriptions", CostadCostAlertSubscriptionsDataSource())
	tfresource.RegisterDatasource("oci_costad_cost_anomaly_event", CostadCostAnomalyEventDataSource())
	tfresource.RegisterDatasource("oci_costad_cost_anomaly_event_analytics", CostadCostAnomalyEventAnalyticsDataSource())
	tfresource.RegisterDatasource("oci_costad_cost_anomaly_events", CostadCostAnomalyEventsDataSource())
	tfresource.RegisterDatasource("oci_costad_cost_anomaly_monitor", CostadCostAnomalyMonitorDataSource())
	tfresource.RegisterDatasource("oci_costad_cost_anomaly_monitors", CostadCostAnomalyMonitorsDataSource())
}
