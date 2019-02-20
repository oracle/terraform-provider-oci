---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_ping_monitors"
sidebar_current: "docs-oci-datasource-health_checks-ping_monitors"
description: |-
  Provides the list of Ping Monitors in Oracle Cloud Infrastructure Health Checks service
---

# Data Source: oci_health_checks_ping_monitors
This data source provides the list of Ping Monitors in Oracle Cloud Infrastructure Health Checks service.

Gets a list of configured ping monitors.

Results are paginated based on `page` and `limit`.  The `opc-next-page` header provides
a URL for fetching the next page.


## Example Usage

```hcl
data "oci_health_checks_ping_monitors" "test_ping_monitors" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.ping_monitor_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Filters results by compartment.
* `display_name` - (Optional) Filters results that exactly match the `displayName` field.


## Attributes Reference

The following attributes are exported:

* `ping_monitors` - The list of ping_monitors.

### PingMonitor Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly and mutable name suitable for display in a user interface.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the resource.
* `interval_in_seconds` - The monitor interval in seconds. Valid values: 10, 30, and 60. 
* `is_enabled` - Enables or disables the monitor. Set to 'true' to launch monitoring. 
* `port` - The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - The protocols for ping probes.
* `results_url` - A URL for fetching the probe results.
* `targets` - An array of A target hostname or IP address of the probe.
* `timeout_in_seconds` - The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - An array of The name of a vantage point from which to execute the probe.

