---
subcategory: "Health Checks"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_http_monitor"
sidebar_current: "docs-oci-datasource-health_checks-http_monitor"
description: |-
  Provides details about a specific Http Monitor in Oracle Cloud Infrastructure Health Checks service
---

# Data Source: oci_health_checks_http_monitor
This data source provides details about a specific Http Monitor resource in Oracle Cloud Infrastructure Health Checks service.

Gets the configuration for the specified monitor.


## Example Usage

```hcl
data "oci_health_checks_http_monitor" "test_http_monitor" {
	#Required
	monitor_id = oci_apm_synthetics_monitor.test_monitor.id
}
```

## Argument Reference

The following arguments are supported:

* `monitor_id` - (Required) The OCID of a monitor.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly and mutable name suitable for display in a user interface.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `headers` - A dictionary of HTTP request headers.

	*Note:* Monitors and probes do not support the use of the `Authorization` HTTP header. 
* `home_region` - The region where updates must be made and where results must be fetched from. 
* `id` - The OCID of the resource.
* `interval_in_seconds` - The monitor interval in seconds. Valid values: 10, 30, and 60. 
* `is_enabled` - Enables or disables the monitor. Set to 'true' to launch monitoring. 
* `method` - The supported HTTP methods available for probes.
* `path` - The optional URL path to probe, including query parameters.
* `port` - The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - The supported protocols available for HTTP probes.
* `results_url` - A URL for fetching the probe results.
* `targets` - A list of targets (hostnames or IP addresses) of the probe.
* `time_created` - The RFC 3339-formatted creation date and time of the probe. 
* `timeout_in_seconds` - The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - A list of names of vantage points from which to execute the probe.

