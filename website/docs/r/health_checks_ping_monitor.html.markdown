---
subcategory: "Health Checks"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_health_checks_ping_monitor"
sidebar_current: "docs-oci-resource-health_checks-ping_monitor"
description: |-
  Provides the Ping Monitor resource in Oracle Cloud Infrastructure Health Checks service
---

# oci_health_checks_ping_monitor
This resource provides the Ping Monitor resource in Oracle Cloud Infrastructure Health Checks service.

Creates a ping monitor. Vantage points will be automatically selected if not specified,
and probes will be initiated from each vantage point to each of the targets at the frequency
specified by `intervalInSeconds`.


## Example Usage

```hcl
resource "oci_health_checks_ping_monitor" "test_ping_monitor" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.ping_monitor_display_name
	interval_in_seconds = var.ping_monitor_interval_in_seconds
	protocol = var.ping_monitor_protocol
	targets = var.ping_monitor_targets

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_enabled = var.ping_monitor_is_enabled
	port = var.ping_monitor_port
	timeout_in_seconds = var.ping_monitor_timeout_in_seconds
	vantage_point_names = var.ping_monitor_vantage_point_names
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly and mutable name suitable for display in a user interface.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `interval_in_seconds` - (Required) (Updatable) The monitor interval in seconds. Valid values: 10, 30, and 60. 
* `is_enabled` - (Optional) (Updatable) Enables or disables the monitor. Set to 'true' to launch monitoring. 
* `port` - (Optional) (Updatable) The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - (Required) (Updatable) The protocols for ping probes.
* `targets` - (Required) (Updatable) A list of targets (hostnames or IP addresses) of the probe.
* `timeout_in_seconds` - (Optional) (Updatable) The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - (Optional) (Updatable) A list of names of vantage points from which to execute the probe.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly and mutable name suitable for display in a user interface.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_region` - The region where updates must be made and where results must be fetched from. 
* `id` - The OCID of the resource.
* `interval_in_seconds` - The monitor interval in seconds. Valid values: 10, 30, and 60. 
* `is_enabled` - Enables or disables the monitor. Set to 'true' to launch monitoring. 
* `port` - The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol. 
* `protocol` - The protocols for ping probes.
* `results_url` - A URL for fetching the probe results.
* `targets` - A list of targets (hostnames or IP addresses) of the probe.
* `time_created` - The RFC 3339-formatted creation date and time of the probe. 
* `timeout_in_seconds` - The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors. 
* `vantage_point_names` - A list of names of vantage points from which to execute the probe.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ping Monitor
	* `update` - (Defaults to 20 minutes), when updating the Ping Monitor
	* `delete` - (Defaults to 20 minutes), when destroying the Ping Monitor


## Import

PingMonitors can be imported using the `id`, e.g.

```
$ terraform import oci_health_checks_ping_monitor.test_ping_monitor "id"
```

