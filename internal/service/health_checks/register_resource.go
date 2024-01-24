// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package health_checks

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_health_checks_http_monitor", HealthChecksHttpMonitorResource())
	tfresource.RegisterResource("oci_health_checks_http_probe", HealthChecksHttpProbeResource())
	tfresource.RegisterResource("oci_health_checks_ping_monitor", HealthChecksPingMonitorResource())
	tfresource.RegisterResource("oci_health_checks_ping_probe", HealthChecksPingProbeResource())
}
