// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_monitoring_alarm", MonitoringAlarmResource())
	tfresource.RegisterResource("oci_monitoring_alarm_suppression", MonitoringAlarmSuppressionResource())
}
